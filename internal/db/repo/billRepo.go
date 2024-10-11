package repo

import (
	"gorm.io/gorm"
	"mine/internal/db/model"
)

type BillRepository struct {
	db *gorm.DB
}

func (br *BillRepository) List() ([]*model.Bill, error) {
	var bills []*model.Bill
	err := br.db.Find(&bills).Error
	return bills, err
}

func (br *BillRepository) Create(bill *model.Bill) error {
	return br.db.Create(bill).Error
}

func NewBillRepo(db *gorm.DB) *BillRepository {
	return &BillRepository{db: db}
}
