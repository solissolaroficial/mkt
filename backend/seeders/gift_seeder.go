package seeders

import (
	"context"
	"log"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GiftSeeder é responsável por criar dados iniciais de gifts no sistema
type GiftSeeder struct {
	giftItemGateway        gateway.GiftItemGateway
	giftTransactionGateway gateway.GiftTransactionGateway
}

// NewGiftSeeder cria uma nova instância do GiftSeeder
func NewGiftSeeder(giftItemGateway gateway.GiftItemGateway, giftTransactionGateway gateway.GiftTransactionGateway) *GiftSeeder {
	return &GiftSeeder{
		giftItemGateway:        giftItemGateway,
		giftTransactionGateway: giftTransactionGateway,
	}
}

// Seed cria dados iniciais de gifts no banco de dados
func (s *GiftSeeder) Seed(ctx context.Context) error {
	log.Println("🌱 Seeding gifts...")

	// Criar Gift Items iniciais
	giftItems := []struct {
		name  string
		stock int
		price float64
	}{
		{"Boné Solis", 50, 25.00},
		{"Caneta Premium", 100, 5.50},
		{"Caderno de Anotações", 30, 15.00},
		{"Garrafa Térmica", 15, 45.00},
		{"Camiseta Solis", 40, 80.00},
	}

	for _, itemData := range giftItems {
		// Verificar se o item já existe
		exists, err := s.giftItemGateway.ExistsByName(itemData.name)
		if err != nil {
			return err
		}

		if exists {
			log.Printf("✅ Gift item %s already exists, skipping...", itemData.name)
			continue
		}

		// Criar gift item
		giftItem, err := entity.NewGiftItem(itemData.name, itemData.stock, itemData.price)
		if err != nil {
			return err
		}

		if err := s.giftItemGateway.Create(giftItem); err != nil {
			return err
		}

		log.Printf("✅ Gift item %s created successfully", itemData.name)
	}

	// Criar Gift Transactions iniciais
	var price1 float64 = 25.00
	var price2 float64 = 5.50
	var price3 float64 = 15.00
	var price4 float64 = 45.00
	var price5 float64 = 80.00

	var andre string = "André"
	var maria string = "Maria"
	var joao string = "João"
	var ana string = "Ana"
	var pedro string = "Pedro"

	giftTransactions := []struct {
		itemName        string
		quantity        int
		transactionType string
		date            string
		time            string
		price           *float64
		representative  *string
		unit            string
	}{
		{"Boné Solis", 50, "in", "2026-01-01", "09:00", &price1, nil, "unid."},
		{"Boné Solis", 5, "out", "2026-01-05", "14:00", nil, &andre, "unid."},
		{"Caneta Premium", 100, "in", "2026-01-02", "10:30", &price2, nil, "unid."},
		{"Caneta Premium", 10, "out", "2026-01-06", "15:00", nil, &maria, "unid."},
		{"Caderno de Anotações", 30, "in", "2026-01-03", "11:00", &price3, nil, "unid."},
		{"Caderno de Anotações", 5, "out", "2026-01-07", "16:00", nil, &joao, "unid."},
		{"Garrafa Térmica", 15, "in", "2026-01-04", "08:00", &price4, nil, "unid."},
		{"Garrafa Térmica", 2, "out", "2026-01-08", "09:00", nil, &ana, "unid."},
		{"Camiseta Solis", 40, "in", "2026-01-09", "10:00", &price5, nil, "unid."},
		{"Camiseta Solis", 3, "out", "2026-01-10", "11:00", nil, &pedro, "unid."},
	}

	for _, txData := range giftTransactions {
		// Buscar o item para ajustar o estoque
		giftItem, err := s.giftItemGateway.FindByName(txData.itemName)
		if err != nil {
			return err
		}
		if giftItem == nil {
			log.Printf("⚠️ Gift item %s not found, skipping transaction", txData.itemName)
			continue
		}

		// Criar gift transaction
		giftTransaction, err := entity.NewGiftTransaction(
			txData.itemName,
			txData.quantity,
			txData.transactionType,
			txData.date,
			txData.time,
			txData.price,
			txData.representative,
			txData.unit,
		)
		if err != nil {
			return err
		}

		if err := s.giftTransactionGateway.Create(giftTransaction); err != nil {
			return err
		}

		// Ajustar o estoque do item
		stockAdjustment := txData.quantity
		if txData.transactionType == "out" {
			stockAdjustment = -txData.quantity
		}
		if err := s.giftItemGateway.AdjustStock(giftItem.ID(), stockAdjustment); err != nil {
			return err
		}

		log.Printf("✅ Gift transaction %s created successfully (stock adjusted for %s)", giftTransaction.ID(), txData.itemName)
	}

	log.Println("✅ Gifts seeded successfully")
	return nil
}
