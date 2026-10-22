package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Comment notification: B comments on A’s published post
	// Expected: one unread notification for A
	t.Fatal("TODO: implement independent assertions for Comment notification")
}

func TestContract02(t *testing.T) {
	// Like/dislike: B sets like then switches to dislike
	// Expected: one notification per actual state change
	t.Fatal("TODO: implement independent assertions for Like/dislike")
}

func TestContract03(t *testing.T) {
	// Retry same like: B repeats set-like
	// Expected: no extra event or notification
	t.Fatal("TODO: implement independent assertions for Retry same like")
}

func TestContract04(t *testing.T) {
	// Self action: A comments on or reacts to own post
	// Expected: no notification
	t.Fatal("TODO: implement independent assertions for Self action")
}

func TestContract05(t *testing.T) {
	// Rollback: force domain transaction failure
	// Expected: no event and no notification
	t.Fatal("TODO: implement independent assertions for Rollback")
}

func TestContract06(t *testing.T) {
	// Dispatcher retry: process same event twice
	// Expected: one notification
	t.Fatal("TODO: implement independent assertions for Dispatcher retry")
}

func TestContract07(t *testing.T) {
	// Recipient isolation: B reads/marks A’s notification
	// Expected: 403/404; unread state unchanged
	t.Fatal("TODO: implement independent assertions for Recipient isolation")
}
