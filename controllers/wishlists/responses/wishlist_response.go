package responses

import (
	"alterra/business/wishlists"
	"time"
)

type WishlistResponse struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdat "`
	UpdatedAt time.Time `json:"updateat "`
}

type SearchResponse struct {
	Wishlist interface{}
}

func FromDomain(domain wishlists.Domain) WishlistResponse {
	return WishlistResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Title:     domain.Name,
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
