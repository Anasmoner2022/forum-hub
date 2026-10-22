package lab

import (
	"context"
	"fmt"
)

// Authorize: Return allow/deny from explicit matrix; default deny.
// TODO: define these fields from the specification and use trusted context.
type AuthorizeInput struct{}
type AuthorizeOutput struct{}

func Authorize(ctx context.Context, input AuthorizeInput) (AuthorizeOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return AuthorizeOutput{}, fmt.Errorf("TODO: Authorize")
}

// BootstrapAdmin: Create initial administrator through an operator-only command.
// TODO: define these fields from the specification and use trusted context.
type BootstrapAdminInput struct{}
type BootstrapAdminOutput struct{}

func BootstrapAdmin(ctx context.Context, input BootstrapAdminInput) (BootstrapAdminOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return BootstrapAdminOutput{}, fmt.Errorf("TODO: BootstrapAdmin")
}
