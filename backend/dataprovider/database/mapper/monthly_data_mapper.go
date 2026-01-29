package mapper

import (
	"encoding/json"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/datatypes"
)

type MonthlyDataMapper struct{}

// ToDomain converte Model (GORM) para Entity (Domínio)
func (m *MonthlyDataMapper) ToDomain(dataModel *model.MonthlyData) (*entity.MonthlyData, error) {
	// Converte o campo JSONB para []byte
	var breakdownBytes []byte
	if dataModel.Breakdown != nil {
		// Marshal do datatypes.JSON para []byte
		var err error
		breakdownBytes, err = json.Marshal(dataModel.Breakdown)
		if err != nil {
			return nil, err
		}
	}

	monthlyData, err := entity.ReconstructMonthlyData(
		dataModel.UUID,
		dataModel.KpiCategoryID,
		dataModel.Year,
		dataModel.Month,
		dataModel.Realized,
		dataModel.Meta,
		breakdownBytes,
		dataModel.CreatedAt,
		dataModel.UpdatedAt,
		dataModel.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	// Converte o campo Logs JSONB para string e define na entidade
	if dataModel.Logs != nil {
		monthlyData.SetLogs(string(dataModel.Logs))
	}

	return monthlyData, nil
}

// ToModel converte Entity (Domínio) para Model (GORM)
func (m *MonthlyDataMapper) ToModel(data *entity.MonthlyData) (*model.MonthlyData, error) {
	// Converte o campo []byte para datatypes.JSON
	var breakdownJSON datatypes.JSON
	if data.Breakdown() != nil {
		var breakdown interface{}
		err := json.Unmarshal(data.Breakdown(), &breakdown)
		if err != nil {
			return nil, err
		}

		// Marshal para datatypes.JSON
		breakdownJSON, err = json.Marshal(breakdown)
		if err != nil {
			return nil, err
		}
	}

	// Converte o campo Logs string para datatypes.JSON
	var logsJSON datatypes.JSON
	if data.Logs() != "" {
		logsJSON = datatypes.JSON([]byte(data.Logs()))
	}

	return &model.MonthlyData{
		UUID:          data.ID(),
		KpiCategoryID: data.KpiCategoryID(),
		Year:          data.Year(),
		Month:         data.Month(),
		Realized:      data.Realized(),
		Meta:          data.Meta(),
		Breakdown:     breakdownJSON,
		Logs:          logsJSON,
		CreatedAt:     data.CreatedAt(),
		UpdatedAt:     data.UpdatedAt(),
		DeletedAt:     data.DeletedAt(),
	}, nil
}
