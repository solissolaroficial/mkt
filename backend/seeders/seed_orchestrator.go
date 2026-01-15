package seeders

import (
	"context"
	"log"
)

// SeedAll runs all seeders in correct order
func SeedAll(
	ctx context.Context,
	userSeeder *UserSeeder,
	kpiSeeder *KpiSeeder,
	socialBenchmarkingSeeder *SocialBenchmarkingSeeder,
	cooperativeSeeder *CooperativeSeeder,
	giftSeeder *GiftSeeder,
	budgetSeeder *BudgetSeeder,
	representativeMonthlyGoalSeeder *RepresentativeMonthlyGoalSeeder,
) error {
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

	// Step 4: Seed Cooperative Data
	if err := cooperativeSeeder.Seed(ctx); err != nil {
		log.Printf("❌ Error seeding cooperative data: %v", err)
		return err
	}
	log.Println("✅ Cooperative data seeded successfully")
	log.Println("----------------------------------------")

	// Step 5: Seed Gifts Data
	if err := giftSeeder.Seed(ctx); err != nil {
		log.Printf("❌ Error seeding gifts data: %v", err)
		return err
	}
	log.Println("✅ Gifts data seeded successfully")
	log.Println("----------------------------------------")

	// Step 6: Seed Budget Data
	if err := budgetSeeder.Seed(ctx); err != nil {
		log.Printf("❌ Error seeding budget data: %v", err)
		return err
	}
	log.Println("✅ Budget data seeded successfully")
	log.Println("----------------------------------------")

	// Step 7: Seed Representative Monthly Goals Data
	if err := representativeMonthlyGoalSeeder.Seed(); err != nil {
		log.Printf("❌ Error seeding representative monthly goals data: %v", err)
		return err
	}
	log.Println("✅ Representative monthly goals data seeded successfully")
	log.Println("----------------------------------------")

	log.Println("========================================")
	log.Println("🎉 All seeders completed successfully!")
	log.Println("========================================")

	return nil
}
