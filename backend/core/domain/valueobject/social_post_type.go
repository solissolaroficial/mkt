package valueobject

import (
	"errors"
	"strings"
)

// Valid social media post types
var validSocialPostTypes = map[string]bool{
	"image":    true,
	"video":    true,
	"carousel": true,
	"reel":     true,
	"story":    true,
}

// SocialPostType represents a social media post type
type SocialPostType struct {
	value string
}

// NewSocialPostType creates a new SocialPostType value object
func NewSocialPostType(postType string) (SocialPostType, error) {
	postType = strings.ToLower(strings.TrimSpace(postType))

	if postType == "" {
		return SocialPostType{}, errors.New("post type cannot be empty")
	}

	if !validSocialPostTypes[postType] {
		return SocialPostType{}, errors.New("invalid post type: must be one of image, video, carousel, reel, story")
	}

	return SocialPostType{value: postType}, nil
}

// String returns string representation of post type
func (p SocialPostType) String() string {
	return p.value
}

// IsValid checks if post type is valid
func (p SocialPostType) IsValid() bool {
	return validSocialPostTypes[p.value]
}

// ReconstructSocialPostType creates a SocialPostType from a string without validation
// This is used when reconstructing entities from database
func ReconstructSocialPostType(postType string) SocialPostType {
	return SocialPostType{value: postType}
}
