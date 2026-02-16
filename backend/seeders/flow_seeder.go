package seeders

import (
	"log"

	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

// FlowSeeder é responsável por criar flows iniciais no sistema
type FlowSeeder struct {
	db *gorm.DB
}

// NewFlowSeeder cria uma nova instância do FlowSeeder
func NewFlowSeeder(db *gorm.DB) *FlowSeeder {
	return &FlowSeeder{db: db}
}

// Seed cria flows iniciais no banco de dados
func (s *FlowSeeder) Seed() error {
	log.Println("🌱 Seeding flows...")

	// Verificar se já existem flows
	var existingFlows []model.Flow
	if err := s.db.Where("name IN ?", []string{"Beatriz", "Larissa", "Jackson Andrade"}).
		Find(&existingFlows).Error; err != nil {
		return err
	}

	if len(existingFlows) > 0 {
		log.Println("✅ Flows already exist, skipping...")
		return nil
	}

	// Criar flows iniciais
	flows := []model.Flow{
		{
			Name:        "Beatriz",
			Description: stringPtr("Fluxo da Beatriz"),
			Color:       stringPtr("#10b981"), // verde
			SortOrder:   0,
		},
		{
			Name:        "Larissa",
			Description: stringPtr("Fluxo da Larissa"),
			Color:       stringPtr("#3b82f6"), // azul
			SortOrder:   1,
		},
		{
			Name:        "Jackson Andrade",
			Description: stringPtr("Fluxo do Jackson Andrade"),
			Color:       stringPtr("#f59e0b"), // laranja
			SortOrder:   2,
		},
	}

	if err := s.db.Create(&flows).Error; err != nil {
		return err
	}

	log.Println("✅ Flows seeded successfully:")
	log.Printf("   - Beatriz (verde)")
	log.Printf("   - Larissa (azul)")
	log.Printf("   - Jackson Andrade (laranja)")

	return nil
}

// stringPtr retorna um ponteiro para uma string
func stringPtr(s string) *string {
	return &s
}
