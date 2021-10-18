package transaction_details

import (
	"alterra/business/transaction_details"
	"time"

	"gorm.io/gorm"
)

type Transaction_Detail struct {
	Id             uint `gorm:"primaryKey"`
	Book_Id        uint
	Transaction_Id uint
	Qty            uint
	Price          uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func FromDomain(domain transaction_details.Domain) Transaction_Detail {
	return Transaction_Detail{
		Id:             domain.Id,
		Book_Id:        domain.Book_Id,
		Transaction_Id: domain.Transaction_Id,
		Qty:            domain.Qty,
		Price:          domain.Price,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}

func (transaction_detail *Transaction_Detail) ToDomain() transaction_details.Domain {
	return transaction_details.Domain{
		Id:             transaction_detail.Id,
		Book_Id:        transaction_detail.Book_Id,
		Transaction_Id: transaction_detail.Transaction_Id,
		Qty:            transaction_detail.Qty,
		Price:          transaction_detail.Price,
		CreatedAt:      transaction_detail.CreatedAt,
		UpdatedAt:      transaction_detail.UpdatedAt,
	}
}

func ToListDomain(data []Transaction_Detail) []transaction_details.Domain {
	list := []transaction_details.Domain{}
	for _, v := range data {
		list = append(list, v.ToDomain())
	}

	return list
}
