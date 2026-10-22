package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Successful round trip: matching browser, state, code and verifier
	// Expected: one token result
	t.Fatal("TODO: implement independent assertions for Successful round trip")
}

func TestContract02(t *testing.T) {
	// Wrong state: alter one state character
	// Expected: 400; no exchange
	t.Fatal("TODO: implement independent assertions for Wrong state")
}

func TestContract03(t *testing.T) {
	// Cross-browser: valid state in a different browser
	// Expected: 400; no exchange
	t.Fatal("TODO: implement independent assertions for Cross-browser")
}

func TestContract04(t *testing.T) {
	// Replay: same callback twice
	// Expected: second rejected; one exchange total
	t.Fatal("TODO: implement independent assertions for Replay")
}

func TestContract05(t *testing.T) {
	// Expiry: advance attempt clock by 5 minutes
	// Expected: rejected
	t.Fatal("TODO: implement independent assertions for Expiry")
}

func TestContract06(t *testing.T) {
	// Bad PKCE: wrong verifier at token endpoint
	// Expected: provider rejection; no app session
	t.Fatal("TODO: implement independent assertions for Bad PKCE")
}
