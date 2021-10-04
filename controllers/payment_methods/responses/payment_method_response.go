package responses

import (
	"alterra/business/payment_methods"
	"time"
)

type Payment_MethodResponse struct {
	Id        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdat "`
	UpdatedAt time.Time `json:"updateat "`
}

type SearchResponse struct {
	Payment_Method interface{}
}

func FromDomain(domain payment_methods.Domain) Payment_MethodResponse {
	return Payment_MethodResponse{
		Id:        domain.Id,
		Type:      domain.Type,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromPayment_MethodsListDomain(domain []payment_methods.Domain) []Payment_MethodResponse {
	var list []Payment_MethodResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}
