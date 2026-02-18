package entity

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/constants"

	"github.com/google/uuid"
)

// MonthlyData represents monthly data for a KPI category
type MonthlyData struct {
	id            uuid.UUID
	kpiCategoryID uuid.UUID
	year          int    // Year (e.g., 2024, 2025)
	month         string // 'JAN', 'FEV', etc
	realized      *float64
	meta          *float64
	breakdown     []byte // JSONB serialized
	logs          string // JSONB array of log entries
	createdAt     time.Time
	updatedAt     time.Time
	deletedAt     *time.Time // Soft delete timestamp
}

// NewMonthlyData creates new monthly data with validation
func NewMonthlyData(kpiCategoryID uuid.UUID, year int, month string) (*MonthlyData, error) {
	data := &MonthlyData{
		id:            uuid.New(),
		kpiCategoryID: kpiCategoryID,
		year:          year,
		month:         month,
		realized:      nil,
		meta:          nil,
		breakdown:     nil,
		createdAt:     time.Now(),
		updatedAt:     time.Now(),
	}

	if err := data.Validate(); err != nil {
		return nil, err
	}

	return data, nil
}

// ReconstructMonthlyData reconstructs monthly data from database (without validation)
func ReconstructMonthlyData(id, kpiCategoryID uuid.UUID, year int, month string, realized, meta *float64, breakdown []byte, createdAt, updatedAt time.Time, deletedAt *time.Time) (*MonthlyData, error) {
	return &MonthlyData{
		id:            id,
		kpiCategoryID: kpiCategoryID,
		year:          year,
		month:         month,
		realized:      realized,
		meta:          meta,
		breakdown:     breakdown,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
		deletedAt:     deletedAt,
	}, nil
}

// Getters (encapsulation)
func (m *MonthlyData) ID() uuid.UUID            { return m.id }
func (m *MonthlyData) KpiCategoryID() uuid.UUID { return m.kpiCategoryID }
func (m *MonthlyData) Year() int                { return m.year }
func (m *MonthlyData) Month() string            { return m.month }
func (m *MonthlyData) Realized() *float64       { return m.realized }
func (m *MonthlyData) Meta() *float64           { return m.meta }
func (m *MonthlyData) Breakdown() []byte        { return m.breakdown }
func (m *MonthlyData) Logs() string             { return m.logs }
func (m *MonthlyData) CreatedAt() time.Time     { return m.createdAt }
func (m *MonthlyData) UpdatedAt() time.Time     { return m.updatedAt }
func (m *MonthlyData) DeletedAt() *time.Time    { return m.deletedAt }

// Validate performs business validation
func (m *MonthlyData) Validate() error {
	if !constants.IsValidMonth(m.month) {
		return errors.New("invalid month: must be one of JAN, FEV, MAR, ABR, MAI, JUN, JUL, AGO, SET, OUT, NOV, DEZ")
	}
	return nil
}

// SetRealized updates the realized value
func (m *MonthlyData) SetRealized(value float64) {
	m.realized = &value
	m.updatedAt = time.Now()
}

// SetMeta updates the meta value
func (m *MonthlyData) SetMeta(value float64) {
	m.meta = &value
	m.updatedAt = time.Now()
}

// SetBreakdown serializes breakdown to JSON and stores it as []byte
func (m *MonthlyData) SetBreakdown(breakdown interface{}) error {
	if breakdown == nil {
		m.breakdown = nil
		m.updatedAt = time.Now()
		return nil
	}

	jsonBytes, err := json.Marshal(breakdown)
	if err != nil {
		return err
	}

	m.breakdown = jsonBytes
	m.updatedAt = time.Now()
	return nil
}

