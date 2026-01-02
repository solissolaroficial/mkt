package entity

import (
	"time"

	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"

	"github.com/google/uuid"
)

// Comment representa um comentário em uma tarefa
type Comment struct {
	id        uuid.UUID
	taskID    uuid.UUID
	userID    uuid.UUID
	text      string
	timestamp time.Time
}

// NewComment cria um novo comentário
func NewComment(
	taskID uuid.UUID,
	userID uuid.UUID,
	content string,
) (*Comment, error) {
	if content == "" {
		return nil, &domainerrors.CommentEmptyContentError{}
	}

	now := time.Now()
	return &Comment{
		id:        uuid.New(),
		taskID:    taskID,
		userID:    userID,
		text:      content,
		timestamp: now,
	}, nil
}

// ReconstructComment reconstrói um comentário a partir de dados persistidos
func ReconstructComment(
	id uuid.UUID,
	taskID uuid.UUID,
	userID uuid.UUID,
	text string,
	timestamp time.Time,
) *Comment {
	return &Comment{
		id:        id,
		taskID:    taskID,
		userID:    userID,
		text:      text,
		timestamp: timestamp,
	}
}

// Getters

func (c *Comment) ID() uuid.UUID {
	return c.id
}

func (c *Comment) TaskID() uuid.UUID {
	return c.taskID
}

func (c *Comment) UserID() uuid.UUID {
	return c.userID
}

func (c *Comment) Text() string {
	return c.text
}

func (c *Comment) Timestamp() time.Time {
	return c.timestamp
}

// Setters (business logic)

// UpdateContent atualiza o conteúdo do comentário
func (c *Comment) UpdateContent(text string) error {
	if text == "" {
		return &domainerrors.CommentEmptyContentError{}
	}
	c.text = text
	c.timestamp = time.Now()
	return nil
}

// Validate valida o comentário
func (c *Comment) Validate() error {
	if c.text == "" {
		return &domainerrors.CommentEmptyContentError{}
	}
	return nil
}
