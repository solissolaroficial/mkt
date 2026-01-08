package response

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
)

// PdvPayloadMapper converte entre Entity e Response DTO para PDV
type PdvPayloadMapper struct{}

// NewPdvPayloadMapper cria uma nova instância do PdvPayloadMapper
func NewPdvPayloadMapper() *PdvPayloadMapper {
	return &PdvPayloadMapper{}
}

// ToPdvPostResponse converte Entity para Response DTO
func (m *PdvPayloadMapper) ToPdvPostResponse(post *entity.PdvPost) PdvPostResponse {
	// Obter status string
	var statusStr string
	if post.Status() != nil {
		statusStr = post.Status().String()
	} else {
		statusStr = "pending"
	}

	return PdvPostResponse{
		ID:        post.ID().String(),
		RepName:   post.RepName(),
		PdvName:   post.PdvName(),
		PostDate:  post.PostDate().String(),
		Month:     post.Month(),
		Platform:  post.Platform(),
		Link:      post.Link(),
		ProofUrl:  post.ProofUrl(),
		Status:    statusStr,
		CreatedAt: post.CreatedAt().Format(time.RFC3339),
		UpdatedAt: post.UpdatedAt().Format(time.RFC3339),
	}
}

// ToPdvPostResponseList converte slice de Entity para slice de Response DTO
func (m *PdvPayloadMapper) ToPdvPostResponseList(posts []*entity.PdvPost) []PdvPostResponse {
	responses := make([]PdvPostResponse, len(posts))
	for i, post := range posts {
		responses[i] = m.ToPdvPostResponse(post)
	}
	return responses
}

// ToRecurrentPdvResponse converte Entity para Response DTO
func (m *PdvPayloadMapper) ToRecurrentPdvResponse(pdv *entity.RecurrentPdv) RecurrentPdvResponse {
	return RecurrentPdvResponse{
		ID:               pdv.ID().String(),
		Name:             pdv.Name(),
		RepName:          pdv.RepName(),
		City:             pdv.City(),
		Followers:        pdv.Followers(),
		InstagramProfile: pdv.InstagramProfile(),
		CreatedAt:        pdv.CreatedAt().Format(time.RFC3339),
		UpdatedAt:        pdv.UpdatedAt().Format(time.RFC3339),
	}
}

// ToRecurrentPdvResponseList converte slice de Entity para slice de Response DTO
func (m *PdvPayloadMapper) ToRecurrentPdvResponseList(pdvs []*entity.RecurrentPdv) []RecurrentPdvResponse {
	responses := make([]RecurrentPdvResponse, len(pdvs))
	for i, pdv := range pdvs {
		responses[i] = m.ToRecurrentPdvResponse(pdv)
	}
	return responses
}

// ToPdvPostsListResponse cria a resposta paginada de posts
func (m *PdvPayloadMapper) ToPdvPostsListResponse(
	posts []*entity.PdvPost,
	total int64,
	page int,
	limit int,
) PdvPostListData {
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	if totalPages < 0 {
		totalPages = 0
	}
	return PdvPostListData{
		Posts: m.ToPdvPostResponseList(posts),
		Meta: MetaResponse{
			Total:      total,
			Page:       page,
			Limit:      limit,
			TotalPages: totalPages,
		},
	}
}

// ToRecurrentPdvsListResponse cria a resposta paginada de PDVs recorrentes
func (m *PdvPayloadMapper) ToRecurrentPdvsListResponse(
	pdvs []*entity.RecurrentPdv,
	total int64,
	page int,
	limit int,
) RecurrentPdvListData {
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	if totalPages < 0 {
		totalPages = 0
	}
	return RecurrentPdvListData{
		Pdvs: m.ToRecurrentPdvResponseList(pdvs),
		Meta: MetaResponse{
			Total:      total,
			Page:       page,
			Limit:      limit,
			TotalPages: totalPages,
		},
	}
}
