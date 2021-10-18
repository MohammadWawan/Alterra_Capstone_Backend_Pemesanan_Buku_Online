package requests

import (
	"alterra/business/wishlists"
)

type WishlistSearch struct {
	User_Id uint `json:"user_id"`
	Book_Id uint `json:"book_id"`
}

func ToDomain(search WishlistSearch) wishlists.Domain {
	return wishlists.Domain{
		User_Id: search.User_Id,
		Book_Id: search.Book_Id,
	}
}
