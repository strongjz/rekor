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

package x509

import (
	"bytes"
	"crypto"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	validator "github.com/go-playground/validator/v10"
	"github.com/sigstore/sigstore/pkg/cryptoutils"
	sigsig "github.com/sigstore/sigstore/pkg/signature"
)

// EmailAddressOID defined by https://oidref.com/1.2.840.113549.1.9.1
var EmailAddressOID asn1.ObjectIdentifier = []int{1, 2, 840, 113549, 1, 9, 1}

type Signature struct {
	signature []byte
}

// NewSignature creates and validates an x509 signature object
func NewSignature(r io.Reader) (*Signature, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return &Signature{
		signature: b,
	}, nil
}

// CanonicalValue implements the pki.Signature interface
func (s Signature) CanonicalValue() ([]byte, error) {
	return s.signature, nil
}

// Verify implements the pki.Signature interface
func (s Signature) Verify(r io.Reader, k interface{}, opts ...sigsig.VerifyOption) error {
	if len(s.signature) == 0 {
		//lint:ignore ST1005 X509 is proper use of term
		return fmt.Errorf("X509 signature has not been initialized")
	}

	key, ok := k.(*PublicKey)
	if !ok {
		return fmt.Errorf("invalid public key type for: %v", k)
	}

	p := key.key
	if p == nil {
		switch {
		case key.cert != nil:
			p = key.cert.c.PublicKey
		case len(key.certs) > 0:
			if err := verifyCertChain(key.certs); err != nil {
				return err
			}
			p = key.certs[0].PublicKey
		default:
			return errors.New("no public key found")
		}
	}

	verifier, err := sigsig.LoadVerifier(p, crypto.SHA256)
	if err != nil {
		return err
	}
	return verifier.VerifySignature(bytes.NewReader(s.signature), r, opts...)
}

// PublicKey Public Key that follows the x509 standard
type PublicKey struct {
	key   interface{}
	cert  *cert
	certs []*x509.Certificate
}

type cert struct {
	c *x509.Certificate
	b []byte
}

// NewPublicKey implements the pki.PublicKey interface
func NewPublicKey(r io.Reader) (*PublicKey, error) {
	rawPub, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	block, rest := pem.Decode(rawPub)
	if block == nil {
		return nil, errors.New("invalid public key: failure decoding PEM")
	}

	// Handle certificate chain, concatenated PEM-encoded certificates
	if len(rest) > 0 {
		// Support up to 10 certificates in a chain, to avoid parsing extremely long chains
		certs, err := cryptoutils.UnmarshalCertificatesFromPEMLimited(rawPub, 10)
		if err != nil {
			return nil, err
		}
		return &PublicKey{certs: certs}, nil
	}

	switch block.Type {
	case string(cryptoutils.PublicKeyPEMType):
		key, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return &PublicKey{key: key}, nil
	case string(cryptoutils.CertificatePEMType):
		c, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, err
		}
		return &PublicKey{
			cert: &cert{
				c: c,
				b: block.Bytes,
			}}, nil
	}
	return nil, fmt.Errorf("invalid public key: cannot handle type %v", block.Type)
}

// CanonicalValue implements the pki.PublicKey interface
func (k PublicKey) CanonicalValue() (encoded []byte, err error) {

	switch {
	case k.key != nil:
		encoded, err = cryptoutils.MarshalPublicKeyToPEM(k.key)
	case k.cert != nil:
		encoded, err = cryptoutils.MarshalCertificateToPEM(k.cert.c)
	case k.certs != nil:
		encoded, err = cryptoutils.MarshalCertificatesToPEM(k.certs)
	default:
		err = fmt.Errorf("x509 public key has not been initialized")
	}

	return
}

func (k PublicKey) CryptoPubKey() crypto.PublicKey {
	if k.cert != nil {
		return k.cert.c.PublicKey
	}
	if len(k.certs) > 0 {
		return k.certs[0].PublicKey
	}
	return k.key
}

// EmailAddresses implements the pki.PublicKey interface
func (k PublicKey) EmailAddresses() []string {
	var names []string
	var cert *x509.Certificate
	if k.cert != nil {
		cert = k.cert.c
	} else if len(k.certs) > 0 {
		cert = k.certs[0]
	}
	if cert != nil {
		for _, name := range cert.EmailAddresses {
			validate := validator.New()
			errs := validate.Var(name, "required,email")
			if errs == nil {
				names = append(names, strings.ToLower(name))
			}
		}
	}
	return names
}

func verifyCertChain(certChain []*x509.Certificate) error {
	if len(certChain) == 0 {
		return errors.New("no certificate chain provided")
	}
	// No certificate chain to verify
	if len(certChain) == 1 {
		return nil
	}
	rootPool := x509.NewCertPool()
	rootPool.AddCert(certChain[len(certChain)-1])
	subPool := x509.NewCertPool()
	for _, c := range certChain[1 : len(certChain)-1] {
		subPool.AddCert(c)
	}
	if _, err := certChain[0].Verify(x509.VerifyOptions{
		Roots:         rootPool,
		Intermediates: subPool,
		// Allow any key usage
		KeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
	}); err != nil {
		return err
	}
	return nil
}

func CertChainToPEM(certChain []*x509.Certificate) ([]byte, error) {
	var pemBytes bytes.Buffer
	for _, cert := range certChain {
		if err := pem.Encode(&pemBytes, &pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw}); err != nil {
			return nil, err
		}
	}
	return pemBytes.Bytes(), nil
}

func ParseTimestampCertChain(pemBytes []byte) ([]*x509.Certificate, error) {
	certChain := []*x509.Certificate{}
	var block *pem.Block
	block, pemBytes = pem.Decode(pemBytes)
	for ; block != nil; block, pemBytes = pem.Decode(pemBytes) {
		if block.Type == "CERTIFICATE" {
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return nil, err
			}
			certChain = append(certChain, cert)
		} else {
			return nil, errors.New("invalid block type")
		}
	}
	if len(certChain) == 0 {
		return nil, errors.New("no valid certificates in chain")
	}
	// Verify cert chain for timestamping
	roots := x509.NewCertPool()
	intermediates := x509.NewCertPool()
	for _, cert := range certChain[1:(len(certChain) - 1)] {
		intermediates.AddCert(cert)
	}
	roots.AddCert(certChain[len(certChain)-1])
	if _, err := certChain[0].Verify(x509.VerifyOptions{
		Roots:         roots,
		KeyUsages:     []x509.ExtKeyUsage{x509.ExtKeyUsageTimeStamping},
		Intermediates: intermediates,
	}); err != nil {
		return nil, err
	}
	return certChain, nil
}
