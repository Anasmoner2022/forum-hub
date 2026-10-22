package lab

import (
	"context"
	"fmt"
)

// HardenSession: Apply rotation/expiry rules for the stated lifecycle event.
// TODO: define these fields from the specification and use trusted context.
type HardenSessionInput struct{}
type HardenSessionOutput struct{}

func HardenSession(ctx context.Context, input HardenSessionInput) (HardenSessionOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return HardenSessionOutput{}, fmt.Errorf("TODO: HardenSession")
}

// SecurityHeaders: Return the reviewed header policy for public versus private content.
// TODO: define these fields from the specification and use trusted context.
type SecurityHeadersInput struct{}
type SecurityHeadersOutput struct{}

func SecurityHeaders(ctx context.Context, input SecurityHeadersInput) (SecurityHeadersOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return SecurityHeadersOutput{}, fmt.Errorf("TODO: SecurityHeaders")
}
