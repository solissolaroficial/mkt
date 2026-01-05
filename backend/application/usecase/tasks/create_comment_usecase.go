package tasks

import (
	"context"
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
}

// NewCreateCommentUseCase cria um novo CreateCommentUseCase
func NewCreateCommentUseCase(
	commentGateway gateway.CommentGateway,
	notificationGateway gateway.NotificationGateway,
	userGateway gateway.UserGateway,
) *CreateCommentUseCase {
	return &CreateCommentUseCase{
		commentGateway:      commentGateway,
		notificationGateway: notificationGateway,
		userGateway:         userGateway,
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

	// Check for mentions and create notifications
	uc.createMentionNotifications(ctx, comment)

	return nil
}

// createMentionNotifications creates notifications for users mentioned in a comment
func (uc *CreateCommentUseCase) createMentionNotifications(ctx context.Context, comment *entity.Comment) {
	// Regex to find @username mentions
	mentionRegex := regexp.MustCompile(`@(\w+)`)
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
			taskID := comment.TaskID()
			notification, err := entity.NewNotification(
				user.ID(),
				&taskID,
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
