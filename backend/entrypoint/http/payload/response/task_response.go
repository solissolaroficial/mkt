package response

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/constants"
)

// TaskResponse representa a resposta de uma tarefa
type TaskResponse struct {
	ID            string                 `json:"id"`
	Title         string                 `json:"title"`
	Description   string                 `json:"description"`
	Category      constants.TaskCategory `json:"category"`
	Priority      constants.TaskPriority `json:"priority"`
	Status        constants.TaskStatus   `json:"status"`
	Flow          constants.TaskFlow     `json:"flow"`
	DueDate       *string                `json:"due_date"`
	AssigneeID    *string                `json:"assignee_id"`
	Flows         []constants.TaskFlow   `json:"flows"`
	Archived      bool                   `json:"archived"`
	Subtasks      []SubtaskResponse      `json:"subtasks,omitempty"`
	Comments      []CommentResponse      `json:"comments,omitempty"`
	Notifications []NotificationResponse `json:"notifications,omitempty"`
	CreatedAt     string                 `json:"created_at"`
	UpdatedAt     string                 `json:"updated_at"`
}

// SubtaskResponse representa a resposta de uma subtarefa
type SubtaskResponse struct {
	ID          string                 `json:"id"`
	TaskID      string                 `json:"task_id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Priority    constants.TaskPriority `json:"priority"`
	Status      constants.TaskStatus   `json:"status"`
	DueDate     *string                `json:"due_date"`
	AssigneeID  *string                `json:"assignee_id"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
}

// CommentResponse representa a resposta de um comentário
type CommentResponse struct {
	ID        string `json:"id"`
	TaskID    string `json:"task_id"`
	UserID    string `json:"user_id"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}

// NotificationResponse representa a resposta de uma notificação
type NotificationResponse struct {
	ID               string                     `json:"id"`
	UserID           string                     `json:"user_id"`
	TaskID           *string                    `json:"task_id"`
	NotificationType constants.NotificationType `json:"notification_type"`
	Title            string                     `json:"title"`
	Message          string                     `json:"message"`
	Read             bool                       `json:"read"`
	Archived         bool                       `json:"archived"`
	Timestamp        string                     `json:"timestamp"`
}

// TasksListResponse representa a resposta de lista de tarefas
type TasksListResponse struct {
	Data       []TaskResponse `json:"data"`
	Total      int64          `json:"total"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
	TotalPages int            `json:"total_pages"`
}

// SubtasksListResponse representa a resposta de lista de subtarefas
type SubtasksListResponse struct {
	Data       []SubtaskResponse `json:"data"`
	Total      int64             `json:"total"`
	Page       int               `json:"page"`
	Limit      int               `json:"limit"`
	TotalPages int               `json:"total_pages"`
}

// CommentsListResponse representa a resposta de lista de comentários
type CommentsListResponse struct {
	Data       []CommentResponse `json:"data"`
	Total      int64             `json:"total"`
	Page       int               `json:"page"`
	Limit      int               `json:"limit"`
	TotalPages int               `json:"total_pages"`
}

// NotificationsListResponse representa a resposta de lista de notificações
type NotificationsListResponse struct {
	Data       []NotificationResponse `json:"data"`
	Total      int64                  `json:"total"`
	Page       int                    `json:"page"`
	Limit      int                    `json:"limit"`
	TotalPages int                    `json:"total_pages"`
}

// formatTime format uma data para string RFC3339
func formatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// formatTimePtr format um ponteiro de data para string RFC3339
func formatTimePtr(t *time.Time) *string {
	if t == nil {
		return nil
	}
	formatted := t.Format(time.RFC3339)
	return &formatted
}
