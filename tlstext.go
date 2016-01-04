package tlstext

import (
	"crypto/tls"
	"fmt"
)

// This package provides simple functions `VersionText` and
// `CipherSuiteText` that provide the raw value to string translations.

// reverse map of binary TLS Version to string
var versionMap = map[uint16]string{
	0x0300: "SSL30",
	0x0301: "TLS10",
	0x0302: "TLS11",
	0x0303: "TLS12",
}

// reverse map of binary TLS Cipher suite to string
var cipherMap = map[uint16]string{
	0x0005: "TLS_RSA_WITH_RC4_128_SHA",
	0x000a: "TLS_RSA_WITH_3DES_EDE_CBC_SHA",
	0x002f: "TLS_RSA_WITH_AES_128_CBC_SHA",
	0x0035: "TLS_RSA_WITH_AES_256_CBC_SHA",
	0xc007: "TLS_ECDHE_ECDSA_WITH_RC4_128_SHA",
	0xc009: "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA",
	0xc00a: "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA",
	0xc011: "TLS_ECDHE_RSA_WITH_RC4_128_SHA",
	0xc012: "TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA",
	0xc013: "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
	0xc014: "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
	0xc02f: "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
	0xc02b: "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
}

// Version maps a TLS version to a string, or the hex
//  representation if unknown.
func Version(x uint16) string {
	s, ok := versionMap[x]
	if !ok {
		return fmt.Sprintf("%04x", x)
	}
	return s
}

// CipherSuite maps a TLS Cipher Suite to a string or the hex
// representation if unknown
func CipherSuite(x uint16) string {
	s, ok := cipherMap[x]
	if !ok {
		return fmt.Sprintf("%04x", x)
	}
	return s
}

// VersionFromConnection retuns a string representation of CipherSuite
//  or empty string if not TLS
//
func VersionFromConnection(t *tls.ConnectionState) string {
	if t == nil {
		return ""
	}
	return Version(t.Version)
}

// CipherSuiteFromConnection retuns a string representation of
// CipherSuite or empty string if not TLS
//
func CipherSuiteFromConnection(t *tls.ConnectionState) string {
	if t == nil {
		return ""
	}
	return CipherSuite(t.CipherSuite)
}
