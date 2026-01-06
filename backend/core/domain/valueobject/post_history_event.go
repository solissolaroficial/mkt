package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// PostHistoryEvent é um value object imutável que representa um evento no histórico
// Campos são públicos para permitir serialização JSON correta
type PostHistoryEvent struct {
	ID        string    `json:"id"`
	Action    string    `json:"action"`
	User      string    `json:"user"`
	Text      *string   `json:"text,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

// NewPostHistoryEvent cria um novo evento de histórico com timestamp atual
func NewPostHistoryEvent(action, user string, text *string) *PostHistoryEvent {
	return &PostHistoryEvent{
		ID:        uuid.New().String(),
		Action:    action,
		User:      user,
		Text:      text,
		Timestamp: time.Now(),
	}
}

// ReconstructPostHistoryEvent reconstrói um evento de histórico a partir de dados persistidos
func ReconstructPostHistoryEvent(id, action, user string, text *string, timestamp time.Time) *PostHistoryEvent {
	return &PostHistoryEvent{
		ID:        id,
		Action:    action,
		User:      user,
		Text:      text,
		Timestamp: timestamp,
	}
}

// FormattedTimestamp retorna o timestamp formatado para exibição
func (e *PostHistoryEvent) FormattedTimestamp() string {
	return e.Timestamp.Format("02/01/2006 15:04")
}
