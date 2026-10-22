package lab

import (
	"context"
	"fmt"
)

// ValidateAndNormalizeImage: Resolve verified output metadata and normalized bytes/path.
// TODO: define these fields from the specification and use trusted context.
type ValidateAndNormalizeImageInput struct{}
type ValidateAndNormalizeImageOutput struct{}

func ValidateAndNormalizeImage(ctx context.Context, input ValidateAndNormalizeImageInput) (ValidateAndNormalizeImageOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return ValidateAndNormalizeImageOutput{}, fmt.Errorf("TODO: ValidateAndNormalizeImage")
}

// PublishPostWithAsset: Commit ownership and post linkage; compensate filesystem failures.
// TODO: define these fields from the specification and use trusted context.
type PublishPostWithAssetInput struct{}
type PublishPostWithAssetOutput struct{}

func PublishPostWithAsset(ctx context.Context, input PublishPostWithAssetInput) (PublishPostWithAssetOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return PublishPostWithAssetOutput{}, fmt.Errorf("TODO: PublishPostWithAsset")
}

// ServeAsset: Authorize parent visibility and return safe verified headers.
// TODO: define these fields from the specification and use trusted context.
type ServeAssetInput struct{}
type ServeAssetOutput struct{}

func ServeAsset(ctx context.Context, input ServeAssetInput) (ServeAssetOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return ServeAssetOutput{}, fmt.Errorf("TODO: ServeAsset")
}
