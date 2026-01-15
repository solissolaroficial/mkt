package representativemonthlygoal

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetRepresentativeGoalsTableDataInput defines input data for getting goals table data
type GetRepresentativeGoalsTableDataInput struct {
	Year  int
	Month *int // Optional: if provided, filter by month
}

// GetRepresentativeGoalsTableDataOutput defines output data for transposed table view
// This matches the old implementation structure with representatives as rows and months as columns
type GetRepresentativeGoalsTableDataOutput struct {
	Year    int
	Months  []MonthData
	Rows    []RepresentativeRow
	Summary TableSummary
}

// MonthData represents a month column in table
type MonthData struct {
	Month       int
	MonthName   string
	TargetSum   float64
	RealizedSum float64
}

// RepresentativeRow represents a row in transposed table (one representative)
type RepresentativeRow struct {
	RepresentativeID   string
	RepresentativeName string
	Company            string
	Region             string
	City               string
	MonthValues        []MonthValue
	TotalTarget        float64
	TotalRealized      float64
	TotalPercentage    float64
}

// MonthValue represents a value for a specific month
type MonthValue struct {
	Month      int
	Target     float64
	Realized   float64
	Percentage float64
	IsMet      bool
}

// TableSummary represents summary statistics for table
type TableSummary struct {
	TotalRepresentatives int
	TotalTarget          float64
	TotalRealized        float64
	OverallPercentage    float64
	GoalsMetCount        int
	GoalsNotMetCount     int
}

type GetRepresentativeGoalsTableDataUseCase struct {
	monthlyGoalGateway gateway.RepresentativeMonthlyGoalGateway
}

func NewGetRepresentativeGoalsTableDataUseCase(
	monthlyGoalGateway gateway.RepresentativeMonthlyGoalGateway,
) *GetRepresentativeGoalsTableDataUseCase {
	return &GetRepresentativeGoalsTableDataUseCase{
		monthlyGoalGateway: monthlyGoalGateway,
	}
}

func (uc *GetRepresentativeGoalsTableDataUseCase) Execute(ctx context.Context, input GetRepresentativeGoalsTableDataInput) (*GetRepresentativeGoalsTableDataOutput, error) {
	// Get raw table data from gateway
	rawData, err := uc.monthlyGoalGateway.GetGoalsTableData(input.Year, input.Month)
	if err != nil {
		return nil, err
	}

	// Process raw data into structured format
	output := &GetRepresentativeGoalsTableDataOutput{
		Year:    input.Year,
		Months:  make([]MonthData, 0),
		Rows:    make([]RepresentativeRow, 0),
		Summary: TableSummary{},
	}

	// Build a map to group goals by representative
	type RepresentativeGoals struct {
		ID           string
		Name         string
		Company      string
		Region       string
		City         string
		GoalsByMonth map[int]struct {
			Target   float64
			Realized float64
		}
	}

	repGoalsMap := make(map[string]*RepresentativeGoals)
	monthsSet := make(map[int]bool)

	// Process raw data to build representative goals map
	for _, row := range rawData {
		repID := row["representative_uuid"].(string)
		repName := row["representative_name"].(string)
		company := row["company"].(string)
		region := row["region"].(string)
		city := row["city"].(string)
		month := int(row["month"].(float64))
		target := row["target"].(float64)
		realized := row["realized"].(float64)

		// Track months
		monthsSet[month] = true

		// Get or create representative goals entry
		if _, exists := repGoalsMap[repID]; !exists {
			repGoalsMap[repID] = &RepresentativeGoals{
				ID:      repID,
				Name:    repName,
				Company: company,
				Region:  region,
				City:    city,
				GoalsByMonth: make(map[int]struct {
					Target   float64
					Realized float64
				}),
			}
		}

		// Add month goal
		repGoalsMap[repID].GoalsByMonth[month] = struct {
			Target   float64
			Realized float64
		}{
			Target:   target,
			Realized: realized,
		}
	}

	// Determine which months to include
	monthsToShow := make([]int, 0)
	if input.Month != nil {
		monthsToShow = append(monthsToShow, *input.Month)
	} else {
		for m := 1; m <= 12; m++ {
			if monthsSet[m] {
				monthsToShow = append(monthsToShow, m)
			}
		}
	}

	// Build month data
	for _, month := range monthsToShow {
		monthNames := []string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}
		monthData := MonthData{
			Month:       month,
			MonthName:   monthNames[month-1],
			TargetSum:   0,
			RealizedSum: 0,
		}
		output.Months = append(output.Months, monthData)
	}

	// Build representative rows
	for _, repGoals := range repGoalsMap {
		monthValues := make([]MonthValue, 0, len(monthsToShow))
		totalTarget := 0.0
		totalRealized := 0.0

		for i, month := range monthsToShow {
			goal, exists := repGoals.GoalsByMonth[month]
			if !exists {
				monthValues[i] = MonthValue{
					Month:      month,
					Target:     0,
					Realized:   0,
					Percentage: 0,
					IsMet:      false,
				}
				continue
			}

			totalTarget += goal.Target
			totalRealized += goal.Realized

			percentage := 0.0
			if goal.Target > 0 {
				percentage = (goal.Realized / goal.Target) * 100
			}

			monthValues[i] = MonthValue{
				Month:      month,
				Target:     goal.Target,
				Realized:   goal.Realized,
				Percentage: percentage,
				IsMet:      goal.Realized >= goal.Target,
			}
		}

		// Calculate totals for this representative
		totalPercentage := 0.0
		if totalTarget > 0 {
			totalPercentage = (totalRealized / totalTarget) * 100
		}

		row := RepresentativeRow{
			RepresentativeID:   repGoals.ID,
			RepresentativeName: repGoals.Name,
			Company:            repGoals.Company,
			Region:             repGoals.Region,
			City:               repGoals.City,
			MonthValues:        monthValues,
			TotalTarget:        totalTarget,
			TotalRealized:      totalRealized,
			TotalPercentage:    totalPercentage,
		}

		output.Rows = append(output.Rows, row)

		// Update month sums
		for i, month := range monthsToShow {
			goal, exists := repGoals.GoalsByMonth[month]
			if exists {
				output.Months[i].TargetSum += goal.Target
				output.Months[i].RealizedSum += goal.Realized
			}
		}
	}

	// Calculate summary
	output.Summary.TotalRepresentatives = len(repGoalsMap)
	output.Summary.TotalTarget = 0
	output.Summary.TotalRealized = 0
	output.Summary.GoalsMetCount = 0
	output.Summary.GoalsNotMetCount = 0

	for _, repGoals := range repGoalsMap {
		for _, month := range monthsToShow {
			goal, exists := repGoals.GoalsByMonth[month]
			if !exists {
				continue
			}

			output.Summary.TotalTarget += goal.Target
			output.Summary.TotalRealized += goal.Realized

			if goal.Realized >= goal.Target {
				output.Summary.GoalsMetCount++
			} else {
				output.Summary.GoalsNotMetCount++
			}
		}
	}

	if output.Summary.TotalTarget > 0 {
		output.Summary.OverallPercentage = (output.Summary.TotalRealized / output.Summary.TotalTarget) * 100
	}

	return output, nil
}
