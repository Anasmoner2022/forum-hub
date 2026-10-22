package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Both providers: fresh Google and GitHub test identities
	// Expected: each can create and later log into a local account
	t.Fatal("TODO: implement independent assertions for Both providers")
}

func TestContract02(t *testing.T) {
	// Repeat login: same provider+subject twice
	// Expected: same local member; no duplicate
	t.Fatal("TODO: implement independent assertions for Repeat login")
}

func TestContract03(t *testing.T) {
	// Bad Google claims: wrong issuer/audience/nonce or expired token
	// Expected: rejected; no session
	t.Fatal("TODO: implement independent assertions for Bad Google claims")
}

func TestContract04(t *testing.T) {
	// Private GitHub email: /user email null; verified /user/emails available
	// Expected: safe onboarding with verified contact
	t.Fatal("TODO: implement independent assertions for Private GitHub email")
}

func TestContract05(t *testing.T) {
	// Email collision: new provider subject matches another member email
	// Expected: no automatic link; explicit existing-account proof required
	t.Fatal("TODO: implement independent assertions for Email collision")
}

func TestContract06(t *testing.T) {
	// Link conflict: identity already linked to member A; member B attempts link
	// Expected: 409; mapping stays A
	t.Fatal("TODO: implement independent assertions for Link conflict")
}
