package tasks

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/constants"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// UpdateCommentUseCase atualiza um comentário existente
type UpdateCommentUseCase struct {
	commentGateway      gateway.CommentGateway
	notificationGateway gateway.NotificationGateway
	userGateway         gateway.UserGateway
	taskGateway         gateway.TaskGateway
}

// NewUpdateCommentUseCase cria um novo UpdateCommentUseCase
func NewUpdateCommentUseCase(
	commentGateway gateway.CommentGateway,
	notificationGateway gateway.NotificationGateway,
	userGateway gateway.UserGateway,
	taskGateway gateway.TaskGateway,
) *UpdateCommentUseCase {
	return &UpdateCommentUseCase{
		commentGateway:      commentGateway,
		notificationGateway: notificationGateway,
		userGateway:         userGateway,
		taskGateway:         taskGateway,
	}
}

// Execute atualiza um comentário
func (uc *UpdateCommentUseCase) Execute(id, content string) error {
	// Parse ID
	commentID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid comment ID: %w", err)
	}

	// Find existing comment
	comment, err := uc.commentGateway.FindByID(commentID)
	if err != nil {
		return err
	}

	// Update comment content
	if err := comment.Validate(); err != nil {
		return err
	}

	// Update content via Reconstruct
	updatedComment := entity.ReconstructComment(
		comment.ID(),
		comment.TaskUUID(),
		comment.UserUUID(),
		content,
		time.Now(), // timestamp
	)

	// Save comment
	if err := uc.commentGateway.Update(updatedComment); err != nil {
		return err
	}

	// Create notifications for mentions in the updated comment
	uc.createMentionNotifications(context.Background(), updatedComment)

	return nil
}

// createMentionNotifications creates notifications for users mentioned in a comment
func (uc *UpdateCommentUseCase) createMentionNotifications(ctx context.Context, comment *entity.Comment) {
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
			_ = uc.notificationGateway.Create(notification)
		}
	}
}
