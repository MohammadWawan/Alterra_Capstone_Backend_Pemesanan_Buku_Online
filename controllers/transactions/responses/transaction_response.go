package responses

import (
	"alterra/business/transactions"
	"time"
)

type TransactionResponse struct {
	Id                uint      `json:"id"`
	Method_Payment_Id uint      `json:"method_payment_Id"`
	User_Id           uint      `json:"user_id"`
	Karyawan_Id       uint      `json:"karyawan_id"`
	Total_Qty         uint      `json:"total_qty"`
	Total_Price       uint      `json:"total_price"`
	CreatedAt         time.Time `json:"createdat "`
	UpdatedAt         time.Time `json:"updateat "`
}

type SearchResponse struct {
	Transaction interface{}
}

func FromDomain(domain transactions.Domain) TransactionResponse {
	return TransactionResponse{
		Id:                domain.Id,
		Method_Payment_Id: domain.Method_Payment_Id,
		User_Id:           domain.User_Id,
		Karyawan_Id:       domain.Karyawan_Id,
		Total_Qty:         domain.Total_Qty,
		Total_Price:       domain.Total_Price,
		CreatedAt:         domain.CreatedAt,
		UpdatedAt:         domain.UpdatedAt,
	}
}

func FromTransactionsListDomain(domain []transactions.Domain) []TransactionResponse {
	var list []TransactionResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}
