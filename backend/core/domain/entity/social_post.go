package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type SocialPost struct {
	id              uuid.UUID
	brandID         uuid.UUID
	brandName       *valueobject.BrandName
	platform        valueobject.SocialPlatform
	postDate        time.Time
	postTime        *time.Time
	likes           int
	comments        int
	shares          *int
	postType        valueobject.SocialPostType
	caption         *string
	followersAtPost *int
	createdAt       time.Time
	updatedAt       time.Time
	deletedAt       *time.Time
}

// NewSocialPost cria uma nova entidade SocialPost
func NewSocialPost(
	brandName string,
	platform string,
	postDate time.Time,
	postTime *time.Time,
	likes int,
	comments int,
	shares *int,
	postType string,
	caption *string,
	followersAtPost *int,
) (*SocialPost, error) {
	// Validar e criar value objects
	brand, err := valueobject.NewBrandName(brandName)
	if err != nil {
		return nil, err
	}

	plat, err := valueobject.NewSocialPlatform(platform)
	if err != nil {
		return nil, err
	}

	pt, err := valueobject.NewSocialPostType(postType)
	if err != nil {
		return nil, err
	}

	// Validar valores numéricos
	if likes < 0 {
		return nil, errors.New("likes cannot be negative")
	}

	if comments < 0 {
		return nil, errors.New("comments cannot be negative")
	}

	if shares != nil && *shares < 0 {
		return nil, errors.New("shares cannot be negative")
	}

	if followersAtPost != nil && *followersAtPost < 0 {
		return nil, errors.New("followersAtPost cannot be negative")
	}

	if postDate.IsZero() {
		return nil, errors.New("postDate is required")
	}

	post := &SocialPost{
		id:              uuid.New(),
		brandID:         uuid.New(), // Will be updated when brand is created/fetched
		brandName:       brand,
		platform:        plat,
		postDate:        postDate,
		postTime:        postTime,
		likes:           likes,
		comments:        comments,
		shares:          shares,
		postType:        pt,
		caption:         caption,
		followersAtPost: followersAtPost,
		createdAt:       time.Now(),
		updatedAt:       time.Now(),
	}

	if err := post.Validate(); err != nil {
		return nil, err
	}

	return post, nil
}

