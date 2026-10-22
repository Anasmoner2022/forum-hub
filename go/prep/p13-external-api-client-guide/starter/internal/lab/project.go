package lab

import (
	"context"
	"fmt"
)

// FetchProfile: Return a validated profile or a sanitized provider error.
// TODO: define these fields from the specification and use trusted context.
type FetchProfileInput struct{}
type FetchProfileOutput struct{}

func FetchProfile(ctx context.Context, input FetchProfileInput) (FetchProfileOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return FetchProfileOutput{}, fmt.Errorf("TODO: FetchProfile")
}

// CreateFakeProvider: Start a disposable local server with an explicit synthetic response.
// TODO: define these fields from the specification and use trusted context.
type CreateFakeProviderInput struct{}
type CreateFakeProviderOutput struct{}

func CreateFakeProvider(ctx context.Context, input CreateFakeProviderInput) (CreateFakeProviderOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return CreateFakeProviderOutput{}, fmt.Errorf("TODO: CreateFakeProvider")
}
