package controller

import (
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/constants"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	taskrequest "github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	taskresponse "github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// TaskMapper converte entre requests/responses e entities
type TaskMapper struct{}

// NewTaskMapper cria um novo TaskMapper
func NewTaskMapper() *TaskMapper {
	return &TaskMapper{}
}

// ToTask converte CreateTaskRequest para entity.Task
func (m *TaskMapper) ToTask(req *taskrequest.CreateTaskRequest) (*entity.Task, error) {
	var dueDate time.Time
	if req.DueDate != nil {
		parsedTime, err := time.Parse(time.RFC3339, *req.DueDate)
		if err != nil {
			return nil, err
		}
		dueDate = parsedTime
	}

	var startDate *time.Time
	if req.StartDate != nil {
		parsedTime, err := time.Parse(time.RFC3339, *req.StartDate)
		if err != nil {
			return nil, err
		}
		startDate = &parsedTime
	}

	// NewTask doesn't accept assigneeID, assigneeID is set via SetAssigneeID
	task, err := entity.NewTask(
		req.Title,
		nil, // description (not in request)
		startDate,
		dueDate,
		req.Priority,
		req.Status,
		req.Category,
		0, // sortOrder (default 0)
	)
	if err != nil {
		return nil, err
	}

	// Set assigneeID after creation
	if req.AssigneeID != nil {
		parsedID, err := uuid.Parse(*req.AssigneeID)
		if err != nil {
			return nil, err
		}
		task.SetAssigneeUUID(&parsedID)
	}

	// Set flows after creation
	flows := make([]string, len(req.Flows))
	for i, flow := range req.Flows {
		flows[i] = string(flow)
	}
	task.SetFlows(flows)

	return task, nil
}

// ToTaskUpdate converte UpdateTaskRequest para entity.Task (via Reconstruct)
func (m *TaskMapper) ToTaskUpdate(id string, req *taskrequest.UpdateTaskRequest) (*entity.Task, error) {
	taskID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	var dueDate time.Time
	if req.DueDate != nil {
		parsedTime, err := time.Parse(time.RFC3339, *req.DueDate)
		if err != nil {
			return nil, err
		}
		dueDate = parsedTime
	}

	var startDate *time.Time
	if req.StartDate != nil {
		parsedTime, err := time.Parse(time.RFC3339, *req.StartDate)
		if err != nil {
			return nil, err
		}
		startDate = &parsedTime
	}

	var assigneeID *uuid.UUID
	if req.AssigneeID != nil {
		parsedID, err := uuid.Parse(*req.AssigneeID)
		if err != nil {
			return nil, err
		}
		assigneeID = &parsedID
	}

	// Convert flows to []string
	flows := make([]string, len(req.Flows))
	for i, flow := range req.Flows {
		flows[i] = string(flow)
	}

	// Convert description to *string
	var description *string
	if req.Description != "" {
		description = &req.Description
	}

	// Use default values if not provided
	var priority constants.TaskPriority
	if req.Priority != nil {
		priority = *req.Priority
	} else {
		priority = constants.TaskPriorityMedium
	}
	var status constants.TaskStatus
	if req.Status != nil {
		status = *req.Status
	} else {
		status = constants.TaskStatusPending
	}
	var category constants.TaskCategory
	if req.Category != nil {
		category = *req.Category
	} else {
		category = constants.TaskCategoryAdministrative
	}
	var archived bool
	if req.Archived != nil {
		archived = *req.Archived
	}

	return entity.ReconstructTask(
		taskID,
		req.Title,
		description,
		startDate,
		dueDate,
		priority,
		status,
		category,
		assigneeID,
		flows,
		archived,
		0,          // sortOrder (default 0)
		time.Now(), // updatedAt
		time.Now(), // createdAt
		nil,        // deletedAt
	), nil
}

// ToTaskResponse converte entity.Task para TaskResponse
func (m *TaskMapper) ToTaskResponse(task *entity.Task) taskresponse.TaskResponse {
	// Convert []string to []constants.TaskFlow
	flows := make([]constants.TaskFlow, len(task.Flows()))
	for i, flow := range task.Flows() {
		flows[i] = constants.TaskFlow(flow)
	}

	return taskresponse.TaskResponse{
		ID:            task.ID().String(),
		Title:         task.Title(),
		Description:   formatStringPtr(task.Description()),
		StartDate:     formatTimePtr(task.StartDate()),
		Category:      task.Category(),
		Priority:      task.Priority(),
		Status:        task.Status(),
		Flow:          "", // Task doesn't have Flow field, use Flows instead
		DueDate:       formatTimePtr(formatTimeToPtr(task.DueDate())),
		AssigneeID:    uuidPtrToString(task.AssigneeUUID()),
		Flows:         flows,
		Archived:      task.Archived(),
		Subtasks:      []taskresponse.SubtaskResponse{},
		Comments:      []taskresponse.CommentResponse{},
		Notifications: []taskresponse.NotificationResponse{},
		CreatedAt:     formatTime(task.CreatedAt()),
		UpdatedAt:     formatTime(task.UpdatedAt()),
	}
}

// ToTasksListResponse converte lista de entity.Task para TasksListResponse
func (m *TaskMapper) ToTasksListResponse(tasks []*entity.Task, total int64, page, limit int) *taskresponse.TasksListResponse {
	data := make([]taskresponse.TaskResponse, len(tasks))
	for i, task := range tasks {
		data[i] = m.ToTaskResponse(task)
	}

	// Prevenir divisão por zero
	var totalPages int
	if limit > 0 {
		totalPages = int(total) / limit
		if totalPages == 0 && total > 0 {
			totalPages = 1
		}
	} else {
		// Se limit for 0 ou negativo, usar total como totalPages
		if total > 0 {
			totalPages = 1
		} else {
			totalPages = 0
		}
	}

	return &taskresponse.TasksListResponse{
		Data:       data,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}
}

// ToSubtask converte CreateSubtaskRequest para entity.Subtask
func (m *TaskMapper) ToSubtask(req *taskrequest.CreateSubtaskRequest) (*entity.Subtask, error) {
	taskID, err := uuid.Parse(req.TaskID)
	if err != nil {
		return nil, err
	}

	return entity.NewSubtask(
		taskID,
		req.Title,
	)
}

// ToSubtaskUpdate converte UpdateSubtaskRequest para entity.Subtask (via Reconstruct)
func (m *TaskMapper) ToSubtaskUpdate(id string, req *taskrequest.UpdateSubtaskRequest) (*entity.Subtask, error) {
	subtaskID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	var assigneeID *uuid.UUID
	if req.AssigneeID != nil {
		parsedID, err := uuid.Parse(*req.AssigneeID)
		if err != nil {
			return nil, err
		}
		assigneeID = &parsedID
	}

	var dueDate *time.Time
	if req.DueDate != nil {
		parsedTime, err := time.Parse(time.RFC3339, *req.DueDate)
		if err != nil {
			return nil, err
		}
		dueDate = &parsedTime
	}

	var completed bool
	if req.Status != nil {
		statusStr := string(*req.Status)
		completed = statusStr == "completed"
	}

	return entity.ReconstructSubtask(
		subtaskID,
		uuid.Nil, // taskID (should be preserved)
		*req.Title,
		completed,
		assigneeID,
		dueDate,
		time.Now(), // updatedAt
		time.Now(), // createdAt
		nil,        // deletedAt
	), nil
}

// ToSubtaskResponse converte entity.Subtask para SubtaskResponse
func (m *TaskMapper) ToSubtaskResponse(subtask *entity.Subtask) taskresponse.SubtaskResponse {
	// Convert completed bool to TaskStatus
	var status constants.TaskStatus
	if subtask.Completed() {
		status = constants.TaskStatusCompleted
	} else {
		status = constants.TaskStatusPending
	}

	return taskresponse.SubtaskResponse{
		ID:          subtask.ID().String(),
		TaskID:      subtask.TaskUUID().String(),
		Title:       subtask.Title(),
		Description: "",                           // Subtask doesn't have Description field
		Priority:    constants.TaskPriorityMedium, // Default priority
		Status:      status,
		DueDate:     formatTimePtr(subtask.DueDate()),
		AssigneeID:  uuidPtrToString(subtask.AssigneeUUID()),
		CreatedAt:   formatTime(subtask.CreatedAt()),
		UpdatedAt:   formatTime(subtask.UpdatedAt()),
	}
}

// ToSubtasksListResponse converte lista de entity.Subtask para SubtasksListResponse
func (m *TaskMapper) ToSubtasksListResponse(subtasks []*entity.Subtask, total int64, page, limit int) *taskresponse.SubtasksListResponse {
	data := make([]taskresponse.SubtaskResponse, len(subtasks))
	for i, subtask := range subtasks {
		data[i] = m.ToSubtaskResponse(subtask)
	}

	// Prevenir divisão por zero
	var totalPages int
	if limit > 0 {
		totalPages = int(total) / limit
		if totalPages == 0 && total > 0 {
			totalPages = 1
		}
	} else {
		// Se limit for 0 ou negativo, usar total como totalPages
		if total > 0 {
			totalPages = 1
		} else {
			totalPages = 0
		}
	}

	return &taskresponse.SubtasksListResponse{
		Data:       data,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}
}

// ToComment converte CreateCommentRequest para entity.Comment
func (m *TaskMapper) ToComment(req *taskrequest.CreateCommentRequest) (*entity.Comment, error) {
	taskID, err := uuid.Parse(req.TaskID)
	if err != nil {
		return nil, err
	}

	// Usar Text se Content estiver vazio (compatibilidade com frontend)
	content := req.Content
	if content == "" && req.Text != "" {
		content = req.Text
	}

	// Parse userID da requisição
	userID := uuid.Nil
	if req.UserID != "" {
		parsedID, err := uuid.Parse(req.UserID)
		if err != nil {
			return nil, err
		}
		userID = parsedID
	}

	return entity.NewComment(
		taskID,
		userID,
		content,
	)
}

// ToCommentUpdate converte UpdateCommentRequest para entity.Comment (via Reconstruct)
func (m *TaskMapper) ToCommentUpdate(id string, req *taskrequest.UpdateCommentRequest) (*entity.Comment, error) {
	commentID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return entity.ReconstructComment(
		commentID,
		uuid.Nil, // taskID (should be preserved)
		uuid.Nil, // userID (should be preserved)
		req.Content,
		time.Now(), // timestamp
	), nil
}

// ToCommentResponse converte entity.Comment para CommentResponse
func (m *TaskMapper) ToCommentResponse(comment *entity.Comment) taskresponse.CommentResponse {
	return taskresponse.CommentResponse{
		ID:        comment.ID().String(),
		TaskID:    comment.TaskUUID().String(),
		UserID:    comment.UserUUID().String(),
		Text:      comment.Text(),
		Timestamp: formatTime(comment.Timestamp()),
	}
}

// ToCommentsListResponse converte lista de entity.Comment para CommentsListResponse
func (m *TaskMapper) ToCommentsListResponse(comments []*entity.Comment, total int64, page, limit int) *taskresponse.CommentsListResponse {
	data := make([]taskresponse.CommentResponse, len(comments))
	for i, comment := range comments {
		data[i] = m.ToCommentResponse(comment)
	}

	// Prevenir divisão por zero
	var totalPages int
	if limit > 0 {
		totalPages = int(total) / limit
		if totalPages == 0 && total > 0 {
			totalPages = 1
		}
	} else {
		// Se limit for 0 ou negativo, usar total como totalPages
		if total > 0 {
			totalPages = 1
		} else {
			totalPages = 0
		}
	}

	return &taskresponse.CommentsListResponse{
		Data:       data,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}
}

// ToNotification converte CreateNotificationRequest para entity.Notification
func (m *TaskMapper) ToNotification(req *taskrequest.CreateNotificationRequest) (*entity.Notification, error) {
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, err
	}

	var taskID *uuid.UUID
	if req.TaskID != "" {
		parsedID, err := uuid.Parse(req.TaskID)
		if err != nil {
			return nil, err
		}
		taskID = &parsedID
	}

	return entity.NewNotification(
		userID,
		taskID,
		constants.NotificationType(req.NotificationType),
		req.Title,
		req.Message,
	)
}

