package service

import (
	"fmt"
	"gorm.io/gorm"
	"mine/internal/db/model"
	"mine/internal/db/repo"
)

type UserRepo interface {
	Create(user *model.User) error
	CheckExistByField(field string, value string) bool
}

type UserServiceImpl struct {
	ur UserRepo
}

func (us *UserServiceImpl) Register(user *model.User) error {
	// check
	if "" == user.Name {
		return fmt.Errorf("用户名不能为空")
	}
	if us.ur.CheckExistByField("name", user.Name) {
		return fmt.Errorf("用户名已存在")
	}
	user.SetID()
	return us.ur.Create(user)
}

func NewUserService(db *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{ur: repo.NewUserRepo(db)}
}
