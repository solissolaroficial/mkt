package entity

import (
	"time"

	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"

	"github.com/google/uuid"
)

// Comment representa um comentário em uma tarefa
type Comment struct {
	id        uuid.UUID
	taskUUID  uuid.UUID
	userUUID  uuid.UUID
	text      string
	timestamp time.Time
}

// NewComment cria um novo comentário
func NewComment(
	taskUUID uuid.UUID,
	userUUID uuid.UUID,
	content string,
) (*Comment, error) {
	if content == "" {
		return nil, &domainerrors.CommentEmptyContentError{}
	}

	now := time.Now()
	return &Comment{
		id:        uuid.New(),
		taskUUID:  taskUUID,
		userUUID:  userUUID,
		text:      content,
		timestamp: now,
	}, nil
}

// ReconstructComment reconstrói um comentário a partir de dados persistidos
func ReconstructComment(
	id uuid.UUID,
	taskUUID uuid.UUID,
	userUUID uuid.UUID,
	text string,
	timestamp time.Time,
) *Comment {
	return &Comment{
		id:        id,
		taskUUID:  taskUUID,
		userUUID:  userUUID,
		text:      text,
		timestamp: timestamp,
	}
}

// Getters

func (c *Comment) ID() uuid.UUID {
	return c.id
}

func (c *Comment) TaskUUID() uuid.UUID {
	return c.taskUUID
}

func (c *Comment) UserUUID() uuid.UUID {
	return c.userUUID
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
