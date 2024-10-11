package service

import (
	"gorm.io/gorm"
	"mine/internal/db/model"
	"mine/internal/db/repo"
)

type BillRepo interface {
	Create(bill *model.Bill) error
	List() ([]*model.Bill, error)
}

type BillServiceImpl struct {
	br BillRepo
}

func (bs *BillServiceImpl) List() ([]*model.Bill, error) {
	return bs.br.List()
}

func (bs *BillServiceImpl) Account(bill *model.Bill) error {
	bill.SetID()
	return bs.br.Create(bill)
}

func NewBillService(db *gorm.DB) *BillServiceImpl {
	return &BillServiceImpl{br: repo.NewBillRepo(db)}
}
