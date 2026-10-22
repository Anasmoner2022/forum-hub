package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Cipher engine: query cipher_version in clean Linux image
	// Expected: nonempty supported cipher version
	t.Fatal("TODO: implement independent assertions for Cipher engine")
}

func TestContract02(t *testing.T) {
	// Correct key: key then read schema and fixture rows
	// Expected: expected schema and counts
	t.Fatal("TODO: implement independent assertions for Correct key")
}

func TestContract03(t *testing.T) {
	// Wrong/missing key: open encrypted target with absent/different key
	// Expected: read fails; no silent plaintext fallback
	t.Fatal("TODO: implement independent assertions for Wrong/missing key")
}

func TestContract04(t *testing.T) {
	// Ordinary reader: open target with ordinary go-sqlite3
	// Expected: cannot read schema/data
	t.Fatal("TODO: implement independent assertions for Ordinary reader")
}

func TestContract05(t *testing.T) {
	// Repository parity: run P06/P07 contract cases on encrypted adapter
	// Expected: same public results and rollback behavior
	t.Fatal("TODO: implement independent assertions for Repository parity")
}

func TestContract06(t *testing.T) {
	// Export integrity: compare all table counts and checks after export
	// Expected: same data; integrity and FK checks pass
	t.Fatal("TODO: implement independent assertions for Export integrity")
}

func TestContract07(t *testing.T) {
	// Rotation/recovery: rekey disposable copy; test old/new keys; restore encrypted snapshot
	// Expected: old fails; new succeeds; recovered counts match
	t.Fatal("TODO: implement independent assertions for Rotation/recovery")
}
