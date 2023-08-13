package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	_const "github.com/xuanziDy/go_book/internal/web/const"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id           int64  `gorm:"primaryKey,autoIncrement"`
	Email        string `gorm:"unique"`
	Password     string
	Nickname     string `json:"nickname"`
	Birthday     string `json:"birthday"`
	Introduction string `json:"introduction"`
	Ctime        int64  //毫秒数
	Utime        int64
}

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Utime = now
	u.Ctime = now

	err := dao.db.WithContext(ctx).Create(&u).Error

	//唯一键冲突的需要特殊文案
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			return errors.New(_const.ErrUserDuplicateEmail)
		}
	}
	return err
}

func (dao *UserDAO) Update(ctx context.Context, u User) error {
	return dao.db.WithContext(ctx).Model(&u).Updates(User{
		Nickname:     u.Nickname,
		Birthday:     u.Birthday,
		Introduction: u.Introduction,
	}).Error
}

func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	return u, err
}

func (dao *UserDAO) FindById(ctx context.Context, id int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&u).Error
	return u, err
}
