package lab

import (
	"context"
	"fmt"
)

// Migrate: Apply pending versions atomically and reject changed applied checksums.
// TODO: define these fields from the specification and use trusted context.
type MigrateInput struct{}
type MigrateOutput struct{}

func Migrate(ctx context.Context, input MigrateInput) (MigrateOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return MigrateOutput{}, fmt.Errorf("TODO: Migrate")
}

// Seed: Insert a repeatable fixture using bound values.
// TODO: define these fields from the specification and use trusted context.
type SeedInput struct{}
type SeedOutput struct{}

func Seed(ctx context.Context, input SeedInput) (SeedOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return SeedOutput{}, fmt.Errorf("TODO: Seed")
}
