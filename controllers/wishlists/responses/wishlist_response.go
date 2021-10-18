package responses

import (
	"alterra/business/wishlists"
	"time"
)

type WishlistResponse struct {
	Id        uint      `json:"id"`
	User_Id   uint      `json:"user_id"`
	Book_Id   uint      `json:"book_id"`
	CreatedAt time.Time `json:"createdat "`
	UpdatedAt time.Time `json:"updateat "`
}

type SearchResponse struct {
	Wishlist interface{}
}

func FromDomain(domain wishlists.Domain) WishlistResponse {
	return WishlistResponse{
		Id:        domain.Id,
		User_Id:   domain.User_Id,
		Book_Id:   domain.Book_Id,
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
