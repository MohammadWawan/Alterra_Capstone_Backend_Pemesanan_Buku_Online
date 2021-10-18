package requests

import (
	"alterra/business/wishlists"
)

type InsertWishlist struct {
	User_Id uint `json:"user_id"`
	Book_Id uint `json:"book_id"`
}

func (wishlist *InsertWishlist) ToDomain() *wishlists.Domain {
	return &wishlists.Domain{
		User_Id: wishlist.User_Id,
		Book_Id: wishlist.Book_Id,
	}
}
