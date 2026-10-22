package lab

import (
	"context"
	"fmt"
)

// Dispatch: Return {status, headers, body}; consume request only for routes requiring a body.
// TODO: define these fields from the specification and use trusted context.
type DispatchInput struct{}
type DispatchOutput struct{}

func Dispatch(ctx context.Context, input DispatchInput) (DispatchOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return DispatchOutput{}, fmt.Errorf("TODO: Dispatch")
}

// ReadBoundedBody: Resolve a byte slice within the byte limit or reject with a body error.
// TODO: define these fields from the specification and use trusted context.
type ReadBoundedBodyInput struct{}
type ReadBoundedBodyOutput struct{}

func ReadBoundedBody(ctx context.Context, input ReadBoundedBodyInput) (ReadBoundedBodyOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return ReadBoundedBodyOutput{}, fmt.Errorf("TODO: ReadBoundedBody")
}
