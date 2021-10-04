package payment_methods

import (
	"context"
	"time"
)

type Domain struct {
	Id        uint
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	InsertPayment_Method(ctx context.Context, domain Domain) (Domain, error)
	GetListPayment_Method(ctx context.Context, search string) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	InsertPayment_Method(ctx context.Context, domain Domain) (Domain, error)
	GetListPayment_Method(ctx context.Context, search string) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}
