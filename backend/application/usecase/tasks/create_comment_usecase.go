package tasks

import (
	"context"
	"fmt"
	"regexp"

	"github.com/seu-usuario/solis-backend/core/domain/constants"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// CreateCommentUseCase cria um novo comentário
type CreateCommentUseCase struct {
	commentGateway      gateway.CommentGateway
	notificationGateway gateway.NotificationGateway
	userGateway         gateway.UserGateway
	taskGateway         gateway.TaskGateway
}

// NewCreateCommentUseCase cria um novo CreateCommentUseCase
func NewCreateCommentUseCase(
	commentGateway gateway.CommentGateway,
	notificationGateway gateway.NotificationGateway,
	userGateway gateway.UserGateway,
	taskGateway gateway.TaskGateway,
) *CreateCommentUseCase {
	return &CreateCommentUseCase{
		commentGateway:      commentGateway,
		notificationGateway: notificationGateway,
		userGateway:         userGateway,
		taskGateway:         taskGateway,
	}
}

// Execute cria um novo comentário
func (uc *CreateCommentUseCase) Execute(ctx context.Context, comment *entity.Comment) error {
	if err := comment.Validate(); err != nil {
		return err
	}

	// Create the comment
	if err := uc.commentGateway.Create(comment); err != nil {
		return err
	}

	// Criar notificação para o assignee da tarefa (assíncrono)
	go func() {
		if err := uc.createTaskCommentNotification(ctx, comment); err != nil {
			fmt.Printf("Erro ao criar notificação: %v\n", err)
		}
	}()

	// Check for mentions and create notifications
	uc.createMentionNotifications(ctx, comment)

	return nil
}

// createMentionNotifications creates notifications for users mentioned in a comment
func (uc *CreateCommentUseCase) createMentionNotifications(ctx context.Context, comment *entity.Comment) {
	// Regex to find @username mentions - [a-zA-Z0-9_]+ captures only alphanumeric characters and underscores
	mentionRegex := regexp.MustCompile(`@([a-zA-Z0-9_]+)`)
	matches := mentionRegex.FindAllStringSubmatch(comment.Text(), -1)

	// Use a map to avoid duplicate notifications
	mentionedUsers := make(map[string]bool)

	for _, match := range matches {
		if len(match) > 1 {
			username := match[1]
			if mentionedUsers[username] {
				continue
			}
			mentionedUsers[username] = true

			// Find user by name
			user, err := uc.userGateway.FindByName(ctx, username)
			if err != nil {
				// User not found, skip this mention
				continue
			}

			// Create notification
			taskUUID := comment.TaskUUID()
			notification, err := entity.NewNotification(
				user.ID(),
				&taskUUID,
				constants.NotificationTypeMention,
				"Você foi mencionado em um comentário",
				comment.Text(),
			)
			if err != nil {
				// Failed to create notification, skip this mention
				continue
			}

			// Save notification
			if err := uc.notificationGateway.Create(notification); err != nil {
				fmt.Printf("Erro ao criar notificação de menção para usuário %s: %v\n", user.ID(), err)
			}
		}
	}
}

// createTaskCommentNotification cria uma notificação quando um comentário é adicionado a uma tarefa
func (uc *CreateCommentUseCase) createTaskCommentNotification(ctx context.Context, comment *entity.Comment) error {
	// Buscar a tarefa
	task, err := uc.taskGateway.FindByID(comment.TaskUUID())
	if err != nil {
		return err
	}

	// Verificar se a tarefa tem assignee
	if task.AssigneeUUID() == nil {
		return nil // Sem assignee, sem notificação
	}

	// Não notificar se o comentador é o mesmo do assignee
	if comment.UserUUID() == *task.AssigneeUUID() {
		return nil
	}

	taskID := task.ID()
	notification, err := entity.NewNotification(
		*task.AssigneeUUID(),
		&taskID,
		constants.NotificationTypeCommentAdded,
		"Novo comentário na sua tarefa",
		fmt.Sprintf("Alguém comentou na tarefa: %s", task.Title()),
	)
	if err != nil {
		return err
	}

	return uc.notificationGateway.Create(notification)
}
