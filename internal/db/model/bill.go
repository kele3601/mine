package model

import gonanoid "github.com/matoous/go-nanoid/v2"

type Bill struct {
	BaseModel
	User   string  `json:"user"`
	Amount float64 `json:"amount"`
	Memo   string  `json:"memo"`
	IsOut  bool    `json:"is_out"`
}

func (b *Bill) TableName() string {
	return "bill"
}

func (b *Bill) SetID() {
	s, _ := gonanoid.New()
	b.ID = s
}
