package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// First like: member 1 sets topic 1 to +1
	// Expected: likes=1; dislikes=0
	t.Fatal("TODO: implement independent assertions for First like")
}

func TestContract02(t *testing.T) {
	// Repeat: same set +1 repeated 20 times
	// Expected: likes=1
	t.Fatal("TODO: implement independent assertions for Repeat")
}

func TestContract03(t *testing.T) {
	// Switch: same member sets -1
	// Expected: likes=0; dislikes=1
	t.Fatal("TODO: implement independent assertions for Switch")
}

func TestContract04(t *testing.T) {
	// Clear: clear twice
	// Expected: likes=0; dislikes=0
	t.Fatal("TODO: implement independent assertions for Clear")
}

func TestContract05(t *testing.T) {
	// Reply target: member 2 likes reply 5
	// Expected: reply counts change; topic counts unchanged
	t.Fatal("TODO: implement independent assertions for Reply target")
}

func TestContract06(t *testing.T) {
	// Missing target: reply 999999
	// Expected: not-found; no reaction row
	t.Fatal("TODO: implement independent assertions for Missing target")
}
