package lab

import (
	"context"
	"fmt"
)

// BeginLoginAttempt: Store a short-lived attempt and return the authorization URL.
// TODO: define these fields from the specification and use trusted context.
type BeginLoginAttemptInput struct{}
type BeginLoginAttemptOutput struct{}

func BeginLoginAttempt(ctx context.Context, input BeginLoginAttemptInput) (BeginLoginAttemptOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return BeginLoginAttemptOutput{}, fmt.Errorf("TODO: BeginLoginAttempt")
}

// ConsumeLoginAttempt: Claim exactly one valid attempt or reject.
// TODO: define these fields from the specification and use trusted context.
type ConsumeLoginAttemptInput struct{}
type ConsumeLoginAttemptOutput struct{}

func ConsumeLoginAttempt(ctx context.Context, input ConsumeLoginAttemptInput) (ConsumeLoginAttemptOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return ConsumeLoginAttemptOutput{}, fmt.Errorf("TODO: ConsumeLoginAttempt")
}

// ExchangeCode: Resolve token response with no automatic replay.
// TODO: define these fields from the specification and use trusted context.
type ExchangeCodeInput struct{}
type ExchangeCodeOutput struct{}

func ExchangeCode(ctx context.Context, input ExchangeCodeInput) (ExchangeCodeOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return ExchangeCodeOutput{}, fmt.Errorf("TODO: ExchangeCode")
}
