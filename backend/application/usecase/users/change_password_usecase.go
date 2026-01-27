package users

import (
	"context"
	stderrors "errors"
	"unicode"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/service"
)

// ChangePasswordInput contém os dados necessários para alterar a senha
type ChangePasswordInput struct {
	UserID          uuid.UUID
	CurrentPassword string
	NewPassword     string
}

// ChangePasswordUseCase é responsável por alterar a senha do usuário
type ChangePasswordUseCase struct {
	userGateway   gateway.UserGateway
	hasherService service.HasherService
}

// NewChangePasswordUseCase cria uma nova instância do ChangePasswordUseCase
func NewChangePasswordUseCase(
	userGateway gateway.UserGateway,
	hasherService service.HasherService,
) *ChangePasswordUseCase {
	return &ChangePasswordUseCase{
		userGateway:   userGateway,
		hasherService: hasherService,
	}
}

// Execute executa a alteração de senha
func (uc *ChangePasswordUseCase) Execute(ctx context.Context, input ChangePasswordInput) error {
	// 1. Buscar usuário pelo ID
	user, err := uc.userGateway.FindByID(ctx, input.UserID)
	if err != nil {
		return errors.ErrUserNotFound
	}

	// 2. Verificar se a senha atual está correta
	if err := uc.hasherService.Compare(user.Password(), input.CurrentPassword); err != nil {
		return errors.ErrCurrentPasswordMismatch
	}

	// 3. Validar que a nova senha é diferente da atual
	if input.CurrentPassword == input.NewPassword {
		return errors.ErrPasswordSameAsCurrent
	}

	// 4. Validar força da nova senha
	if err := validatePasswordStrength(input.NewPassword); err != nil {
		return errors.ErrPasswordTooWeak
	}

	// 5. Hash da nova senha
	hashedPassword, err := uc.hasherService.Hash(input.NewPassword)
	if err != nil {
		return err
	}

	// 6. Reconstruir usuário com nova senha
	updatedUser := entity.ReconstructUser(
		user.ID(),
		user.Email(),
		hashedPassword,
		user.Name(),
		user.Role(),
		user.IsActive(),
		user.ProfilePhotoKey(),
		user.CreatedAt(),
		user.UpdatedAt(),
	)

	// 7. Atualizar usuário no banco de dados
	if err := uc.userGateway.Update(ctx, updatedUser); err != nil {
		return err
	}

	return nil
}

// validatePasswordStrength valida se a senha atende aos requisitos mínimos
func validatePasswordStrength(password string) error {
	if len(password) < 8 {
		return stderrors.New("password must be at least 8 characters long")
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	// Requer pelo menos 3 dos 4 tipos de caracteres
	strengthCount := 0
	if hasUpper {
		strengthCount++
	}
	if hasLower {
		strengthCount++
	}
	if hasNumber {
		strengthCount++
	}
	if hasSpecial {
		strengthCount++
	}

	if strengthCount < 3 {
		return stderrors.New("password must contain at least 3 of: uppercase, lowercase, numbers, special characters")
	}

	return nil
}
