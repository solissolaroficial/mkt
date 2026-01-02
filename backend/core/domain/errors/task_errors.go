package errors

import (
	"fmt"
)

// TaskNotFoundError é retornado quando uma tarefa não é encontrada
type TaskNotFoundError struct {
	ID string
}

func (e *TaskNotFoundError) Error() string {
	return fmt.Sprintf("task with ID %s not found", e.ID)
}

// TaskInvalidStatusError é retornado quando o status da tarefa é inválido
type TaskInvalidStatusError struct {
	Status string
}

func (e *TaskInvalidStatusError) Error() string {
	return fmt.Sprintf("invalid task status: %s", e.Status)
}

// TaskInvalidPriorityError é retornado quando a prioridade da tarefa é inválida
type TaskInvalidPriorityError struct {
	Priority string
}

func (e *TaskInvalidPriorityError) Error() string {
	return fmt.Sprintf("invalid task priority: %s", e.Priority)
}

// TaskInvalidCategoryError é retornado quando a categoria da tarefa é inválida
type TaskInvalidCategoryError struct {
	Category string
}

func (e *TaskInvalidCategoryError) Error() string {
	return fmt.Sprintf("invalid task category: %s", e.Category)
}

// TaskInvalidFlowError é retornado quando o fluxo da tarefa é inválido
type TaskInvalidFlowError struct {
	Flow string
}

func (e *TaskInvalidFlowError) Error() string {
	return fmt.Sprintf("invalid task flow: %s", e.Flow)
}

// TaskEmptyTitleError é retornado quando o título da tarefa está vazio
type TaskEmptyTitleError struct{}

func (e *TaskEmptyTitleError) Error() string {
	return "task title cannot be empty"
}

// TaskInvalidAssigneeError é retornado quando o assignee não é válido
type TaskInvalidAssigneeError struct {
	AssigneeID string
}

func (e *TaskInvalidAssigneeError) Error() string {
	return fmt.Sprintf("invalid assignee ID: %s", e.AssigneeID)
}

// SubtaskNotFoundError é retornado quando uma subtarefa não é encontrada
type SubtaskNotFoundError struct {
	ID string
}

func (e *SubtaskNotFoundError) Error() string {
	return fmt.Sprintf("subtask with ID %s not found", e.ID)
}

// SubtaskEmptyTitleError é retornado quando o título da subtarefa está vazio
type SubtaskEmptyTitleError struct{}

func (e *SubtaskEmptyTitleError) Error() string {
	return "subtask title cannot be empty"
}

// SubtaskTitleTooShortError é retornado quando o título da subtarefa é muito curto
type SubtaskTitleTooShortError struct{}

func (e *SubtaskTitleTooShortError) Error() string {
	return "subtask title must be at least 3 characters"
}

// SubtaskTitleTooLongError é retornado quando o título da subtarefa é muito longo
type SubtaskTitleTooLongError struct{}

func (e *SubtaskTitleTooLongError) Error() string {
	return "subtask title must be at most 200 characters"
}

// SubtaskInvalidStatusError é retornado quando o status da subtarefa é inválido
type SubtaskInvalidStatusError struct {
	Status string
}

func (e *SubtaskInvalidStatusError) Error() string {
	return fmt.Sprintf("invalid subtask status: %s", e.Status)
}

// SubtaskInvalidPriorityError é retornado quando a prioridade da subtarefa é inválida
type SubtaskInvalidPriorityError struct {
	Priority string
}

func (e *SubtaskInvalidPriorityError) Error() string {
	return fmt.Sprintf("invalid subtask priority: %s", e.Priority)
}

// SubtaskInvalidAssigneeError é retornado quando o assignee da subtarefa não é válido
type SubtaskInvalidAssigneeError struct {
	AssigneeID string
}

func (e *SubtaskInvalidAssigneeError) Error() string {
	return fmt.Sprintf("invalid subtask assignee ID: %s", e.AssigneeID)
}

// CommentNotFoundError é retornado quando um comentário não é encontrado
type CommentNotFoundError struct {
	ID string
}

func (e *CommentNotFoundError) Error() string {
	return fmt.Sprintf("comment with ID %s not found", e.ID)
}

// CommentEmptyContentError é retornado quando o conteúdo do comentário está vazio
type CommentEmptyContentError struct{}

func (e *CommentEmptyContentError) Error() string {
	return "comment content cannot be empty"
}

// NotificationNotFoundError é retornado quando uma notificação não é encontrada
type NotificationNotFoundError struct {
	ID string
}

func (e *NotificationNotFoundError) Error() string {
	return fmt.Sprintf("notification with ID %s not found", e.ID)
}

// NotificationInvalidTypeError é retornado quando o tipo da notificação é inválido
type NotificationInvalidTypeError struct {
	Type string
}

func (e *NotificationInvalidTypeError) Error() string {
	return fmt.Sprintf("invalid notification type: %s", e.Type)
}
