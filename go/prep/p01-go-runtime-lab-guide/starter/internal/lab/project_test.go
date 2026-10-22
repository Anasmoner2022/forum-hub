package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Valid file: {"appName":"inspector","port":3000}
	// Expected: JSON includes appName=inspector, port=3000 and runtimeMajor="go1.26"; exit 0
	t.Fatal("TODO: implement independent assertions for Valid file")
}

func TestContract02(t *testing.T) {
	// Environment override: PORT=4000 with file port 3000
	// Expected: port=4000
	t.Fatal("TODO: implement independent assertions for Environment override")
}

func TestContract03(t *testing.T) {
	// Malformed JSON: a file containing {
	// Expected: stderr identifies invalid JSON; exit 1
	t.Fatal("TODO: implement independent assertions for Malformed JSON")
}

func TestContract04(t *testing.T) {
	// Missing file: fixtures/missing.json
	// Expected: stderr identifies missing configuration; exit 1
	t.Fatal("TODO: implement independent assertions for Missing file")
}

func TestContract05(t *testing.T) {
	// Bad port: PORT=abc or port=0
	// Expected: validation error; exit 1
	t.Fatal("TODO: implement independent assertions for Bad port")
}

func TestContract06(t *testing.T) {
	// Read failure: inject a failing file reader
	// Expected: returned error reaches main once; exit 1
	t.Fatal("TODO: implement independent assertions for Read failure")
}
