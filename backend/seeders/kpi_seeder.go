package seeders

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// KpiSeeder is responsible for creating initial KPI categories and their monthly data in the system
type KpiSeeder struct {
	kpiGateway         gateway.KpiGateway
	monthlyDataGateway gateway.MonthlyDataGateway
}

// NewKpiSeeder creates a new instance of KpiSeeder
func NewKpiSeeder(kpiGateway gateway.KpiGateway, monthlyDataGateway gateway.MonthlyDataGateway) *KpiSeeder {
	return &KpiSeeder{
		kpiGateway:         kpiGateway,
		monthlyDataGateway: monthlyDataGateway,
	}
}

// Seed creates initial KPI categories and their monthly data in the database
func (s *KpiSeeder) Seed(ctx context.Context) error {
	log.Println("🌱 Seeding KPI categories...")

	// Step 1: Create KPI categories
	kpiCategories, err := s.SeedKpiCategories(ctx)
	if err != nil {
		return err
	}

	// Step 2: Seed monthly data for each KPI
	if err := s.SeedMonthlyData(ctx, kpiCategories); err != nil {
		return err
	}

	return nil
}

// SeedKpiCategories creates KPI categories and returns a map of title -> KpiCategory
func (s *KpiSeeder) SeedKpiCategories(ctx context.Context) (map[string]*entity.KpiCategory, error) {
	log.Println("📊 Creating KPI categories...")

	// Define KPI categories based on legacy.constants.ts
	kpiDefinitions := []struct {
		title string
		slug  string
		color string
		unit  string
	}{
		{"Novos Inscritos no Youtube", "novos_inscritos_no_youtube", "#FF0000", "number"},
		{"Treinamentos Realizados (Reps)", "treinamentos_realizados_reps", "#3B82F6", "number"},
		{"Pessoas Treinadas (Solis)", "pessoas_treinadas_solis", "#10B981", "number"},
		{"Ações de Marketing (Reps)", "acoes_de_marketing_reps", "#8B5CF6", "number"},
		{"Visitantes Inbound", "visitantes_inbound", "#F59E0B", "number"},
		{"Taxa de Oportunidades", "taxa_de_oportunidades", "#EC4899", "number"},
		{"Autoridade na Internet (DA)", "autoridade_na_internet_da", "#6366F1", "number"},
		{"Investimento em Ads (R$)", "investimento_em_ads", "#85bb65", "currency"},
		{"Custo por Lead (CPL)", "custo_por_lead_cpl", "#eab308", "currency"},
		{"Leads Qualificados (MQLs)", "leads_qualificados_mqls", "#06b6d4", "number"},
		{"ROAS (Retorno em Ads)", "roas_retorno_em_ads", "#6366f1", "number"},
		{"Taxa de Conversão Global (%)", "taxa_de_conversao_global", "#f43f5e", "percent"},
	}

	kpiMap := make(map[string]*entity.KpiCategory)
	createdCount := 0
	for _, def := range kpiDefinitions {
		// Check if KPI category already exists by title
		existing, err := s.kpiGateway.FindByTitle(ctx, def.title)
		if err == nil && existing != nil {
			log.Printf("✅ KPI category '%s' already exists, using existing...", def.title)
			kpiMap[def.title] = existing
			continue
		}

		// Create new KPI category with explicit slug
		kpi, err := entity.NewKpiCategoryWithSlug(def.title, def.color, def.unit, def.slug)
		if err != nil {
			log.Printf("❌ Error creating KPI category '%s': %v", def.title, err)
			continue
		}

		// Save to database
		if err := s.kpiGateway.Save(ctx, kpi); err != nil {
			log.Printf("❌ Error saving KPI category '%s': %v", def.title, err)
			continue
		}

		log.Printf("✅ Created KPI category: %s", def.title)
		kpiMap[def.title] = kpi
		createdCount++
	}

	log.Printf("🎉 KPI categories seeding completed. Created %d categories.", createdCount)
	return kpiMap, nil
}

