package books

import (
	"alterra/business"
	"context"
	"errors"
	"time"
)

type BookUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewBookUsecase(repo Repository, timeout time.Duration) *BookUsecase {
	return &BookUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *BookUsecase) InsertBook(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Title == "" {
		return Domain{}, errors.New("title empty")
	}

	if domain.Author == "" {
		return Domain{}, errors.New("author empty")
	}
	if domain.Publisher == "" {
		return Domain{}, errors.New("publisher empty")
	}
	book, err := uc.Repo.InsertBook(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return book, nil
}

func (uc *BookUsecase) GetListBook(ctx context.Context, search string) ([]Domain, error) {
	book, err := uc.Repo.GetListBook(ctx, search)
	if err != nil {
		return []Domain{}, err
	}

	return book, nil
}

func (uc *BookUsecase) GetById(ctx context.Context, id uint) (Domain, error) {
	book, err := uc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if book.Id == 0 {
		return Domain{}, business.ErrIDNotFound
	}
	return book, nil
}

func (uc *BookUsecase) Update(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	book, err := uc.Repo.Update(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}

	return book, nil
}

func (uc *BookUsecase) Delete(ctx context.Context, id uint) error {
	err := uc.Repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
