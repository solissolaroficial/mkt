package seeders

import (
	"log"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// RepresentativeMonthlyGoalSeeder handles seeding of representative monthly goals
type RepresentativeMonthlyGoalSeeder struct {
	monthlyGoalGateway    gateway.RepresentativeMonthlyGoalGateway
	representativeGateway gateway.RepresentativeGateway
}

// NewRepresentativeMonthlyGoalSeeder creates a new seeder instance
func NewRepresentativeMonthlyGoalSeeder(
	monthlyGoalGateway gateway.RepresentativeMonthlyGoalGateway,
	representativeGateway gateway.RepresentativeGateway,
) *RepresentativeMonthlyGoalSeeder {
	return &RepresentativeMonthlyGoalSeeder{
		monthlyGoalGateway:    monthlyGoalGateway,
		representativeGateway: representativeGateway,
	}
}

// Seed creates sample monthly goals for representatives
func (s *RepresentativeMonthlyGoalSeeder) Seed() error {
	// Get all representatives
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

	if len(representatives) == 0 {
		log.Println("No representatives found, skipping monthly goals seeding")
		return nil
	}

	// Create monthly goals for each representative for current year
	currentYear := 2025
	months := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	for _, rep := range representatives {
		for _, month := range months {
			// Check if goal already exists
			existing, err := s.monthlyGoalGateway.GetByRepresentativeAndMonth(rep.UUID(), month, currentYear)
			if err == nil && existing != nil {
				continue // Skip if already exists
			}

			// Create a sample goal with random target values
			target := float64(100000 + (month * 10000))
			realized := float64(float64(target) * (0.7 + float64(month%3)*0.1))

			goal, err := entity.NewRepresentativeMonthlyGoal(
				rep.UUID(),
				month,
				currentYear,
				target,
			)
			if err != nil {
				log.Printf("Error creating monthly goal: %v", err)
				continue
			}

			// Set realized value
			goal.UpdateRealized(realized)

			// Save to database
			if err := s.monthlyGoalGateway.Create(goal); err != nil {
				log.Printf("Error saving monthly goal: %v", err)
				continue
			}
		}
	}

	log.Printf("Seeded %d monthly goals for %d representatives", len(months)*len(representatives), len(representatives))
	return nil
}

// Clear removes all monthly goals
func (s *RepresentativeMonthlyGoalSeeder) Clear() error {
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

	for _, rep := range representatives {
		goals, err := s.monthlyGoalGateway.GetGoalsByRepresentative(rep.UUID())
		if err != nil {
			continue
		}

		for _, goal := range goals {
			if err := s.monthlyGoalGateway.Delete(goal.ID()); err != nil {
				log.Printf("Error deleting monthly goal %s: %v", goal.ID(), err)
			}
		}
	}

	log.Println("Cleared all monthly goals")
	return nil
}

// SeedForRepresentative creates monthly goals for a specific representative
func (s *RepresentativeMonthlyGoalSeeder) SeedForRepresentative(representativeID uuid.UUID, year int) error {
	// Verify representative exists
	rep, err := s.representativeGateway.FindByID(representativeID)
	if err != nil {
		return err
	}

	months := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for _, month := range months {
		// Check if goal already exists
		existing, err := s.monthlyGoalGateway.GetByRepresentativeAndMonth(representativeID, month, year)
		if err == nil && existing != nil {
			continue
		}

		// Create a sample goal
		target := float64(100000 + (month * 10000))
		realized := float64(float64(target) * 0.8)

		goal, err := entity.NewRepresentativeMonthlyGoal(
			representativeID,
			month,
			year,
			target,
		)
		if err != nil {
			log.Printf("Error creating monthly goal: %v", err)
			continue
		}

		goal.UpdateRealized(realized)

		if err := s.monthlyGoalGateway.Create(goal); err != nil {
			log.Printf("Error saving monthly goal: %v", err)
			continue
		}
	}

	log.Printf("Seeded monthly goals for representative %s (%s)", representativeID, rep.Name())
	return nil
}
