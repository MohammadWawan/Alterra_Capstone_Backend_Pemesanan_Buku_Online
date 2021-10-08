package transaction_details

import (
	"alterra/business/books"
	"alterra/business/transactions"
	"context"
	"time"
)

type Domain struct {
	Id             uint
	Book_Id        uint
	Transaction_Id uint
	Book           books.Domain
	Transaction    transactions.Domain
	Qty            uint
	Price          uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Usecase interface {
	InsertTransaction_Details(ctx context.Context, domain Domain) (Domain, error)
	GetListTransaction_Details(ctx context.Context, search string) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	InsertTransaction_Details(ctx context.Context, domain Domain) (Domain, error)
	GetListTransaction_Details(ctx context.Context, search string) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}
