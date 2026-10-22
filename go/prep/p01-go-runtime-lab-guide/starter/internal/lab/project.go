package lab

import (
	"context"
	"fmt"
)

// InspectConfig: Return the validated report; reject with a classified error on bad input.
// TODO: define these fields from the specification and use trusted context.
type InspectConfigInput struct{}
type InspectConfigOutput struct{}

func InspectConfig(ctx context.Context, input InspectConfigInput) (InspectConfigOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return InspectConfigOutput{}, fmt.Errorf("TODO: InspectConfig")
}
