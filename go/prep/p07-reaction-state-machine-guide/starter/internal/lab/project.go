package lab

import (
	"context"
	"fmt"
)

// SetReaction: Set the requested current value atomically; target={kind,id}.
// TODO: define these fields from the specification and use trusted context.
type SetReactionInput struct{}
type SetReactionOutput struct{}

func SetReaction(ctx context.Context, input SetReactionInput) (SetReactionOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return SetReactionOutput{}, fmt.Errorf("TODO: SetReaction")
}

// ClearReaction: Remove the current value if present.
// TODO: define these fields from the specification and use trusted context.
type ClearReactionInput struct{}
type ClearReactionOutput struct{}

func ClearReaction(ctx context.Context, input ClearReactionInput) (ClearReactionOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return ClearReactionOutput{}, fmt.Errorf("TODO: ClearReaction")
}

// ReactionCounts: Return {likes, dislikes} with zero defaults.
// TODO: define these fields from the specification and use trusted context.
type ReactionCountsInput struct{}
type ReactionCountsOutput struct{}

func ReactionCounts(ctx context.Context, input ReactionCountsInput) (ReactionCountsOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return ReactionCountsOutput{}, fmt.Errorf("TODO: ReactionCounts")
}
