package transactions

import (
	"context"
	"time"
)

type Domain struct {
	Id                uint
	Method_Payment_Id uint
	User_Id           uint
	Karyawan_Id       uint
	Total_Qty         uint
	Total_Price       uint
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Usecase interface {
	InsertTransactions(ctx context.Context, domain Domain) (Domain, error)
	GetListTransactions(ctx context.Context, search string) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	InsertTransactions(ctx context.Context, domain Domain) (Domain, error)
	GetListTransactions(ctx context.Context, search string) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}
