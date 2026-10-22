package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Trusted local request: curl --cacert certs/localhost.crt https://localhost:3443/health
	// Expected: 200 with certificate verification
	t.Fatal("TODO: implement independent assertions for Trusted local request")
}

func TestContract02(t *testing.T) {
	// Wrong hostname: use an unlisted hostname with same certificate
	// Expected: verification failure
	t.Fatal("TODO: implement independent assertions for Wrong hostname")
}

func TestContract03(t *testing.T) {
	// Missing key: configured nonexistent private-key path
	// Expected: startup failure before listen
	t.Fatal("TODO: implement independent assertions for Missing key")
}

func TestContract04(t *testing.T) {
	// TLS floor: attempt TLS 1.0 handshake
	// Expected: rejected
	t.Fatal("TODO: implement independent assertions for TLS floor")
}

func TestContract05(t *testing.T) {
	// Cookie: login via HTTPS
	// Expected: Secure; HttpOnly; SameSite=Lax
	t.Fatal("TODO: implement independent assertions for Cookie")
}

func TestContract06(t *testing.T) {
	// Container mount: read-only key/cert mounts
	// Expected: HTTPS succeeds; keys absent from image layers
	t.Fatal("TODO: implement independent assertions for Container mount")
}
