package requests

import (
	"alterra/business/books"
)

type InsertBook struct {
	Title          string `json:"title"`
	Price          uint   `json:"price"`
	Author         string `json:"author"`
	Publisher      string `json:"publisher"`
	Category_Id    uint   `json:"category_id"`
	Description_Id uint   `json:"description_id"`
}

func (book *InsertBook) ToDomain() *books.Domain {
	return &books.Domain{
		Title:          book.Title,
		Price:          book.Price,
		Author:         book.Author,
		Publisher:      book.Publisher,
		Category_Id:    book.Category_Id,
		Description_Id: book.Category_Id,
	}
}
