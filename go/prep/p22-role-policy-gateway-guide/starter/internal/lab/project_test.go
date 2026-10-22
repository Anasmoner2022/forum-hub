package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Guest write: direct POST create without session
	// Expected: 401; no row
	t.Fatal("TODO: implement independent assertions for Guest write")
}

func TestContract02(t *testing.T) {
	// Role spoof: member sends role=admin
	// Expected: 403 on admin operation
	t.Fatal("TODO: implement independent assertions for Role spoof")
}

func TestContract03(t *testing.T) {
	// Moderator action: current moderator deletes post
	// Expected: allowed by matrix
	t.Fatal("TODO: implement independent assertions for Moderator action")
}

func TestContract04(t *testing.T) {
	// Demotion: admin demotes moderator then old cookie acts
	// Expected: 403 on moderator action
	t.Fatal("TODO: implement independent assertions for Demotion")
}

func TestContract05(t *testing.T) {
	// Pending detail: guest requests guessed pending ID
	// Expected: 404 without content leakage
	t.Fatal("TODO: implement independent assertions for Pending detail")
}

func TestContract06(t *testing.T) {
	// Last admin: attempt to demote sole admin
	// Expected: 409; one admin remains
	t.Fatal("TODO: implement independent assertions for Last admin")
}
