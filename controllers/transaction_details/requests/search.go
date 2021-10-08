package requests

import (
	"alterra/business/books"
	"alterra/business/transaction_details"
	"alterra/business/transactions"
)

type Transaction_Detail_Search struct {
	Book        books.Domain        `json:"book"`
	Transaction transactions.Domain `json:"transaction"`
	Qty         uint                `json:"qty"`
	Price       uint                `json:"price"`
}

func ToDomain(search Transaction_Detail_Search) transaction_details.Domain {
	return transaction_details.Domain{
		Book:        books.Domain{},
		Transaction: transactions.Domain{},
		Qty:         search.Qty,
		Price:       search.Price,
	}
}
