// Package tlstext provides simple functions VersionText and
// CipherSuiteText that provide the raw value to string translations.
package tlstext

//go:generate go run generator/main.go
import (
	"crypto/tls"
	"fmt"
)

// reverse map of binary TLS Version to string
var versionMap = map[uint16]string{
	0x0300: "SSL30",
	0x0301: "TLS10",
	0x0302: "TLS11",
	0x0303: "TLS12",
	0x0304: "TLS13",
}

// map of TLS token name to hex value
var versionStringMap = map[string]uint16{
	"SSL30": 0x0300,
	"TLS10": 0x0301,
	"TLS11": 0x0302,
	"TLS12": 0x0303,
	"TLS13": 0x0304,
}

// Version maps a TLS version to a string, or the hex
// representation if unknown.
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

// CipherFromString returns the uint16 value of the cipher name or 0 if unknown
func CipherFromString(s string) uint16 {
	return cipherStringMap[s]
}

// VersionFromString returns the uint16 value of a TLS version string, or 0 if unknown
func VersionFromString(s string) uint16 {
	return versionStringMap[s]
}

// VersionFromConnection returns a string representation of CipherSuite
// or empty string if not TLS
func VersionFromConnection(t *tls.ConnectionState) string {
	if t == nil {
		return ""
	}
	return Version(t.Version)
}

// CipherSuiteFromConnection returns a string representation of
// CipherSuite or empty string if not TLS
//
func CipherSuiteFromConnection(t *tls.ConnectionState) string {
	if t == nil {
		return ""
	}
	return CipherSuite(t.CipherSuite)
}
