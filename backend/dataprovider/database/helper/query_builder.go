package helper

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"gorm.io/gorm"
)

// ApplyMonthYearFilters aplica filtros de mês e ano à query
func ApplyMonthYearFilters(query *gorm.DB, month *string, year *int) *gorm.DB {
	if month != nil && *month != "---" {
		query = query.Where("month = ?", *month)
	}
	if year != nil {
		query = query.Where("year = ?", *year)
	}
	return query
}

// ApplyPagination aplica paginação à query
func ApplyPagination(query *gorm.DB, page, limit *int) *gorm.DB {
	if page != nil && limit != nil {
		offset := (*page - 1) * *limit
		query = query.Offset(offset).Limit(*limit)
	}
	return query
}

// ApplySort aplica ordenação à query
func ApplySort(query *gorm.DB, sortBy, sortOrder *string) *gorm.DB {
	if sortBy != nil && *sortBy != "" {
		order := *sortBy
		if sortOrder != nil && *sortOrder != "" {
			order = order + " " + *sortOrder
		}
		query = query.Order(order)
	}
	return query
}

// FilterMonthlyData filtra um slice de MonthlyData baseado em mês e ano
// Se month for "---", retorna todos os dados do ano especificado
// Se ambos month e year forem nil, retorna todos os dados
func FilterMonthlyData(monthlyData []*entity.MonthlyData, month *string, year *int) []*entity.MonthlyData {
	if month == nil && year == nil {
		return monthlyData
	}

	filtered := make([]*entity.MonthlyData, 0)
	for _, data := range monthlyData {
		matchMonth := month == nil || *month == "---" || data.Month() == *month
		matchYear := year == nil || data.Year() == *year

		if matchMonth && matchYear {
			filtered = append(filtered, data)
		}
	}

	return filtered
}
