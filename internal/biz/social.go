package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-realworld/api/realworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type ArticleRepo interface {
}

type CommentRepo interface {
}

type TagRepo interface {
}

type SocialUseCase struct {
	ar ArticleRepo
	cr CommentRepo
	tr TagRepo

	log *log.Helper
}

func NewSocialUseCase(ar ArticleRepo, cr CommentRepo, tr CommentRepo, logger log.Logger) *SocialUseCase {
	return &SocialUseCase{
		ar:  ar,
		cr:  cr,
		tr:  tr,
		log: log.NewHelper(logger),
	}
}

func (uc *SocialUseCase) CreateArticle(ctx context.Context) error {
	return nil
}
