package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// JPEG/PNG/GIF: upload one valid fixture each
	// Expected: post and image visible to guest if public
	t.Fatal("TODO: implement independent assertions for JPEG/PNG/GIF")
}

func TestContract02(t *testing.T) {
	// Text only: normal post without image
	// Expected: success with no asset
	t.Fatal("TODO: implement independent assertions for Text only")
}

func TestContract03(t *testing.T) {
	// File size: 20,000,000 versus 20,000,001 bytes
	// Expected: intake accept versus 413; image validity separately checked
	t.Fatal("TODO: implement independent assertions for File size")
}

func TestContract04(t *testing.T) {
	// Corrupt/spoof: invalid bytes claimed as PNG
	// Expected: 422; cleanup
	t.Fatal("TODO: implement independent assertions for Corrupt/spoof")
}

func TestContract05(t *testing.T) {
	// Pixel/frame cap: over stated dimension or frame budget
	// Expected: 422
	t.Fatal("TODO: implement independent assertions for Pixel/frame cap")
}

func TestContract06(t *testing.T) {
	// Abort/CSRF: disconnect or bad CSRF field
	// Expected: no post/public asset; temp cleanup
	t.Fatal("TODO: implement independent assertions for Abort/CSRF")
}

func TestContract07(t *testing.T) {
	// Storage failure: final SQL linkage fails
	// Expected: compensated file; no partial post
	t.Fatal("TODO: implement independent assertions for Storage failure")
}

func TestContract08(t *testing.T) {
	// Visibility/restart: replace container; request restricted asset in integrated mode
	// Expected: published retained; restricted denied
	t.Fatal("TODO: implement independent assertions for Visibility/restart")
}
