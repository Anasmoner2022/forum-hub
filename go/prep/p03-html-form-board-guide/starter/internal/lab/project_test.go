package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Creation: POST title=Hello&body=World
	// Expected: 303; detail page shows Hello and World
	t.Fatal("TODO: implement independent assertions for Creation")
}

func TestContract02(t *testing.T) {
	// Blank title: title spaces only
	// Expected: 422; field error; no message
	t.Fatal("TODO: implement independent assertions for Blank title")
}

func TestContract03(t *testing.T) {
	// Boundary title: 120 then 121 Unicode code points
	// Expected: accepted then rejected
	t.Fatal("TODO: implement independent assertions for Boundary title")
}

func TestContract04(t *testing.T) {
	// Hostile body: <img src=x onerror=alert(1)>
	// Expected: visible literal text; no execution
	t.Fatal("TODO: implement independent assertions for Hostile body")
}

func TestContract05(t *testing.T) {
	// Refresh: refresh the redirected detail page
	// Expected: no duplicate
	t.Fatal("TODO: implement independent assertions for Refresh")
}

func TestContract06(t *testing.T) {
	// Missing item: GET /messages/999999
	// Expected: 404 page; usable navigation
	t.Fatal("TODO: implement independent assertions for Missing item")
}
