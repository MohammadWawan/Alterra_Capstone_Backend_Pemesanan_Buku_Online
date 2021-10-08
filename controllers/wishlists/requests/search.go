package requests

import (
	"alterra/business/books"
	"alterra/business/users"
	"alterra/business/wishlists"
)

type WishlistSearch struct {
	User users.Domain `json:"user"`
	Book books.Domain `json:"book"`
}

func ToDomain(search WishlistSearch) wishlists.Domain {
	return wishlists.Domain{
		User: search.User,
		Book: search.Book,
	}
}
