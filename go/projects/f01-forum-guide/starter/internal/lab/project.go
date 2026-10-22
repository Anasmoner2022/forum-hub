package lab

import (
	"context"
	"fmt"
)

// Dispatch: Map the documented routes to trusted guards and prepared services.
// TODO: define these fields from the specification and use trusted context.
type DispatchInput struct{}
type DispatchOutput struct{}

func Dispatch(ctx context.Context, input DispatchInput) (DispatchOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return DispatchOutput{}, fmt.Errorf("TODO: Dispatch")
}

// CreatePost: Reuse P06 atomic creation with a session-derived author.
// TODO: define these fields from the specification and use trusted context.
type CreatePostInput struct{}
type CreatePostOutput struct{}

func CreatePost(ctx context.Context, input CreatePostInput) (CreatePostOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return CreatePostOutput{}, fmt.Errorf("TODO: CreatePost")
}

// CreateComment: Reuse prepared reply service with authorization.
// TODO: define these fields from the specification and use trusted context.
type CreateCommentInput struct{}
type CreateCommentOutput struct{}

func CreateComment(ctx context.Context, input CreateCommentInput) (CreateCommentOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return CreateCommentOutput{}, fmt.Errorf("TODO: CreateComment")
}

// SetForumReaction: Apply P07 contract to posts and comments.
// TODO: define these fields from the specification and use trusted context.
type SetForumReactionInput struct{}
type SetForumReactionOutput struct{}

func SetForumReaction(ctx context.Context, input SetForumReactionInput) (SetForumReactionOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return SetForumReactionOutput{}, fmt.Errorf("TODO: SetForumReaction")
}

// GetForumFeed: Return P08 public cards with current-user filter semantics.
// TODO: define these fields from the specification and use trusted context.
type GetForumFeedInput struct{}
type GetForumFeedOutput struct{}

func GetForumFeed(ctx context.Context, input GetForumFeedInput) (GetForumFeedOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return GetForumFeedOutput{}, fmt.Errorf("TODO: GetForumFeed")
}
