package seeders

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// CooperativeSeeder é responsável por criar dados iniciais do módulo cooperative
type CooperativeSeeder struct {
	offlineActionGateway      gateway.OfflineActionGateway
	showroomItemGateway       gateway.ShowroomItemGateway
	repMarketingActionGateway gateway.RepMarketingActionGateway
	representativeGateway     gateway.RepresentativeGateway
}

// NewCooperativeSeeder cria uma nova instância do CooperativeSeeder
func NewCooperativeSeeder(
	offlineActionGateway gateway.OfflineActionGateway,
	showroomItemGateway gateway.ShowroomItemGateway,
	repMarketingActionGateway gateway.RepMarketingActionGateway,
	representativeGateway gateway.RepresentativeGateway,
) *CooperativeSeeder {
	return &CooperativeSeeder{
		offlineActionGateway:      offlineActionGateway,
		showroomItemGateway:       showroomItemGateway,
		repMarketingActionGateway: repMarketingActionGateway,
		representativeGateway:     representativeGateway,
	}
}

// Seed cria dados iniciais do módulo cooperative no banco de dados
func (s *CooperativeSeeder) Seed(ctx context.Context) error {
	log.Println("🌱 Seeding cooperative data...")

	// Obter todos os representantes e criar um mapa de nome para UUID
	pagination := valueobject.NewPagination(1, 1000)
	sortOrder, err := valueobject.NewSortOrder("name", valueobject.SortDirectionAsc)
	if err != nil {
		return err
	}
	sortOrders := []*valueobject.SortOrder{sortOrder}
	representatives, _, err := s.representativeGateway.FindAll(&pagination, sortOrders)
	if err != nil {
		return err
	}

	// Criar mapa de nome para UUID
	repNameToUUID := make(map[string]uuid.UUID)
	for _, rep := range representatives {
		repNameToUUID[rep.Name()] = rep.UUID()
	}

	// Seed Offline Actions
	if err := s.seedOfflineActions(ctx, repNameToUUID); err != nil {
		return err
	}

	// Seed Showroom Items
	if err := s.seedShowroomItems(ctx, repNameToUUID); err != nil {
		return err
	}

	// Seed Rep Marketing Actions
	if err := s.seedRepMarketingActions(ctx, repNameToUUID); err != nil {
		return err
	}

	return nil
}

// seedOfflineActions cria ações offline iniciais
func (s *CooperativeSeeder) seedOfflineActions(ctx context.Context, repNameToUUID map[string]uuid.UUID) error {
	log.Println("🌱 Seeding offline actions...")

	offlineActions := []struct {
		requestedAmount float64
		actionDate      string
		category        string
		pdv             string
		repName         string
		observation     string
	}{
		{
			requestedAmount: 5000.00,
			actionDate:      "2024-01-15",
			category:        "PARCERIA",
			pdv:             "PDV Centro",
			repName:         "Jackson",
			observation:     "Parceria para promoção de verão",
		},
		{
			requestedAmount: 3000.00,
			actionDate:      "2024-02-20",
			category:        "AÇÃO COOPERADA",
			pdv:             "PDV Norte",
			repName:         "Beatriz",
			observation:     "Ação cooperada para lançamento de produto",
		},
		{
			requestedAmount: 7500.00,
			actionDate:      "2024-03-10",
			category:        "ENTREGA DE BRINDES EXCLUSIVOS",
			pdv:             "PDV Sul",
			repName:         "Larissa",
			observation:     "Entrega de brindes exclusivos para clientes VIP",
		},
	}

	for _, data := range offlineActions {
		// Criar ActionDate value object
		actionDate, err := valueobject.NewActionDate(data.actionDate)
		if err != nil {
			log.Printf("❌ Error creating ActionDate: %v", err)
			continue
		}

		// Criar OfflineCategory value object
		category := valueobject.OfflineCategory(data.category)

		// Obter UUID do representante
		repUUID, ok := repNameToUUID[data.repName]
		if !ok {
			log.Printf("❌ Representative not found: %s", data.repName)
			continue
		}

		// Criar ação offline
		action, err := entity.NewOfflineAction(
			data.requestedAmount,
			actionDate,
			category,
			data.pdv,
			repUUID,
			data.observation,
		)
		if err != nil {
			log.Printf("❌ Error creating offline action: %v", err)
			continue
		}

		// Salvar no banco
		if err := s.offlineActionGateway.Save(ctx, action); err != nil {
			log.Printf("❌ Error saving offline action: %v", err)
			continue
		}

		log.Printf("✅ Offline action created: %s - %s", data.pdv, data.actionDate)
	}

	return nil
}

