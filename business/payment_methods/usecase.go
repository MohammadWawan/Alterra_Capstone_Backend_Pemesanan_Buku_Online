package payment_methods

import (
	"alterra/business"
	"context"
	"errors"
	"time"
)

type Payment_MethodUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewPayment_MethodUsecase(repo Repository, timeout time.Duration) *Payment_MethodUsecase {
	return &Payment_MethodUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *Payment_MethodUsecase) InsertPayment_Method(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Type == "" {
		return Domain{}, errors.New("method payment empty")
	}

	payment_method, err := uc.Repo.InsertPayment_Method(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return payment_method, nil
}

func (uc *Payment_MethodUsecase) GetListPayment_Method(ctx context.Context, search string) ([]Domain, error) {
	payment_method, err := uc.Repo.GetListPayment_Method(ctx, search)
	if err != nil {
		return []Domain{}, err
	}
	return payment_method, nil
}

func (uc *Payment_MethodUsecase) GetById(ctx context.Context, id uint) (Domain, error) {
	payment_method, err := uc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if payment_method.Id == 0 {
		return Domain{}, business.ErrIDNotFound
	}
	return payment_method, nil
}

func (uc *Payment_MethodUsecase) Update(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	payment_method, err := uc.Repo.Update(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}
	return payment_method, nil
}

func (uc *Payment_MethodUsecase) Delete(ctx context.Context, id uint) error {
	err := uc.Repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