// SeedMonthlyData creates monthly data for each KPI category
func (s *KpiSeeder) SeedMonthlyData(ctx context.Context, kpiCategories map[string]*entity.KpiCategory) error {
	log.Println("📅 Creating MonthlyData for KPIs...")

	// Get current year for seeding
	currentYear := time.Now().Year()

	// Get legacy KPI data from frontend constants
	legacyData := s.getLegacyKpiData()

	createdCount := 0
	for _, legacy := range legacyData {
		kpi, exists := kpiCategories[legacy.title]
		if !exists {
			log.Printf("⚠️  KPI '%s' not found in categories, skipping...", legacy.title)
			continue
		}

		// Check if monthly data already exists for this KPI
		existingData, err := s.monthlyDataGateway.FindAllByKpi(ctx, kpi.ID())
		if err == nil && len(existingData) > 0 {
			log.Printf("✅ MonthlyData for KPI '%s' already exists, skipping...", legacy.title)
			continue
		}

		// Create monthly data for each month
		for _, monthData := range legacy.monthlyData {
			// Create MonthlyData
			monthlyData, err := entity.NewMonthlyData(kpi.ID(), currentYear, monthData.month)
			if err != nil {
				log.Printf("❌ Error creating MonthlyData for KPI '%s', month '%s': %v", legacy.title, monthData.month, err)
				continue
			}

			// Set realized value if not nil
			if monthData.realized != nil {
				monthlyData.SetRealized(*monthData.realized)
			}

			// Set meta value if not nil
			if monthData.meta != nil {
				monthlyData.SetMeta(*monthData.meta)
			}

			// Set breakdown if exists
			if monthData.breakdown != nil {
				if err := monthlyData.SetBreakdown(monthData.breakdown); err != nil {
					log.Printf("❌ Error setting breakdown for KPI '%s', month '%s': %v", legacy.title, monthData.month, err)
				}
			}

			// Add sample log entry
			if monthData.realized != nil {
				logEntry := entity.KpiLogEntry{
					ID:        uuid.New().String(),
					Date:      time.Now().Format(time.RFC3339),
					Timestamp: "10:30",
					User:      "Jackson",
					Month:     monthData.month,
					OldValue:  nil,
					NewValue:  *monthData.realized,
					Action:    "create",
					Context:   "Lote 01",
				}
				if err := monthlyData.AddLog(logEntry); err != nil {
					log.Printf("❌ Error adding log for KPI '%s', month '%s': %v", legacy.title, monthData.month, err)
				}
			}

			// Save to database
			if err := s.monthlyDataGateway.Save(ctx, monthlyData); err != nil {
				log.Printf("❌ Error saving MonthlyData for KPI '%s', month '%s': %v", legacy.title, monthData.month, err)
				continue
			}

			createdCount++
		}

		log.Printf("✅ Created MonthlyData for KPI: %s", legacy.title)
	}

	log.Printf("🎉 MonthlyData seeding completed. Created %d records.", createdCount)
	return nil
}

