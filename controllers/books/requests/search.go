package requests

import (
	"alterra/business/books"
)

type BookSearch struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func ToDomain(search BookSearch) books.Domain {
	return books.Domain{
		Title:     search.Title,
		Author:    search.Author,
		Publisher: search.Publisher,
	}
}
