package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Label: label=node
	// Expected: A and C
	t.Fatal("TODO: implement independent assertions for Label")
}

func TestContract02(t *testing.T) {
	// Mine: viewer 1; mine=true
	// Expected: A and C
	t.Fatal("TODO: implement independent assertions for Mine")
}

func TestContract03(t *testing.T) {
	// Liked: viewer 2; liked=true
	// Expected: A only
	t.Fatal("TODO: implement independent assertions for Liked")
}

func TestContract04(t *testing.T) {
	// Combined: viewer 2; label=sql; liked=true
	// Expected: A only
	t.Fatal("TODO: implement independent assertions for Combined")
}

func TestContract05(t *testing.T) {
	// Guest ownership spoof: viewer null; mine=true; viewerId=1
	// Expected: authentication error
	t.Fatal("TODO: implement independent assertions for Guest ownership spoof")
}

func TestContract06(t *testing.T) {
	// Counts: A has 2 labels and 2 replies
	// Expected: replyCount=2; one A card
	t.Fatal("TODO: implement independent assertions for Counts")
}
