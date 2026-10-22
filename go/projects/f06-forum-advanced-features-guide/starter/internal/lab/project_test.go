package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Comment/like/dislike: B acts on A’s post
	// Expected: A receives correct durable notifications
	t.Fatal("TODO: implement independent assertions for Comment/like/dislike")
}

func TestContract02(t *testing.T) {
	// No-op/self: repeat same reaction; A acts on own post
	// Expected: no extra notification
	t.Fatal("TODO: implement independent assertions for No-op/self")
}

func TestContract03(t *testing.T) {
	// Offline/retry: A offline; dispatch same event twice
	// Expected: one notification available on return
	t.Fatal("TODO: implement independent assertions for Offline/retry")
}

func TestContract04(t *testing.T) {
	// Privacy: B requests A inbox/activity by supplied IDs
	// Expected: no A data; no state changes
	t.Fatal("TODO: implement independent assertions for Privacy")
}

func TestContract05(t *testing.T) {
	// Activity context: authored post/comment plus both target reactions
	// Expected: correct current sections and parent context
	t.Fatal("TODO: implement independent assertions for Activity context")
}

func TestContract06(t *testing.T) {
	// Edit ownership/staleness: nonowner edit then stale owner edit
	// Expected: 403 then 409; latest content intact
	t.Fatal("TODO: implement independent assertions for Edit ownership/staleness")
}

func TestContract07(t *testing.T) {
	// Remove both types: owner removes post and separate comment
	// Expected: public content hidden; lifecycle followed
	t.Fatal("TODO: implement independent assertions for Remove both types")
}

func TestContract08(t *testing.T) {
	// Installed moderation: edit published post when F05 installed
	// Expected: pending; no public/asset leak
	t.Fatal("TODO: implement independent assertions for Installed moderation")
}

func TestContract09(t *testing.T) {
	// Rollback/index regression: failed domain write; optimize activity indexes
	// Expected: no phantom event; exact results retained
	t.Fatal("TODO: implement independent assertions for Rollback/index regression")
}
