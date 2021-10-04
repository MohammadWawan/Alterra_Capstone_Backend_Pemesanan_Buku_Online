package categories

import (
	"alterra/business"
	"context"
	"errors"
	"time"
)

type CategoryUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewCategoryUsecase(repo Repository, timeout time.Duration) *CategoryUsecase {
	return &CategoryUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *CategoryUsecase) InsertCategory(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name_Category == "" {
		return Domain{}, errors.New("category empty")
	}

	category, err := uc.Repo.InsertCategory(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return category, nil
}

func (uc *CategoryUsecase) GetListCategory(ctx context.Context, search string) ([]Domain, error) {
	category, err := uc.Repo.GetListCategory(ctx, search)
	if err != nil {
		return []Domain{}, err
	}

	return category, nil
}

func (uc *CategoryUsecase) GetById(ctx context.Context, id uint) (Domain, error) {
	category, err := uc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if category.Id == 0 {
		return Domain{}, business.ErrIDNotFound
	}
	return category, nil
}

func (uc *CategoryUsecase) Update(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	category, err := uc.Repo.Update(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}

	return category, nil
}

func (uc *CategoryUsecase) Delete(ctx context.Context, id uint) error {
	err := uc.Repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
