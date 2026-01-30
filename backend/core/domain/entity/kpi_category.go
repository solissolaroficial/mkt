package entity

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

// KpiCategory represents a category of KPI with monthly data
type KpiCategory struct {
	id           uuid.UUID
	title        string
	slug         string
	color        string
	unit         string     // 'currency', 'percent', 'number'
	createdBy    *uuid.UUID // nullable para KPIs do sistema
	monthlyDatas []*MonthlyData
	createdAt    time.Time
	updatedAt    time.Time
}

// NewKpiCategory creates a new KPI category with validation
func NewKpiCategory(title, color, unit string, createdBy *uuid.UUID) (*KpiCategory, error) {
	return NewKpiCategoryWithSlug(title, color, unit, "", createdBy)
}

// NewKpiCategoryWithSlug creates a new KPI category with an optional custom slug
// If slug is empty, it will be generated automatically from the title
func NewKpiCategoryWithSlug(title, color, unit, slug string, createdBy *uuid.UUID) (*KpiCategory, error) {
	// Generate slug if not provided
	if slug == "" {
		slug = generateSlug(title)
	}

	kpi := &KpiCategory{
		id:           uuid.New(),
		title:        title,
		slug:         slug,
		color:        color,
		unit:         unit,
		createdBy:    createdBy,
		monthlyDatas: []*MonthlyData{},
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}

	if err := kpi.Validate(); err != nil {
		return nil, err
	}

	return kpi, nil
}

// ReconstructKpiCategory reconstructs a KPI category from database (without validation)
func ReconstructKpiCategory(id uuid.UUID, title, color, unit string, createdBy *uuid.UUID, createdAt, updatedAt time.Time) (*KpiCategory, error) {
	return &KpiCategory{
		id:           id,
		title:        title,
		slug:         generateSlug(title),
		color:        color,
		unit:         unit,
		createdBy:    createdBy,
		monthlyDatas: []*MonthlyData{},
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}, nil
}

// ReconstructKpiCategoryWithMonthlyData reconstructs a KPI category with monthly data from database
func ReconstructKpiCategoryWithMonthlyData(id uuid.UUID, title, color, unit string, createdBy *uuid.UUID, monthlyDatas []*MonthlyData, createdAt, updatedAt time.Time) (*KpiCategory, error) {
	return &KpiCategory{
		id:           id,
		title:        title,
		slug:         generateSlug(title),
		color:        color,
		unit:         unit,
		createdBy:    createdBy,
		monthlyDatas: monthlyDatas,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}, nil
}

// Getters (encapsulation)
func (k *KpiCategory) ID() uuid.UUID                { return k.id }
func (k *KpiCategory) Title() string                { return k.title }
func (k *KpiCategory) Slug() string                 { return k.slug }
func (k *KpiCategory) Color() string                { return k.color }
func (k *KpiCategory) Unit() string                 { return k.unit }
func (k *KpiCategory) CreatedBy() *uuid.UUID        { return k.createdBy }
func (k *KpiCategory) MonthlyDatas() []*MonthlyData { return k.monthlyDatas }
func (k *KpiCategory) CreatedAt() time.Time         { return k.createdAt }
func (k *KpiCategory) UpdatedAt() time.Time         { return k.updatedAt }

// Validate performs business validation
func (k *KpiCategory) Validate() error {
	if k.title == "" {
		return errors.New("title is required")
	}
	if k.unit != "currency" && k.unit != "percent" && k.unit != "number" {
		return errors.New("invalid unit: must be 'currency', 'percent', or 'number'")
	}
	return nil
}

// UpdateTitle updates the KPI category title with validation
func (k *KpiCategory) UpdateTitle(newTitle string) error {
	if newTitle == "" {
		return errors.New("title cannot be empty")
	}
	k.title = newTitle
	k.slug = generateSlug(newTitle)
	k.updatedAt = time.Now()
	return nil
}

// UpdateColor updates the KPI category color
func (k *KpiCategory) UpdateColor(newColor string) error {
	if newColor == "" {
		return errors.New("color cannot be empty")
	}
	k.color = newColor
	k.updatedAt = time.Now()
	return nil
}

// UpdateUnit updates the KPI category unit with validation
func (k *KpiCategory) UpdateUnit(newUnit string) error {
	if newUnit != "currency" && newUnit != "percent" && newUnit != "number" {
		return errors.New("invalid unit: must be 'currency', 'percent', or 'number'")
	}
	k.unit = newUnit
	k.updatedAt = time.Now()
	return nil
}

// AddMonthlyData adds monthly data to the KPI category
func (k *KpiCategory) AddMonthlyData(monthlyData *MonthlyData) {
	k.monthlyDatas = append(k.monthlyDatas, monthlyData)
	k.updatedAt = time.Now()
}

// SetMonthlyDatas sets the monthly data for the KPI category
func (k *KpiCategory) SetMonthlyDatas(monthlyDatas []*MonthlyData) {
	k.monthlyDatas = monthlyDatas
	k.updatedAt = time.Now()
}

// generateSlug creates a URL-friendly slug from a title
// Converts to lowercase, removes accents and special characters, replaces spaces with underscores
func generateSlug(title string) string {
	// Convert to lowercase
	slug := strings.ToLower(title)

	// Remove accents and special characters
	re := regexp.MustCompile(`[áéíóúçâêô%()]`)
	slug = re.ReplaceAllString(slug, "")

	// Replace spaces and special characters with underscores
	re = regexp.MustCompile(`[\s\(\)\%]`)
	slug = re.ReplaceAllString(slug, "_")

	// Remove multiple consecutive underscores
	re = regexp.MustCompile(`_+`)
	slug = re.ReplaceAllString(slug, "_")

	// Remove leading/trailing underscores
	slug = strings.Trim(slug, "_")

	return slug
}
