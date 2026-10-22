package lab

import (
	"context"
	"fmt"
)

// ResolveGoogleIdentity: Verify Google claims and return {provider,subject,email,emailVerified}.
// TODO: define these fields from the specification and use trusted context.
type ResolveGoogleIdentityInput struct{}
type ResolveGoogleIdentityOutput struct{}

func ResolveGoogleIdentity(ctx context.Context, input ResolveGoogleIdentityInput) (ResolveGoogleIdentityOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return ResolveGoogleIdentityOutput{}, fmt.Errorf("TODO: ResolveGoogleIdentity")
}

// ResolveGithubIdentity: Return stable GitHub ID with explicitly verified contact data.
// TODO: define these fields from the specification and use trusted context.
type ResolveGithubIdentityInput struct{}
type ResolveGithubIdentityOutput struct{}

func ResolveGithubIdentity(ctx context.Context, input ResolveGithubIdentityInput) (ResolveGithubIdentityOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return ResolveGithubIdentityOutput{}, fmt.Errorf("TODO: ResolveGithubIdentity")
}

// FindOrCreateMember: Resolve a unique mapping; never auto-link by email.
// TODO: define these fields from the specification and use trusted context.
type FindOrCreateMemberInput struct{}
type FindOrCreateMemberOutput struct{}

func FindOrCreateMember(ctx context.Context, input FindOrCreateMemberInput) (FindOrCreateMemberOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return FindOrCreateMemberOutput{}, fmt.Errorf("TODO: FindOrCreateMember")
}

// LinkIdentity: Attach only after recent reauthentication and conflict checks.
// TODO: define these fields from the specification and use trusted context.
type LinkIdentityInput struct{}
type LinkIdentityOutput struct{}

func LinkIdentity(ctx context.Context, input LinkIdentityInput) (LinkIdentityOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return LinkIdentityOutput{}, fmt.Errorf("TODO: LinkIdentity")
}