// ToNotificationUpdate converte UpdateNotificationRequest para entity.Notification (via Reconstruct)
func (m *TaskMapper) ToNotificationUpdate(id string, req *taskrequest.UpdateNotificationRequest) (*entity.Notification, error) {
	notificationID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return entity.ReconstructNotification(
		notificationID,
		uuid.Nil,                       // userID (should be preserved)
		nil,                            // taskID (should be preserved)
		constants.NotificationType(""), // notificationType (should be preserved)
		req.Title,
		req.Message,
		false,      // read (should be preserved)
		false,      // archived (should be preserved)
		time.Now(), // timestamp
	), nil
}

// ToNotificationResponse converte entity.Notification para NotificationResponse
func (m *TaskMapper) ToNotificationResponse(notification *entity.Notification) taskresponse.NotificationResponse {
	return taskresponse.NotificationResponse{
		ID:               notification.ID().String(),
		UserID:           notification.UserUUID().String(),
		TaskID:           uuidPtrToString(notification.TaskUUID()),
		NotificationType: notification.Type(), // Using Type() getter
		Title:            notification.Title(),
		Message:          notification.Message(),
		Read:             notification.Read(),
		Archived:         notification.Archived(),
		Timestamp:        formatTime(notification.Timestamp()),
	}
}

