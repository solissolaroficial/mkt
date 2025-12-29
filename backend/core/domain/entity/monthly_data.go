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
	month         string // 'JAN', 'FEV', etc
	realized      *float64
	meta          *float64
	breakdown     []byte // JSONB serialized
	logs          string // JSONB array of log entries
	createdAt     time.Time
	updatedAt     time.Time
}

// NewMonthlyData creates new monthly data with validation
func NewMonthlyData(kpiCategoryID uuid.UUID, month string) (*MonthlyData, error) {
	data := &MonthlyData{
		id:            uuid.New(),
		kpiCategoryID: kpiCategoryID,
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
func ReconstructMonthlyData(id, kpiCategoryID uuid.UUID, month string, realized, meta *float64, breakdown []byte, createdAt, updatedAt time.Time) (*MonthlyData, error) {
	return &MonthlyData{
		id:            id,
		kpiCategoryID: kpiCategoryID,
		month:         month,
		realized:      realized,
		meta:          meta,
		breakdown:     breakdown,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
	}, nil
}

// Getters (encapsulation)
func (m *MonthlyData) ID() uuid.UUID            { return m.id }
func (m *MonthlyData) KpiCategoryID() uuid.UUID { return m.kpiCategoryID }
func (m *MonthlyData) Month() string            { return m.month }
func (m *MonthlyData) Realized() *float64       { return m.realized }
func (m *MonthlyData) Meta() *float64           { return m.meta }
func (m *MonthlyData) Breakdown() []byte        { return m.breakdown }
func (m *MonthlyData) Logs() string             { return m.logs }
func (m *MonthlyData) CreatedAt() time.Time     { return m.createdAt }
func (m *MonthlyData) UpdatedAt() time.Time     { return m.updatedAt }

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
