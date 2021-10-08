package requests

import (
	"alterra/business/karyawans"
	"alterra/business/payment_methods"
	"alterra/business/transactions"
	"alterra/business/users"
)

type InsertTransaction struct {
	Type        payment_methods.Domain `json:"type"`
	User        users.Domain           `json:"user"`
	Karyawan    karyawans.Domain       `json:"karyawan"`
	Total_Qty   uint                   `json:"total_qty"`
	Total_Price uint                   `json:"total_price"`
}

func (transaction *InsertTransaction) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		Payment_Method: payment_methods.Domain{},
		User:           users.Domain{},
		Karyawan:       karyawans.Domain{},
		Total_Qty:      transaction.Total_Price,
		Total_Price:    transaction.Total_Price,
	}
}
