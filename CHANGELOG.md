# v0.6.0

Notice: The server side remote fetching of resources will be removed in the next release

## Enhancements

* Create EntryID for new artifacts and return EntryID to user (#623)
* Add search through inactive shards for GET by UUID (#750)
* Add in configmap to release for sharding config (#766)
* set p.Block after parsing; other cleanup (#759)
* Add index to hashed intoto envelope (#761)
* Add the SHA256 digest of the intoto payload into the rekor entry (#764)
* Add support for providing certificate chain for X509 signature types (#747)
* Specify public key for inactive shards in shard config (#746)
* Use active tree on server startup (#727)
* Require tlog_id when inactive shard config file is passed in (#739)
* Replace `trillian_log_server.log_id_ranges` flag with a config file (#742)
* Update loginfo API endpoint to return information about inactive shards (#738)
* Refactor rekor-cli loginfo (#734)
* Get log proofs by Tree ID (#733)
* Return virtual index when creating and getting a log entry (#725)
* Clearer logging for createAndInitTree (#724)
* Change TreeID to be of type `string` instead of `int64` (#712)
* Switch to using the swag library for pointer manipulation. (#719)
* Make the loginfo command a bit more future/backwards proof. (#718)
* Use logRangesFlag in API, route reads based on TreeID (#671)
* Set rekor-cli User-Agent header on requests (#684)
* create namespace for rekor config in yaml. (#680)
* add securityContext to deployment. (#678)
* Move k8s objects out of the default namespace (#674)

## Bug Fixes

* Fix search without sha prefix (#767)
* Fix link in types README (#765)
* fix typo in filename (#758)
* fix build date format for version command (#745)
* fix merge conflict (#720)

## Documentation

* Add documentation about Alpine type (#697)
* update security process link (#685)
* Add intoto type documentation (#679)
* Add docs about API stabilitly and deprecation policy (#661)

## Others

* Bump github.com/go-openapi/spec from 0.20.4 to 0.20.5 (#768)
* Bump anchore/sbom-action from 0.9.0 to 0.10.0 (#763)
* Bump github/codeql-action from 2.1.7 to 2.1.8 (#762)
* Update release jobs and trillian images (#756)
* Bump sigstore/cosign-installer from 2.1.0 to 2.2.0 (#757)
* Bump anchore/sbom-action from 0.8.0 to 0.9.0 (#754)
* Bump codecov/codecov-action from 2.1.0 to 3 (#753)
* Bump github/codeql-action from 2.1.6 to 2.1.7 (#752)
* Bump google-github-actions/auth from 0.6.0 to 0.7.0 (#751)
* Bump github/codeql-action from 1.1.5 to 2.1.6 (#748)
* Bump anchore/sbom-action from 0.7.0 to 0.8.0 (#743)
* Bump google.golang.org/protobuf from 1.27.1 to 1.28.0 (#744)
* Bump github.com/go-openapi/runtime from 0.23.2 to 0.23.3 (#740)
* Bump github/codeql-action from 1.1.4 to 1.1.5 (#736)
* Use reusuable release workflow in sigstore/sigstore (#729)
* Fix copy/paste mistake in repo name. (#730)
* Bump github.com/spf13/cobra from 1.3.0 to 1.4.0 (#728)
* Bump golang from `ca70980` to `c7c9458` (#722)
* Bump google.golang.org/grpc from 1.44.0 to 1.45.0 (#723)
* Add sharding e2e test to Github Actions (#714)
* Bump github.com/go-playground/validator/v10 from 10.10.0 to 10.10.1 (#717)
* Bump github/codeql-action from 1.1.3 to 1.1.4 (#716)
* Add trillian container to existing release. (#715)
* Bump golang from `0168c35` to `ca70980` (#707)
* Mirror signed release images from GCR to GHCR as part of release (#701)
* Bump anchore/sbom-action from 0.6.0 to 0.7.0 (#709)
* Bump github.com/go-openapi/runtime from 0.23.1 to 0.23.2 (#710)
* Bump sigstore/cosign-installer from 2.0.1 to 2.1.0 (#708)
* Generate release yaml artifact. (#702)
* Bump actions/upload-artifact from 2.3.1 to 3 (#704)
* Go update to 1.17.8 and cosign to 1.6.0 (#705)
* Consistent parenthesis use in Makefile (#700)
* add code coverage to pull request. (#676)
* Bump actions/checkout from 2.4.0 to 3 (#698)
* Bump goreleaser/goreleaser-action from 2.9.0 to 2.9.1 (#696)
* Bump actions/setup-go from 2.2.0 to 3.0.0 (#694)
* Bump github.com/secure-systems-lab/go-securesystemslib (#695)
* Bump golangci/golangci-lint-action from 3.0.0 to 3.1.0 (#693)
* Bump goreleaser/goreleaser-action from 2.8.1 to 2.9.0 (#692)
* Bump golangci/golangci-lint-action from 2.5.2 to 3 (#691)
* Bump github/codeql-action from 1.1.2 to 1.1.3 (#690)
* Bump github.com/go-openapi/runtime from 0.23.0 to 0.23.1 (#689)
* explicitly set permissions for github actions (#687)
* Bump sigstore/cosign-installer from 2.0.0 to 2.0.1 (#686)
* Bump ossf/scorecard-action from 1.0.3 to 1.0.4 (#683)
* Bump github/codeql-action from 1.1.0 to 1.1.2 (#682)
* Bump actions/github-script from 5.1.0 to 6 (#669)
* Bump github/codeql-action from 1.0.32 to 1.1.0 (#668)
* update cross-build and dockerfile to use go 1.17.7 (#666)
* Bump gopkg.in/ini.v1 from 1.66.3 to 1.66.4 (#664)
* Bump actions/setup-go from 2.1.5 to 2.2.0 (#663)
* Bump golang from `301609e` to `fff998d` (#662)
* use upstream k8s version lib (#657)
* Bump github/codeql-action from 1.0.31 to 1.0.32 (#659)
* Bump go.uber.org/zap from 1.20.0 to 1.21.0 (#660)
* Bump github.com/go-openapi/strfmt from 0.21.1 to 0.21.2 (#656)
* Bump github.com/go-openapi/runtime from 0.22.0 to 0.23.0 (#655)
* Update the warning text for the GA release. (#654)
* attempting to fix codeowners file (#653)
* update release job (#651)
* Bump google-github-actions/auth from 0.5.0 to 0.6.0 (#652)

## Contributors

* Asra Ali (@asraa)
* Bob Callaway (@bobcallaway)
* Carlos Tadeu Panato Junior (@cpanato)
* Dan Lorenc (@dlorenc)
* Eddie Zaneski (@eddiezane)
* Hayden Blauzvern (@haydentherapper)
* John Speed Meyers
* Kenny Leung (@k4leung4)
* Lily Sturmann (@lkatalin)
* Priya Wadhwa (@priyawadhwa)
* Scott Nichols (@n3wscott)

# v0.5.0

## Highlights

* Add Rekor logo to README (https://github.com/sigstore/rekor/pull/650)
* update API calls to v5 (https://github.com/sigstore/rekor/pull/591)
* Refactor helm type to remove intermediate state. (https://github.com/sigstore/rekor/pull/575)
* Refactor the shard map parsing so we can pass it down into the API object. (https://github.com/sigstore/rekor/pull/564)
* Refactor the alpine type to reduce intermediate state. (https://github.com/sigstore/rekor/pull/573)

## Enhancements

* Add logic to GET artifacts via old or new UUID (https://github.com/sigstore/rekor/pull/587)
* helpful error message for hashedrekord types (https://github.com/sigstore/rekor/pull/605)
* Set Accept header in dynamic counter requests (https://github.com/sigstore/rekor/pull/594)
* Add sharding package and update validators (https://github.com/sigstore/rekor/pull/583)
* rekor-cli: show the url in case of error (https://github.com/sigstore/rekor/pull/581)
* Enable parsing of incomplete minisign keys, to enable re-indexing. (https://github.com/sigstore/rekor/pull/567)
* Cleanups on the TUF pluggable type. (https://github.com/sigstore/rekor/pull/563)
* Refactor the RPM type to remove more intermediate state. (https://github.com/sigstore/rekor/pull/566)
* Do some cleanups of the jar type to remove intermediate state. (https://github.com/sigstore/rekor/pull/561)

## Others

* Update Makefile (https://github.com/sigstore/rekor/pull/621)
* update version comments since dependabot doesn't do it (https://github.com/sigstore/rekor/pull/617)
* Use workload identity provider instead of GitHub Secret for GCR access (https://github.com/sigstore/rekor/pull/600)
* add OSSF scorecard action (https://github.com/sigstore/rekor/pull/599)
* enable the sbom for rekor releases (https://github.com/sigstore/rekor/pull/586)
* Point to the official website (instead of a 404) (https://github.com/sigstore/rekor/pull/580)
* add milestone to closed prs (https://github.com/sigstore/rekor/pull/574)
* Add a Makefile target for the "ko apply" step. (https://github.com/sigstore/rekor/pull/572)
* types/README.md: Corrected documentation link (https://github.com/sigstore/rekor/pull/568)

## Dependencies Updates

* Bump github.com/prometheus/client_golang from 1.12.0 to 1.12.1 (https://github.com/sigstore/rekor/pull/636)
* Bump github.com/go-openapi/runtime from 0.21.1 to 0.22.0 (https://github.com/sigstore/rekor/pull/635)
* Bump github.com/go-openapi/swag from 0.19.15 to 0.20.0 (https://github.com/sigstore/rekor/pull/634)
* Bump golang from `f71d4ca` to `301609e` (https://github.com/sigstore/rekor/pull/627)
* Bump golang from `0fa6504` to `f71d4ca` (https://github.com/sigstore/rekor/pull/624)
* Bump google.golang.org/grpc from 1.43.0 to 1.44.0 (https://github.com/sigstore/rekor/pull/622)
* Bump github/codeql-action from 1.0.29 to 1.0.30 (https://github.com/sigstore/rekor/pull/619)
* Bump ossf/scorecard-action from 1.0.1 to 1.0.2 (https://github.com/sigstore/rekor/pull/618)
* bump swagger and go mod tidy (https://github.com/sigstore/rekor/pull/616)
* Bump github.com/go-openapi/runtime from 0.21.0 to 0.21.1 (https://github.com/sigstore/rekor/pull/614)
* Bump github.com/go-openapi/errors from 0.20.1 to 0.20.2 (https://github.com/sigstore/rekor/pull/613)
* Bump google-github-actions/auth from 0.4.4 to 0.5.0 (https://github.com/sigstore/rekor/pull/612)
* Bump github/codeql-action from 1.0.28 to 1.0.29 (https://github.com/sigstore/rekor/pull/611)
* Bump gopkg.in/ini.v1 from 1.66.2 to 1.66.3 (https://github.com/sigstore/rekor/pull/608)
* Bump github.com/google/go-cmp from 0.5.6 to 0.5.7 (https://github.com/sigstore/rekor/pull/609)
* Update github/codeql-action requirement to 8a4b243fbf9a03a93e93a71c1ec257347041f9c4 (https://github.com/sigstore/rekor/pull/606)
* Bump github.com/prometheus/client_golang from 1.11.0 to 1.12.0 (https://github.com/sigstore/rekor/pull/607)
* Bump ossf/scorecard-action from 0fe1afdc40f536c78e3dc69147b91b3ecec2cc8a to 1.0.1 (https://github.com/sigstore/rekor/pull/603)
* Bump goreleaser/goreleaser-action from 2.8.0 to 2.8.1 (https://github.com/sigstore/rekor/pull/602)
* Bump golang from `8c0269d` to `0fa6504` (https://github.com/sigstore/rekor/pull/597)
* Pin dependencies in github action workflows and Dockerfile (https://github.com/sigstore/rekor/pull/595)
* update release image to use go 1.17.6 (https://github.com/sigstore/rekor/pull/589)
* Bump golang from 1.17.5 to 1.17.6 (https://github.com/sigstore/rekor/pull/588)
* Bump go.uber.org/goleak from 1.1.11 to 1.1.12 (https://github.com/sigstore/rekor/pull/585)
* Bump go.uber.org/zap from 1.19.1 to 1.20.0 (https://github.com/sigstore/rekor/pull/584)
* Bump github.com/go-playground/validator/v10 from 10.9.0 to 10.10.0 (https://github.com/sigstore/rekor/pull/579)
* Bump actions/github-script from 4 to 5 (https://github.com/sigstore/rekor/pull/577)

## Contributors

* Asra Ali (@asraa)
* Bob Callaway (@bobcallaway)
* Carlos Tadeu Panato Junior (@cpanato)
* Dan Lorenc (@dlorenc)
* Jason Hall (@imjasonh)
* Lily Sturmann (@lkatalin)
* Morten Linderud (@Foxboron)
* Nathan Smith (@nsmith5)
* Sylvestre Ledru (@sylvestre)
* Trishank Karthik Kuppusamy (@trishankatdatadog)

# v0.4.0

## Highlights

* Adds hashed rekord type that can be used to upload signatures along with the hashed content signed (https://github.com/sigstore/rekor/pull/501)

## Enhancements

* Update the schema to match that of Trillian repo. The map specific (https://github.com/sigstore/rekor/pull/528)
* allow setting the user-agent string sent from the client (https://github.com/sigstore/rekor/pull/521)
* update key usage for ts cert (https://github.com/sigstore/rekor/pull/504)
* api/index/retrieve: allow searching on indicies with sha1 hashes (https://github.com/sigstore/rekor/pull/499)
* Only include Attestation data if attestation storage enabled (https://github.com/sigstore/rekor/pull/494)
* Fuzzing RequestFromRekor API (https://github.com/sigstore/rekor/pull/488)
* Included pprof for profiling the application. (https://github.com/sigstore/rekor/pull/485)
* refactor release and add signing (https://github.com/sigstore/rekor/pull/483)
* More verbose error message for redis connection failure (https://github.com/sigstore/rekor/pull/479) (https://github.com/sigstore/rekor/pull/480)
* Fixed modtime for reproducible goreleaser (https://github.com/sigstore/rekor/pull/473)
* add goreleaser and cloudbuild for releases (https://github.com/sigstore/rekor/pull/443)
* Add dynamic JS tree size counter (https://github.com/sigstore/rekor/pull/468)
* check that entry UUID == leafHash of returned entry (https://github.com/sigstore/rekor/pull/469)
* chore: upgrade cosign version (https://github.com/sigstore/rekor/pull/465)
* Reproducible builds with trimpath (https://github.com/sigstore/rekor/pull/464)
* correct links, add Table of Contents of sorts (https://github.com/sigstore/rekor/pull/449)
* update go tuf for rsa key impl (https://github.com/sigstore/rekor/pull/446)
* Canonicalize JSON before inserting into trillian (https://github.com/sigstore/rekor/pull/445)
* Export search UUIDs field (https://github.com/sigstore/rekor/pull/438)
* Add a flag to start specifying log index ranges for virtual indices. (https://github.com/sigstore/rekor/pull/435)
* Cleanup some initialization/flag parsing in rekor-server. (https://github.com/sigstore/rekor/pull/433)
* Drop 404 errors down to a warning. (https://github.com/sigstore/rekor/pull/426)
* Cleanup the output of search (the text goes to stderr not stdout). (https://github.com/sigstore/rekor/pull/421)
* remove extradata field from types (https://github.com/sigstore/rekor/pull/418)
* Update usage of ./cmd/rekor-cli/ from `rekor` to `rekor-cli` (https://github.com/sigstore/rekor/pull/417)
* Add TUF type (https://github.com/sigstore/rekor/pull/383)
* Updates to INSTALLATION.md notes (https://github.com/sigstore/rekor/pull/415)
* Update snippets to use `console` type for snippets (https://github.com/sigstore/rekor/pull/410)
* version: add way to display a version when using go get or go install (https://github.com/sigstore/rekor/pull/405)
* Use an in memory timestamping key (https://github.com/sigstore/rekor/pull/402)
* Links are case sensitive (https://github.com/sigstore/rekor/pull/401)
* Installation guide (https://github.com/sigstore/rekor/pull/400)
* Add a SignedTimestampNote (https://github.com/sigstore/rekor/pull/397)
* Provide instructions on verifying releases (https://github.com/sigstore/rekor/pull/399)
* rekor-server: add html page when humans reach the server via the browser (https://github.com/sigstore/rekor/pull/394)
* use go modules to track tools (https://github.com/sigstore/rekor/pull/395)

## Bug Fixes

* bug: fix minisign prehashed entries (https://github.com/sigstore/rekor/pull/639)
* fix timestamp addition and unmarshal (https://github.com/sigstore/rekor/pull/525)
* Correct & parallelize tests (https://github.com/sigstore/rekor/pull/522)
* Fix fuzz go.sum issue (https://github.com/sigstore/rekor/pull/509)
* fix validation error (https://github.com/sigstore/rekor/pull/503)
* Correct Helm index keys (https://github.com/sigstore/rekor/pull/474)
* Fix a bug in x509 certificate handling. (https://github.com/sigstore/rekor/pull/461)
* Fix a conflict from parallel dependabot merges. (https://github.com/sigstore/rekor/pull/456)
* fix tuf metadata marshalling (https://github.com/sigstore/rekor/pull/447)
* Switch DSSE provider to go-securesystemslib (https://github.com/sigstore/rekor/pull/442)
* fix unmarshalling sth (https://github.com/sigstore/rekor/pull/409)
* Fix port flag override (https://github.com/sigstore/rekor/pull/396)
* makefile: small fix on the makefile for the rekor-server (https://github.com/sigstore/rekor/pull/393)

## Dependencies Updates

* Bump github.com/spf13/viper from 1.9.0 to 1.10.0 (https://github.com/sigstore/rekor/pull/531)
* Bump sigstore/cosign-installer from 1.3.1 to 1.4.1 (https://github.com/sigstore/rekor/pull/530)
* Bump the DSSE signing library. (https://github.com/sigstore/rekor/pull/529)
* Bump golang from 1.17.4 to 1.17.5 (https://github.com/sigstore/rekor/pull/527)
* Bump golang from 1.17.3 to 1.17.4 (https://github.com/sigstore/rekor/pull/523)
* Bump gopkg.in/ini.v1 from 1.66.0 to 1.66.2 (https://github.com/sigstore/rekor/pull/520)
* Bump github.com/mitchellh/mapstructure from 1.4.2 to 1.4.3 (https://github.com/sigstore/rekor/pull/517)
* Bump github.com/secure-systems-lab/go-securesystemslib (https://github.com/sigstore/rekor/pull/516)
* Bump gopkg.in/ini.v1 from 1.64.0 to 1.66.0 (https://github.com/sigstore/rekor/pull/513)
* Upgraded go-playground/validator module to v10 (https://github.com/sigstore/rekor/pull/507)
* Bump gopkg.in/ini.v1 from 1.63.2 to 1.64.0 (https://github.com/sigstore/rekor/pull/495)
* Bump github.com/go-openapi/strfmt from 0.21.0 to 0.21.1 (https://github.com/sigstore/rekor/pull/510)
* Bump the trillian import to v1.4.0. (https://github.com/sigstore/rekor/pull/502)
* Bump the trillian versions to v1.4.0 in our docker-compose setup. (https://github.com/sigstore/rekor/pull/500)
* update go.mod for go-fuzz (https://github.com/sigstore/rekor/pull/496)
* Bump sigstore/cosign-installer from 1.3.0 to 1.3.1 (https://github.com/sigstore/rekor/pull/491)
* Bump golang from 1.17.2 to 1.17.3 (https://github.com/sigstore/rekor/pull/482)
* Bump google.golang.org/grpc from 1.41.0 to 1.42.0 (https://github.com/sigstore/rekor/pull/478)
* Bump actions/checkout from 2.3.5 to 2.4.0 (https://github.com/sigstore/rekor/pull/477)
* Bump github.com/go-openapi/runtime from 0.20.0 to 0.21.0 (https://github.com/sigstore/rekor/pull/470)
* bump go-swagger to v0.28.0 (https://github.com/sigstore/rekor/pull/463)
* Bump github.com/in-toto/in-toto-golang from 0.3.2 to 0.3.3 (https://github.com/sigstore/rekor/pull/459)
* Bump actions/checkout from 2.3.4 to 2.3.5 (https://github.com/sigstore/rekor/pull/458)
* Bump github.com/mediocregopher/radix/v4 from 4.0.0-beta.1 to 4.0.0 (https://github.com/sigstore/rekor/pull/460)
* Bump github.com/go-openapi/runtime from 0.19.31 to 0.20.0 (https://github.com/sigstore/rekor/pull/451)
* Bump github.com/go-openapi/spec from 0.20.3 to 0.20.4 (https://github.com/sigstore/rekor/pull/454)
* Bump github.com/go-openapi/validate from 0.20.2 to 0.20.3 (https://github.com/sigstore/rekor/pull/453)
* Bump github.com/go-openapi/strfmt from 0.20.2 to 0.20.3 (https://github.com/sigstore/rekor/pull/452)
* Bump github.com/go-openapi/loads from 0.20.2 to 0.20.3 (https://github.com/sigstore/rekor/pull/450)
* Bump golang from 1.17.1 to 1.17.2 (https://github.com/sigstore/rekor/pull/448)
* Bump google.golang.org/grpc from 1.40.0 to 1.41.0 (https://github.com/sigstore/rekor/pull/441)
* Bump golang.org/x/mod from 0.5.0 to 0.5.1 (https://github.com/sigstore/rekor/pull/440)
* Bump github.com/spf13/viper from 1.8.1 to 1.9.0 (https://github.com/sigstore/rekor/pull/439)
* Bump gopkg.in/ini.v1 from 1.63.0 to 1.63.2 (https://github.com/sigstore/rekor/pull/437)
* Bump github.com/mitchellh/mapstructure from 1.4.1 to 1.4.2 (https://github.com/sigstore/rekor/pull/436)
* Bump gocloud to v0.24.0. (https://github.com/sigstore/rekor/pull/434)
* Bump golang from 1.17.0 to 1.17.1 (https://github.com/sigstore/rekor/pull/432)
* Bump go.uber.org/zap from 1.19.0 to 1.19.1 (https://github.com/sigstore/rekor/pull/431)
* Bump gopkg.in/ini.v1 from 1.62.0 to 1.63.0 (https://github.com/sigstore/rekor/pull/429)
* Bump github.com/go-openapi/runtime from 0.19.30 to 0.19.31 (https://github.com/sigstore/rekor/pull/425)
* Bump github.com/go-openapi/errors from 0.20.0 to 0.20.1 (https://github.com/sigstore/rekor/pull/423)
* Bump github.com/go-openapi/strfmt from 0.20.1 to 0.20.2 (https://github.com/sigstore/rekor/pull/422)
* Bump golang from 1.16.7 to 1.17.0 (https://github.com/sigstore/rekor/pull/413)
* Bump golang.org/x/mod from 0.4.2 to 0.5.0 (https://github.com/sigstore/rekor/pull/412)
* Bump google.golang.org/grpc from 1.39.1 to 1.40.0 (https://github.com/sigstore/rekor/pull/411)
* Bump github.com/go-openapi/runtime from 0.19.29 to 0.19.30 (https://github.com/sigstore/rekor/pull/408)
* Bump go.uber.org/zap from 1.18.1 to 1.19.0 (https://github.com/sigstore/rekor/pull/407)
* Bump golang from 1.16.6 to 1.16.7 (https://github.com/sigstore/rekor/pull/403)
* Bump google.golang.org/grpc from 1.39.0 to 1.39.1 (https://github.com/sigstore/rekor/pull/404)


## Contributors

* Aditya Sirish (@adityasaky)
* Andrew Block (@sabre1041)
* Asra Ali (@asraa)
* Axel Simon (@axelsimon)
* Batuhan Apaydın (@developer-guy)
* Bob Callaway (@bobcallaway)
* Carlos Panato (@cpanato)
* Dan Lorenc (@dlorenc)
* Dan Luhring (@luhring)
* Harry Fallows (@harryfallows)
* Hector Fernandez (@hectorj2f)
* Jake Sanders (@dekkagaijin)
* Jason Hall (@imjasonh)
* Lily Sturmann (@lkatalin)
* Luke Hinds (@lukehinds)
* Marina Moore (@mnm678)
* Mikhail Swift (@mikhailswift)
* Naveen Srinivasan (@naveensrinivasan)
* Robert James Hernandez (@sarcasticadmin)
* Santiago Torres (@SantiagoTorres)
* Tiziano Santoro (@tiziano88)
* Trishank Karthik Kuppusamy (@trishankatdatadog)
* Ville Aikas (@vaikas)
* kpcyrd (@kpcyrd)
