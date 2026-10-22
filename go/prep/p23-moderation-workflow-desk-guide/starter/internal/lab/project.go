package lab

import (
	"context"
	"fmt"
)

// ReviewPost: Apply one valid review and audit it atomically.
// TODO: define these fields from the specification and use trusted context.
type ReviewPostInput struct{}
type ReviewPostOutput struct{}

func ReviewPost(ctx context.Context, input ReviewPostInput) (ReviewPostOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return ReviewPostOutput{}, fmt.Errorf("TODO: ReviewPost")
}

// DecideRoleRequest: Record decision and change role when approved.
// TODO: define these fields from the specification and use trusted context.
type DecideRoleRequestInput struct{}
type DecideRoleRequestOutput struct{}

func DecideRoleRequest(ctx context.Context, input DecideRoleRequestInput) (DecideRoleRequestOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return DecideRoleRequestOutput{}, fmt.Errorf("TODO: DecideRoleRequest")
}

// RespondToReport: Persist admin response and resolution state.
// TODO: define these fields from the specification and use trusted context.
type RespondToReportInput struct{}
type RespondToReportOutput struct{}

func RespondToReport(ctx context.Context, input RespondToReportInput) (RespondToReportOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return RespondToReportOutput{}, fmt.Errorf("TODO: RespondToReport")
}

// DeleteCategory: Reject retained-post references with conflict.
// TODO: define these fields from the specification and use trusted context.
type DeleteCategoryInput struct{}
type DeleteCategoryOutput struct{}

func DeleteCategory(ctx context.Context, input DeleteCategoryInput) (DeleteCategoryOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return DeleteCategoryOutput{}, fmt.Errorf("TODO: DeleteCategory")
}
