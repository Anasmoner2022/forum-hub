package lab

import (
	"context"
	"fmt"
)

// ClassifyError: Return a safe status/code without exposing the internal error.
// TODO: define these fields from the specification and use trusted context.
type ClassifyErrorInput struct{}
type ClassifyErrorOutput struct{}

func ClassifyError(ctx context.Context, input ClassifyErrorInput) (ClassifyErrorOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return ClassifyErrorOutput{}, fmt.Errorf("TODO: ClassifyError")
}

// CreateTestContext: Create disposable storage, clock and cleanup handles.
// TODO: define these fields from the specification and use trusted context.
type CreateTestContextInput struct{}
type CreateTestContextOutput struct{}

func CreateTestContext(ctx context.Context, input CreateTestContextInput) (CreateTestContextOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return CreateTestContextOutput{}, fmt.Errorf("TODO: CreateTestContext")
}
