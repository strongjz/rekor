//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package intoto

import (
	"bytes"
	"context"
	"crypto"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/in-toto/in-toto-golang/in_toto"
	"github.com/secure-systems-lab/go-securesystemslib/dsse"
	"github.com/spf13/viper"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/pkg/errors"

	"github.com/sigstore/rekor/pkg/generated/models"
	"github.com/sigstore/rekor/pkg/log"
	"github.com/sigstore/rekor/pkg/pki"
	"github.com/sigstore/rekor/pkg/pki/x509"
	"github.com/sigstore/rekor/pkg/types"
	"github.com/sigstore/rekor/pkg/types/intoto"
	"github.com/sigstore/sigstore/pkg/signature"
	"github.com/sigstore/sigstore/pkg/signature/options"
)

const (
	APIVERSION = "0.0.1"
)

func init() {
	if err := intoto.VersionMap.SetEntryFactory(APIVERSION, NewEntry); err != nil {
		log.Logger.Panic(err)
	}
}

type V001Entry struct {
	IntotoObj models.IntotoV001Schema
	keyObj    pki.PublicKey
	env       dsse.Envelope
}

func (v V001Entry) APIVersion() string {
	return APIVERSION
}

func NewEntry() types.EntryImpl {
	return &V001Entry{}
}

func (v V001Entry) IndexKeys() ([]string, error) {
	var result []string

	h := sha256.Sum256([]byte(v.env.Payload))
	payloadKey := "sha256:" + hex.EncodeToString(h[:])
	result = append(result, payloadKey)

	switch v.env.PayloadType {
	case in_toto.PayloadType:
		if v.IntotoObj.Content == nil || v.IntotoObj.Content.Hash == nil {
			log.Logger.Info("IntotoObj content or hash is nil")
			return result, nil
		}
		hashkey := strings.ToLower(fmt.Sprintf("%s:%s", *v.IntotoObj.Content.Hash.Algorithm, *v.IntotoObj.Content.Hash.Value))
		result = append(result, hashkey)

		statement, err := parseStatement(v.env.Payload)
		if err != nil {
			return result, err
		}
		for _, s := range statement.Subject {
			for alg, ds := range s.Digest {
				result = append(result, alg+":"+ds)
			}
		}
		// Not all in-toto statements will contain a SLSA provenance predicate.
		// See https://github.com/in-toto/attestation/blob/main/spec/README.md#predicate
		// for other predicates.
		if predicate, err := parseSlsaPredicate(v.env.Payload); err == nil {
			if predicate.Predicate.Materials != nil {
				for _, s := range predicate.Predicate.Materials {
					for alg, ds := range s.Digest {
						result = append(result, alg+":"+ds)
					}
				}
			}
		}
	default:
		log.Logger.Infof("Unknown in_toto Statement Type: %s", v.env.PayloadType)
	}
	return result, nil
}

func parseStatement(p string) (*in_toto.Statement, error) {
	ps := in_toto.Statement{}
	payload, err := base64.StdEncoding.DecodeString(p)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(payload, &ps); err != nil {
		return nil, err
	}
	return &ps, nil
}

func parseSlsaPredicate(p string) (*in_toto.ProvenanceStatement, error) {
	predicate := in_toto.ProvenanceStatement{}
	payload, err := base64.StdEncoding.DecodeString(p)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(payload, &predicate); err != nil {
		return nil, err
	}
	return &predicate, nil
}

func (v *V001Entry) Unmarshal(pe models.ProposedEntry) error {
	it, ok := pe.(*models.Intoto)
	if !ok {
		return errors.New("cannot unmarshal non Intoto v0.0.1 type")
	}

	var err error
	if err := types.DecodeEntry(it.Spec, &v.IntotoObj); err != nil {
		return err
	}

	// field validation
	if err := v.IntotoObj.Validate(strfmt.Default); err != nil {
		return err
	}

	v.keyObj, err = x509.NewPublicKey(bytes.NewReader(*v.IntotoObj.PublicKey))
	if err != nil {
		return err
	}

	return v.validate()
}

