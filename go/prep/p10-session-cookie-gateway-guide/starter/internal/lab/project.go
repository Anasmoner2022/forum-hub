package lab

import (
	"context"
	"fmt"
)

// CreateSession: Return fresh token and cookie metadata after replacing old sessions.
// TODO: define these fields from the specification and use trusted context.
type CreateSessionInput struct{}
type CreateSessionOutput struct{}

func CreateSession(ctx context.Context, input CreateSessionInput) (CreateSessionOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return CreateSessionOutput{}, fmt.Errorf("TODO: CreateSession")
}

// Authenticate: Return trusted member context or null.
// TODO: define these fields from the specification and use trusted context.
type AuthenticateInput struct{}
type AuthenticateOutput struct{}

func Authenticate(ctx context.Context, input AuthenticateInput) (AuthenticateOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return AuthenticateOutput{}, fmt.Errorf("TODO: Authenticate")
}

// RequireCsrf: Reject unsafe requests without the matching token and allowed origin.
// TODO: define these fields from the specification and use trusted context.
type RequireCsrfInput struct{}
type RequireCsrfOutput struct{}

func RequireCsrf(ctx context.Context, input RequireCsrfInput) (RequireCsrfOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return RequireCsrfOutput{}, fmt.Errorf("TODO: RequireCsrf")
}

// Logout: Revoke server state and expire the client cookie.
// TODO: define these fields from the specification and use trusted context.
type LogoutInput struct{}
type LogoutOutput struct{}

func Logout(ctx context.Context, input LogoutInput) (LogoutOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return LogoutOutput{}, fmt.Errorf("TODO: Logout")
}
