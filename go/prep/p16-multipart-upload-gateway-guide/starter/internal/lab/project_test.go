package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// At limit: one 20,000,000-byte synthetic file
	// Expected: accepted by byte intake; image validation still required in P17
	t.Fatal("TODO: implement independent assertions for At limit")
}

func TestContract02(t *testing.T) {
	// Over limit: one 20,000,001-byte file
	// Expected: 413 image too big; no temp residue
	t.Fatal("TODO: implement independent assertions for Over limit")
}

func TestContract03(t *testing.T) {
	// Extra file: two image parts
	// Expected: rejected; cleanup
	t.Fatal("TODO: implement independent assertions for Extra file")
}

func TestContract04(t *testing.T) {
	// Wrong CSRF: valid cookie; bad csrf field
	// Expected: 403; no public record
	t.Fatal("TODO: implement independent assertions for Wrong CSRF")
}

func TestContract05(t *testing.T) {
	// Abort: disconnect after partial file
	// Expected: temp removed; server healthy
	t.Fatal("TODO: implement independent assertions for Abort")
}

func TestContract06(t *testing.T) {
	// Guest: upload without session
	// Expected: 401 before intake
	t.Fatal("TODO: implement independent assertions for Guest")
}
