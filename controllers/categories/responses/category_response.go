package responses

import (
	"alterra/business/categories"
	"time"
)

type CategoryResponse struct {
	Id        uint      `json:"id"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updateat"`
}

type SearchResponse struct {
	Category interface{}
}

func FromDomain(domain categories.Domain) CategoryResponse {
	return CategoryResponse{
		Id:        domain.Id,
		Category:  domain.Category,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromCategoriesListDomain(domain []categories.Domain) []CategoryResponse {
	var list []CategoryResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}
