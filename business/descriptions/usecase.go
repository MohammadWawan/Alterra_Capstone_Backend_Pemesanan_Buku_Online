package descriptions

import (
	"alterra/business"
	"context"
	"errors"
	"time"
)

type DescriptionUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewDescriptionUsecase(repo Repository, timeout time.Duration) *DescriptionUsecase {
	return &DescriptionUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *DescriptionUsecase) InsertDescription(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Description == "" {
		return Domain{}, errors.New("description empty")
	}

	description, err := uc.Repo.InsertDescription(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return description, nil
}

func (uc *DescriptionUsecase) GetListDescription(ctx context.Context, search string) ([]Domain, error) {
	description, err := uc.Repo.GetListDescription(ctx, search)
	if err != nil {
		return []Domain{}, err
	}
	return description, nil
}

func (uc *DescriptionUsecase) GetById(ctx context.Context, id uint) (Domain, error) {
	description, err := uc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if description.Id == 0 {
		return Domain{}, business.ErrIDNotFound
	}
	return description, nil
}

func (uc *DescriptionUsecase) Update(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	description, err := uc.Repo.Update(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}
	return description, nil
}

func (uc *DescriptionUsecase) Delete(ctx context.Context, id uint) error {
	err := uc.Repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
