package entity

import (
	"time"

	"github.com/google/uuid"
	budgeterrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type BudgetItem struct {
	id           uuid.UUID
	objectUUID   *uuid.UUID
	objectName   string // Populado via JOIN com BudgetObject
	groupUUID    *uuid.UUID
	groupName    string // Populado via JOIN com BudgetGroup
	cod          *valueobject.BudgetCode
	desc         string
	vals         []float64 // 12 valores orçados (JAN-DEZ)
	realizedVals []float64 // 12 valores realizados (JAN-DEZ)
	year         int       // Ano do orçamento (ex: 2025)
	createdAt    time.Time
	updatedAt    time.Time
	deletedAt    *time.Time
}

// NewBudgetItem cria uma nova entidade BudgetItem com validação completa
// Validações:
// - Arrays vals e realizedVals devem ter exatamente 12 elementos
// - Todos os valores devem ser não-negativos
// - Year deve estar entre 2000 e 2100
// - Desc não pode ser vazio
// - objectUUID e groupUUID são opcionais (podem ser nil)
func NewBudgetItem(
	objectUUID *uuid.UUID,
	objectName string,
	groupUUID *uuid.UUID,
	groupName string,
	cod string,
	desc string,
	vals []float64,
	realizedVals []float64,
	year int,
) (*BudgetItem, error) {
	// Validar arrays de 12 meses
	if len(vals) != 12 {
		return nil, budgeterrors.ErrInvalidMonthCount
	}
	if len(realizedVals) != 12 {
		return nil, budgeterrors.ErrInvalidMonthCount
	}

	// Validar valores não-negativos
	for _, v := range vals {
		if v < 0 {
			return nil, budgeterrors.ErrNegativeValue
		}
	}
	for _, v := range realizedVals {
		if v < 0 {
			return nil, budgeterrors.ErrNegativeValue
		}
	}

	// Validar ano
	if year < 2000 || year > 2100 {
		return nil, budgeterrors.ErrInvalidYear
	}

	// Validar campos obrigatórios
	if desc == "" {
		return nil, budgeterrors.ErrInvalidDescription
	}

	// Validar código
	codVO, err := valueobject.NewBudgetCode(cod)
	if err != nil {
		return nil, budgeterrors.ErrInvalidCod
	}

	now := time.Now()

	return &BudgetItem{
		id:           uuid.New(),
		objectUUID:   objectUUID,
		objectName:   objectName,
		groupUUID:    groupUUID,
		groupName:    groupName,
		cod:          codVO,
		desc:         desc,
		vals:         vals,
		realizedVals: realizedVals,
		year:         year,
		createdAt:    now,
		updatedAt:    now,
		deletedAt:    nil,
	}, nil
}

// ReconstructBudgetItem reconstrói a entidade do banco de dados
// Não realiza validações, assume que os dados são válidos
func ReconstructBudgetItem(
	id uuid.UUID,
	objectUUID *uuid.UUID,
	objectName string,
	groupUUID *uuid.UUID,
	groupName string,
	cod string,
	desc string,
	vals []float64,
	realizedVals []float64,
	year int,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *BudgetItem {
	return &BudgetItem{
		id:           id,
		objectUUID:   objectUUID,
		objectName:   objectName,
		groupUUID:    groupUUID,
		groupName:    groupName,
		cod:          valueobject.ReconstructBudgetCode(cod),
		desc:         desc,
		vals:         vals,
		realizedVals: realizedVals,
		year:         year,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
		deletedAt:    deletedAt,
	}
}

// Getters
func (b *BudgetItem) ID() uuid.UUID {
	return b.id
}

func (b *BudgetItem) ObjectUUID() *uuid.UUID {
	return b.objectUUID
}

func (b *BudgetItem) ObjectName() string {
	return b.objectName
}

func (b *BudgetItem) GroupUUID() *uuid.UUID {
	return b.groupUUID
}

func (b *BudgetItem) GroupName() string {
	return b.groupName
}

func (b *BudgetItem) Cod() string {
	if b.cod == nil {
		return ""
	}
	return b.cod.String()
}

func (b *BudgetItem) Desc() string {
	return b.desc
}

func (b *BudgetItem) Vals() []float64 {
	return b.vals
}

func (b *BudgetItem) RealizedVals() []float64 {
	return b.realizedVals
}

func (b *BudgetItem) Year() int {
	return b.year
}

func (b *BudgetItem) CreatedAt() time.Time {
	return b.createdAt
}

func (b *BudgetItem) UpdatedAt() time.Time {
	return b.updatedAt
}

func (b *BudgetItem) DeletedAt() *time.Time {
	return b.deletedAt
}

// Métodos de Negócio

// Validate realiza validações completas da entidade
func (b *BudgetItem) Validate() error {
	if len(b.vals) != 12 {
		return budgeterrors.ErrInvalidMonthCount
	}
	if len(b.realizedVals) != 12 {
		return budgeterrors.ErrInvalidMonthCount
	}
	for _, v := range b.vals {
		if v < 0 {
			return budgeterrors.ErrNegativeValue
		}
	}
	for _, v := range b.realizedVals {
		if v < 0 {
			return budgeterrors.ErrNegativeValue
		}
	}
	if b.year < 2000 || b.year > 2100 {
		return budgeterrors.ErrInvalidYear
	}
	if b.desc == "" {
		return budgeterrors.ErrInvalidDescription
	}
	return nil
}

// UpdateVals atualiza os valores orçados
func (b *BudgetItem) UpdateVals(vals []float64) error {
	if len(vals) != 12 {
		return budgeterrors.ErrInvalidMonthCount
	}
	for _, v := range vals {
		if v < 0 {
			return budgeterrors.ErrNegativeValue
		}
	}
	b.vals = vals
	b.updatedAt = time.Now()
	return nil
}

// UpdateRealizedVals atualiza os valores realizados
func (b *BudgetItem) UpdateRealizedVals(realizedVals []float64) error {
	if len(realizedVals) != 12 {
		return budgeterrors.ErrInvalidMonthCount
	}
	for _, v := range realizedVals {
		if v < 0 {
			return budgeterrors.ErrNegativeValue
		}
	}
	b.realizedVals = realizedVals
	b.updatedAt = time.Now()
	return nil
}

// UpdateMonthValue atualiza o valor orçado de um mês específico
// monthIndex: 0 = Janeiro, 11 = Dezembro
func (b *BudgetItem) UpdateMonthValue(monthIndex int, val float64) error {
	if monthIndex < 0 || monthIndex > 11 {
		return budgeterrors.ErrInvalidMonthIndex
	}
	if val < 0 {
		return budgeterrors.ErrNegativeValue
	}
	b.vals[monthIndex] = val
	b.updatedAt = time.Now()
	return nil
}

// UpdateMonthRealized atualiza o valor realizado de um mês específico
// monthIndex: 0 = Janeiro, 11 = Dezembro
func (b *BudgetItem) UpdateMonthRealized(monthIndex int, realizedVal float64) error {
	if monthIndex < 0 || monthIndex > 11 {
		return budgeterrors.ErrInvalidMonthIndex
	}
	if realizedVal < 0 {
		return budgeterrors.ErrNegativeValue
	}
	b.realizedVals[monthIndex] = realizedVal
	b.updatedAt = time.Now()
	return nil
}

// SoftDelete marca a entidade como deletada
func (b *BudgetItem) SoftDelete() {
	now := time.Now()
	b.deletedAt = &now
	b.updatedAt = now
}

// IsActive retorna true se a entidade não foi deletada
func (b *BudgetItem) IsActive() bool {
	return b.deletedAt == nil
}

// UpdateObject atualiza o objeto (UUID e nome)
func (b *BudgetItem) UpdateObject(objectUUID *uuid.UUID, objectName string) {
	b.objectUUID = objectUUID
	b.objectName = objectName
	b.updatedAt = time.Now()
}

// UpdateGroup atualiza o grupo (UUID e nome)
func (b *BudgetItem) UpdateGroup(groupUUID *uuid.UUID, groupName string) {
	b.groupUUID = groupUUID
	b.groupName = groupName
	b.updatedAt = time.Now()
}

// UpdateCod atualiza o código do item
func (b *BudgetItem) UpdateCod(cod string) error {
	codVO, err := valueobject.NewBudgetCode(cod)
	if err != nil {
		return budgeterrors.ErrInvalidCod
	}
	b.cod = codVO
	b.updatedAt = time.Now()
	return nil
}

// UpdateDesc atualiza a descrição
func (b *BudgetItem) UpdateDesc(desc string) error {
	if desc == "" {
		return budgeterrors.ErrInvalidDescription
	}
	b.desc = desc
	b.updatedAt = time.Now()
	return nil
}

// GetTotalBudget retorna a soma de todos os valores orçados
func (b *BudgetItem) GetTotalBudget() float64 {
	total := 0.0
	for _, v := range b.vals {
		total += v
	}
	return total
}

// GetTotalRealized retorna a soma de todos os valores realizados
func (b *BudgetItem) GetTotalRealized() float64 {
	total := 0.0
	for _, v := range b.realizedVals {
		total += v
	}
	return total
}

// GetBudgetVariance retorna a diferença entre orçado e realizado
// Negativo = acima do orçamento (gastou mais do que planejado)
// Positivo = abaixo do orçamento (economizou)
func (b *BudgetItem) GetBudgetVariance() float64 {
	return b.GetTotalBudget() - b.GetTotalRealized()
}

// GetMonthBudget retorna o valor orçado de um mês específico
// monthIndex: 0 = Janeiro, 11 = Dezembro
func (b *BudgetItem) GetMonthBudget(monthIndex int) (float64, error) {
	if monthIndex < 0 || monthIndex > 11 {
		return 0, budgeterrors.ErrInvalidMonthIndex
	}
	return b.vals[monthIndex], nil
}

// GetMonthRealized retorna o valor realizado de um mês específico
// monthIndex: 0 = Janeiro, 11 = Dezembro
func (b *BudgetItem) GetMonthRealized(monthIndex int) (float64, error) {
	if monthIndex < 0 || monthIndex > 11 {
		return 0, budgeterrors.ErrInvalidMonthIndex
	}
	return b.realizedVals[monthIndex], nil
}

// GetMonthVariance retorna a variância de um mês específico
// monthIndex: 0 = Janeiro, 11 = Dezembro
func (b *BudgetItem) GetMonthVariance(monthIndex int) (float64, error) {
	if monthIndex < 0 || monthIndex > 11 {
		return 0, budgeterrors.ErrInvalidMonthIndex
	}
	return b.vals[monthIndex] - b.realizedVals[monthIndex], nil
}
