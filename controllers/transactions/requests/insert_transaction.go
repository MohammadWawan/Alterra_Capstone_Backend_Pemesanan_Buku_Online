package requests

import "alterra/business/transactions"

type InsertTransaction struct {
	Total_Qty   uint `json:"total_qty"`
	Total_Price uint `json:"total_price"`
}

func (transaction *InsertTransaction) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		Total_Qty:   transaction.Total_Price,
		Total_Price: transaction.Total_Price,
	}
}
