package responses

import (
	"alterra/business/descriptions"
	"time"
)

type DescriptionResponse struct {
	Id          uint      `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdat "`
	UpdatedAt   time.Time `json:"updateat "`
}

type SearchResponse struct {
	Description interface{}
}

func FromDomain(domain descriptions.Domain) DescriptionResponse {
	return DescriptionResponse{
		Id:          domain.Id,
		Description: domain.Description,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func FromDescriptionsListDomain(domain []descriptions.Domain) []DescriptionResponse {
	var list []DescriptionResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}
