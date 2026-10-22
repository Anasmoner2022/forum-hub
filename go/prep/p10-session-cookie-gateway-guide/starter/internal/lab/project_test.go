package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Valid login: valid credentials in browser A
	// Expected: authenticated session; required cookie attributes
	t.Fatal("TODO: implement independent assertions for Valid login")
}

func TestContract02(t *testing.T) {
	// Second login: same member in browser B
	// Expected: A rejected; B accepted
	t.Fatal("TODO: implement independent assertions for Second login")
}

func TestContract03(t *testing.T) {
	// Server expiry: advance clock by 8 hours
	// Expected: token rejected despite copied cookie
	t.Fatal("TODO: implement independent assertions for Server expiry")
}

func TestContract04(t *testing.T) {
	// Logout: valid session and CSRF; POST /logout
	// Expected: row revoked; cookie expired
	t.Fatal("TODO: implement independent assertions for Logout")
}

func TestContract05(t *testing.T) {
	// Guest mutation: POST protected create without session
	// Expected: 401; no write
	t.Fatal("TODO: implement independent assertions for Guest mutation")
}

func TestContract06(t *testing.T) {
	// CSRF: valid cookie but wrong/missing form token
	// Expected: 403; no write
	t.Fatal("TODO: implement independent assertions for CSRF")
}
