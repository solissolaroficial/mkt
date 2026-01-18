package mapper

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

type CalendarPostMapper struct{}

func NewCalendarPostMapper() *CalendarPostMapper {
	return &CalendarPostMapper{}
}

func (m *CalendarPostMapper) ModelToEntity(model *model.CalendarPostModel) (*entity.CalendarPost, error) {
	// Converter plataformas
	var platforms []string
	if model.Platforms != nil {
		if err := json.Unmarshal(model.Platforms, &platforms); err != nil {
			return nil, fmt.Errorf("failed to unmarshal platforms: %w", err)
		}
	}

	// Converter plataformas publicadas
	var publishedPlatforms []string
	if model.PublishedPlatforms != nil {
		if err := json.Unmarshal(model.PublishedPlatforms, &publishedPlatforms); err != nil {
			return nil, fmt.Errorf("failed to unmarshal published platforms: %w", err)
		}
	}

	// Converter histórico
	var history []*valueobject.PostHistoryEvent
	if model.History != nil {
		if err := json.Unmarshal(model.History, &history); err != nil {
			return nil, fmt.Errorf("failed to unmarshal history: %w", err)
		}
	}

	// Criar value objects (usar ReconstructPostDate para dados do banco)
	postDate := valueobject.ReconstructPostDate(model.Date.Format("2006-01-02"))
	postTime := valueobject.ReconstructPostTime(model.Time)
	category, err := valueobject.NewPostCategory(model.Category)
	if err != nil {
		return nil, err
	}
	postType, err := valueobject.NewPostType(model.Type)
	if err != nil {
		return nil, err
	}
	status, err := valueobject.NewPostStatus(model.Status)
	if err != nil {
		return nil, err
	}

	// Converter deletedAt
	var deletedAt *time.Time
	if model.DeletedAt.Valid {
		deletedAt = &model.DeletedAt.Time
	}

	// Reconstruir entidade
	calendarPost := entity.ReconstructCalendarPost(
		model.UUID,
		model.Title,
		model.Description,
		postDate,
		postTime,
		model.Caption,
		category,
		postType,
		status,
		model.AssigneeUUID,
		platforms,
		publishedPlatforms,
		model.ImageURL,
		history,
		model.CreatedAt,
		model.UpdatedAt,
		deletedAt,
	)

	return calendarPost, nil
}

func (m *CalendarPostMapper) EntityToModel(entity *entity.CalendarPost) (*model.CalendarPostModel, error) {
	// Converter plataformas para JSON
	platformsJSON, err := json.Marshal(entity.Platforms())
	if err != nil {
		return nil, err
	}

	// Converter plataformas publicadas para JSON
	publishedPlatformsJSON, err := json.Marshal(entity.PublishedPlatforms())
	if err != nil {
		return nil, err
	}

	// Converter histórico para JSON
	historyJSON, err := json.Marshal(entity.History())
	if err != nil {
		return nil, err
	}

	// Converter data
	date, err := time.Parse("2006-01-02", entity.Date().String())
	if err != nil {
		return nil, err
	}

	return &model.CalendarPostModel{
		UUID:               entity.ID(),
		Title:              entity.Title(),
		Description:        entity.Description(),
		Date:               date,
		Time:               entity.PostTime().String(),
		Caption:            entity.Caption(),
		Category:           entity.Category().String(),
		Type:               entity.Type().String(),
		Status:             entity.Status().String(),
		AssigneeUUID:       entity.AssigneeUUID(),
		Platforms:          platformsJSON,
		PublishedPlatforms: publishedPlatformsJSON,
		ImageURL:           entity.ImageURL(),
		History:            historyJSON,
		CreatedAt:          entity.CreatedAt(),
		UpdatedAt:          entity.UpdatedAt(),
	}, nil
}

func (m *CalendarPostMapper) EntitiesToModels(entities []*entity.CalendarPost) ([]*model.CalendarPostModel, error) {
	models := make([]*model.CalendarPostModel, 0, len(entities))
	for _, entity := range entities {
		model, err := m.EntityToModel(entity)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}
	return models, nil
}

func (m *CalendarPostMapper) ModelsToEntities(models []*model.CalendarPostModel) ([]*entity.CalendarPost, error) {
	entities := make([]*entity.CalendarPost, 0, len(models))
	for _, model := range models {
		entity, err := m.ModelToEntity(model)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	return entities, nil
}
