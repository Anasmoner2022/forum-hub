package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Two providers: Google and GitHub first+repeat login
	// Expected: working registration and stable returning account
	t.Fatal("TODO: implement independent assertions for Two providers")
}

func TestContract02(t *testing.T) {
	// Replay: reuse completed callback
	// Expected: 400; no new session
	t.Fatal("TODO: implement independent assertions for Replay")
}

func TestContract03(t *testing.T) {
	// Cross browser: callback state used without originating browser binding
	// Expected: 400
	t.Fatal("TODO: implement independent assertions for Cross browser")
}

func TestContract04(t *testing.T) {
	// Bad ID token: wrong Google audience or nonce
	// Expected: rejected
	t.Fatal("TODO: implement independent assertions for Bad ID token")
}

func TestContract05(t *testing.T) {
	// Collision: provider email equals another local member
	// Expected: explicit proof required; no takeover
	t.Fatal("TODO: implement independent assertions for Collision")
}

func TestContract06(t *testing.T) {
	// Link conflict: provider identity owned by another member
	// Expected: 409
	t.Fatal("TODO: implement independent assertions for Link conflict")
}

func TestContract07(t *testing.T) {
	// Session parity: SSO login after local login in another browser
	// Expected: old session revoked
	t.Fatal("TODO: implement independent assertions for Session parity")
}

func TestContract08(t *testing.T) {
	// Provider failure: denial, timeout, malformed profile
	// Expected: safe error; base forum remains usable
	t.Fatal("TODO: implement independent assertions for Provider failure")
}