// GetBreakdown deserializes breakdown from []byte to interface{}
func (m *MonthlyData) GetBreakdown() (interface{}, error) {
	if m.breakdown == nil {
		return nil, nil
	}

	var result interface{}
	err := json.Unmarshal(m.breakdown, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// KpiLogEntry represents a single log entry for KPI changes
type KpiLogEntry struct {
	ID        string   `json:"id"`
	Date      string   `json:"date"`      // ISO date format
	Timestamp string   `json:"timestamp"` // "HH:MM" format
	User      string   `json:"user"`      // User name
	Month     string   `json:"month"`     // Month abbreviation
	OldValue  *float64 `json:"oldValue"`
	NewValue  float64  `json:"newValue"`
	Action    string   `json:"action"`  // "create" | "update"
	Context   string   `json:"context"` // Context of the change
}

// GetLogs returns the logs as a slice of KpiLogEntry
func (m *MonthlyData) GetLogs() []KpiLogEntry {
	if m.logs == "" {
		return []KpiLogEntry{}
	}
	var logs []KpiLogEntry
	err := json.Unmarshal([]byte(m.logs), &logs)
	if err != nil {
		return []KpiLogEntry{}
	}
	return logs
}

// SetLogs sets the logs from a JSON string
func (m *MonthlyData) SetLogs(logs string) {
	m.logs = logs
	m.updatedAt = time.Now()
}

// AddLog adds a new log entry to the logs array
func (m *MonthlyData) AddLog(log KpiLogEntry) error {
	logs := m.GetLogs()
	// Add new log at the beginning
	logs = append([]KpiLogEntry{log}, logs...)
	data, err := json.Marshal(logs)
	if err != nil {
		return err
	}
	m.logs = string(data)
	m.updatedAt = time.Now()
	return nil
}

// SoftDelete marks the monthly data as deleted
func (m *MonthlyData) SoftDelete() {
	now := time.Now()
	m.deletedAt = &now
	m.updatedAt = time.Now()
}

// IsActive returns true if the monthly data is not deleted
func (m *MonthlyData) IsActive() bool {
	return m.deletedAt == nil
}

// DailyEntry represents a single daily entry in the breakdown
type DailyEntry struct {
	Date      string    `json:"date"`      // ISO date format (e.g., "2024-11-01")
	Value     float64   `json:"value"`     // The value for this day
	Context   string    `json:"context"`   // Context description
	User      string    `json:"user"`      // User who made the entry
	CreatedAt time.Time `json:"createdAt"` // Creation timestamp
}

// GetDailyEntries retrieves all daily entries from the breakdown
func (m *MonthlyData) GetDailyEntries() ([]DailyEntry, error) {
	if m.breakdown == nil {
		return []DailyEntry{}, nil
	}

	var breakdown map[string]interface{}
	if err := json.Unmarshal(m.breakdown, &breakdown); err != nil {
		return nil, err
	}

	dailyData, ok := breakdown["daily"].([]interface{})
	if !ok || dailyData == nil {
		return []DailyEntry{}, nil
	}

	var entries []DailyEntry
	for _, item := range dailyData {
		entryMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		// Convert to DailyEntry struct with proper type checking
		date, ok := entryMap["date"].(string)
		if !ok {
			continue
		}
		value, ok := entryMap["value"].(float64)
		if !ok {
			continue
		}
		ctx, ok := entryMap["context"].(string)
		if !ok {
			continue
		}
		user, ok := entryMap["user"].(string)
		if !ok {
			continue
		}

		entries = append(entries, DailyEntry{
			Date:      date,
			Value:     value,
			Context:   ctx,
			User:      user,
			CreatedAt: time.Now(),
		})
	}

	return entries, nil
}

// AddDailyEntry adds a new daily entry to the breakdown
func (m *MonthlyData) AddDailyEntry(date string, value float64, context, user string) error {
	// Parse date to check validity
	if date == "" {
		return errors.New("date cannot be empty")
	}

	// Get existing daily entries
	dailyEntries, err := m.GetDailyEntries()
	if err != nil {
		return err
	}

	// Check if entry for this date already exists and update, or add new entry
	found := false
	for i, entry := range dailyEntries {
		if entry.Date == date {
			// Update existing entry
			dailyEntries[i] = DailyEntry{
				Date:      date,
				Value:     value,
				Context:   context,
				User:      user,
				CreatedAt: time.Now(),
			}
			found = true
			break
		}
	}

	// If not found, add new entry
	if !found {
		dailyEntries = append(dailyEntries, DailyEntry{
			Date:      date,
			Value:     value,
			Context:   context,
			User:      user,
			CreatedAt: time.Now(),
		})
	}

	// Update breakdown with daily entries
	breakdown := make(map[string]interface{})
	breakdown["daily"] = dailyEntries

	// Preserve existing breakdown data if any
	if m.breakdown != nil {
		var existingBreakdown map[string]interface{}
		if err := json.Unmarshal(m.breakdown, &existingBreakdown); err == nil {
			// Keep non-daily data
			for k, v := range existingBreakdown {
				if k != "daily" {
					breakdown[k] = v
				}
			}
		}
	}

	return m.SetBreakdown(breakdown)
}

// UpdateDailyEntry updates an existing daily entry
func (m *MonthlyData) UpdateDailyEntry(date string, value float64, context, user string) error {
	dailyEntries, err := m.GetDailyEntries()
	if err != nil {
		return err
	}

	found := false
	for i, entry := range dailyEntries {
		if entry.Date == date {
			dailyEntries[i] = DailyEntry{
				Date:      date,
				Value:     value,
				Context:   context,
				User:      user,
				CreatedAt: time.Now(),
			}
			found = true
			break
		}
	}

	if !found {
		return errors.New("daily entry not found")
	}

	breakdown := make(map[string]interface{})
	breakdown["daily"] = dailyEntries

	if m.breakdown != nil {
		var existingBreakdown map[string]interface{}
		if err := json.Unmarshal(m.breakdown, &existingBreakdown); err == nil {
			for k, v := range existingBreakdown {
				if k != "daily" {
					breakdown[k] = v
				}
			}
		}
	}

	return m.SetBreakdown(breakdown)
}

// DeleteDailyEntry removes a daily entry by date
func (m *MonthlyData) DeleteDailyEntry(date string) error {
	dailyEntries, err := m.GetDailyEntries()
	if err != nil {
		return err
	}

	found := false
	for i, entry := range dailyEntries {
		if entry.Date == date {
			dailyEntries = append(dailyEntries[:i], dailyEntries[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return errors.New("daily entry not found")
	}

	breakdown := make(map[string]interface{})
	breakdown["daily"] = dailyEntries

	if m.breakdown != nil {
		var existingBreakdown map[string]interface{}
		if err := json.Unmarshal(m.breakdown, &existingBreakdown); err == nil {
			for k, v := range existingBreakdown {
				if k != "daily" {
					breakdown[k] = v
				}
			}
		}
	}

	return m.SetBreakdown(breakdown)
}

// RecalculateFromDaily recalculates the realized value by summing all daily entries
func (m *MonthlyData) RecalculateFromDaily() error {
	dailyEntries, err := m.GetDailyEntries()
	if err != nil {
		return err
	}

	if len(dailyEntries) == 0 {
		// If no daily entries, clear the realized value
		m.realized = nil
		return nil
	}

	// Sum all daily values
	var total float64
	for _, entry := range dailyEntries {
		total += entry.Value
	}

	// Update realized value
	m.SetRealized(total)
	return nil
}
