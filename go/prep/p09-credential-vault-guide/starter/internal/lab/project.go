package lab

import (
	"context"
	"fmt"
)

// RegisterMember: Create one member with a hash record; duplicate identity is a conflict.
// TODO: define these fields from the specification and use trusted context.
type RegisterMemberInput struct{}
type RegisterMemberOutput struct{}

func RegisterMember(ctx context.Context, input RegisterMemberInput) (RegisterMemberOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return RegisterMemberOutput{}, fmt.Errorf("TODO: RegisterMember")
}

// HashPassword: Resolve a versioned record, never the plaintext.
// TODO: define these fields from the specification and use trusted context.
type HashPasswordInput struct{}
type HashPasswordOutput struct{}

func HashPassword(ctx context.Context, input HashPasswordInput) (HashPasswordOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return HashPasswordOutput{}, fmt.Errorf("TODO: HashPassword")
}

// VerifyCredentials: Resolve a public member context or a generic invalid-credentials result.
// TODO: define these fields from the specification and use trusted context.
type VerifyCredentialsInput struct{}
type VerifyCredentialsOutput struct{}

func VerifyCredentials(ctx context.Context, input VerifyCredentialsInput) (VerifyCredentialsOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return VerifyCredentialsOutput{}, fmt.Errorf("TODO: VerifyCredentials")
}
