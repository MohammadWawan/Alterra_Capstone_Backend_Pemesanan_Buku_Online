package requests

import "alterra/business/wishlists"

type InsertWishlist struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

func (wishlist *InsertWishlist) ToDomain() *wishlists.Domain {
	return &wishlists.Domain{
		Name:  wishlist.Name,
		Title: wishlist.Title,
	}
}
