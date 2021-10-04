package requests

import "alterra/business/wishlists"

type WishlistSearch struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

func ToDomain(search WishlistSearch) wishlists.Domain {
	return wishlists.Domain{
		Name:  search.Name,
		Title: search.Title,
	}
}
