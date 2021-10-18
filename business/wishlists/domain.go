package wishlists

import (
	"alterra/business/books"
	"alterra/business/users"
	"context"
	"time"
)

type Domain struct {
	Id        uint
	User_Id   uint
	Book_Id   uint
	User      users.Domain
	Book      books.Domain
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	InsertWishlist(ctx context.Context, domain *Domain) (Domain, error)
	GetListWishlist(ctx context.Context, User_Id uint, Book_Id uint) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	InsertWishlist(ctx context.Context, domain *Domain) (Domain, error)
	GetListWishlist(ctx context.Context, User_Id uint, Book_Id uint) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}
