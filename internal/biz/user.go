package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/pkg/middleware/auth"
)

type User struct {
	Username     string
	Token        string
	Email        string
	Bio          string
	Image        string
	PasswordHash string
}

type UserLogin struct {
	Email    string
	Username string
	Token    string
	Bio      string
	Image    string
}

func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("pwd: ", pwd, "b: ", b, "err ", err)
		panic(err)
	}
	return string(b)
}
func verifyPassword(rawPassword, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword)); err != nil {
		return false
	}
	return true
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type ProfileRepo interface {
}

type UserUseCase struct {
	ur UserRepo
	pr ProfileRepo

	log *log.Helper
	jwt *conf.JWT
}

func NewUserUseCase(ur UserRepo, pr ProfileRepo, logger log.Logger, jwt *conf.JWT) *UserUseCase {
	return &UserUseCase{
		ur:  ur,
		pr:  pr,
		log: log.NewHelper(logger),
		jwt: jwt,
	}
}

func (uc *UserUseCase) generateToken(username string) string {
	return auth.GenerateToken(uc.jwt.Token, username)
}

func (uc *UserUseCase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	u := &User{
		Username:     username,
		Email:        email,
		PasswordHash: hashPassword(password),
	}
	if err := uc.ur.CreateUser(ctx, u); err != nil {
		return &UserLogin{
			Email:    u.Email,
			Username: u.Username,
			Token:    u.Token,
			Bio:      u.Bio,
			Image:    u.Image,
		}, err
	}
	return &UserLogin{
		Email:    u.Email,
		Username: u.Username,
		Token:    uc.generateToken(u.Username),
	}, nil
}

func (uc *UserUseCase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !verifyPassword(u.PasswordHash, password) {
		return nil, errors.New(0, "error", "login failed")
	}
	return &UserLogin{
		Email:    u.Email,
		Username: u.Username,
		Bio:      u.Bio,
		Image:    u.Image,
		Token:    "abc",
	}, nil
}
