package handler

import (
	"gorm.io/gorm"
	"mine/internal/web/service"
)

type Handler struct {
	us UserService
	bs BillService
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		us: service.NewUserService(db),
		bs: service.NewBillService(db),
	}
}
