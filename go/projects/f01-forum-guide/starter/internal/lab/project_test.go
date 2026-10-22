package lab

import "testing"

// These failing TODO tests are not evidence of implemented behavior.
// Replace them with specification-based assertions; no required skips remain.
func TestContract01(t *testing.T) {
	// Registration race: same normalized email registered concurrently
	// Expected: one account; other request 409
	t.Fatal("TODO: implement independent assertions for Registration race")
}

func TestContract02(t *testing.T) {
	// Session replacement: two browser logins for one member
	// Expected: only latest active token usable
	t.Fatal("TODO: implement independent assertions for Session replacement")
}

func TestContract03(t *testing.T) {
	// Public access: guest opens post with comments and counts
	// Expected: 200; correct public content
	t.Fatal("TODO: implement independent assertions for Public access")
}

func TestContract04(t *testing.T) {
	// Write protection: guest or forged author submits post/comment/reaction
	// Expected: 401/403; no write
	t.Fatal("TODO: implement independent assertions for Write protection")
}

func TestContract05(t *testing.T) {
	// Atomic post: categories [valid,missing]
	// Expected: validation/not-found response; no partial post
	t.Fatal("TODO: implement independent assertions for Atomic post")
}

func TestContract06(t *testing.T) {
	// Post/comment reactions: repeat like, switch dislike, clear on each target
	// Expected: exact counts; one current reaction
	t.Fatal("TODO: implement independent assertions for Post/comment reactions")
}

func TestContract07(t *testing.T) {
	// Filters: category+liked; mine from another memberId param
	// Expected: correct AND result; member comes from session
	t.Fatal("TODO: implement independent assertions for Filters")
}

func TestContract08(t *testing.T) {
	// Hostile input: HTML and SQL-like text
	// Expected: safe literal rendering; schema intact
	t.Fatal("TODO: implement independent assertions for Hostile input")
}

func TestContract09(t *testing.T) {
	// Fault recovery: malformed body and injected storage fault
	// Expected: safe status/message; next health request succeeds
	t.Fatal("TODO: implement independent assertions for Fault recovery")
}

func TestContract10(t *testing.T) {
	// Container persistence: build, create data, replace container
	// Expected: data unchanged; non-root process; successful separate restore
	t.Fatal("TODO: implement independent assertions for Container persistence")
}
