package wishlists

import (
	"alterra/business"
	"context"
	"time"
)

type WishlistUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewWishlistUsecase(repo Repository, timeout time.Duration) *WishlistUsecase {
	return &WishlistUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *WishlistUsecase) InsertWishlist(ctx context.Context, domain *Domain) (Domain, error) {
	wishlist, err := uc.Repo.InsertWishlist(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return wishlist, nil
}

func (uc *WishlistUsecase) GetListWishlist(ctx context.Context, User_Id uint, Book_Id uint) ([]Domain, error) {
	wishlist, err := uc.Repo.GetListWishlist(ctx, 0, 0)
	if err != nil {
		return []Domain{}, err
	}

	return wishlist, nil
}

func (uc *WishlistUsecase) GetById(ctx context.Context, id uint) (Domain, error) {
	wishlist, err := uc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if wishlist.Id == 0 {
		return Domain{}, business.ErrIDNotFound
	}
	return wishlist, nil
}

func (uc *WishlistUsecase) Update(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	wishlist, err := uc.Repo.Update(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}

	return wishlist, nil
}

func (uc *WishlistUsecase) Delete(ctx context.Context, id uint) error {
	err := uc.Repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
