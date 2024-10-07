package repo

import (
	"gorm.io/gorm"
	"mine/internal/db/model"
)

type UserRepository struct {
	db *gorm.DB
}

func (ur *UserRepository) CheckExistByField(field string, value string) bool {
	var user model.User
	return ur.db.Where(field+"=?", value).First(&user).RowsAffected > 0
}

func (ur *UserRepository) Create(user *model.User) error {
	return ur.db.Create(user).Error
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
