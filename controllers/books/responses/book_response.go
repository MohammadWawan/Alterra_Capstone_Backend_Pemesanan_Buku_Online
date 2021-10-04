package responses

import (
	"alterra/business/books"
	"time"
)

type BookResponse struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Price     uint      `json:"price"`
	Author    string    `json:"author"`
	Publisher string    `json:"publisher"`
	CreatedAt time.Time `json:"createdat "`
	UpdatedAt time.Time `json:"updateat "`
}
type SearchResponse struct {
	Book interface{}
}

func FromDomain(domain books.Domain) BookResponse {
	return BookResponse{
		Id:        domain.Id,
		Title:     domain.Title,
		Price:     domain.Price,
		Author:    domain.Author,
		Publisher: domain.Publisher,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromBooksListDomain(domain []books.Domain) []BookResponse {
	var list []BookResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}
