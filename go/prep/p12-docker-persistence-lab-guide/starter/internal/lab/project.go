package lab

import (
	"context"
	"fmt"
)

// CheckRuntime: Validate startup paths, configured port and writable storage before serving.
// TODO: define these fields from the specification and use trusted context.
type CheckRuntimeInput struct{}
type CheckRuntimeOutput struct{}

func CheckRuntime(ctx context.Context, input CheckRuntimeInput) (CheckRuntimeOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return CheckRuntimeOutput{}, fmt.Errorf("TODO: CheckRuntime")
}
