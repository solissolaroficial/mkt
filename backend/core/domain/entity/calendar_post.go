package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type CalendarPost struct {
	id                 uuid.UUID
	title              string
	description        *string
	date               *valueobject.PostDate
	postTime           *valueobject.PostTime
	caption            *string
	category           valueobject.PostCategory
	postType           valueobject.PostType
	status             valueobject.PostStatus
	assigneeID         *uuid.UUID
	platforms          []string
	publishedPlatforms []string
	imageURL           *string
	history            []*valueobject.PostHistoryEvent
	createdAt          time.Time
	updatedAt          time.Time
	deletedAt          *time.Time
}

// NewCalendarPost cria uma nova entidade CalendarPost
func NewCalendarPost(
	title string,
	description *string,
	date *valueobject.PostDate,
	postTime *valueobject.PostTime,
	caption *string,
	category valueobject.PostCategory,
	postType valueobject.PostType,
	assigneeID *uuid.UUID,
	platforms []string,
	imageURL *string,
) (*CalendarPost, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	if assigneeID == nil {
		return nil, errors.New("assignee is required")
	}

	post := &CalendarPost{
		id:                 uuid.New(),
		title:              title,
		description:        description,
		date:               date,
		postTime:           postTime,
		caption:            caption,
		category:           category,
		postType:           postType,
		status:             valueobject.StatusInProgress, // status inicial
		assigneeID:         assigneeID,
		platforms:          platforms,
		publishedPlatforms: []string{},
		imageURL:           imageURL,
		history:            []*valueobject.PostHistoryEvent{},
		createdAt:          time.Now(),
		updatedAt:          time.Now(),
	}

	if err := post.Validate(); err != nil {
		return nil, err
	}

	return post, nil
}

// ReconstructCalendarPost reconstrói a entidade do banco de dados
func ReconstructCalendarPost(
	id uuid.UUID,
	title string,
	description *string,
	date *valueobject.PostDate,
	postTime *valueobject.PostTime,
	caption *string,
	category valueobject.PostCategory,
	postType valueobject.PostType,
	status valueobject.PostStatus,
	assigneeID *uuid.UUID,
	platforms []string,
	publishedPlatforms []string,
	imageURL *string,
	history []*valueobject.PostHistoryEvent,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *CalendarPost {
	return &CalendarPost{
		id:                 id,
		title:              title,
		description:        description,
		date:               date,
		postTime:           postTime,
		caption:            caption,
		category:           category,
		postType:           postType,
		status:             status,
		assigneeID:         assigneeID,
		platforms:          platforms,
		publishedPlatforms: publishedPlatforms,
		imageURL:           imageURL,
		history:            history,
		createdAt:          createdAt,
		updatedAt:          updatedAt,
		deletedAt:          deletedAt,
	}
}

// Getters
func (p *CalendarPost) ID() uuid.UUID                            { return p.id }
func (p *CalendarPost) Title() string                            { return p.title }
func (p *CalendarPost) Description() *string                     { return p.description }
func (p *CalendarPost) Date() *valueobject.PostDate              { return p.date }
func (p *CalendarPost) PostTime() *valueobject.PostTime          { return p.postTime }
func (p *CalendarPost) Caption() *string                         { return p.caption }
func (p *CalendarPost) Category() valueobject.PostCategory       { return p.category }
func (p *CalendarPost) Type() valueobject.PostType               { return p.postType }
func (p *CalendarPost) Status() valueobject.PostStatus           { return p.status }
func (p *CalendarPost) AssigneeID() *uuid.UUID                   { return p.assigneeID }
func (p *CalendarPost) Platforms() []string                      { return p.platforms }
func (p *CalendarPost) PublishedPlatforms() []string             { return p.publishedPlatforms }
func (p *CalendarPost) ImageURL() *string                        { return p.imageURL }
func (p *CalendarPost) History() []*valueobject.PostHistoryEvent { return p.history }
func (p *CalendarPost) CreatedAt() time.Time                     { return p.createdAt }
func (p *CalendarPost) UpdatedAt() time.Time                     { return p.updatedAt }
func (p *CalendarPost) DeletedAt() *time.Time                    { return p.deletedAt }

// Métodos de Negócio

// Validate valida os dados da entidade
func (p *CalendarPost) Validate() error {
	if p.title == "" {
		return errors.New("title is required")
	}

	if len(p.title) > 500 {
		return errors.New("title must be at most 500 characters")
	}

	if p.assigneeID == nil {
		return errors.New("assignee is required")
	}

	if !p.category.IsValid() {
		return errors.New("invalid category")
	}

	if !p.postType.IsValid() {
		return errors.New("invalid type")
	}

	if !p.status.IsValid() {
		return errors.New("invalid status")
	}

	return nil
}

// UpdateTitle atualiza o título
func (p *CalendarPost) UpdateTitle(title string) error {
	if title == "" {
		return errors.New("title is required")
	}

	if len(title) > 500 {
		return errors.New("title must be at most 500 characters")
	}

	p.title = title
	p.updatedAt = time.Now()
	return nil
}

// UpdateDescription atualiza a descrição
func (p *CalendarPost) UpdateDescription(description *string) {
	p.description = description
	p.updatedAt = time.Now()
}

// UpdateCaption atualiza a legenda
func (p *CalendarPost) UpdateCaption(caption *string) {
	p.caption = caption
	p.updatedAt = time.Now()
}

// UpdateImageURL atualiza a URL da imagem
func (p *CalendarPost) UpdateImageURL(imageURL *string) {
	p.imageURL = imageURL
	p.updatedAt = time.Now()
}

// UpdateStatus atualiza o status com validação de transição
func (p *CalendarPost) UpdateStatus(newStatus valueobject.PostStatus) error {
	if !newStatus.IsValid() {
		return errors.New("invalid status")
	}

	// Validar transição de status
	if !p.status.CanTransitionTo(newStatus) {
		return errors.New("invalid status transition")
	}

	p.status = newStatus
	p.updatedAt = time.Now()
	return nil
}

// UpdatePlatforms atualiza as plataformas
func (p *CalendarPost) UpdatePlatforms(platforms []string) error {
	if err := valueobject.ValidatePlatforms(platforms); err != nil {
		return err
	}

	p.platforms = platforms
	p.updatedAt = time.Now()
	return nil
}

// UpdatePublishedPlatforms atualiza as plataformas publicadas
func (p *CalendarPost) UpdatePublishedPlatforms(platforms []string) error {
	if err := valueobject.ValidatePlatforms(platforms); err != nil {
		return err
	}

	p.publishedPlatforms = platforms
	p.updatedAt = time.Now()
	return nil
}

// AddHistoryEvent adiciona um evento ao histórico
func (p *CalendarPost) AddHistoryEvent(event *valueobject.PostHistoryEvent) {
	p.history = append(p.history, event)
	p.updatedAt = time.Now()
}

// SoftDelete marca o post como deletado
func (p *CalendarPost) SoftDelete() {
	now := time.Now()
	p.deletedAt = &now
	p.updatedAt = now
}

// IsActive verifica se o post está ativo (não deletado)
func (p *CalendarPost) IsActive() bool {
	return p.deletedAt == nil
}

// IsPublished verifica se o post está publicado
func (p *CalendarPost) IsPublished() bool {
	return p.status == valueobject.StatusPublished
}