// getLegacyKpiData returns KPI data from frontend legacy constants
func (s *KpiSeeder) getLegacyKpiData() []struct {
	title       string
	monthlyData []struct {
		month     string
		realized  *float64
		meta      *float64
		breakdown interface{}
	}
} {
	return []struct {
		title       string
		monthlyData []struct {
			month     string
			realized  *float64
			meta      *float64
			breakdown interface{}
		}
	}{
		{
			title: "Novos Inscritos no Youtube",
			monthlyData: []struct {
				month     string
				realized  *float64
				meta      *float64
				breakdown interface{}
			}{
				{month: "JAN", realized: floatPtr(14), meta: floatPtr(15), breakdown: nil},
				{month: "FEV", realized: floatPtr(66), meta: floatPtr(25), breakdown: nil},
				{month: "MAR", realized: floatPtr(270), meta: floatPtr(25), breakdown: nil},
				{month: "ABR", realized: floatPtr(81), meta: floatPtr(50), breakdown: nil},
				{month: "MAI", realized: floatPtr(57), meta: floatPtr(75), breakdown: nil},
				{month: "JUN", realized: floatPtr(45), meta: floatPtr(75), breakdown: nil},
				{month: "JUL", realized: floatPtr(55), meta: floatPtr(100), breakdown: nil},
				{month: "AGO", realized: floatPtr(66), meta: floatPtr(120), breakdown: nil},
				{month: "SET", realized: floatPtr(33), meta: floatPtr(120), breakdown: nil},
				{month: "OUT", realized: floatPtr(66), meta: floatPtr(115), breakdown: nil},
				{month: "NOV", realized: floatPtr(45), meta: floatPtr(75), breakdown: nil},
				{month: "DEZ", realized: nil, meta: floatPtr(45), breakdown: nil},
			},
		},
		{
			title: "Treinamentos Realizados (Reps)",
			monthlyData: []struct {
				month     string
				realized  *float64
				meta      *float64
				breakdown interface{}
			}{
				{month: "JAN", realized: floatPtr(0), meta: floatPtr(0), breakdown: nil},
				{month: "FEV", realized: floatPtr(3), meta: floatPtr(15), breakdown: nil},
				{month: "MAR", realized: floatPtr(5), meta: floatPtr(30), breakdown: nil},
				{month: "ABR", realized: floatPtr(10), meta: floatPtr(30), breakdown: nil},
				{month: "MAI", realized: floatPtr(5), meta: floatPtr(30), breakdown: nil},
				{month: "JUN", realized: floatPtr(12), meta: floatPtr(30), breakdown: nil},
				{month: "JUL", realized: floatPtr(23), meta: floatPtr(30), breakdown: nil},
				{month: "AGO", realized: floatPtr(22), meta: floatPtr(30), breakdown: nil},
				{month: "SET", realized: floatPtr(22), meta: floatPtr(30), breakdown: nil},
				{month: "OUT", realized: floatPtr(21), meta: floatPtr(30), breakdown: nil},
				{month: "NOV", realized: floatPtr(11), meta: floatPtr(15), breakdown: nil},
				{month: "DEZ", realized: nil, meta: floatPtr(15), breakdown: nil},
			},
		},
		{
			title: "Pessoas Treinadas (Solis)",
			monthlyData: []struct {
				month     string
				realized  *float64
				meta      *float64
				breakdown interface{}
			}{
				{month: "JAN", realized: floatPtr(5), meta: floatPtr(0), breakdown: nil},
				{month: "FEV", realized: floatPtr(44), meta: floatPtr(30), breakdown: nil},
				{month: "MAR", realized: floatPtr(38), meta: floatPtr(40), breakdown: nil},
				{month: "ABR", realized: floatPtr(0), meta: floatPtr(20), breakdown: nil},
				{month: "MAI", realized: floatPtr(77), meta: floatPtr(55), breakdown: nil},
				{month: "JUN", realized: floatPtr(43), meta: floatPtr(65), breakdown: nil},
				{month: "JUL", realized: floatPtr(97), meta: floatPtr(65), breakdown: nil},
				{month: "AGO", realized: floatPtr(85), meta: floatPtr(55), breakdown: nil},
				{month: "SET", realized: floatPtr(24), meta: floatPtr(20), breakdown: nil},
				{month: "OUT", realized: floatPtr(58), meta: floatPtr(50), breakdown: nil},
				{month: "NOV", realized: floatPtr(56), meta: floatPtr(40), breakdown: nil},
				{month: "DEZ", realized: nil, meta: floatPtr(30), breakdown: nil},
			},
		},
		{
			title: "Ações de Marketing (Reps)",
			monthlyData: []struct {
				month     string
				realized  *float64
				meta      *float64
				breakdown interface{}
			}{
				{month: "JAN", realized: floatPtr(1), meta: floatPtr(0), breakdown: nil},
				{month: "FEV", realized: floatPtr(10), meta: floatPtr(30), breakdown: nil},
				{month: "MAR", realized: floatPtr(13), meta: floatPtr(60), breakdown: nil},
				{month: "ABR", realized: floatPtr(82), meta: floatPtr(60), breakdown: nil},
				{month: "MAI", realized: floatPtr(18), meta: floatPtr(60), breakdown: nil},
				{month: "JUN", realized: floatPtr(33), meta: floatPtr(60), breakdown: nil},
				{month: "JUL", realized: floatPtr(48), meta: floatPtr(60), breakdown: nil},
				{month: "AGO", realized: floatPtr(64), meta: floatPtr(60), breakdown: nil},
				{month: "SET", realized: floatPtr(62), meta: floatPtr(60), breakdown: nil},
				{month: "OUT", realized: floatPtr(48), meta: floatPtr(60), breakdown: nil},
				{month: "NOV", realized: floatPtr(35), meta: floatPtr(30), breakdown: nil},
				{month: "DEZ", realized: nil, meta: floatPtr(30), breakdown: nil},
			},
		},
		{
			title: "Visitantes Inbound",
			monthlyData: []struct {
				month     string
				realized  *float64
				meta      *float64
				breakdown interface{}
			}{
				{month: "JAN", realized: floatPtr(3417), meta: floatPtr(12500), breakdown: nil},
				{month: "FEV", realized: floatPtr(3266), meta: floatPtr(14000), breakdown: nil},
				{month: "MAR", realized: floatPtr(3437), meta: floatPtr(19000), breakdown: nil},
				{month: "ABR", realized: floatPtr(3396), meta: floatPtr(19000), breakdown: nil},
				{month: "MAI", realized: floatPtr(3604), meta: floatPtr(18000), breakdown: nil},
				{month: "JUN", realized: floatPtr(3838), meta: floatPtr(18000), breakdown: nil},
				{month: "JUL", realized: floatPtr(4590), meta: floatPtr(18000), breakdown: nil},
				{month: "AGO", realized: floatPtr(4650), meta: floatPtr(18000), breakdown: nil},
				{month: "SET", realized: floatPtr(3968), meta: floatPtr(18000), breakdown: nil},
				{month: "OUT", realized: floatPtr(4473), meta: floatPtr(18000), breakdown: nil},
				{month: "NOV", realized: floatPtr(3201), meta: floatPtr(15000), breakdown: nil},
				{month: "DEZ", realized: nil, meta: floatPtr(12500), breakdown: nil},
			},
		},
		{
			title: "Taxa de Oportunidades",
			monthlyData: []struct {
				month     string
				realized  *float64
				meta      *float64
				breakdown interface{}
			}{
				{month: "JAN", realized: floatPtr(147), meta: floatPtr(100), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 66.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 44.0},
						{"label": "RJ", "value": 10.0},
						{"label": "Sul de MG", "value": 5.0},
						{"label": "Espírito Santo", "value": 7.0},
					},
				}}},
				{month: "FEV", realized: floatPtr(69), meta: floatPtr(100), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 31.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 21.0},
						{"label": "RJ", "value": 5.0},
						{"label": "Sul de MG", "value": 2.0},
						{"label": "Espírito Santo", "value": 3.0},
					},
				}}},
				{month: "MAR", realized: floatPtr(270), meta: floatPtr(170), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 121.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 81.0},
						{"label": "RJ", "value": 18.0},
						{"label": "Sul de MG", "value": 9.0},
						{"label": "Espírito Santo", "value": 13.0},
					},
				}}},
				{month: "ABR", realized: floatPtr(778), meta: floatPtr(200), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 350.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 233.0},
						{"label": "RJ", "value": 52.0},
						{"label": "Sul de MG", "value": 26.0},
						{"label": "Espírito Santo", "value": 39.0},
					},
				}}},
				{month: "MAI", realized: floatPtr(160), meta: floatPtr(170), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 72.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 48.0},
						{"label": "RJ", "value": 11.0},
						{"label": "Sul de MG", "value": 5.0},
						{"label": "Espírito Santo", "value": 8.0},
					},
				}}},
				{month: "JUN", realized: floatPtr(152), meta: floatPtr(250), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 68.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 45.0},
						{"label": "RJ", "value": 10.0},
						{"label": "Sul de MG", "value": 5.0},
						{"label": "Espírito Santo", "value": 8.0},
					},
				}}},
				{month: "JUL", realized: floatPtr(224), meta: floatPtr(250), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 101.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 67.0},
						{"label": "RJ", "value": 15.0},
						{"label": "Sul de MG", "value": 8.0},
						{"label": "Espírito Santo", "value": 11.0},
					},
				}}},
				{month: "AGO", realized: floatPtr(278), meta: floatPtr(200), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 125.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 83.0},
						{"label": "RJ", "value": 19.0},
						{"label": "Sul de MG", "value": 9.0},
						{"label": "Espírito Santo", "value": 14.0},
					},
				}}},
				{month: "SET", realized: floatPtr(324), meta: floatPtr(180), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 146.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 97.0},
						{"label": "RJ", "value": 22.0},
						{"label": "Sul de MG", "value": 11.0},
						{"label": "Espírito Santo", "value": 16.0},
					},
				}}},
				{month: "OUT", realized: floatPtr(392), meta: floatPtr(300), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 176.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 117.0},
						{"label": "RJ", "value": 26.0},
						{"label": "Sul de MG", "value": 13.0},
						{"label": "Espírito Santo", "value": 20.0},
					},
				}}},
				{month: "NOV", realized: floatPtr(344), meta: floatPtr(200), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 178.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 118.0},
						{"label": "RJ", "value": 26.0},
						{"label": "Sul de MG", "value": 11.0},
						{"label": "Espírito Santo", "value": 23.0},
					},
				}}},
				{month: "DEZ", realized: nil, meta: floatPtr(80), breakdown: nil},
			},
		},
		{
			title: "Autoridade na Internet (DA)",
			monthlyData: []struct {
				month     string
				realized  *float64
				meta      *float64
				breakdown interface{}
			}{
				{month: "JAN", realized: floatPtr(30), meta: floatPtr(30), breakdown: nil},
				{month: "FEV", realized: floatPtr(30), meta: floatPtr(30), breakdown: nil},
				{month: "MAR", realized: floatPtr(30), meta: floatPtr(30), breakdown: nil},
				{month: "ABR", realized: floatPtr(30), meta: floatPtr(30), breakdown: nil},
				{month: "MAI", realized: floatPtr(33), meta: floatPtr(31), breakdown: nil},
				{month: "JUN", realized: floatPtr(33), meta: floatPtr(31), breakdown: nil},
				{month: "JUL", realized: floatPtr(33), meta: floatPtr(33), breakdown: nil},
				{month: "AGO", realized: floatPtr(35), meta: floatPtr(33), breakdown: nil},
				{month: "SET", realized: floatPtr(37), meta: floatPtr(33), breakdown: nil},
				{month: "OUT", realized: floatPtr(38), meta: floatPtr(33), breakdown: nil},
				{month: "NOV", realized: floatPtr(38), meta: floatPtr(33), breakdown: nil},
				{month: "DEZ", realized: floatPtr(38), meta: floatPtr(33), breakdown: nil},
			},
		},
		{
			title: "Investimento em Ads (R$)",
			monthlyData: []struct {
				month     string
				realized  *float64
				meta      *float64
				breakdown interface{}
			}{
				{month: "JAN", realized: floatPtr(12500), meta: floatPtr(12000), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 7500.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 4500.0},
						{"label": "RJ", "value": 1125.0},
						{"label": "Sul de MG", "value": 750.0},
						{"label": "Espírito Santo", "value": 1125.0},
					},
				}}},
				{month: "FEV", realized: floatPtr(13200), meta: floatPtr(13000), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 7920.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 4752.0},
						{"label": "RJ", "value": 1188.0},
						{"label": "Sul de MG", "value": 792.0},
						{"label": "Espírito Santo", "value": 1188.0},
					},
				}}},
				{month: "MAR", realized: floatPtr(15800), meta: floatPtr(15000), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 9480.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 5688.0},
						{"label": "RJ", "value": 1422.0},
						{"label": "Sul de MG", "value": 948.0},
						{"label": "Espírito Santo", "value": 1422.0},
					},
				}}},
				{month: "ABR", realized: floatPtr(14900), meta: floatPtr(15000), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 8940.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 5364.0},
						{"label": "RJ", "value": 1341.0},
						{"label": "Sul de MG", "value": 894.0},
						{"label": "Espírito Santo", "value": 1341.0},
					},
				}}},
				{month: "MAI", realized: floatPtr(18500), meta: floatPtr(18000), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 11100.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 6660.0},
						{"label": "RJ", "value": 1665.0},
						{"label": "Sul de MG", "value": 1110.0},
						{"label": "Espírito Santo", "value": 1665.0},
					},
				}}},
				{month: "JUN", realized: floatPtr(19200), meta: floatPtr(20000), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 11520.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 6912.0},
						{"label": "RJ", "value": 1728.0},
						{"label": "Sul de MG", "value": 1152.0},
						{"label": "Espírito Santo", "value": 1728.0},
					},
				}}},
				{month: "JUL", realized: floatPtr(21500), meta: floatPtr(20000), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 12900.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 7740.0},
						{"label": "RJ", "value": 1935.0},
						{"label": "Sul de MG", "value": 1290.0},
						{"label": "Espírito Santo", "value": 1935.0},
					},
				}}},
				{month: "AGO", realized: floatPtr(22000), meta: floatPtr(20000), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 13200.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 7920.0},
						{"label": "RJ", "value": 1980.0},
						{"label": "Sul de MG", "value": 1320.0},
						{"label": "Espírito Santo", "value": 1980.0},
					},
				}}},
				{month: "SET", realized: floatPtr(19800), meta: floatPtr(20000), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 11880.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 7128.0},
						{"label": "RJ", "value": 1782.0},
						{"label": "Sul de MG", "value": 1188.0},
						{"label": "Espírito Santo", "value": 1782.0},
					},
				}}},
				{month: "OUT", realized: floatPtr(23500), meta: floatPtr(22000), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 14100.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 8460.0},
						{"label": "RJ", "value": 2115.0},
						{"label": "Sul de MG", "value": 1410.0},
						{"label": "Espírito Santo", "value": 2115.0},
					},
				}}},
				{month: "NOV", realized: floatPtr(18000), meta: floatPtr(20000), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 12500.0,
					"subItems": []map[string]interface{}{
						{"label": "Brasil", "value": 7000.0},
						{"label": "RJ", "value": 2500.0},
						{"label": "Sul de MG", "value": 1000.0},
						{"label": "Espírito Santo", "value": 2000.0},
					},
				}}},
				{month: "DEZ", realized: nil, meta: floatPtr(15000), breakdown: nil},
			},
		},
		{
			title: "Custo por Lead (CPL)",
			monthlyData: []struct {
				month     string
				realized  *float64
				meta      *float64
				breakdown interface{}
			}{
				{month: "JAN", realized: floatPtr(45), meta: floatPtr(40), breakdown: nil},
				{month: "FEV", realized: floatPtr(42), meta: floatPtr(40), breakdown: nil},
				{month: "MAR", realized: floatPtr(38), meta: floatPtr(38), breakdown: nil},
				{month: "ABR", realized: floatPtr(35), meta: floatPtr(38), breakdown: nil},
				{month: "MAI", realized: floatPtr(39), meta: floatPtr(35), breakdown: nil},
				{month: "JUN", realized: floatPtr(41), meta: floatPtr(35), breakdown: nil},
				{month: "JUL", realized: floatPtr(36), meta: floatPtr(35), breakdown: nil},
				{month: "AGO", realized: floatPtr(34), meta: floatPtr(35), breakdown: nil},
				{month: "SET", realized: floatPtr(37), meta: floatPtr(35), breakdown: nil},
				{month: "OUT", realized: floatPtr(33), meta: floatPtr(32), breakdown: nil},
				{month: "NOV", realized: floatPtr(35), meta: floatPtr(32), breakdown: nil},
				{month: "DEZ", realized: nil, meta: floatPtr(32), breakdown: nil},
			},
		},
		{
			title: "Leads Qualificados (MQLs)",
			monthlyData: []struct {
				month     string
				realized  *float64
				meta      *float64
				breakdown interface{}
			}{
				{month: "JAN", realized: floatPtr(85), meta: floatPtr(100), breakdown: nil},
				{month: "FEV", realized: floatPtr(110), meta: floatPtr(120), breakdown: nil},
				{month: "MAR", realized: floatPtr(145), meta: floatPtr(150), breakdown: nil},
				{month: "ABR", realized: floatPtr(180), meta: floatPtr(160), breakdown: nil},
				{month: "MAI", realized: floatPtr(210), meta: floatPtr(200), breakdown: nil},
				{month: "JUN", realized: floatPtr(195), meta: floatPtr(200), breakdown: nil},
				{month: "JUL", realized: floatPtr(230), meta: floatPtr(220), breakdown: nil},
				{month: "AGO", realized: floatPtr(255), meta: floatPtr(240), breakdown: nil},
				{month: "SET", realized: floatPtr(240), meta: floatPtr(240), breakdown: nil},
				{month: "OUT", realized: floatPtr(280), meta: floatPtr(260), breakdown: nil},
				{month: "NOV", realized: floatPtr(190), meta: floatPtr(200), breakdown: nil},
				{month: "DEZ", realized: nil, meta: floatPtr(150), breakdown: nil},
			},
		},
		{
			title: "ROAS (Retorno em Ads)",
			monthlyData: []struct {
				month     string
				realized  *float64
				meta      *float64
				breakdown interface{}
			}{
				{month: "JAN", realized: floatPtr(3.2), meta: floatPtr(4.0), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 3.52,
				}}},
				{month: "FEV", realized: floatPtr(3.5), meta: floatPtr(4.0), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 3.85,
				}}},
				{month: "MAR", realized: floatPtr(4.2), meta: floatPtr(4.5), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 4.62,
				}}},
				{month: "ABR", realized: floatPtr(4.8), meta: floatPtr(4.5), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 5.28,
				}}},
				{month: "MAI", realized: floatPtr(4.5), meta: floatPtr(5.0), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 4.95,
				}}},
				{month: "JUN", realized: floatPtr(4.1), meta: floatPtr(5.0), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 4.51,
				}}},
				{month: "JUL", realized: floatPtr(5.2), meta: floatPtr(5.0), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 5.72,
				}}},
				{month: "AGO", realized: floatPtr(5.5), meta: floatPtr(5.5), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 6.05,
				}}},
				{month: "SET", realized: floatPtr(5.1), meta: floatPtr(5.5), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 5.61,
				}}},
				{month: "OUT", realized: floatPtr(5.8), meta: floatPtr(6.0), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 6.38,
				}}},
				{month: "NOV", realized: floatPtr(4.9), meta: floatPtr(6.0), breakdown: []map[string]interface{}{{
					"label": "Campanhas Ads Meta (Instagram/Facebook)",
					"value": 5.2,
				}}},
				{month: "DEZ", realized: nil, meta: floatPtr(5.0), breakdown: nil},
			},
		},
		{
			title: "Taxa de Conversão Global (%)",
			monthlyData: []struct {
				month     string
				realized  *float64
				meta      *float64
				breakdown interface{}
			}{
				{month: "JAN", realized: floatPtr(2.1), meta: floatPtr(2.5), breakdown: nil},
				{month: "FEV", realized: floatPtr(2.3), meta: floatPtr(2.5), breakdown: nil},
				{month: "MAR", realized: floatPtr(2.8), meta: floatPtr(3.0), breakdown: nil},
				{month: "ABR", realized: floatPtr(3.2), meta: floatPtr(3.0), breakdown: nil},
				{month: "MAI", realized: floatPtr(3.0), meta: floatPtr(3.5), breakdown: nil},
				{month: "JUN", realized: floatPtr(2.9), meta: floatPtr(3.5), breakdown: nil},
				{month: "JUL", realized: floatPtr(3.5), meta: floatPtr(3.5), breakdown: nil},
				{month: "AGO", realized: floatPtr(3.8), meta: floatPtr(4.0), breakdown: nil},
				{month: "SET", realized: floatPtr(3.4), meta: floatPtr(4.0), breakdown: nil},
				{month: "OUT", realized: floatPtr(4.1), meta: floatPtr(4.0), breakdown: nil},
				{month: "NOV", realized: floatPtr(3.2), meta: floatPtr(4.0), breakdown: nil},
				{month: "DEZ", realized: nil, meta: floatPtr(3.5), breakdown: nil},
			},
		},
	}
}

// floatPtr is a helper to create a pointer to a float64
func floatPtr(f float64) *float64 {
	return &f
}
