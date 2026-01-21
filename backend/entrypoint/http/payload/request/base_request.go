package request

import (
	"fmt"
	"time"
)

// Constantes para validação
const (
	MinYear      = 2020
	MaxYear      = 2100
	MinPage      = 1
	MinLimit     = 1
	MaxLimit     = 100
	DefaultPage  = 1
	DefaultLimit = 50
)

// BaseQueryParams define parâmetros de query comuns para filtros temporais
type BaseQueryParams struct {
	Month *string `query:"month" validate:"omitempty,oneof=JAN FEV MAR ABR MAI JUN JUL AGO SET OUT NOV DEZ ---"`
	// Nota: A struct tag max=2100 é mais permissiva que o ValidateYear(), que limita a ano atual + 5
	Year  *int `query:"year" validate:"omitempty,min=2000,max=2100"`
	Page  *int `query:"page" validate:"omitempty,min=1"`
	Limit *int `query:"limit" validate:"omitempty,min=1,max=100"`
}

// GetMonth retorna o mês selecionado
func (b *BaseQueryParams) GetMonth() *string {
	return b.Month
}

// GetYear retorna o ano selecionado
func (b *BaseQueryParams) GetYear() *int {
	return b.Year
}

// GetPage retorna a página selecionada (com valor padrão)
func (b *BaseQueryParams) GetPage() int {
	if b.Page == nil {
		return DefaultPage
	}
	return *b.Page
}

// GetLimit retorna o limite selecionado (com valor padrão)
func (b *BaseQueryParams) GetLimit() int {
	if b.Limit == nil {
		return DefaultLimit
	}
	return *b.Limit
}

// IsFullYear retorna true se o mês selecionado for "---" (Ano Completo)
func (b *BaseQueryParams) IsFullYear() bool {
	return b.Month != nil && *b.Month == "---"
}

// ValidateYear valida se o ano está dentro do intervalo permitido
func (b *BaseQueryParams) ValidateYear() error {
	if b.Year == nil {
		return nil // Ano opcional, não valida se não fornecido
	}

	currentYear := time.Now().Year()
	year := *b.Year

	// Validação mais robusta
	if year < MinYear {
		return fmt.Errorf("ano deve ser maior ou igual a %d", MinYear)
	}

	if year > currentYear+5 {
		return fmt.Errorf("ano não pode ser mais de 5 anos no futuro (ano atual: %d)", currentYear)
	}

	return nil
}

// ValidatePagination valida se paginação está dentro dos limites
func (b *BaseQueryParams) ValidatePagination() error {
	page := b.GetPage()
	limit := b.GetLimit()

	if page < MinPage {
		return fmt.Errorf("página deve ser maior ou igual a %d", MinPage)
	}

	if limit < MinLimit {
		return fmt.Errorf("limite deve ser maior ou igual a %d", MinLimit)
	}

	if limit > MaxLimit {
		return fmt.Errorf("limite não pode ser maior que %d (para evitar ataques de DoS)", MaxLimit)
	}

	return nil
}

// Validate valida todos os parâmetros de query
func (b *BaseQueryParams) Validate() error {
	if err := b.ValidateYear(); err != nil {
		return err
	}

	if err := b.ValidatePagination(); err != nil {
		return err
	}

	return nil
}
