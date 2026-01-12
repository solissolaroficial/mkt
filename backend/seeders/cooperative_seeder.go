package seeders

import (
	"context"
	"log"

	"solis/backend/core/domain/entity"
	"solis/backend/core/domain/gateway"
)

// CooperativeSeeder é responsável por criar dados iniciais do módulo cooperative
type CooperativeSeeder struct {
	offlineActionGateway      gateway.OfflineActionGateway
	showroomItemGateway       gateway.ShowroomItemGateway
	repMarketingActionGateway gateway.RepMarketingActionGateway
}

// NewCooperativeSeeder cria uma nova instância do CooperativeSeeder
func NewCooperativeSeeder(
	offlineActionGateway gateway.OfflineActionGateway,
	showroomItemGateway gateway.ShowroomItemGateway,
	repMarketingActionGateway gateway.RepMarketingActionGateway,
) *CooperativeSeeder {
	return &CooperativeSeeder{
		offlineActionGateway:      offlineActionGateway,
		showroomItemGateway:       showroomItemGateway,
		repMarketingActionGateway: repMarketingActionGateway,
	}
}

// Seed cria dados iniciais do módulo cooperative no banco de dados
func (s *CooperativeSeeder) Seed(ctx context.Context) error {
	log.Println("🌱 Seeding cooperative data...")

	// Seed Offline Actions
	if err := s.seedOfflineActions(ctx); err != nil {
		return err
	}

	// Seed Showroom Items
	if err := s.seedShowroomItems(ctx); err != nil {
		return err
	}

	// Seed Rep Marketing Actions
	if err := s.seedRepMarketingActions(ctx); err != nil {
		return err
	}

	return nil
}

// seedOfflineActions cria ações offline iniciais
func (s *CooperativeSeeder) seedOfflineActions(ctx context.Context) error {
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
		// Verificar se já existe uma ação com mesma PDV e data
		criteria := entity.NewOfflineActionCriteria().
			WithPDV(data.pdv).
			WithActionDate(data.actionDate)

		existingActions, _, err := s.offlineActionGateway.FindByCriteria(ctx, criteria, nil, nil)
		if err != nil {
			return err
		}

		if len(existingActions) > 0 {
			log.Printf("✅ Offline action for PDV %s on %s already exists, skipping...", data.pdv, data.actionDate)
			continue
		}

		// Criar ação offline
		action, err := entity.NewOfflineAction(
			data.requestedAmount,
			data.actionDate,
			data.category,
			data.pdv,
			data.repName,
			data.observation,
		)
		if err != nil {
			return err
		}

		// Salvar no banco
		if err := s.offlineActionGateway.Save(ctx, action); err != nil {
			return err
		}

		log.Printf("✅ Offline action created: %s - %s", data.pdv, data.actionDate)
	}

	return nil
}

// seedShowroomItems cria itens de showroom iniciais
func (s *CooperativeSeeder) seedShowroomItems(ctx context.Context) error {
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
		},
	}

	for _, data := range showroomItems {
		// Verificar se já existe um item com mesma PDV
		criteria := entity.NewShowroomItemCriteria().WithPDV(data.pdv)

		existingItems, _, err := s.showroomItemGateway.FindByCriteria(ctx, criteria, nil, nil)
		if err != nil {
			return err
		}

		if len(existingItems) > 0 {
			log.Printf("✅ Showroom item for PDV %s already exists, skipping...", data.pdv)
			continue
		}

		// Criar item de showroom
		item, err := entity.NewShowroomItem(
			data.pdv,
			data.city,
			data.contact,
			data.repName,
			data.deliveryForecast,
		)
		if err != nil {
			return err
		}

		// Atualizar data de workshop se fornecida
		if data.workshopDate != "" {
			if err := item.UpdateWorkshopDate(data.workshopDate); err != nil {
				return err
			}
		}

		// Salvar no banco
		if err := s.showroomItemGateway.Save(ctx, item); err != nil {
			return err
		}

		log.Printf("✅ Showroom item created: %s - %s", data.pdv, data.city)
	}

	return nil
}

// seedRepMarketingActions cria ações de marketing de representantes iniciais
func (s *CooperativeSeeder) seedRepMarketingActions(ctx context.Context) error {
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
		// Verificar se já existe uma ação com mesmo representante e data
		criteria := entity.NewRepMarketingActionCriteria().
			WithRepName(data.repName).
			WithDate(data.date)

		existingActions, _, err := s.repMarketingActionGateway.FindByCriteria(ctx, criteria, nil, nil)
		if err != nil {
			return err
		}

		if len(existingActions) > 0 {
			log.Printf("✅ Rep marketing action for %s on %s already exists, skipping...", data.repName, data.date)
			continue
		}

		// Criar ação de marketing de representante
		action, err := entity.NewRepMarketingAction(
			data.repName,
			data.date,
			data.description,
		)
		if err != nil {
			return err
		}

		// Salvar no banco
		if err := s.repMarketingActionGateway.Save(ctx, action); err != nil {
			return err
		}

		log.Printf("✅ Rep marketing action created: %s - %s", data.repName, data.date)
	}

	return nil
}
