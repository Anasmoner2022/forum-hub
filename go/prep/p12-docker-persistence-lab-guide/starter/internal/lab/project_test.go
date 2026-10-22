package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Clean build: docker compose build
	// Expected: image builds using lockfile
	t.Fatal("TODO: implement independent assertions for Clean build")
}

func TestContract02(t *testing.T) {
	// Run: docker compose up -d; GET /health
	// Expected: 200 through mapped port
	t.Fatal("TODO: implement independent assertions for Run")
}

func TestContract03(t *testing.T) {
	// Persistence: create topic, docker compose down then up
	// Expected: topic remains; no -v flag used
	t.Fatal("TODO: implement independent assertions for Persistence")
}

func TestContract04(t *testing.T) {
	// Permissions: inspect process UID and create data
	// Expected: non-root; directory writable
	t.Fatal("TODO: implement independent assertions for Permissions")
}

func TestContract05(t *testing.T) {
	// Architecture: build on supported team host
	// Expected: native dependencies target container platform
	t.Fatal("TODO: implement independent assertions for Architecture")
}

func TestContract06(t *testing.T) {
	// Restore: restore snapshot to separate path
	// Expected: row counts match; no FK violations
	t.Fatal("TODO: implement independent assertions for Restore")
}
