# tlstext

Simple mapping of TLS Versions and Cipher Suites to Strings

The [Go TLS cipher suites and TLS versions](http://golang.org/pkg/crypto/tls/#pkg-constants) are untyped or `uint16` and without a string representation.

This package provides simple functions `tlstxt.Version` and
`tlstext.CipherSuite` that provide the raw value to string translations.

Unfortuantely the tool
[stringer](https://godoc.org/golang.org/x/tools/cmd/stringer)
can not be used since neither cipher suite not tls version are a
`type` and don't satisfy the `Stringer` interface.

This intentionally does not use the constants in
[tls/cipher_suites.go](https://golang.org/src/crypto/tls/cipher_suites.go)
since they are dependant on the version of Go used.


