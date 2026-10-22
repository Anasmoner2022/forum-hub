package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Required formats: valid JPEG, PNG and GIF including animation
	// Expected: all accepted with correct decoded types
	t.Fatal("TODO: implement independent assertions for Required formats")
}

func TestContract02(t *testing.T) {
	// MIME spoof: text bytes named image.jpg
	// Expected: 422; no asset
	t.Fatal("TODO: implement independent assertions for MIME spoof")
}

func TestContract03(t *testing.T) {
	// Corrupt image: truncated PNG compressed data
	// Expected: 422 after decode; cleanup
	t.Fatal("TODO: implement independent assertions for Corrupt image")
}

func TestContract04(t *testing.T) {
	// Dimension bound: width 8193 or 25,000,001 pixels/frame
	// Expected: 422
	t.Fatal("TODO: implement independent assertions for Dimension bound")
}

func TestContract05(t *testing.T) {
	// Animation bound: 101 GIF frames or >100,000,000 total pixels
	// Expected: 422
	t.Fatal("TODO: implement independent assertions for Animation bound")
}

func TestContract06(t *testing.T) {
	// Database failure: file placed then link insert fails
	// Expected: no post; compensation or recorded reconciliation cleanup
	t.Fatal("TODO: implement independent assertions for Database failure")
}
