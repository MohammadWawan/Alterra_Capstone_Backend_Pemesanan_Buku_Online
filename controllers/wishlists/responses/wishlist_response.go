package responses

import (
	"alterra/business/books"
	"alterra/business/users"
	"alterra/business/wishlists"
	"time"
)

type WishlistResponse struct {
	Id        uint         `json:"id"`
	User      users.Domain `json:"user"`
	Book      books.Domain `json:"book"`
	CreatedAt time.Time    `json:"createdat "`
	UpdatedAt time.Time    `json:"updateat "`
}

type SearchResponse struct {
	Wishlist interface{}
}

func FromDomain(domain wishlists.Domain) WishlistResponse {
	return WishlistResponse{
		Id:        domain.Id,
		User:      domain.User,
		Book:      domain.Book,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromWishlistListDomain(domain []wishlists.Domain) []WishlistResponse {
	var list []WishlistResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}
