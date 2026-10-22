package lab

import (
	"context"
	"fmt"
)

// StartProviderLogin: Begin only an allowlisted login or link flow.
// TODO: define these fields from the specification and use trusted context.
type StartProviderLoginInput struct{}
type StartProviderLoginOutput struct{}

func StartProviderLogin(ctx context.Context, input StartProviderLoginInput) (StartProviderLoginOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return StartProviderLoginOutput{}, fmt.Errorf("TODO: StartProviderLogin")
}

// FinishProviderLogin: Consume attempt, verify identity and issue forum session.
// TODO: define these fields from the specification and use trusted context.
type FinishProviderLoginInput struct{}
type FinishProviderLoginOutput struct{}

func FinishProviderLogin(ctx context.Context, input FinishProviderLoginInput) (FinishProviderLoginOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return FinishProviderLoginOutput{}, fmt.Errorf("TODO: FinishProviderLogin")
}
