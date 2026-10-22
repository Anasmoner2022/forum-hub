package lab

import (
	"context"
	"fmt"
)

// RenderMessagePage: Return HTML with every user-controlled value safely rendered.
// TODO: define these fields from the specification and use trusted context.
type RenderMessagePageInput struct{}
type RenderMessagePageOutput struct{}

func RenderMessagePage(ctx context.Context, input RenderMessagePageInput) (RenderMessagePageOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return RenderMessagePageOutput{}, fmt.Errorf("TODO: RenderMessagePage")
}

// ValidateMessage: Return normalized data or a field-error object.
// TODO: define these fields from the specification and use trusted context.
type ValidateMessageInput struct{}
type ValidateMessageOutput struct{}

func ValidateMessage(ctx context.Context, input ValidateMessageInput) (ValidateMessageOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return ValidateMessageOutput{}, fmt.Errorf("TODO: ValidateMessage")
}
