package service

import (
	"context"
	"errors"
	"github.com/xuanziDy/go_book/internal/domain"
	"github.com/xuanziDy/go_book/internal/repository"
	_const "github.com/xuanziDy/go_book/internal/web/const"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (srv *UserService) SignUp(ctx context.Context, user domain.User) error {
	//对密码进行加密存储
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return srv.repo.Create(ctx, user)
}

func (srv *UserService) Login(ctx context.Context, user domain.User) (domain.User, error) {
	u, err := srv.repo.FindByEmail(ctx, user.Email)
	if err == gorm.ErrRecordNotFound {
		return domain.User{}, gorm.ErrRecordNotFound
	}
	if err != nil {
		return domain.User{}, err
	}
	// 比较密码了
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		return domain.User{}, errors.New(_const.ErrInvalidUserOrPassword)
	}
	return u, nil
}

func (srv *UserService) Edit(ctx context.Context, user domain.User) error {
	u, err := srv.repo.FindById(ctx, user.Id)
	if err != nil {
		return err
	}

	u.Nickname = user.Nickname
	u.Birthday = user.Birthday
	u.Introduction = user.Introduction

	return srv.repo.Update(ctx, u)
}

func (srv *UserService) Profile(ctx context.Context, userId int64) (domain.User, error) {
	u, err := srv.repo.FindById(ctx, userId)
	if err != nil {
		return domain.User{}, err
	}
	return u, nil
}
