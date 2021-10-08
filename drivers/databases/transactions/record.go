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
	Payment_Method    payment_methods.Payment_Method `gorm:"foreignKey:Payment_Method_Id"`
	User              users.Users                    `gorm:"foreignKey:User_Id"`
	Karyawan          karyawans.Karyawans            `gorm:"foreignKey:Karyawan_Id"`
	Total_Qty         uint
	Total_Price       uint
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

func FromDomain(domain transactions.Domain) Transaction {
	return Transaction{
		Id:             domain.Id,
		Payment_Method: payment_methods.FromDomain(domain.Payment_Method),
		User:           users.FromDomain(domain.User),
		Karyawan:       karyawans.FromDomain(domain.Karyawan),
		Total_Qty:      domain.Total_Qty,
		Total_Price:    domain.Total_Price,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}

func (transaction *Transaction) ToDomain() transactions.Domain {
	return transactions.Domain{
		Id:                transaction.Id,
		Method_Payment_Id: transaction.Payment_Method_Id,
		User_Id:           transaction.User_Id,
		Karyawan_Id:       transaction.Karyawan_Id,
		Total_Qty:         transaction.Total_Qty,
		Total_Price:       transaction.Total_Price,
		CreatedAt:         transaction.CreatedAt,
		UpdatedAt:         transaction.UpdatedAt,
	}
}

func ToListDomain(data []Transaction) []transactions.Domain {
	list := []transactions.Domain{}
	for _, v := range data {
		list = append(list, v.ToDomain())
	}

	return list
}