// seedShowroomItems cria itens de showroom iniciais
func (s *CooperativeSeeder) seedShowroomItems(ctx context.Context, repNameToUUID map[string]uuid.UUID) error {
	log.Println("🌱 Seeding showroom items...")

	showroomItems := []struct {
		pdv              string
		city             string
		contact          string
		repName          string
		deliveryForecast string
		workshopDate     string
	}{
		{
			pdv:              "PDV Centro",
			city:             "São Paulo",
			contact:          "João Silva",
			repName:          "Jackson",
			deliveryForecast: "2024-02-01",
			workshopDate:     "2024-02-15",
		},
		{
			pdv:              "PDV Norte",
			city:             "Rio de Janeiro",
			contact:          "Maria Santos",
			repName:          "Beatriz",
			deliveryForecast: "2024-03-01",
			workshopDate:     "2024-03-10",
		},
		{
			pdv:              "PDV Sul",
			city:             "Belo Horizonte",
			contact:          "Pedro Costa",
			repName:          "Larissa",
			deliveryForecast: "2024-04-01",
			workshopDate:     "2024-04-10",
		},
	}

	for _, data := range showroomItems {
		// Obter UUID do representante
		repUUID, ok := repNameToUUID[data.repName]
		if !ok {
			log.Printf("❌ Representative not found: %s", data.repName)
			continue
		}

		// Criar item de showroom com apenas 2 parâmetros
		item, err := entity.NewShowroomItem(
			data.pdv,
			repUUID,
		)
		if err != nil {
			log.Printf("❌ Error creating showroom item: %v", err)
			continue
		}

		// Atualizar cidade
		if data.city != "" {
			if err := item.UpdateCity(&data.city); err != nil {
				log.Printf("❌ Error updating city: %v", err)
			}
		}

		// Atualizar contato
		if data.contact != "" {
			if err := item.UpdateContact(&data.contact); err != nil {
				log.Printf("❌ Error updating contact: %v", err)
			}
		}

		// Atualizar previsão de entrega
		if data.deliveryForecast != "" {
			if err := item.UpdateDeliveryForecast(&data.deliveryForecast); err != nil {
				log.Printf("❌ Error updating delivery forecast: %v", err)
			}
		}

		// Atualizar data de workshop se fornecida
		if data.workshopDate != "" {
			if err := item.UpdateWorkshopDate(&data.workshopDate); err != nil {
				log.Printf("❌ Error updating workshop date: %v", err)
			}
		}

		// Salvar no banco
		if err := s.showroomItemGateway.Save(ctx, item); err != nil {
			log.Printf("❌ Error saving showroom item: %v", err)
			continue
		}

		log.Printf("✅ Showroom item created: %s - %s", data.pdv, data.city)
	}

	return nil
}

// seedRepMarketingActions cria ações de marketing de representantes iniciais
func (s *CooperativeSeeder) seedRepMarketingActions(ctx context.Context, repNameToUUID map[string]uuid.UUID) error {
	log.Println("🌱 Seeding rep marketing actions...")

	repMarketingActions := []struct {
		repName     string
		date        string
		description string
	}{
		{
			repName:     "Jackson",
			date:        "2024-01-20",
			description: "Visita técnica aos PDVs da região Centro",
		},
		{
			repName:     "Beatriz",
			date:        "2024-02-15",
			description: "Apresentação de novos produtos para equipe de vendas",
		},
		{
			repName:     "Larissa",
			date:        "2024-03-10",
			description: "Treinamento de vendas para novos colaboradores",
		},
		{
			repName:     "Jackson",
			date:        "2024-04-05",
			description: "Participação em feira de negócios",
		},
	}

	for _, data := range repMarketingActions {
		// Parse date
		parsedDate, err := time.Parse("2006-01-02", data.date)
		if err != nil {
			log.Printf("❌ Error parsing date %s: %v", data.date, err)
			continue
		}

		// Obter UUID do representante
		repUUID, ok := repNameToUUID[data.repName]
		if !ok {
			log.Printf("❌ Representative not found: %s", data.repName)
			continue
		}

		// Criar ação de marketing de representante
		action, err := entity.NewRepMarketingAction(
			repUUID,
			parsedDate,
			data.description,
		)
		if err != nil {
			log.Printf("❌ Error creating rep marketing action: %v", err)
			continue
		}

		// Salvar no banco
		if err := s.repMarketingActionGateway.Save(ctx, action); err != nil {
			log.Printf("❌ Error saving rep marketing action: %v", err)
			continue
		}

		log.Printf("✅ Rep marketing action created: %s - %s", data.repName, data.date)
	}

	return nil
}
