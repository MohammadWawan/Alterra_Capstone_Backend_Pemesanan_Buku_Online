package wishlists

import (
	"alterra/business/wishlists"
	"time"

	"gorm.io/gorm"
)

type Wishlist struct {
	Id        uint `gorm:"primaryKey"`
	User_Id   uint
	Book_Id   uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomain(domain wishlists.Domain) Wishlist {
	return Wishlist{
		Id:        domain.Id,
		User_Id:   domain.User_Id,
		Book_Id:   domain.Book_Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func (wishlist *Wishlist) ToDomain() wishlists.Domain {
	return wishlists.Domain{
		Id:        wishlist.Id,
		User_Id:   wishlist.User_Id,
		Book_Id:   wishlist.Book_Id,
		CreatedAt: wishlist.CreatedAt,
		UpdatedAt: wishlist.UpdatedAt,
	}
}

func ToListDomain(data []Wishlist) []wishlists.Domain {
	list := []wishlists.Domain{}
	for _, v := range data {
		list = append(list, v.ToDomain())
	}

	return list
}
