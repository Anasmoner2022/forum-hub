package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Profile: fake /profile returns subject 101 and displayName Ada
	// Expected: validated public profile
	t.Fatal("TODO: implement independent assertions for Profile")
}

func TestContract02(t *testing.T) {
	// HTTP failure: fake /unavailable returns 503
	// Expected: provider-unavailable error; no profile
	t.Fatal("TODO: implement independent assertions for HTTP failure")
}

func TestContract03(t *testing.T) {
	// Timeout: fake /slow delays >2 seconds
	// Expected: controlled timeout
	t.Fatal("TODO: implement independent assertions for Timeout")
}

func TestContract04(t *testing.T) {
	// Malformed JSON: fake /bad-json
	// Expected: controlled parse error
	t.Fatal("TODO: implement independent assertions for Malformed JSON")
}

func TestContract05(t *testing.T) {
	// Oversized response: 262145 response bytes
	// Expected: controlled size error; bounded read
	t.Fatal("TODO: implement independent assertions for Oversized response")
}

func TestContract06(t *testing.T) {
	// Secret leak: token=SYNTHETIC_SECRET_MARKER
	// Expected: absent from all public responses and logs
	t.Fatal("TODO: implement independent assertions for Secret leak")
}
