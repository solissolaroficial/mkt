package seeders

import (
	"context"
	"log"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/service"
)

// UserSeeder é responsável por criar usuários iniciais no sistema
type UserSeeder struct {
	userGateway   gateway.UserGateway
	hasherService service.HasherService
}

// NewUserSeeder cria uma nova instância do UserSeeder
func NewUserSeeder(userGateway gateway.UserGateway, hasherService service.HasherService) *UserSeeder {
	return &UserSeeder{
		userGateway:   userGateway,
		hasherService: hasherService,
	}
}

// Seed cria usuários iniciais no banco de dados
func (s *UserSeeder) Seed(ctx context.Context) error {
	log.Println("🌱 Seeding users...")

	// Dados dos usuários a criar
	usersToCreate := []struct {
		email    string
		password string
		name     string
		role     string
	}{
		{"admin@gmail.com", "admin123", "Admin User", "admin"},
		{"jackson@solis.com", "jackson123", "Jackson", "marketing"},
		{"beatriz@solis.com", "beatriz123", "Beatriz", "commercial"},
		{"larissa@solis.com", "larissa123", "Larissa", "admin"},
	}

	// Criar cada usuário
	for _, userData := range usersToCreate {
		// Verificar se o usuário já existe
		exists, err := s.userGateway.ExistsByEmail(ctx, userData.email)
		if err != nil {
			return err
		}

		if exists {
			log.Printf("✅ User %s already exists, skipping...", userData.name)
			continue
		}

		// Hash da senha
		hashedPassword, err := s.hasherService.Hash(userData.password)
		if err != nil {
			return err
		}

		// Criar usuário
		user, err := entity.NewUser(
			userData.email,
			hashedPassword,
			userData.name,
			userData.role,
		)
		if err != nil {
			return err
		}

		// Salvar no banco
		if err := s.userGateway.Save(ctx, user); err != nil {
			return err
		}

		log.Printf("✅ User %s created successfully:", userData.name)
		log.Printf("   Email: %s", userData.email)
	}

	return nil
}
