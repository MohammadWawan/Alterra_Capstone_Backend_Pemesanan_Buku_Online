package responses

import (
	"alterra/business/karyawans"
	"alterra/business/payment_methods"
	"alterra/business/transactions"
	"alterra/business/users"
	"time"
)

type TransactionResponse struct {
	Id             uint                   `json:"id"`
	Payment_Method payment_methods.Domain `json:"payment_method"`
	User           users.Domain           `json:"user"`
	Karyawan       karyawans.Domain       `json:"karyawan"`
	Total_Qty      uint                   `json:"total_qty"`
	Total_Price    uint                   `json:"total_price"`
	CreatedAt      time.Time              `json:"createdat "`
	UpdatedAt      time.Time              `json:"updateat "`
}

type SearchResponse struct {
	Transaction interface{}
}

func FromDomain(domain transactions.Domain) TransactionResponse {
	return TransactionResponse{
		Id:             domain.Id,
		Payment_Method: payment_methods.Domain{},
		User:           users.Domain{},
		Karyawan:       karyawans.Domain{},
		Total_Qty:      domain.Total_Qty,
		Total_Price:    domain.Total_Price,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}

func FromTransactionsListDomain(domain []transactions.Domain) []TransactionResponse {
	var list []TransactionResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}
