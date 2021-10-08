package books

import (
	"alterra/business/categories"
	"alterra/business/descriptions"
	"context"
	"time"
)

type Domain struct {
	Id             uint
	Category_Id    uint
	Description_Id uint
	Title          string
	Price          uint
	Author         string
	Publisher      string
	Category       categories.Domain
	Description    descriptions.Domain
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Usecase interface {
	InsertBook(ctx context.Context, domain Domain) (Domain, error)
	GetListBook(ctx context.Context, search string) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	InsertBook(ctx context.Context, domain Domain) (Domain, error)
	GetListBook(ctx context.Context, search string) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}
