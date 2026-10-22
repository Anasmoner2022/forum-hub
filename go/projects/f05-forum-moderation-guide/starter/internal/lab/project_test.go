package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Guest/member separation: direct privileged request from guest/member
	// Expected: 401/403; unchanged state
	t.Fatal("TODO: implement independent assertions for Guest/member separation")
}

func TestContract02(t *testing.T) {
	// Prepublication: create then approve post
	// Expected: pending hidden; approved public
	t.Fatal("TODO: implement independent assertions for Prepublication")
}

func TestContract03(t *testing.T) {
	// Self approval: moderator submits and reviews own post
	// Expected: 403
	t.Fatal("TODO: implement independent assertions for Self approval")
}

func TestContract04(t *testing.T) {
	// Role lifecycle: request, approve, demote, reuse old cookie
	// Expected: audit persists; demoted action denied
	t.Fatal("TODO: implement independent assertions for Role lifecycle")
}

func TestContract05(t *testing.T) {
	// Reports: moderator report; admin response/resolution
	// Expected: reporter sees response
	t.Fatal("TODO: implement independent assertions for Reports")
}

func TestContract06(t *testing.T) {
	// Admin deletion: admin deletes post and comment
	// Expected: both hidden publicly
	t.Fatal("TODO: implement independent assertions for Admin deletion")
}

func TestContract07(t *testing.T) {
	// Categories: admin creates, deletes unused, attempts attached delete
	// Expected: success, success, 409
	t.Fatal("TODO: implement independent assertions for Categories")
}

func TestContract08(t *testing.T) {
	// Conflict/leak audit: concurrent review and guessed pending asset when installed
	// Expected: one review wins; restricted content denied
	t.Fatal("TODO: implement independent assertions for Conflict/leak audit")
}
