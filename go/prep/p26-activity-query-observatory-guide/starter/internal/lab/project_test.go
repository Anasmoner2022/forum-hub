package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Created posts: A has two authored posts
	// Expected: A section lists exactly two
	t.Fatal("TODO: implement independent assertions for Created posts")
}

func TestContract02(t *testing.T) {
	// Comment context: A comments on B’s post
	// Expected: comment text plus B post title/link
	t.Fatal("TODO: implement independent assertions for Comment context")
}

func TestContract03(t *testing.T) {
	// Both reaction targets: A likes post and dislikes comment
	// Expected: each in correct section with target context
	t.Fatal("TODO: implement independent assertions for Both reaction targets")
}

func TestContract04(t *testing.T) {
	// Cleared reaction: A clears post like
	// Expected: absent from current likes
	t.Fatal("TODO: implement independent assertions for Cleared reaction")
}

func TestContract05(t *testing.T) {
	// Actor spoof: A requests memberId=B
	// Expected: still only A activity
	t.Fatal("TODO: implement independent assertions for Actor spoof")
}

func TestContract06(t *testing.T) {
	// Deleted parent: comment parent soft-deleted
	// Expected: neutral tombstone; no retained private text
	t.Fatal("TODO: implement independent assertions for Deleted parent")
}

func TestContract07(t *testing.T) {
	// Optimization: 1000-record fixture before/after indexes
	// Expected: same results; documented plan/query-count change
	t.Fatal("TODO: implement independent assertions for Optimization")
}