// ReconstructSocialPost reconstrói a entidade do banco de dados
func ReconstructSocialPost(
	id uuid.UUID,
	brandID uuid.UUID,
	brandName string,
	platform string,
	postDate time.Time,
	postTime *time.Time,
	likes int,
	comments int,
	shares *int,
	postType string,
	caption *string,
	followersAtPost *int,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *SocialPost {
	brand := valueobject.ReconstructBrandName(brandName)
	plat := valueobject.ReconstructSocialPlatform(platform)
	pt := valueobject.ReconstructSocialPostType(postType)

	return &SocialPost{
		id:              id,
		brandID:         brandID,
		brandName:       brand,
		platform:        plat,
		postDate:        postDate,
		postTime:        postTime,
		likes:           likes,
		comments:        comments,
		shares:          shares,
		postType:        pt,
		caption:         caption,
		followersAtPost: followersAtPost,
		createdAt:       createdAt,
		updatedAt:       updatedAt,
		deletedAt:       deletedAt,
	}
}

// Getters
func (s *SocialPost) ID() uuid.UUID                        { return s.id }
func (s *SocialPost) BrandID() uuid.UUID                   { return s.brandID }
func (s *SocialPost) BrandName() *valueobject.BrandName    { return s.brandName }
func (s *SocialPost) Platform() valueobject.SocialPlatform { return s.platform }
func (s *SocialPost) PostDate() time.Time                  { return s.postDate }
func (s *SocialPost) PostTime() *time.Time                 { return s.postTime }
func (s *SocialPost) Likes() int                           { return s.likes }
func (s *SocialPost) Comments() int                        { return s.comments }
func (s *SocialPost) Shares() *int                         { return s.shares }
func (s *SocialPost) PostType() valueobject.SocialPostType { return s.postType }
func (s *SocialPost) Caption() *string                     { return s.caption }
func (s *SocialPost) FollowersAtPost() *int                { return s.followersAtPost }
func (s *SocialPost) CreatedAt() time.Time                 { return s.createdAt }
func (s *SocialPost) UpdatedAt() time.Time                 { return s.updatedAt }
func (s *SocialPost) DeletedAt() *time.Time                { return s.deletedAt }

// Métodos de Negócio

// Validate valida os dados da entidade
func (s *SocialPost) Validate() error {
	if s.brandID == uuid.Nil {
		return errors.New("brandID is required")
	}

	if s.brandName == nil {
		return errors.New("brandName is required")
	}

	if !s.platform.IsValid() {
		return errors.New("platform is required")
	}

	if s.postDate.IsZero() {
		return errors.New("postDate is required")
	}

	if s.likes < 0 {
		return errors.New("likes cannot be negative")
	}

	if s.comments < 0 {
		return errors.New("comments cannot be negative")
	}

	if s.shares != nil && *s.shares < 0 {
		return errors.New("shares cannot be negative")
	}

	if s.followersAtPost != nil && *s.followersAtPost < 0 {
		return errors.New("followersAtPost cannot be negative")
	}

	return nil
}

// UpdateMetrics atualiza as métricas (likes, comments, shares)
func (s *SocialPost) UpdateMetrics(likes, comments int, shares *int) error {
	if likes < 0 {
		return errors.New("likes cannot be negative")
	}

	if comments < 0 {
		return errors.New("comments cannot be negative")
	}

	if shares != nil && *shares < 0 {
		return errors.New("shares cannot be negative")
	}

	s.likes = likes
	s.comments = comments
	s.shares = shares
	s.updatedAt = time.Now()
	return nil
}

// UpdateCaption atualiza a legenda do post
func (s *SocialPost) UpdateCaption(caption *string) {
	s.caption = caption
	s.updatedAt = time.Now()
}

// UpdateFollowersAtPost atualiza o número de seguidores no momento do post
func (s *SocialPost) UpdateFollowersAtPost(followers *int) error {
	if followers != nil && *followers < 0 {
		return errors.New("followers cannot be negative")
	}

	s.followersAtPost = followers
	s.updatedAt = time.Now()
	return nil
}

// UpdateBrandID atualiza o ID da marca
func (s *SocialPost) UpdateBrandID(brandID uuid.UUID) error {
	if brandID == uuid.Nil {
		return errors.New("brandID cannot be nil")
	}

	s.brandID = brandID
	s.updatedAt = time.Now()
	return nil
}

// UpdatePlatform atualiza a plataforma
func (s *SocialPost) UpdatePlatform(platform valueobject.SocialPlatform) error {
	if !platform.IsValid() {
		return errors.New("invalid platform")
	}

	s.platform = platform
	s.updatedAt = time.Now()
	return nil
}

// UpdatePostDate atualiza a data do post
func (s *SocialPost) UpdatePostDate(postDate time.Time) error {
	if postDate.IsZero() {
		return errors.New("postDate is required")
	}

	s.postDate = postDate
	s.updatedAt = time.Now()
	return nil
}

// SetPostTime define o horário do post
func (s *SocialPost) SetPostTime(postTime *time.Time) {
	s.postTime = postTime
	s.updatedAt = time.Now()
}

// SetPostType define o tipo do post
func (s *SocialPost) SetPostType(postType valueobject.SocialPostType) {
	s.postType = postType
	s.updatedAt = time.Now()
}

// SetFollowersAtPost define o número de seguidores no momento do post
func (s *SocialPost) SetFollowersAtPost(followers *int) error {
	if followers != nil && *followers < 0 {
		return errors.New("followers cannot be negative")
	}

	s.followersAtPost = followers
	s.updatedAt = time.Now()
	return nil
}

// SoftDelete marca o post como deletado
func (s *SocialPost) SoftDelete() {
	now := time.Now()
	s.deletedAt = &now
	s.updatedAt = now
}

// IsActive verifica se o post está ativo (não deletado)
func (s *SocialPost) IsActive() bool {
	return s.deletedAt == nil
}
