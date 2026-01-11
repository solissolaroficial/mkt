package valueobject

import (
	"errors"
	"strings"
)

// Valid social media platforms
var validSocialPlatforms = map[string]bool{
	"instagram": true,
	"facebook":  true,
	"linkedin":  true,
	"tiktok":    true,
	"twitter":   true,
}

// SocialPlatform represents a social media platform
type SocialPlatform struct {
	value string
}

// NewSocialPlatform creates a new SocialPlatform value object
func NewSocialPlatform(platform string) (SocialPlatform, error) {
	platform = strings.ToLower(strings.TrimSpace(platform))

	if platform == "" {
		return SocialPlatform{}, errors.New("platform cannot be empty")
	}

	if !validSocialPlatforms[platform] {
		return SocialPlatform{}, errors.New("invalid platform: must be one of instagram, facebook, linkedin, tiktok, twitter")
	}

	return SocialPlatform{value: platform}, nil
}

// String returns string representation of platform
func (p SocialPlatform) String() string {
	return p.value
}

// IsValid checks if platform is valid
func (p SocialPlatform) IsValid() bool {
	return validSocialPlatforms[p.value]
}

// ReconstructSocialPlatform creates a SocialPlatform from a string without validation
// This is used when reconstructing entities from database
func ReconstructSocialPlatform(platform string) SocialPlatform {
	return SocialPlatform{value: platform}
}
