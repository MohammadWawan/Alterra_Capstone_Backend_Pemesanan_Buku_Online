package transactions

import (
	"alterra/business/transactions"
	"alterra/drivers/databases/karyawans"
	"alterra/drivers/databases/payment_methods"
	"alterra/drivers/databases/users"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Id                uint `gorm:"primaryKey"`
	Payment_Method_Id uint
	User_Id           uint
	Karyawan_Id       uint
	Payment_Method    payment_methods.Payment_Method `gorm:"foreignKey:payment_method_id"`
	User              users.Users                    `gorm:"foreignKey:user_id"`
	Karyawan          karyawans.Karyawans            `gorm:"foreignKey:karyawan_id"`
	Total_Qty         uint
	Total_Price       uint
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

func FromDomain(domain transactions.Domain) Transaction {
	return Transaction{
		Id:             domain.Id,
		Payment_Method: payment_methods.Payment_Method{},
		User:           users.Users{},
		Karyawan:       karyawans.Karyawans{},
		Total_Qty:      domain.Total_Qty,
		Total_Price:    domain.Total_Price,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}

func (transaction *Transaction) ToDomain() transactions.Domain {
	return transactions.Domain{
		Id:          transaction.Id,
		Total_Qty:   transaction.Total_Qty,
		Total_Price: transaction.Total_Price,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	}
}

func ToListDomain(data []Transaction) []transactions.Domain {
	list := []transactions.Domain{}
	for _, v := range data {
		list = append(list, v.ToDomain())
	}

	return list
}
