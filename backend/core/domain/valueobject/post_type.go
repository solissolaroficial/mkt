package valueobject

import (
	"errors"
)

var ErrInvalidType = errors.New("invalid type")

type PostType string

const (
	TypeVideo           PostType = "video"
	TypeStatic          PostType = "static"
	TypeCarousel        PostType = "carousel"
	TypeStory           PostType = "story"
	TypeArticleLinkedin PostType = "article_linkedin"
	TypeArticleBlog     PostType = "article_blog"
)

func NewPostType(postType string) (PostType, error) {
	switch PostType(postType) {
	case TypeVideo, TypeStatic, TypeCarousel, TypeStory, TypeArticleLinkedin, TypeArticleBlog:
		return PostType(postType), nil
	default:
		return "", ErrInvalidType
	}
}

func (t PostType) String() string {
	return string(t)
}

func (t PostType) IsValid() bool {
	switch t {
	case TypeVideo, TypeStatic, TypeCarousel, TypeStory, TypeArticleLinkedin, TypeArticleBlog:
		return true
	default:
		return false
	}
}
