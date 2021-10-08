package requests

import (
	"alterra/business/books"
	"alterra/business/transaction_details"
	"alterra/business/transactions"
)

type InsertTransaction_detail struct {
	Book        books.Domain        `json:"book"`
	Transaction transactions.Domain `json:"transaction"`
	Qty         uint                `json:"qty"`
	Price       uint                `json:"price"`
}

func (transaction_detail *InsertTransaction_detail) ToDomain() *transaction_details.Domain {
	return &transaction_details.Domain{
		Book:        books.Domain{},
		Transaction: transactions.Domain{},
		Qty:         transaction_detail.Qty,
		Price:       transaction_detail.Price,
	}
}
