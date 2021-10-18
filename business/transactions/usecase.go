package transactions

import (
	"alterra/business"
	"context"
	"time"
)

type TransactionUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewTransactionUsecase(repo Repository, timeout time.Duration) *TransactionUsecase {
	return &TransactionUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *TransactionUsecase) InsertTransactions(ctx context.Context, domain *Domain) (Domain, error) {
	transaction, err := uc.Repo.InsertTransactions(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}

func (uc *TransactionUsecase) GetListTransactions(ctx context.Context, Method_Payment_Id uint, User_Id uint, Karyawan_Id uint) ([]Domain, error) {
	transaction, err := uc.Repo.GetListTransactions(ctx, 0, 0, 0)
	if err != nil {
		return []Domain{}, err
	}
	return transaction, nil
}

func (uc *TransactionUsecase) GetById(ctx context.Context, id uint) (Domain, error) {
	transaction, err := uc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if transaction.Id == 0 {
		return Domain{}, business.ErrIDNotFound
	}
	return transaction, nil
}

func (uc *TransactionUsecase) Update(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	transaction, err := uc.Repo.Update(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}
	return transaction, nil
}

func (uc *TransactionUsecase) Delete(ctx context.Context, id uint) error {
	err := uc.Repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
