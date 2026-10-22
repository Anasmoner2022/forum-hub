package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Fresh migration: empty data/lab.sqlite; migrate
	// Expected: v1 and v2 recorded once
	t.Fatal("TODO: implement independent assertions for Fresh migration")
}

func TestContract02(t *testing.T) {
	// Second run: same database; migrate again
	// Expected: no changes; no duplicates
	t.Fatal("TODO: implement independent assertions for Second run")
}

func TestContract03(t *testing.T) {
	// Integrity: insert reply for nonexistent topic
	// Expected: foreign-key error
	t.Fatal("TODO: implement independent assertions for Integrity")
}

func TestContract04(t *testing.T) {
	// Bad reaction: insert reaction value 0
	// Expected: CHECK violation
	t.Fatal("TODO: implement independent assertions for Bad reaction")
}

func TestContract05(t *testing.T) {
	// Broken v2: invalid SQL midway through a disposable v2
	// Expected: no partial column or applied-version row
	t.Fatal("TODO: implement independent assertions for Broken v2")
}

func TestContract06(t *testing.T) {
	// Changed migration: edit an already-applied migration
	// Expected: checksum error; no mutation
	t.Fatal("TODO: implement independent assertions for Changed migration")
}
