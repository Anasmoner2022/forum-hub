package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// HTTPS: verified localhost request with --cacert
	// Expected: 200; SAN/time verified
	t.Fatal("TODO: implement independent assertions for HTTPS")
}

func TestContract02(t *testing.T) {
	// Bad configuration: missing private key
	// Expected: startup refuses before listen
	t.Fatal("TODO: implement independent assertions for Bad configuration")
}

func TestContract03(t *testing.T) {
	// Login limits: sixth identifier attempt inside 60 seconds
	// Expected: 429 with Retry-After
	t.Fatal("TODO: implement independent assertions for Login limits")
}

func TestContract04(t *testing.T) {
	// Work concurrency: burst across many identifiers
	// Expected: configured active hash bound
	t.Fatal("TODO: implement independent assertions for Work concurrency")
}

func TestContract05(t *testing.T) {
	// Session fixation: chosen pre-login token then successful login
	// Expected: fresh random token; chosen token fails
	t.Fatal("TODO: implement independent assertions for Session fixation")
}

func TestContract06(t *testing.T) {
	// Revocation/expiry: logout replay and idle/absolute clock boundaries
	// Expected: no authentication
	t.Fatal("TODO: implement independent assertions for Revocation/expiry")
}

func TestContract07(t *testing.T) {
	// CSRF/injection: hostile form/text/SQL-like payload
	// Expected: no unauthorized write or executable content
	t.Fatal("TODO: implement independent assertions for CSRF/injection")
}

func TestContract08(t *testing.T) {
	// Regression: all installed forum features over HTTPS
	// Expected: same permitted behavior; no secret leakage
	t.Fatal("TODO: implement independent assertions for Regression")
}
