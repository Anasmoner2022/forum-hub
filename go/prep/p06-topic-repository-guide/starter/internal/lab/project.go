package lab

import (
	"context"
	"fmt"
)

// CreateTopic: Return created ID after the topic and all label links commit.
// TODO: define these fields from the specification and use trusted context.
type CreateTopicInput struct{}
type CreateTopicOutput struct{}

func CreateTopic(ctx context.Context, input CreateTopicInput) (CreateTopicOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return CreateTopicOutput{}, fmt.Errorf("TODO: CreateTopic")
}

// CreateReply: Return reply ID or a domain error.
// TODO: define these fields from the specification and use trusted context.
type CreateReplyInput struct{}
type CreateReplyOutput struct{}

func CreateReply(ctx context.Context, input CreateReplyInput) (CreateReplyOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return CreateReplyOutput{}, fmt.Errorf("TODO: CreateReply")
}

// GetTopic: Return a public detail DTO or null.
// TODO: define these fields from the specification and use trusted context.
type GetTopicInput struct{}
type GetTopicOutput struct{}

func GetTopic(ctx context.Context, input GetTopicInput) (GetTopicOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return GetTopicOutput{}, fmt.Errorf("TODO: GetTopic")
}
