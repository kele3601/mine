package repo

import (
	"gorm.io/gorm"
	"mine/internal/db/model"
)

type UserRepository struct {
	db *gorm.DB
}

func (ur *UserRepository) List() ([]*model.User, error) {
	var users []*model.User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (ur *UserRepository) FindByName(name string) (*model.User, error) {
	return ur.FindByField("name", name)
}

func (ur *UserRepository) CheckExistByField(field string, value string) bool {
	var user model.User
	return ur.db.Where(field+"=?", value).First(&user).RowsAffected > 0
}

func (ur *UserRepository) Create(user *model.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) FindByField(field string, value any) (*model.User, error) {
	var user model.User
	if err := ur.db.Where(field+"=?", value).First(&user).Error; err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
