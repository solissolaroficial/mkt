package controller

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// ProgramCredentialPayloadMapper maps entity to response DTO
type ProgramCredentialPayloadMapper struct{}

func NewProgramCredentialPayloadMapper() *ProgramCredentialPayloadMapper {
	return &ProgramCredentialPayloadMapper{}
}

func (m *ProgramCredentialPayloadMapper) ToResponse(credential *entity.ProgramCredential) response.ProgramCredentialResponse {
	return response.ProgramCredentialResponse{
		ID:        credential.ID().String(),
		Name:      credential.Name(),
		User:      credential.User(),
		Password:  credential.Password(),
		Access:    credential.Access(),
		Notes:     credential.Notes(),
		Active:    credential.Active(),
		CreatedAt: credential.CreatedAt(),
		UpdatedAt: credential.UpdatedAt(),
	}
}

func (m *ProgramCredentialPayloadMapper) ToResponseList(credentials []*entity.ProgramCredential) []response.ProgramCredentialResponse {
	result := make([]response.ProgramCredentialResponse, len(credentials))
	for i, cred := range credentials {
		result[i] = m.ToResponse(cred)
	}
	return result
}
