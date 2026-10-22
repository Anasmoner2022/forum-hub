package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Valid member: ada@example.test, Ada_1, LongPassword_123
	// Expected: one member; public result excludes hash
	t.Fatal("TODO: implement independent assertions for Valid member")
}

func TestContract02(t *testing.T) {
	// Normalized duplicate:  Ada@Example.test after ada@example.test
	// Expected: conflict; one member
	t.Fatal("TODO: implement independent assertions for Normalized duplicate")
}

func TestContract03(t *testing.T) {
	// Salt: hash identical password twice
	// Expected: different records; both verify true
	t.Fatal("TODO: implement independent assertions for Salt")
}

func TestContract04(t *testing.T) {
	// Wrong password: valid email; incorrect password
	// Expected: generic invalid credentials
	t.Fatal("TODO: implement independent assertions for Wrong password")
}

func TestContract05(t *testing.T) {
	// Malformed record: missing salt or invalid derived length
	// Expected: false/error handled; process survives
	t.Fatal("TODO: implement independent assertions for Malformed record")
}

func TestContract06(t *testing.T) {
	// Oversized password: 1025 UTF-8 bytes
	// Expected: validation error before scrypt
	t.Fatal("TODO: implement independent assertions for Oversized password")
}
