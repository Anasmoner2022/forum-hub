package lab

import (
	"context"
	"fmt"
)

// AllowAttempt: Return {allowed, retryAfterSeconds} without extending blocked windows.
// TODO: define these fields from the specification and use trusted context.
type AllowAttemptInput struct{}
type AllowAttemptOutput struct{}

func AllowAttempt(ctx context.Context, input AllowAttemptInput) (AllowAttemptOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return AllowAttemptOutput{}, fmt.Errorf("TODO: AllowAttempt")
}

// PruneLimiter: Remove stale keys while retaining live policy state.
// TODO: define these fields from the specification and use trusted context.
type PruneLimiterInput struct{}
type PruneLimiterOutput struct{}

func PruneLimiter(ctx context.Context, input PruneLimiterInput) (PruneLimiterOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return PruneLimiterOutput{}, fmt.Errorf("TODO: PruneLimiter")
}

// RunExpensiveJob: Enforce a bounded concurrent job count and safe overload behavior.
// TODO: define these fields from the specification and use trusted context.
type RunExpensiveJobInput struct{}
type RunExpensiveJobOutput struct{}

func RunExpensiveJob(ctx context.Context, input RunExpensiveJobInput) (RunExpensiveJobOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return RunExpensiveJobOutput{}, fmt.Errorf("TODO: RunExpensiveJob")
}
