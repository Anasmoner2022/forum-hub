package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Owner edit: author edits post at current version
	// Expected: content changes; version increments
	t.Fatal("TODO: implement independent assertions for Owner edit")
}

func TestContract02(t *testing.T) {
	// Nonowner: other member supplies authorId=owner
	// Expected: 403; unchanged
	t.Fatal("TODO: implement independent assertions for Nonowner")
}

func TestContract03(t *testing.T) {
	// Stale editor: save version 1 after version 2 exists
	// Expected: 409; version 2 remains
	t.Fatal("TODO: implement independent assertions for Stale editor")
}

func TestContract04(t *testing.T) {
	// Comment edit: author changes comment body
	// Expected: updated body visible under published parent
	t.Fatal("TODO: implement independent assertions for Comment edit")
}

func TestContract05(t *testing.T) {
	// Deletion: author deletes published post
	// Expected: public post and child comments disappear
	t.Fatal("TODO: implement independent assertions for Deletion")
}

func TestContract06(t *testing.T) {
	// Moderated edit: published post edited in moderation mode
	// Expected: pending; publicly hidden until reapproved
	t.Fatal("TODO: implement independent assertions for Moderated edit")
}