func (v *V001Entry) Canonicalize(ctx context.Context) ([]byte, error) {
	if v.keyObj == nil {
		return nil, errors.New("cannot canonicalze empty key")
	}
	pk, err := v.keyObj.CanonicalValue()
	if err != nil {
		return nil, err
	}
	pkb := strfmt.Base64(pk)

	h := sha256.Sum256([]byte(v.IntotoObj.Content.Envelope))

	canonicalEntry := models.IntotoV001Schema{
		PublicKey: &pkb,
		Content: &models.IntotoV001SchemaContent{
			Hash: &models.IntotoV001SchemaContentHash{
				Algorithm: swag.String(models.IntotoV001SchemaContentHashAlgorithmSha256),
				Value:     swag.String(hex.EncodeToString(h[:])),
			},
		},
	}
	attestation := v.Attestation()
	if attestation != nil {
		decodedAttestation, err := base64.StdEncoding.DecodeString(string(attestation))
		if err != nil {
			return nil, errors.Wrap(err, "decoding attestation")
		}
		attH := sha256.Sum256(decodedAttestation)
		canonicalEntry.Content.PayloadHash = &models.IntotoV001SchemaContentPayloadHash{
			Algorithm: swag.String(models.IntotoV001SchemaContentHashAlgorithmSha256),
			Value:     swag.String(hex.EncodeToString(attH[:])),
		}
	}

	itObj := models.Intoto{}
	itObj.APIVersion = swag.String(APIVERSION)
	itObj.Spec = &canonicalEntry

	return json.Marshal(&itObj)
}

// validate performs cross-field validation for fields in object
func (v *V001Entry) validate() error {
	// TODO handle multiple
	pk := v.keyObj.(*x509.PublicKey)

	// This also gets called in the CLI, where we won't have this data
	if v.IntotoObj.Content.Envelope == "" {
		return nil
	}
	vfr, err := signature.LoadVerifier(pk.CryptoPubKey(), crypto.SHA256)
	if err != nil {
		return err
	}
	dsseVerifier, err := dsse.NewEnvelopeSigner(&verifier{
		v:   vfr,
		pub: pk,
	})
	if err != nil {
		return err
	}

	if v.IntotoObj.Content.Envelope == "" {
		return nil
	}

	if err := json.Unmarshal([]byte(v.IntotoObj.Content.Envelope), &v.env); err != nil {
		return err
	}

	if _, err := dsseVerifier.Verify(&v.env); err != nil {
		return err
	}
	return nil
}

func (v *V001Entry) Attestation() []byte {
	storageSize := base64.StdEncoding.DecodedLen(len(v.env.Payload))
	if storageSize > viper.GetInt("max_attestation_size") {
		log.Logger.Infof("Skipping attestation storage, size %d is greater than max %d", storageSize, viper.GetInt("max_attestation_size"))
		return nil
	}
	return []byte(v.env.Payload)
}

type verifier struct {
	s   signature.Signer
	v   signature.Verifier
	pub crypto.PublicKey
}

func (v *verifier) KeyID() (string, error) {
	return "", nil
}

func (v *verifier) Public() crypto.PublicKey {
	return v.pub
}

func (v *verifier) Sign(data []byte) (sig []byte, err error) {
	if v.s == nil {
		return nil, errors.New("nil signer")
	}
	sig, err = v.s.SignMessage(bytes.NewReader(data), options.WithCryptoSignerOpts(crypto.SHA256))
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func (v *verifier) Verify(data, sig []byte) error {
	if v.v == nil {
		return errors.New("nil verifier")
	}
	return v.v.VerifySignature(bytes.NewReader(sig), bytes.NewReader(data))
}

func (v V001Entry) CreateFromArtifactProperties(_ context.Context, props types.ArtifactProperties) (models.ProposedEntry, error) {
	returnVal := models.Intoto{}

	var err error
	artifactBytes := props.ArtifactBytes
	if artifactBytes == nil {
		if props.ArtifactPath == nil {
			return nil, errors.New("path to artifact file must be specified")
		}
		if props.ArtifactPath.IsAbs() {
			return nil, errors.New("intoto envelopes cannot be fetched over HTTP(S)")
		}
		artifactBytes, err = ioutil.ReadFile(filepath.Clean(props.ArtifactPath.Path))
		if err != nil {
			return nil, err
		}
	}
	publicKeyBytes := props.PublicKeyBytes
	if publicKeyBytes == nil {
		if props.PublicKeyPath == nil {
			return nil, errors.New("public key must be provided to verify signature")
		}
		publicKeyBytes, err = ioutil.ReadFile(filepath.Clean(props.PublicKeyPath.Path))
		if err != nil {
			return nil, fmt.Errorf("error reading public key file: %w", err)
		}
	}
	kb := strfmt.Base64(publicKeyBytes)

	re := V001Entry{
		IntotoObj: models.IntotoV001Schema{
			Content: &models.IntotoV001SchemaContent{
				Envelope: string(artifactBytes),
			},
			PublicKey: &kb,
		},
	}
	h := sha256.Sum256([]byte(re.IntotoObj.Content.Envelope))
	re.IntotoObj.Content.Hash = &models.IntotoV001SchemaContentHash{
		Algorithm: swag.String(models.IntotoV001SchemaContentHashAlgorithmSha256),
		Value:     swag.String(hex.EncodeToString(h[:])),
	}

	returnVal.Spec = re.IntotoObj
	returnVal.APIVersion = swag.String(re.APIVersion())

	return &returnVal, nil
}
