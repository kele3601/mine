package model

import "gorm.io/gorm"

type BaseModel struct {
	gorm.Model
	ID    string `json:"id" gorm:"primary_key;unique;size:10"`
	IsDel bool   `json:"is_del" gorm:"default:false"`
}
