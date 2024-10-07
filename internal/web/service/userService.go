package service

import (
	"fmt"
	"gorm.io/gorm"
	"mine/common/utils"
	"mine/internal/db/model"
	"mine/internal/db/repo"
)

type UserRepo interface {
	Create(user *model.User) error
	CheckExistByField(field string, value string) bool
	FindByName(name string) (*model.User, error)
}

type UserServiceImpl struct {
	ur UserRepo
}

func (us *UserServiceImpl) Login(user *model.User) (string, error) {
	if "" == user.Name {
		return "", fmt.Errorf("用户名不能为空")
	}
	if userTemp, err := us.ur.FindByName(user.Name); err != nil {
		return "", err
	} else {
		if userTemp.Name == user.Name && userTemp.Pass != user.Pass {
			return "", fmt.Errorf("密码错误")
		}
		// jwt处理逻辑
		if token, err := utils.GenerateJwtToken(*userTemp); nil != err {
			return "", err
		} else {
			return token, nil
		}
	}
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
