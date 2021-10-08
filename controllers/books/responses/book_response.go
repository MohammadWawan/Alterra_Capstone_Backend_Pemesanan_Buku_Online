package responses

import (
	"alterra/business/books"
	"alterra/business/categories"
	"alterra/business/descriptions"
	"time"
)

type BookResponse struct {
	Id          uint                `json:"id"`
	Title       string              `json:"title"`
	Price       uint                `json:"price"`
	Author      string              `json:"author"`
	Publisher   string              `json:"publisher"`
	Category    categories.Domain   `json:"category"`
	Description descriptions.Domain `json:"description"`
	CreatedAt   time.Time           `json:"createdat "`
	UpdatedAt   time.Time           `json:"updateat "`
}
type SearchResponse struct {
	Book interface{}
}

func FromDomain(domain books.Domain) BookResponse {
	return BookResponse{
		Id:          domain.Id,
		Title:       domain.Title,
		Price:       domain.Price,
		Author:      domain.Author,
		Publisher:   domain.Publisher,
		Category:    categories.Domain{},
		Description: descriptions.Domain{},
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func FromBooksListDomain(domain []books.Domain) []BookResponse {
	var list []BookResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}
