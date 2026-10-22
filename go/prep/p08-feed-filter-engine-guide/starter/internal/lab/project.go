package lab

import (
	"context"
	"fmt"
)

// ListFeed: Return {items, page, hasNext}; viewer is null or trusted member context.
// TODO: define these fields from the specification and use trusted context.
type ListFeedInput struct{}
type ListFeedOutput struct{}

func ListFeed(ctx context.Context, input ListFeedInput) (ListFeedOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return ListFeedOutput{}, fmt.Errorf("TODO: ListFeed")
}
