package valueobject

import (
	"errors"
)

var ErrInvalidCategory = errors.New("invalid category")

type PostCategory string

const (
	CategoryOfficial  PostCategory = "official"
	CategorySolisVoce PostCategory = "solis_voce"
	CategoryLeonardo  PostCategory = "leonardo"
	CategoryLuiz      PostCategory = "luiz"
)

func NewPostCategory(category string) (PostCategory, error) {
	switch PostCategory(category) {
	case CategoryOfficial, CategorySolisVoce, CategoryLeonardo, CategoryLuiz:
		return PostCategory(category), nil
	default:
		return "", ErrInvalidCategory
	}
}

func (c PostCategory) String() string {
	return string(c)
}

func (c PostCategory) IsValid() bool {
	switch c {
	case CategoryOfficial, CategorySolisVoce, CategoryLeonardo, CategoryLuiz:
		return true
	default:
		return false
	}
}
