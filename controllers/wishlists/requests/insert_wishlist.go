package requests

import (
	"alterra/business/books"
	"alterra/business/users"
	"alterra/business/wishlists"
)

type InsertWishlist struct {
	User users.Domain `json:"user"`
	Book books.Domain `json:"book"`
}

func (wishlist *InsertWishlist) ToDomain() *wishlists.Domain {
	return &wishlists.Domain{
		User: wishlist.User,
		Book: wishlist.Book,
	}
}
