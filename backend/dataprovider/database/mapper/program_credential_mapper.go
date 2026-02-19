package mapper

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

type ProgramCredentialMapper struct{}

func NewProgramCredentialMapper() *ProgramCredentialMapper {
	return &ProgramCredentialMapper{}
}

func (m *ProgramCredentialMapper) ModelToEntity(model *model.ProgramCredentialModel) *entity.ProgramCredential {
	return entity.ReconstructProgramCredential(
		model.UUID,
		model.Name,
		model.User,
		model.Password,
		model.Access,
		model.Notes,
		model.Active,
		model.CreatedAt,
		model.UpdatedAt,
	)
}

func (m *ProgramCredentialMapper) ModelsToEntities(models []model.ProgramCredentialModel) []*entity.ProgramCredential {
	credentials := make([]*entity.ProgramCredential, len(models))
	for i, model := range models {
		credentials[i] = m.ModelToEntity(&model)
	}
	return credentials
}

func (m *ProgramCredentialMapper) EntityToModel(credential *entity.ProgramCredential) *model.ProgramCredentialModel {
	return &model.ProgramCredentialModel{
		UUID:      credential.ID(),
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
