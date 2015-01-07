# tlstext

Simple mapping of TLS Versions and Cipher Suites to Strings

The [Go TLS cipher suites and TLS versions](http://golang.org/pkg/crypto/tls/#pkg-constants) are untyped or `uint16` and without a string representation.

This package provides simple functions `tlstxt.Version` and
`tlstext.CipherSuite` that provide the raw value to string translations.


