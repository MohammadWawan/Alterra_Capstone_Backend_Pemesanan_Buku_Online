package wishlist

import (
	"alterra/business/wishlists"
	"alterra/drivers/databases/books"
	"alterra/drivers/databases/users"
	"time"

	"gorm.io/gorm"
)

type Wishlist struct {
	Id        uint `gorm:"primaryKey"`
	User_Id   uint
	Book_Id   uint
	Name      users.Users `gorm:"foreignKey:user_id"`
	Title     books.Books `gorm:"many2many:wishlist_books;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomain(domain wishlists.Domain) Wishlist {
	return Wishlist{
		Id:        domain.Id,
		Name:      users.Users{},
		Title:     books.Books{},
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func (wishlist *Wishlist) ToDomain() wishlists.Domain {
	return wishlists.Domain{
		Id:        wishlist.Id,
		Name:      wishlist.Name.Name,
		Title:     wishlist.Title.Title,
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
