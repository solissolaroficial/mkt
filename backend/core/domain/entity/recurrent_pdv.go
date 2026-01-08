package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

// RecurrentPdv representa um PDV recorrente
type RecurrentPdv struct {
	id               uuid.UUID
	name             string
	repName          string
	city             *string
	followers        *int
	instagramProfile *string
	createdAt        time.Time
	updatedAt        time.Time
	deletedAt        *time.Time
}

// NewRecurrentPdv cria uma nova entidade RecurrentPdv
func NewRecurrentPdv(
	name string,
	repName string,
) (*RecurrentPdv, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("name is required")
	}

	repName = strings.TrimSpace(repName)
	if repName == "" {
		return nil, errors.New("repName is required")
	}

	pdv := &RecurrentPdv{
		id:        uuid.New(),
		name:      name,
		repName:   repName,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	if err := pdv.Validate(); err != nil {
		return nil, err
	}

	return pdv, nil
}

// ReconstructRecurrentPdv reconstrói a entidade do banco de dados
func ReconstructRecurrentPdv(
	id uuid.UUID,
	name string,
	repName string,
	city *string,
	followers *int,
	instagramProfile *string,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *RecurrentPdv {
	return &RecurrentPdv{
		id:               id,
		name:             name,
		repName:          repName,
		city:             city,
		followers:        followers,
		instagramProfile: instagramProfile,
		createdAt:        createdAt,
		updatedAt:        updatedAt,
		deletedAt:        deletedAt,
	}
}

// Getters

// ID retorna o ID do PDV
func (r *RecurrentPdv) ID() uuid.UUID {
	return r.id
}

// Name retorna o nome do PDV
func (r *RecurrentPdv) Name() string {
	return r.name
}

// RepName retorna o nome do representante
func (r *RecurrentPdv) RepName() string {
	return r.repName
}

// City retorna a cidade do PDV
func (r *RecurrentPdv) City() *string {
	return r.city
}

// Followers retorna o número de seguidores
func (r *RecurrentPdv) Followers() *int {
	return r.followers
}

// InstagramProfile retorna o perfil do Instagram
func (r *RecurrentPdv) InstagramProfile() *string {
	return r.instagramProfile
}

// CreatedAt retorna a data de criação
func (r *RecurrentPdv) CreatedAt() time.Time {
	return r.createdAt
}

// UpdatedAt retorna a data de atualização
func (r *RecurrentPdv) UpdatedAt() time.Time {
	return r.updatedAt
}

// DeletedAt retorna a data de exclusão (soft delete)
func (r *RecurrentPdv) DeletedAt() *time.Time {
	return r.deletedAt
}

// Métodos de Negócio

// Validate valida os dados da entidade
func (r *RecurrentPdv) Validate() error {
	if r.name == "" {
		return errors.New("name is required")
	}

	if r.repName == "" {
		return errors.New("repName is required")
	}

	if len(r.name) > 200 {
		return errors.New("name must be at most 200 characters")
	}

	if r.city != nil && len(*r.city) > 100 {
		return errors.New("city must be at most 100 characters")
	}

	if r.instagramProfile != nil && len(*r.instagramProfile) > 100 {
		return errors.New("instagramProfile must be at most 100 characters")
	}

	return nil
}

// UpdateCity atualiza a cidade
func (r *RecurrentPdv) UpdateCity(city *string) error {
	if city != nil && len(*city) > 100 {
		return errors.New("city must be at most 100 characters")
	}

	r.city = city
	r.updatedAt = time.Now()
	return nil
}

// UpdateName atualiza o nome do PDV
func (r *RecurrentPdv) UpdateName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("name is required")
	}
	if len(name) > 200 {
		return errors.New("name must be at most 200 characters")
	}
	r.name = name
	r.updatedAt = time.Now()
	return nil
}

// UpdateRepName atualiza o nome do representante
func (r *RecurrentPdv) UpdateRepName(repName string) error {
	repName = strings.TrimSpace(repName)
	if repName == "" {
		return errors.New("repName is required")
	}
	if len(repName) > 100 {
		return errors.New("repName must be at most 100 characters")
	}
	r.repName = repName
	r.updatedAt = time.Now()
	return nil
}

// UpdateFollowers atualiza o número de seguidores
func (r *RecurrentPdv) UpdateFollowers(followers *int) error {
	if followers != nil && *followers < 1 {
		return errors.New("followers must be at least 1")
	}

	r.followers = followers
	r.updatedAt = time.Now()
	return nil
}

// UpdateInstagramProfile atualiza o perfil do Instagram
func (r *RecurrentPdv) UpdateInstagramProfile(profile *string) error {
	if profile != nil && len(*profile) > 100 {
		return errors.New("instagramProfile must be at most 100 characters")
	}

	r.instagramProfile = profile
	r.updatedAt = time.Now()
	return nil
}

// SoftDelete marca o PDV como deletado
func (r *RecurrentPdv) SoftDelete() {
	now := time.Now()
	r.deletedAt = &now
	r.updatedAt = now
}

// IsActive verifica se o PDV está ativo (não deletado)
func (r *RecurrentPdv) IsActive() bool {
	return r.deletedAt == nil
}
