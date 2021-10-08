package requests

import (
	"alterra/business/books"
	"alterra/business/categories"
	"alterra/business/descriptions"
)

type InsertBook struct {
	Title       string              `json:"title"`
	Price       uint                `json:"price"`
	Author      string              `json:"author"`
	Publisher   string              `json:"publisher"`
	Category    categories.Domain   `json:"category"`
	Description descriptions.Domain `json:"description"`
}

func (book *InsertBook) ToDomain() *books.Domain {
	return &books.Domain{
		Title:       book.Title,
		Price:       book.Price,
		Author:      book.Author,
		Publisher:   book.Publisher,
		Category:    categories.Domain{},
		Description: descriptions.Domain{},
	}
}
