package response

import (
	"time"

	"github.com/google/uuid"
)

// RepresentativeResponse represents a representative response
type RepresentativeResponse struct {
	UUID      uuid.UUID  `json:"uuid"`
	Code      int        `json:"code"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Company   string     `json:"company"`
	Region    string     `json:"region"`
	City      string     `json:"city"`
	Attendant string     `json:"attendant"`
	Active    bool       `json:"active"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

// RepresentativeStatsResponse represents statistics for a representative
type RepresentativeStatsResponse struct {
	UUID               string `json:"uuid"`
	OnlineCount        int64  `json:"onlineCount"` // Alias para compatibilidade com implementação antiga
	OnlineActionCount  int64  `json:"onlineActionCount"`
	OfflineCount       int64  `json:"offlineCount"` // Alias para compatibilidade com implementação antiga
	OfflineActionCount int64  `json:"offlineActionCount"`
	OfflineValue       int64  `json:"offlineValue"` // Alias para compatibilidade com implementação antiga
	OfflineActionValue int64  `json:"offlineActionValue"`
	ShowroomItemCount  int64  `json:"showroomItemCount"`
	RepMarketingCount  int64  `json:"repMarketingCount"`
	TotalActions       int64  `json:"totalActions"`
}

// RepresentativeProfileResponse represents a representative profile with stats
type RepresentativeProfileResponse struct {
	UUID              string `json:"uuid"`
	Code              int    `json:"code"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	Company           string `json:"company"`
	Region            string `json:"region"`
	City              string `json:"city"`
	Attendant         string `json:"attendant"`
	Active            bool   `json:"active"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
	TrainingCount     int64  `json:"trainingCount"`
	OnlineCount       int64  `json:"onlineCount"`
	OfflineCount      int64  `json:"offlineCount"`
	OfflineValue      int64  `json:"offlineValue"`
	ShowroomItemCount int64  `json:"showroomItemCount"`
	RepMarketingCount int64  `json:"repMarketingCount"`
	TotalActions      int64  `json:"totalActions"`
}

// GetAllRepresentativeProfilesResponse represents response for all representative profiles
type GetAllRepresentativeProfilesResponse struct {
	Data []RepresentativeProfileResponse `json:"data"`
}

// ListRepresentativesResponse represents a paginated response of representatives
type ListRepresentativesResponse struct {
	Data       []RepresentativeResponse `json:"data"`
	Total      int64                    `json:"total"`
	Page       int                      `json:"page"`
	PageSize   int                      `json:"pageSize"`
	TotalPages int                      `json:"totalPages"`
}
