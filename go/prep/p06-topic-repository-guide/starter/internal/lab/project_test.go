package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Topic creation: author 1, title First, body Text, labels [1,2]
	// Expected: one topic; two links
	t.Fatal("TODO: implement independent assertions for Topic creation")
}

func TestContract02(t *testing.T) {
	// Empty labels: labels []
	// Expected: validation error; no new topic
	t.Fatal("TODO: implement independent assertions for Empty labels")
}

func TestContract03(t *testing.T) {
	// Invalid label: labels [1,9999]
	// Expected: no partial topic
	t.Fatal("TODO: implement independent assertions for Invalid label")
}

func TestContract04(t *testing.T) {
	// Reply: valid member and existing topic
	// Expected: one reply returned with public author name
	t.Fatal("TODO: implement independent assertions for Reply")
}

func TestContract05(t *testing.T) {
	// SQL payload: title containing apostrophe and SQL-like text
	// Expected: literal stored text; members table intact
	t.Fatal("TODO: implement independent assertions for SQL payload")
}

func TestContract06(t *testing.T) {
	// Join multiplication: two labels, three replies
	// Expected: two labels; three replies, no duplicate rows
	t.Fatal("TODO: implement independent assertions for Join multiplication")
}
