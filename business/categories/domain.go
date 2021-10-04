package categories

import (
	"context"
	"time"
)

type Domain struct {
	Id            uint
	Name_Category string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Usecase interface {
	InsertCategory(ctx context.Context, domain Domain) (Domain, error)
	GetListCategory(ctx context.Context, search string) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	InsertCategory(ctx context.Context, domain Domain) (Domain, error)
	GetListCategory(ctx context.Context, search string) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}
