package wishlists

import (
	"context"
	"time"
)

type Domain struct {
	Id        uint
	User_Id   uint
	Book_Id   uint
	Name      string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	InsertWishlist(ctx context.Context, domain Domain) (Domain, error)
	GetListWishlist(ctx context.Context, search string) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	InsertWishlist(ctx context.Context, domain Domain) (Domain, error)
	GetListWishlist(ctx context.Context, search string) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}
