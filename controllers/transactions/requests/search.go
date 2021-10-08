package requests

import (
	"alterra/business/karyawans"
	"alterra/business/payment_methods"
	"alterra/business/transactions"
	"alterra/business/users"
)

type Transaction_Search struct {
	Payment_Method payment_methods.Domain `json:"type"`
	User           users.Domain           `json:"user"`
	Karyawan       karyawans.Domain       `json:"karyawan"`
	Total_Qty      uint                   `json:"total_qty"`
	Total_Price    uint                   `json:"total_price"`
}

func ToDomain(search Transaction_Search) transactions.Domain {
	return transactions.Domain{
		Payment_Method: payment_methods.Domain{},
		User:           users.Domain{},
		Karyawan:       karyawans.Domain{},
		Total_Qty:      search.Total_Qty,
		Total_Price:    search.Total_Price,
	}
}
