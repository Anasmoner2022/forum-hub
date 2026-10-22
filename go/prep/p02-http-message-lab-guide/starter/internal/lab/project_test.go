package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Health: GET /health
	// Expected: 200; JSON {ok:true}
	t.Fatal("TODO: implement independent assertions for Health")
}

func TestContract02(t *testing.T) {
	// Echo: POST /echo; Content-Type application/json; {"message":"hello"}
	// Expected: 200; same object
	t.Fatal("TODO: implement independent assertions for Echo")
}

func TestContract03(t *testing.T) {
	// Wrong method: PUT /echo
	// Expected: 405; Allow includes POST
	t.Fatal("TODO: implement independent assertions for Wrong method")
}

func TestContract04(t *testing.T) {
	// Bad JSON: POST /echo with {
	// Expected: 400
	t.Fatal("TODO: implement independent assertions for Bad JSON")
}

func TestContract05(t *testing.T) {
	// Too large: 65537 streamed body bytes without Content-Length
	// Expected: 413; bounded memory
	t.Fatal("TODO: implement independent assertions for Too large")
}

func TestContract06(t *testing.T) {
	// Wrong content type: POST /echo text/plain
	// Expected: 415
	t.Fatal("TODO: implement independent assertions for Wrong content type")
}