// ToNotificationsListResponse converte lista de entity.Notification para NotificationsListResponse
func (m *TaskMapper) ToNotificationsListResponse(notifications []*entity.Notification, total int64, page, limit int) *taskresponse.NotificationsListResponse {
	data := make([]taskresponse.NotificationResponse, len(notifications))
	for i, notification := range notifications {
		data[i] = m.ToNotificationResponse(notification)
	}

	// Prevenir divisão por zero
	var totalPages int
	if limit > 0 {
		totalPages = int(total) / limit
		if totalPages == 0 && total > 0 {
			totalPages = 1
		}
	} else {
		// Se limit for 0 ou negativo, usar total como totalPages
		if total > 0 {
			totalPages = 1
		} else {
			totalPages = 0
		}
	}

	return &taskresponse.NotificationsListResponse{
		Data:       data,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}
}

// Helper functions

// formatTime formata uma data para string RFC3339
func formatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// formatTimePtr formata um ponteiro de data para string RFC3339
func formatTimePtr(t *time.Time) *string {
	if t == nil {
		return nil
	}
	formatted := t.Format(time.RFC3339)
	return &formatted
}

// formatTimeToPtr converte time.Time para *time.Time
func formatTimeToPtr(t time.Time) *time.Time {
	return &t
}

// formatStringPtr formata um ponteiro de string para *string
func formatStringPtr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// uuidPtrToString converte *uuid.UUID para *string
func uuidPtrToString(id *uuid.UUID) *string {
	if id == nil {
		return nil
	}
	str := id.String()
	return &str
}
