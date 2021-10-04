package requests

import "alterra/business/transaction_details"

type Transaction_Detail_Search struct {
	Qty   uint `json:"qty"`
	Price uint `json:"price"`
}

func ToDomain(search Transaction_Detail_Search) transaction_details.Domain {
	return transaction_details.Domain{
		Qty:   search.Qty,
		Price: search.Price,
	}
}
