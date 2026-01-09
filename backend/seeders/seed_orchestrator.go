package seeders

import (
	"context"
	"log"
)

// SeedAll runs all seeders in the correct order
func SeedAll(ctx context.Context, userSeeder *UserSeeder, kpiSeeder *KpiSeeder, socialBenchmarkingSeeder *SocialBenchmarkingSeeder) error {
	log.Println("🌱 Starting database seeding...")
	log.Println("========================================")

	// Step 1: Seed Users
	if err := userSeeder.Seed(ctx); err != nil {
		log.Printf("❌ Error seeding users: %v", err)
		return err
	}
	log.Println("✅ Users seeded successfully")
	log.Println("----------------------------------------")

	// Step 2: Seed KPI Categories and Monthly Data
	if err := kpiSeeder.Seed(ctx); err != nil {
		log.Printf("❌ Error seeding KPIs: %v", err)
		return err
	}
	log.Println("✅ KPIs seeded successfully")
	log.Println("----------------------------------------")

	// Step 3: Seed Social Benchmarkings
	if err := socialBenchmarkingSeeder.Seed(ctx); err != nil {
		log.Printf("❌ Error seeding social benchmarkings: %v", err)
		return err
	}
	log.Println("✅ Social benchmarkings seeded successfully")
	log.Println("----------------------------------------")

	log.Println("========================================")
	log.Println("🎉 All seeders completed successfully!")
	log.Println("========================================")

	return nil
}
