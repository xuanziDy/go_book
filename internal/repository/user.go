package repository

import (
	"context"
	"github.com/xuanziDy/go_book/internal/domain"
	"github.com/xuanziDy/go_book/internal/repository/dao"
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	user := domain.User{}

	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return user, err
	}

	user.Email = u.Email
	user.Password = u.Password
	return user, err
}

func (r *UserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	user := domain.User{}
	u, err := r.dao.FindById(ctx, id)
	if err != nil {
		return user, err
	}
	user.Nickname = u.Nickname
	user.Birthday = u.Birthday
	user.Introduction = u.Introduction
	return user, err
}

func (r *UserRepository) Update(ctx context.Context, u domain.User) error {
	return r.dao.Update(ctx, dao.User{
		Id:           u.Id,
		Nickname:     u.Nickname,
		Birthday:     u.Birthday,
		Introduction: u.Introduction,
	})
}
