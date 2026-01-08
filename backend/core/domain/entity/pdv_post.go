package entity

import (
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/valueobject"

	"github.com/google/uuid"
)

// isValidURL verifica se uma string é uma URL válida
func isValidURL(s string) bool {
	if s == "" {
		return false
	}
	_, err := url.ParseRequestURI(s)
	return err == nil
}

// PdvPost representa um post de PDV (Ponto de Venda)
type PdvPost struct {
	id        uuid.UUID
	repName   string
	pdvName   string
	postDate  *valueobject.PdvPostDate
	month     string // Derivado da data (JAN, FEV, MAR, etc.)
	platform  string
	link      *string
	proofUrl  *string
	status    *valueobject.PdvStatus
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

// NewPdvPost cria uma nova entidade PdvPost
func NewPdvPost(
	repName string,
	pdvName string,
	postDate *valueobject.PdvPostDate,
	platform string,
	link *string,
	proofUrl *string,
) (*PdvPost, error) {
	repName = strings.TrimSpace(repName)
	if repName == "" {
		return nil, errors.New("repName is required")
	}

	pdvName = strings.TrimSpace(pdvName)
	if pdvName == "" {
		return nil, errors.New("pdvName is required")
	}

	if postDate == nil {
		return nil, errors.New("postDate is required")
	}

	// Validar plataforma
	if !valueobject.IsValidPlatform(platform) {
		return nil, errors.New("invalid platform")
	}

	// Validar link se fornecido
	if link != nil && *link != "" {
		if len(*link) > 500 {
			return nil, errors.New("link must be at most 500 characters")
		}
		if !isValidURL(*link) {
			return nil, errors.New("link must be a valid URL")
		}
	}

	// Validar proofUrl se fornecido
	if proofUrl != nil && *proofUrl != "" {
		if len(*proofUrl) > 500 {
			return nil, errors.New("proofUrl must be at most 500 characters")
		}
		if !isValidURL(*proofUrl) {
			return nil, errors.New("proofUrl must be a valid URL")
		}
	}

	// Derivar mês da data
	month := deriveMonthFromDate(postDate.Value())

	// Criar status inicial
	status, err := valueobject.NewPdvStatus(valueobject.PdvStatusPending)
	if err != nil {
		return nil, err
	}

	post := &PdvPost{
		id:        uuid.New(),
		repName:   repName,
		pdvName:   pdvName,
		postDate:  postDate,
		month:     month,
		platform:  platform,
		link:      link,
		proofUrl:  proofUrl,
		status:    status,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	if err := post.Validate(); err != nil {
		return nil, err
	}

	return post, nil
}

// ReconstructPdvPost reconstrói a entidade do banco de dados
func ReconstructPdvPost(
	id uuid.UUID,
	repName string,
	pdvName string,
	postDate *valueobject.PdvPostDate,
	month string,
	platform string,
	link *string,
	proofUrl *string,
	status *valueobject.PdvStatus,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) (*PdvPost, error) {
	post := &PdvPost{
		id:        id,
		repName:   repName,
		pdvName:   pdvName,
		postDate:  postDate,
		month:     month,
		platform:  platform,
		link:      link,
		proofUrl:  proofUrl,
		status:    status,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
	}
	if err := post.Validate(); err != nil {
		return nil, err
	}
	return post, nil
}

// deriveMonthFromDate deriva o mês abreviado da data
func deriveMonthFromDate(date time.Time) string {
	months := []string{"JAN", "FEV", "MAR", "ABR", "MAI", "JUN", "JUL", "AGO", "SET", "OUT", "NOV", "DEZ"}
	return months[date.Month()-1]
}

// Getters

// ID retorna o ID do post
func (p *PdvPost) ID() uuid.UUID {
	return p.id
}

// RepName retorna o nome do representante
func (p *PdvPost) RepName() string {
	return p.repName
}

// PdvName retorna o nome do PDV
func (p *PdvPost) PdvName() string {
	return p.pdvName
}

// PostDate retorna a data do post
func (p *PdvPost) PostDate() *valueobject.PdvPostDate {
	return p.postDate
}

// Month retorna o mês do post
func (p *PdvPost) Month() string {
	return p.month
}

// Platform retorna a plataforma do post
func (p *PdvPost) Platform() string {
	return p.platform
}

// Link retorna o link do post
func (p *PdvPost) Link() *string {
	return p.link
}

// ProofUrl retorna a URL da prova
func (p *PdvPost) ProofUrl() *string {
	return p.proofUrl
}

// Status retorna o status do post
func (p *PdvPost) Status() *valueobject.PdvStatus {
	return p.status
}

// CreatedAt retorna a data de criação
func (p *PdvPost) CreatedAt() time.Time {
	return p.createdAt
}

// UpdatedAt retorna a data de atualização
func (p *PdvPost) UpdatedAt() time.Time {
	return p.updatedAt
}

// DeletedAt retorna a data de exclusão (soft delete)
func (p *PdvPost) DeletedAt() *time.Time {
	return p.deletedAt
}

// Métodos de Negócio

// Validate valida os dados da entidade
func (p *PdvPost) Validate() error {
	if p.repName == "" {
		return errors.New("repName is required")
	}

	if p.pdvName == "" {
		return errors.New("pdvName is required")
	}

	if p.postDate == nil {
		return errors.New("postDate is required")
	}

	if !valueobject.IsValidPlatform(p.platform) {
		return errors.New("invalid platform")
	}

	if p.status == nil {
		return errors.New("status is required")
	}

	return nil
}

// UpdateLink atualiza o link
func (p *PdvPost) UpdateLink(link *string) error {
	if link != nil && *link != "" {
		if len(*link) > 500 {
			return errors.New("link must be at most 500 characters")
		}
		if !isValidURL(*link) {
			return errors.New("link must be a valid URL")
		}
	}

	p.link = link
	p.updatedAt = time.Now()
	return nil
}

// UpdateRepName atualiza o nome do representante
func (p *PdvPost) UpdateRepName(repName string) error {
	repName = strings.TrimSpace(repName)
	if repName == "" {
		return errors.New("repName is required")
	}
	if len(repName) > 100 {
		return errors.New("repName must be at most 100 characters")
	}
	p.repName = repName
	p.updatedAt = time.Now()
	return nil
}

// UpdatePdvName atualiza o nome do PDV
func (p *PdvPost) UpdatePdvName(pdvName string) error {
	pdvName = strings.TrimSpace(pdvName)
	if pdvName == "" {
		return errors.New("pdvName is required")
	}
	if len(pdvName) > 200 {
		return errors.New("pdvName must be at most 200 characters")
	}
	p.pdvName = pdvName
	p.updatedAt = time.Now()
	return nil
}

// UpdatePostDate atualiza a data do post
func (p *PdvPost) UpdatePostDate(postDate *valueobject.PdvPostDate) error {
	if postDate == nil {
		return errors.New("postDate is required")
	}

	// Derivar novo mês da data
	p.postDate = postDate
	p.month = deriveMonthFromDate(postDate.Value())
	p.updatedAt = time.Now()
	return nil
}

// UpdatePlatform atualiza a plataforma
func (p *PdvPost) UpdatePlatform(platform string) error {
	if !valueobject.IsValidPlatform(platform) {
		return errors.New("invalid platform")
	}
	p.platform = platform
	p.updatedAt = time.Now()
	return nil
}

// UpdateProofUrl atualiza a URL da prova
func (p *PdvPost) UpdateProofUrl(proofUrl *string) error {
	if proofUrl != nil && *proofUrl != "" {
		if len(*proofUrl) > 500 {
			return errors.New("proofUrl must be at most 500 characters")
		}
		if !isValidURL(*proofUrl) {
			return errors.New("proofUrl must be a valid URL")
		}
	}

	p.proofUrl = proofUrl
	p.updatedAt = time.Now()
	return nil
}

// UpdateStatus atualiza o status com validação de transição
func (p *PdvPost) UpdateStatus(newStatus *valueobject.PdvStatus) error {
	if newStatus == nil {
		return errors.New("new status is required")
	}

	// Validar transição de status
	if !p.status.CanTransitionTo(newStatus) {
		return errors.New("invalid status transition")
	}

	p.status = newStatus
	p.updatedAt = time.Now()
	return nil
}

// SoftDelete marca o post como deletado
func (p *PdvPost) SoftDelete() {
	now := time.Now()
	p.deletedAt = &now
	p.updatedAt = now
	// Atualizar status para cancelled
	cancelledStatus, _ := valueobject.NewPdvStatus(valueobject.PdvStatusCancelled)
	p.status = cancelledStatus
}

// IsActive verifica se o post está ativo (não deletado)
func (p *PdvPost) IsActive() bool {
	return p.deletedAt == nil
}

// IsVerified verifica se o post está verificado
func (p *PdvPost) IsVerified() bool {
	return p.status != nil && p.status.IsApproved()
}
