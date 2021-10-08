package responses

import (
	"alterra/business/books"
	"alterra/business/transaction_details"
	"alterra/business/transactions"
	"time"
)

type Transaction_detail_response struct {
	Id          uint                `json:"id"`
	Book        books.Domain        `json:"book"`
	Transaction transactions.Domain `json:"transaction"`
	Qty         uint                `json:"qty"`
	Price       uint                `json:"price"`
	CreatedAt   time.Time           `json:"createdat "`
	UpdatedAt   time.Time           `json:"updateat "`
}

type SearchResponse struct {
	Transaction_detail interface{}
}

func FromDomain(domain transaction_details.Domain) Transaction_detail_response {
	return Transaction_detail_response{
		Id:          domain.Id,
		Book:        books.Domain{},
		Transaction: transactions.Domain{},
		Qty:         domain.Qty,
		Price:       domain.Price,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func FromTransaction_DetailsListDomain(domain []transaction_details.Domain) []Transaction_detail_response {
	var list []Transaction_detail_response
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}
