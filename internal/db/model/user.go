package model

import gonanoid "github.com/matoous/go-nanoid/v2"

type User struct {
	BaseModel
	Name  string `json:"name" gorm:"size:10;not null"`
	Pass  string `json:"pass" gorm:"not null"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Level int    `json:"level"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) SetID() {
	s, _ := gonanoid.New()
	u.ID = s
}
