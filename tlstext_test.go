package tlstext

import (
	"testing"
	"crypto/tls"
)

func TestCipherSuite(t *testing.T) {
	value := uint16(0xc02b)
	expected := "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"
	actual := CipherSuite(value)
	if expected != actual {
		t.Errorf("For %d, expected %q got %q", value, expected, actual)
	}

	value = uint16(1234)
	expected = "04d2"
	actual = CipherSuite(value)
	if expected != actual {
		t.Errorf("For %d, expected %q got %q", value, expected, actual)
	}
}

func TestVersion(t *testing.T) {
	value := uint16(0x0303)
	expected := "TLS12"
	actual := Version(value)
	if expected != actual {
		t.Errorf("For %d, expected %q got %q", value, expected, actual)
	}

	value = uint16(1234)
	expected = "04d2"
	actual = Version(value)
	if expected != actual {
		t.Errorf("For %d, expected %q got %q", value, expected, actual)
	}
}

func TestFromConnection(t *testing.T) {
	actual := VersionFromConnection(nil)
	if actual != "" {
		t.Errorf("Expected empty version from nil input, got %q", actual)
	}
	actual = CipherSuiteFromConnection(nil)
	if actual != "" {
		t.Errorf("Expected empty cipher suite from nil input, got %q", actual)
	}

	c := tls.ConnectionState{
		Version: uint16(0x0303),
		CipherSuite: uint16(0xc02b),
	}
	expected := "TLS12"
	actual = VersionFromConnection(&c)
	if expected != actual {
		t.Errorf("For %d, expected %q got %q", c.Version, expected, actual)
	}

	expected = "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"
	actual = CipherSuiteFromConnection(&c)
	if expected != actual {
		t.Errorf("For %d, expected %q got %q", c.CipherSuite, expected, actual)
	}
}

