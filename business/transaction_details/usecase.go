package transaction_details

import (
	"alterra/business"
	"context"
	"time"
)

type Transaction_DetailUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewTransaction_DetailUsecase(repo Repository, timeout time.Duration) *Transaction_DetailUsecase {
	return &Transaction_DetailUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *Transaction_DetailUsecase) InsertTransaction_Details(ctx context.Context, domain *Domain) (Domain, error) {
	transaction_detail, err := uc.Repo.InsertTransaction_Details(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return transaction_detail, nil
}

func (uc *Transaction_DetailUsecase) GetListTransaction_Details(ctx context.Context, Book_Id uint, Transaction_Id uint) ([]Domain, error) {
	transaction_detail, err := uc.Repo.GetListTransaction_Details(ctx, Book_Id, Transaction_Id)
	if err != nil {
		return []Domain{}, err
	}
	return transaction_detail, nil
}

func (uc *Transaction_DetailUsecase) GetById(ctx context.Context, id uint) (Domain, error) {
	transaction_detail, err := uc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if transaction_detail.Id == 0 {
		return Domain{}, business.ErrIDNotFound
	}
	return transaction_detail, nil
}

func (uc *Transaction_DetailUsecase) Update(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	transaction_detail, err := uc.Repo.Update(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}
	return transaction_detail, nil
}

func (uc *Transaction_DetailUsecase) Delete(ctx context.Context, id uint) error {
	err := uc.Repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
