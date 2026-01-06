package valueobject

import (
	"errors"
)

var ErrInvalidPlatform = errors.New("invalid platform")

const (
	PlatformInstagram = "instagram"
	PlatformFacebook  = "facebook"
	PlatformLinkedIn  = "linkedin"
	PlatformYouTube   = "youtube"
	PlatformTikTok    = "tiktok"
)

// ValidPlatforms contém todas as plataformas válidas
var ValidPlatforms = map[string]bool{
	PlatformInstagram: true,
	PlatformFacebook:  true,
	PlatformLinkedIn:  true,
	PlatformYouTube:   true,
	PlatformTikTok:    true,
}

// IsValidPlatform verifica se a plataforma é válida
func IsValidPlatform(platform string) bool {
	return ValidPlatforms[platform]
}

// ValidatePlatforms verifica se todas as plataformas na lista são válidas
func ValidatePlatforms(platforms []string) error {
	for _, platform := range platforms {
		if !IsValidPlatform(platform) {
			return errors.New("invalid platform: " + platform)
		}
	}
	return nil
}

// GetValidPlatforms retorna a lista de todas as plataformas válidas
func GetValidPlatforms() []string {
	return []string{
		PlatformInstagram,
		PlatformFacebook,
		PlatformLinkedIn,
		PlatformYouTube,
		PlatformTikTok,
	}
}
