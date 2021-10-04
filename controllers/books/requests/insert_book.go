package requests

import "alterra/business/books"

type InsertBook struct {
	Title     string `json:"title"`
	Price     uint   `json:"price"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func (user *InsertBook) ToDomain() *books.Domain {
	return &books.Domain{
		Title:     user.Title,
		Price:     user.Price,
		Author:    user.Author,
		Publisher: user.Publisher,
	}
}
