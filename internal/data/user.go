package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-realworld/internal/biz"
)

// data中小写repo，是一种具体的实现，biz中大写repo，是一种interface，接口定义方法
type userRepo struct {
	data *Data
	log  *log.Helper
}

type articleRepo struct {
	data *Data
	log  *log.Helper
}

type commentRepo struct {
	data *Data
	log  *log.Helper
}

type tagRepo struct {
	data *Data
	log  *log.Helper
}

type profileRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &commentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	//TODO implement me
	return nil
}
