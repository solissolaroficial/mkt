package seeders

import (
	"context"
	"log"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// BudgetSeeder é responsável por criar itens de orçamento iniciais no sistema
type BudgetSeeder struct {
	budgetGateway gateway.BudgetGateway
}

// NewBudgetSeeder cria uma nova instância do BudgetSeeder
func NewBudgetSeeder(budgetGateway gateway.BudgetGateway) *BudgetSeeder {
	return &BudgetSeeder{
		budgetGateway: budgetGateway,
	}
}

// Seed cria itens de orçamento iniciais no banco de dados
func (s *BudgetSeeder) Seed(ctx context.Context) error {
	log.Println("🌱 Seeding budget items...")

	// Verificar se já existem itens de orçamento
	count, err := s.budgetGateway.Count(ctx, nil)
	if err != nil {
		return err
	}

	if count > 0 {
		log.Printf("✅ Budget items already exist (%d items), skipping...", count)
		return nil
	}

	// Dados dos itens de orçamento a criar
	budgetItemsToCreate := s.getMockBudgetItems()

	// Criar cada item de orçamento
	for _, itemData := range budgetItemsToCreate {
		// Criar item de orçamento
		budgetItem, err := entity.NewBudgetItem(
			itemData.codObj,
			itemData.obj,
			itemData.codGrp,
			itemData.grp,
			itemData.cod,
			itemData.desc,
			itemData.vals[:],
			itemData.realizedVals[:],
			itemData.year,
		)
		if err != nil {
			log.Printf("❌ Failed to create budget item %s: %v", itemData.desc, err)
			continue
		}

		// Salvar no banco
		if err := s.budgetGateway.Create(ctx, budgetItem); err != nil {
			log.Printf("❌ Failed to save budget item %s: %v", itemData.desc, err)
			continue
		}

		log.Printf("✅ Budget item created: %s", itemData.desc)
	}

	log.Printf("✅ Budget seeding completed successfully")
	return nil
}

// getMockBudgetItems retorna os dados mockados dos itens de orçamento
func (s *BudgetSeeder) getMockBudgetItems() []struct {
	codObj       string
	obj          string
	codGrp       string
	grp          string
	cod          string
	desc         string
	vals         [12]float64
	realizedVals [12]float64
	year         int
} {
	return []struct {
		codObj       string
		obj          string
		codGrp       string
		grp          string
		cod          string
		desc         string
		vals         [12]float64
		realizedVals [12]float64
		year         int
	}{
		{
			codObj:       "1",
			obj:          "Receitas",
			codGrp:       "1.1",
			grp:          "Receitas Operacionais",
			cod:          "1.1.1",
			desc:         "Vendas de Produtos",
			vals:         [12]float64{100000, 110000, 120000, 130000, 140000, 150000, 160000, 170000, 180000, 190000, 200000, 210000},
			realizedVals: [12]float64{95000, 105000, 115000, 125000, 135000, 145000, 155000, 165000, 175000, 185000, 195000, 205000},
			year:         2025,
		},
		{
			codObj:       "1",
			obj:          "Receitas",
			codGrp:       "1.1",
			grp:          "Receitas Operacionais",
			cod:          "1.1.2",
			desc:         "Vendas de Serviços",
			vals:         [12]float64{50000, 55000, 60000, 65000, 70000, 75000, 80000, 85000, 90000, 95000, 100000, 105000},
			realizedVals: [12]float64{48000, 53000, 58000, 63000, 68000, 73000, 78000, 83000, 88000, 93000, 98000, 103000},
			year:         2025,
		},
		{
			codObj:       "2",
			obj:          "Despesas",
			codGrp:       "2.1",
			grp:          "Despesas Operacionais",
			cod:          "2.1.1",
			desc:         "Custo dos Produtos Vendidos",
			vals:         [12]float64{60000, 66000, 72000, 78000, 84000, 90000, 96000, 102000, 108000, 114000, 120000, 126000},
			realizedVals: [12]float64{58000, 64000, 70000, 76000, 82000, 88000, 94000, 100000, 106000, 112000, 118000, 124000},
			year:         2025,
		},
		{
			codObj:       "2",
			obj:          "Despesas",
			codGrp:       "2.1",
			grp:          "Despesas Operacionais",
			cod:          "2.1.2",
			desc:         "Custo dos Serviços Prestados",
			vals:         [12]float64{30000, 33000, 36000, 39000, 42000, 45000, 48000, 51000, 54000, 57000, 60000, 63000},
			realizedVals: [12]float64{29000, 32000, 35000, 38000, 41000, 44000, 47000, 50000, 53000, 56000, 59000, 62000},
			year:         2025,
		},
		{
			codObj:       "2",
			obj:          "Despesas",
			codGrp:       "2.2",
			grp:          "Despesas Administrativas",
			cod:          "2.2.1",
			desc:         "Salários e Encargos",
			vals:         [12]float64{40000, 40000, 40000, 40000, 40000, 40000, 40000, 40000, 40000, 40000, 40000, 40000},
			realizedVals: [12]float64{40000, 40000, 40000, 40000, 40000, 40000, 40000, 40000, 40000, 40000, 40000, 40000},
			year:         2025,
		},
		{
			codObj:       "2",
			obj:          "Despesas",
			codGrp:       "2.2",
			grp:          "Despesas Administrativas",
			cod:          "2.2.2",
			desc:         "Aluguel",
			vals:         [12]float64{15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000},
			realizedVals: [12]float64{15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000, 15000},
			year:         2025,
		},
		{
			codObj:       "2",
			obj:          "Despesas",
			codGrp:       "2.2",
			grp:          "Despesas Administrativas",
			cod:          "2.2.3",
			desc:         "Marketing e Publicidade",
			vals:         [12]float64{8000, 8000, 8000, 8000, 8000, 8000, 8000, 8000, 8000, 8000, 8000, 8000},
			realizedVals: [12]float64{7500, 7500, 7500, 7500, 7500, 7500, 7500, 7500, 7500, 7500, 7500, 7500},
			year:         2025,
		},
		{
			codObj:       "2",
			obj:          "Despesas",
			codGrp:       "2.2",
			grp:          "Despesas Administrativas",
			cod:          "2.2.4",
			desc:         "Utilidades (Água, Luz, Telefone)",
			vals:         [12]float64{3000, 3000, 3000, 3000, 3000, 3000, 3000, 3000, 3000, 3000, 3000, 3000},
			realizedVals: [12]float64{2800, 2800, 2800, 2800, 2800, 2800, 2800, 2800, 2800, 2800, 2800, 2800},
			year:         2025,
		},
		{
			codObj:       "2",
			obj:          "Despesas",
			codGrp:       "2.2",
			grp:          "Despesas Administrativas",
			cod:          "2.2.5",
			desc:         "Manutenção e Reparos",
			vals:         [12]float64{2000, 2000, 2000, 2000, 2000, 2000, 2000, 2000, 2000, 2000, 2000, 2000},
			realizedVals: [12]float64{1800, 1800, 1800, 1800, 1800, 1800, 1800, 1800, 1800, 1800, 1800, 1800},
			year:         2025,
		},
		{
			codObj:       "2",
			obj:          "Despesas",
			codGrp:       "2.3",
			grp:          "Despesas Financeiras",
			cod:          "2.3.1",
			desc:         "Juros Pagos",
			vals:         [12]float64{1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000},
			realizedVals: [12]float64{950, 950, 950, 950, 950, 950, 950, 950, 950, 950, 950, 950},
			year:         2025,
		},
		{
			codObj:       "3",
			obj:          "Investimentos",
			codGrp:       "3.1",
			grp:          "Investimentos em Ativos Fixos",
			cod:          "3.1.1",
			desc:         "Compra de Equipamentos",
			vals:         [12]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 50000},
			realizedVals: [12]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			year:         2025,
		},
		{
			codObj:       "3",
			obj:          "Investimentos",
			codGrp:       "3.1",
			grp:          "Investimentos em Ativos Fixos",
			cod:          "3.1.2",
			desc:         "Compra de Veículos",
			vals:         [12]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			realizedVals: [12]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			year:         2025,
		},
		{
			codObj:       "3",
			obj:          "Investimentos",
			codGrp:       "3.2",
			grp:          "Investimentos em Tecnologia",
			cod:          "3.2.1",
			desc:         "Software e Sistemas",
			vals:         [12]float64{5000, 5000, 5000, 5000, 5000, 5000, 5000, 5000, 5000, 5000, 5000, 5000},
			realizedVals: [12]float64{4800, 4800, 4800, 4800, 4800, 4800, 4800, 4800, 4800, 4800, 4800, 4800},
			year:         2025,
		},
		{
			codObj:       "3",
			obj:          "Investimentos",
			codGrp:       "3.2",
			grp:          "Investimentos em Tecnologia",
			cod:          "3.2.2",
			desc:         "Hardware e Equipamentos de TI",
			vals:         [12]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 30000},
			realizedVals: [12]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			year:         2025,
		},
		{
			codObj:       "4",
			obj:          "Impostos e Taxas",
			codGrp:       "4.1",
			grp:          "Impostos Federais",
			cod:          "4.1.1",
			desc:         "IRPJ",
			vals:         [12]float64{5000, 5500, 6000, 6500, 7000, 7500, 8000, 8500, 9000, 9500, 10000, 10500},
			realizedVals: [12]float64{4800, 5300, 5800, 6300, 6800, 7300, 7800, 8300, 8800, 9300, 9800, 10300},
			year:         2025,
		},
		{
			codObj:       "4",
			obj:          "Impostos e Taxas",
			codGrp:       "4.1",
			grp:          "Impostos Federais",
			cod:          "4.1.2",
			desc:         "CSLL",
			vals:         [12]float64{2500, 2750, 3000, 3250, 3500, 3750, 4000, 4250, 4500, 4750, 5000, 5250},
			realizedVals: [12]float64{2400, 2650, 2900, 3150, 3400, 3650, 3900, 4150, 4400, 4650, 4900, 5150},
			year:         2025,
		},
		{
			codObj:       "4",
			obj:          "Impostos e Taxas",
			codGrp:       "4.2",
			grp:          "Impostos Estaduais",
			cod:          "4.2.1",
			desc:         "ICMS",
			vals:         [12]float64{8000, 8800, 9600, 10400, 11200, 12000, 12800, 13600, 14400, 15200, 16000, 16800},
			realizedVals: [12]float64{7800, 8600, 9400, 10200, 11000, 11800, 12600, 13400, 14200, 15000, 15800, 16600},
			year:         2025,
		},
		{
			codObj:       "4",
			obj:          "Impostos e Taxas",
			codGrp:       "4.3",
			grp:          "Impostos Municipais",
			cod:          "4.3.1",
			desc:         "ISS",
			vals:         [12]float64{2000, 2200, 2400, 2600, 2800, 3000, 3200, 3400, 3600, 3800, 4000, 4200},
			realizedVals: [12]float64{1900, 2100, 2300, 2500, 2700, 2900, 3100, 3300, 3500, 3700, 3900, 4100},
			year:         2025,
		},
		{
			codObj:       "5",
			obj:          "Outros",
			codGrp:       "5.1",
			grp:          "Outras Receitas",
			cod:          "5.1.1",
			desc:         "Rendimento de Investimentos",
			vals:         [12]float64{1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000},
			realizedVals: [12]float64{950, 950, 950, 950, 950, 950, 950, 950, 950, 950, 950, 950},
			year:         2025,
		},
		{
			codObj:       "5",
			obj:          "Outros",
			codGrp:       "5.1",
			grp:          "Outras Receitas",
			cod:          "5.1.2",
			desc:         "Venda de Ativos",
			vals:         [12]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			realizedVals: [12]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			year:         2025,
		},
		{
			codObj:       "5",
			obj:          "Outros",
			codGrp:       "5.2",
			grp:          "Outras Despesas",
			cod:          "5.2.1",
			desc:         "Provisões",
			vals:         [12]float64{2000, 2000, 2000, 2000, 2000, 2000, 2000, 2000, 2000, 2000, 2000, 2000},
			realizedVals: [12]float64{1800, 1800, 1800, 1800, 1800, 1800, 1800, 1800, 1800, 1800, 1800, 1800},
			year:         2025,
		},
	}
}
