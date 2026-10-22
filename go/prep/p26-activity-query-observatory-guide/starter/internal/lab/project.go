package lab

import (
	"context"
	"fmt"
)

// ListActivity: Return only the actor’s paginated public/authorized activity DTOs.
// TODO: define these fields from the specification and use trusted context.
type ListActivityInput struct{}
type ListActivityOutput struct{}

func ListActivity(ctx context.Context, input ListActivityInput) (ListActivityOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return ListActivityOutput{}, fmt.Errorf("TODO: ListActivity")
}

// InspectActivityPlans: Return query-plan evidence for the documented queries.
// TODO: define these fields from the specification and use trusted context.
type InspectActivityPlansInput struct{}
type InspectActivityPlansOutput struct{}

func InspectActivityPlans(ctx context.Context, input InspectActivityPlansInput) (InspectActivityPlansOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return InspectActivityPlansOutput{}, fmt.Errorf("TODO: InspectActivityPlans")
}
