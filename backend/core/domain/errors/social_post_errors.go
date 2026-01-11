package errors

import "errors"

var (
	// ErrSocialPostIDRequired is returned when the social post ID is required
	ErrSocialPostIDRequired = errors.New("social post ID is required")

	// ErrSocialPostNotFound is returned when a social post is not found
	ErrSocialPostNotFound = errors.New("social post not found")

	// ErrInvalidSocialPostBrandName is returned when the brand name is invalid
	ErrInvalidSocialPostBrandName = errors.New("invalid brand name")

	// ErrInvalidSocialPostPlatform is returned when the platform is invalid
	ErrInvalidSocialPostPlatform = errors.New("invalid platform")

	// ErrInvalidSocialPostType is returned when the post type is invalid
	ErrInvalidSocialPostType = errors.New("invalid post type")

	// ErrNegativeSocialPostLikes is returned when likes are negative
	ErrNegativeSocialPostLikes = errors.New("likes cannot be negative")

	// ErrNegativeSocialPostComments is returned when comments are negative
	ErrNegativeSocialPostComments = errors.New("comments cannot be negative")

	// ErrNegativeSocialPostShares is returned when shares are negative
	ErrNegativeSocialPostShares = errors.New("shares cannot be negative")

	// ErrNegativeSocialPostFollowers is returned when followers are negative
	ErrNegativeSocialPostFollowers = errors.New("followers cannot be negative")

	// ErrSocialPostAlreadyDeleted is returned when trying to update a deleted post
	ErrSocialPostAlreadyDeleted = errors.New("social post is already deleted")
)
