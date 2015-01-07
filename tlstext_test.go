package tlstext

import (
	"testing"
)

func TestCipherSuite(t *testing.T) {
	value := uint16(0xc02b)
	expected := "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"
	actual := CipherSuite(value)
	if expected != actual {
		t.Errorf("For %d, expected %q got %q", value, expected, actual)
	}

	value = uint16(1234)
	expected = ""
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
	expected = ""
	actual = Version(value)
	if expected != actual {
		t.Errorf("For %d, expected %q got %q", value, expected, actual)
	}
}
