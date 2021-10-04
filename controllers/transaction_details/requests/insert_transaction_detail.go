package requests

import "alterra/business/transaction_details"

type InsertTransaction_detail struct {
	Qty   uint `json:"qty"`
	Price uint `json:"price"`
}

func (transaction_detail *InsertTransaction_detail) ToDomain() *transaction_details.Domain {
	return &transaction_details.Domain{
		Qty:   transaction_detail.Qty,
		Price: transaction_detail.Price,
	}
}
