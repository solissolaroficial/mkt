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

	// Verificar se o usuário admin já existe
	exists, err := s.userGateway.ExistsByEmail(ctx, "admin@gmail.com")
	if err != nil {
		return err
	}

	if exists {
		log.Println("✅ Admin user already exists, skipping...")
		return nil
	}

	// Hash da senha
	hashedPassword, err := s.hasherService.Hash("admin123")
	if err != nil {
		return err
	}

	// Criar usuário admin
	adminUser, err := entity.NewUser(
		"admin@gmail.com",
		hashedPassword,
		"Admin User",
		"admin",
	)
	if err != nil {
		return err
	}

	// Salvar no banco
	if err := s.userGateway.Save(ctx, adminUser); err != nil {
		return err
	}

	log.Println("✅ Admin user created successfully:")
	log.Println("   Email: admin@gmail.com")
	log.Println("   Password: admin123")

	return nil
}
