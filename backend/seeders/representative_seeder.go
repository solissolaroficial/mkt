package seeders

import (
	"log"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

// RepresentativeSeeder é responsável por criar representantes iniciais no sistema
type RepresentativeSeeder struct {
	db     *gorm.DB
	mapper *mapper.RepresentativeMapper
}

// NewRepresentativeSeeder cria uma nova instância do RepresentativeSeeder
func NewRepresentativeSeeder(db *gorm.DB) *RepresentativeSeeder {
	return &RepresentativeSeeder{
		db:     db,
		mapper: mapper.NewRepresentativeMapper(),
	}
}

// Seed cria representantes iniciais no banco de dados
func (s *RepresentativeSeeder) Seed() error {
	log.Println("🌱 Seeding representatives...")

	// Dados dos representantes a criar
	repsToCreate := []struct {
		code      int
		name      string
		email     string
		phone     string
		company   string
		region    string
		city      string
		attendant string
	}{
		{code: 101, name: "Jackson", email: "jackson@solis.com", phone: "11999999999", company: "Solis", region: "Sudeste", city: "São Paulo", attendant: ""},
		{code: 102, name: "Beatriz", email: "beatriz@solis.com", phone: "11999999998", company: "Solis", region: "Sudeste", city: "Rio de Janeiro", attendant: ""},
		{code: 103, name: "Larissa", email: "larissa@solis.com", phone: "11999999997", company: "Solis", region: "Sudeste", city: "Belo Horizonte", attendant: ""},
	}

	// Criar cada representante
	for _, repData := range repsToCreate {
		// Verificar se o representante já existe (usando model, não entity)
		var existingModel model.RepresentativeModel
		if err := s.db.Where("name = ?", repData.name).First(&existingModel).Error; err == nil {
			log.Printf("✅ Representative %s already exists, skipping...", repData.name)
			continue
		}

		// Criar código do representante
		repCode, err := valueobject.NewRepresentativeCode(repData.code)
		if err != nil {
			log.Printf("❌ Error creating representative code for %s: %v", repData.name, err)
			continue
		}

		// Criar representante
		rep, err := entity.NewRepresentative(
			repCode,
			repData.name,
			repData.email,
			repData.phone,
			repData.company,
			repData.region,
			repData.city,
			repData.attendant,
		)
		if err != nil {
			log.Printf("❌ Error creating representative %s: %v", repData.name, err)
			continue
		}

		// Converter entity para model
		model := s.mapper.EntityToModel(rep)

		// Salvar no banco
		if err := s.db.Create(model).Error; err != nil {
			log.Printf("❌ Error saving representative %s: %v", repData.name, err)
			continue
		}

		log.Printf("✅ Representative %s created successfully (Code: %d)", repData.name, repData.code)
	}

	log.Println("✅ Representatives seeded successfully:")
	log.Printf("   - Jackson")
	log.Printf("   - Beatriz")
	log.Printf("   - Larissa")

	return nil
}
