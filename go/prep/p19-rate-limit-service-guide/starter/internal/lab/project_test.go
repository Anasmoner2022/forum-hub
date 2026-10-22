package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Identifier limit: 6 attempts for one identifier inside 60s
	// Expected: sixth 429
	t.Fatal("TODO: implement independent assertions for Identifier limit")
}

func TestContract02(t *testing.T) {
	// IP limit: 21 identifiers from one IP inside 60s
	// Expected: twenty-first 429
	t.Fatal("TODO: implement independent assertions for IP limit")
}

func TestContract03(t *testing.T) {
	// Retry header: blocked at 59001ms after start
	// Expected: Retry-After=1
	t.Fatal("TODO: implement independent assertions for Retry header")
}

func TestContract04(t *testing.T) {
	// Boundary: attempt at start+60000
	// Expected: new window; allowed
	t.Fatal("TODO: implement independent assertions for Boundary")
}

func TestContract05(t *testing.T) {
	// Proxy spoof: change untrusted X-Forwarded-For
	// Expected: same socket IP budget
	t.Fatal("TODO: implement independent assertions for Proxy spoof")
}

func TestContract06(t *testing.T) {
	// Work bound: many allowed requests trigger scrypt
	// Expected: at most configured active jobs; overload handled
	t.Fatal("TODO: implement independent assertions for Work bound")
}
