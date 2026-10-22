package lab

import (
	"context"
	"fmt"
)

// EditOwnedPost: Apply validated owner edit conditionally.
// TODO: define these fields from the specification and use trusted context.
type EditOwnedPostInput struct{}
type EditOwnedPostOutput struct{}

func EditOwnedPost(ctx context.Context, input EditOwnedPostInput) (EditOwnedPostOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return EditOwnedPostOutput{}, fmt.Errorf("TODO: EditOwnedPost")
}

// EditOwnedComment: Apply owner edit or forbidden/conflict.
// TODO: define these fields from the specification and use trusted context.
type EditOwnedCommentInput struct{}
type EditOwnedCommentOutput struct{}

func EditOwnedComment(ctx context.Context, input EditOwnedCommentInput) (EditOwnedCommentOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return EditOwnedCommentOutput{}, fmt.Errorf("TODO: EditOwnedComment")
}

// DeleteOwnedContent: Soft-delete and enforce child visibility consistently.
// TODO: define these fields from the specification and use trusted context.
type DeleteOwnedContentInput struct{}
type DeleteOwnedContentOutput struct{}

func DeleteOwnedContent(ctx context.Context, input DeleteOwnedContentInput) (DeleteOwnedContentOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return DeleteOwnedContentOutput{}, fmt.Errorf("TODO: DeleteOwnedContent")
}
