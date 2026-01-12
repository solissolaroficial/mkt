package migration

import (
	"log"

	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

// RunMigrations executes all database migrations using GORM AutoMigrate
func RunMigrations(db *gorm.DB) error {
	log.Println("📦 Creating/updating database tables...")

	// GORM AutoMigrate automatically creates tables, adds missing columns,
	// creates indexes, and foreign keys based on model tags
	err := db.AutoMigrate(
		&model.User{},
		&model.KpiCategory{},
		&model.MonthlyData{},
		&model.Task{},
		&model.Subtask{},
		&model.Comment{},
		&model.Notification{},
		&model.CalendarPostModel{},
		&model.PdvPostModel{},
		&model.RecurrentPdvModel{},
		&model.SocialBenchmarkingModel{},
		&model.SocialPostModel{},
		&model.SocialDailyAggregationModel{},
		&model.OfflineActionModel{},
		&model.ShowroomItemModel{},
		&model.RepMarketingActionModel{},
	)

	if err != nil {
		return err
	}

	log.Println("✅ Tables created/updated successfully")

	// Create additional indexes for performance
	if err := createAdditionalIndexes(db); err != nil {
		return err
	}

	log.Println("✅ Additional indexes created successfully")

	return nil
}

// createAdditionalIndexes creates performance indexes not covered by GORM tags
func createAdditionalIndexes(db *gorm.DB) error {
	// Composite index for monthly_data queries (kpi_category_id + month)
	// This speeds up queries like: "SELECT * FROM monthly_data WHERE kpi_category_id = ? AND month = ?"
	if err := db.Exec(`
		CREATE INDEX IF NOT EXISTS idx_monthly_data_kpi_month 
		ON monthly_data(kpi_category_id, month)
	`).Error; err != nil {
		log.Printf("⚠️  Warning: Failed to create index idx_monthly_data_kpi_month: %v", err)
		// Don't return error, just log warning
	}

	// Index for sorting by creation date
	if err := db.Exec(`
		CREATE INDEX IF NOT EXISTS idx_kpi_categories_created_at 
		ON kpi_categories(created_at DESC)
	`).Error; err != nil {
		log.Printf("⚠️  Warning: Failed to create index idx_kpi_categories_created_at: %v", err)
	}

	return nil
}
