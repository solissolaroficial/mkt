package entity

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/constants"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"

	"github.com/google/uuid"
)

// Notification representa uma notificação do sistema
type Notification struct {
	id               uuid.UUID
	userUUID         uuid.UUID
	taskUUID         *uuid.UUID
	notificationType constants.NotificationType
	title            string
	message          string
	read             bool
	archived         bool
	timestamp        time.Time
}

// NewNotification cria uma nova notificação
func NewNotification(
	userUUID uuid.UUID,
	taskUUID *uuid.UUID,
	notificationType constants.NotificationType,
	title string,
	message string,
) (*Notification, error) {
	if !constants.IsValidNotificationType(notificationType) {
		return nil, &domainerrors.NotificationInvalidTypeError{Type: string(notificationType)}
	}

	now := time.Now()
	return &Notification{
		id:               uuid.New(),
		userUUID:         userUUID,
		taskUUID:         taskUUID,
		notificationType: notificationType,
		title:            title,
		message:          message,
		read:             false,
		archived:         false,
		timestamp:        now,
	}, nil
}

// ReconstructNotification reconstrói uma notificação a partir de dados persistidos
func ReconstructNotification(
	id uuid.UUID,
	userUUID uuid.UUID,
	taskUUID *uuid.UUID,
	notificationType constants.NotificationType,
	title string,
	message string,
	read bool,
	archived bool,
	timestamp time.Time,
) *Notification {
	return &Notification{
		id:               id,
		userUUID:         userUUID,
		taskUUID:         taskUUID,
		notificationType: notificationType,
		title:            title,
		message:          message,
		read:             read,
		archived:         archived,
		timestamp:        timestamp,
	}
}

// Getters

func (n *Notification) ID() uuid.UUID {
	return n.id
}

func (n *Notification) UserUUID() uuid.UUID {
	return n.userUUID
}

func (n *Notification) TaskUUID() *uuid.UUID {
	return n.taskUUID
}

func (n *Notification) Type() constants.NotificationType {
	return n.notificationType
}

func (n *Notification) Title() string {
	return n.title
}

func (n *Notification) Message() string {
	return n.message
}

func (n *Notification) Read() bool {
	return n.read
}

func (n *Notification) Archived() bool {
	return n.archived
}

func (n *Notification) Timestamp() time.Time {
	return n.timestamp
}

// Setters (business logic)

// MarkAsRead marca a notificação como lida
func (n *Notification) MarkAsRead() {
	n.read = true
}

// MarkAsUnread marca a notificação como não lida
func (n *Notification) MarkAsUnread() {
	n.read = false
}

// Archive arquiva a notificação
func (n *Notification) Archive() {
	n.archived = true
}

// Unarchive desarquiva a notificação
func (n *Notification) Unarchive() {
	n.archived = false
}

// Validate valida a notificação
func (n *Notification) Validate() error {
	if !constants.IsValidNotificationType(n.notificationType) {
		return &domainerrors.NotificationInvalidTypeError{Type: string(n.notificationType)}
	}
	return nil
}
