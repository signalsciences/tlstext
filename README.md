[![Build Status](https://travis-ci.org/client9/tlstext.svg?branch=master)](https://travis-ci.org/client9/tlstext) [![Go Report Card](http://goreportcard.com/badge/client9/tlstext)](http://goreportcard.com/report/client9/tlstext) [![GoDoc](https://godoc.org/github.com/client9/tlstext?status.svg)](https://godoc.org/github.com/client9/tlstext) [![Coverage](http://gocover.io/_badge/github.com/client9/tlstext)](http://gocover.io/github.com/client9/tlstext) [![license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://raw.githubusercontent.com/client9/tlstext/master/LICENSE)

tlstext - simple mapping of TLS Versions and Cipher Suites to strings

The [Go TLS cipher suites and TLS versions](http://golang.org/pkg/crypto/tls/#pkg-constants) are untyped
or are `uint16` and without a string representation.  This also means
the tool [stringer](https://godoc.org/golang.org/x/tools/cmd/stringer)
can not be used.

This package provides simple functions `tlstxt.Version` and
`tlstext.CipherSuite` that provide the raw value to string translations.

This intentionally does not use the constants in
[tls/cipher_suites.go](https://golang.org/src/crypto/tls/cipher_suites.go)
since they are dependent on the version of Go used.

The values are generated directly from the [IANA assignments](http://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-4)

