package requests

import "alterra/business/transactions"

type Transaction_Search struct {
	Total_Qty   uint `json:"total_qty"`
	Total_Price uint `json:"total_price"`
}

func ToDomain(search Transaction_Search) transactions.Domain {
	return transactions.Domain{
		Total_Qty:   search.Total_Qty,
		Total_Price: search.Total_Price,
	}
}
