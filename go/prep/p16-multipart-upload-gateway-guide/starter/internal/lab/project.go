package lab

import (
	"context"
	"fmt"
)

// ReceiveUpload: Return a private temp path and bounded fields only after all checks pass.
// TODO: define these fields from the specification and use trusted context.
type ReceiveUploadInput struct{}
type ReceiveUploadOutput struct{}

func ReceiveUpload(ctx context.Context, input ReceiveUploadInput) (ReceiveUploadOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return ReceiveUploadOutput{}, fmt.Errorf("TODO: ReceiveUpload")
}

// CleanupUpload: Remove provisional files safely and idempotently.
// TODO: define these fields from the specification and use trusted context.
type CleanupUploadInput struct{}
type CleanupUploadOutput struct{}

func CleanupUpload(ctx context.Context, input CleanupUploadInput) (CleanupUploadOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return CleanupUploadOutput{}, fmt.Errorf("TODO: CleanupUpload")
}
