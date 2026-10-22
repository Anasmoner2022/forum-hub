package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Approval: member creates pending post; moderator approves
	// Expected: public only after approval
	t.Fatal("TODO: implement independent assertions for Approval")
}

func TestContract02(t *testing.T) {
	// Self review: moderator tries to approve own post
	// Expected: 403; stays pending
	t.Fatal("TODO: implement independent assertions for Self review")
}

func TestContract03(t *testing.T) {
	// Role request: member requests role; admin approves then demotes
	// Expected: request retained; role changes recorded
	t.Fatal("TODO: implement independent assertions for Role request")
}

func TestContract04(t *testing.T) {
	// Report response: moderator reports; admin responds/resolves
	// Expected: reporter sees persisted response
	t.Fatal("TODO: implement independent assertions for Report response")
}

func TestContract05(t *testing.T) {
	// Comment deletion: admin deletes existing comment
	// Expected: comment hidden; audit retained
	t.Fatal("TODO: implement independent assertions for Comment deletion")
}

func TestContract06(t *testing.T) {
	// Category conflict: delete category referenced by a retained post
	// Expected: 409; membership intact
	t.Fatal("TODO: implement independent assertions for Category conflict")
}

func TestContract07(t *testing.T) {
	// Review race: approve and reject same version concurrently
	// Expected: one success; one 409
	t.Fatal("TODO: implement independent assertions for Review race")
}
