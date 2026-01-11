package social

import (
	"context"

	"github.com/google/uuid"
	socialerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// DeleteSocialPostUseCase handles deleting a social post
type DeleteSocialPostUseCase struct {
	postGateway                    gateway.SocialPostGateway
	recalculateAggregationsUseCase *RecalculateDailyAggregationsUseCase
}

// NewDeleteSocialPostUseCase creates a new delete social post use case
func NewDeleteSocialPostUseCase(
	postGateway gateway.SocialPostGateway,
	recalculateAggregationsUseCase *RecalculateDailyAggregationsUseCase,
) *DeleteSocialPostUseCase {
	return &DeleteSocialPostUseCase{
		postGateway:                    postGateway,
		recalculateAggregationsUseCase: recalculateAggregationsUseCase,
	}
}

// Execute deletes a social post
func (uc *DeleteSocialPostUseCase) Execute(ctx context.Context, id string) error {
	if id == "" {
		return socialerrors.ErrSocialPostIDRequired
	}

	// Parse UUID
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return socialerrors.ErrSocialPostNotFound
	}

	// Get the post to check if it exists
	post, err := uc.postGateway.GetByID(uuidID)
	if err != nil {
		if err == socialerrors.ErrSocialPostNotFound {
			return err
		}
		return err
	}

	// Check if post is already deleted
	if !post.IsActive() {
		return socialerrors.ErrSocialPostAlreadyDeleted
	}

	// Soft delete the post
	post.SoftDelete()

	// Update in database
	if err := uc.postGateway.Update(post); err != nil {
		return err
	}

	// Recalculate daily aggregations
	_, err = uc.recalculateAggregationsUseCase.Execute(post.BrandName().Value(), post.PostDate())
	if err != nil {
		// Log error but don't fail delete
		// In production, this should be asynchronous
		// We still return success even if recalculation fails
	}

	return nil
}
