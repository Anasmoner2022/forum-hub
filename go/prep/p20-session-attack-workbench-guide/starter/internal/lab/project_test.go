package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Fixation: chosen cookie before successful login
	// Expected: fresh unrelated token; chosen token rejected
	t.Fatal("TODO: implement independent assertions for Fixation")
}

func TestContract02(t *testing.T) {
	// Concurrent login: two valid logins at once
	// Expected: only one active server session afterward
	t.Fatal("TODO: implement independent assertions for Concurrent login")
}

func TestContract03(t *testing.T) {
	// Idle expiry: 31 minutes inactivity
	// Expected: 401; server rejects token
	t.Fatal("TODO: implement independent assertions for Idle expiry")
}

func TestContract04(t *testing.T) {
	// Forged CSRF: session cookie with attacker form
	// Expected: 403; no write
	t.Fatal("TODO: implement independent assertions for Forged CSRF")
}

func TestContract05(t *testing.T) {
	// XSS/SQL input: hostile text and SQL-like string
	// Expected: literal text; data and schema intact
	t.Fatal("TODO: implement independent assertions for XSS/SQL input")
}

func TestContract06(t *testing.T) {
	// Private response: GET account page
	// Expected: no-store and reviewed browser headers; no secrets
	t.Fatal("TODO: implement independent assertions for Private response")
}
