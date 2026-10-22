package lab

import (
	"context"
	"fmt"
)

// LoadTlsConfig: Validate readable certificate/key paths and return server options.
// TODO: define these fields from the specification and use trusted context.
type LoadTlsConfigInput struct{}
type LoadTlsConfigOutput struct{}

func LoadTlsConfig(ctx context.Context, input LoadTlsConfigInput) (LoadTlsConfigOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return LoadTlsConfigOutput{}, fmt.Errorf("TODO: LoadTlsConfig")
}
