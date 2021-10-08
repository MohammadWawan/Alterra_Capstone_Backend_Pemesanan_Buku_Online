package wishlists_test

import (
	"alterra/business/books"
	"alterra/business/users"
	"alterra/business/wishlists"
	"alterra/business/wishlists/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var wishlistRepository = mocks.Repository{Mock: mock.Mock{}}
var wishlistService wishlists.Usecase
var wishlistDomain wishlists.Domain
var listWishlistDomain []wishlists.Domain

func setup() {
	wishlistService = wishlists.NewWishlistUsecase(&wishlistRepository, time.Hour*10)
	wishlistDomain = wishlists.Domain{
		Id:      1,
		User_Id: 1,
		Book_Id: 1,
		User: users.Domain{
			Id:      1,
			Name:    "wawan",
			Address: "pasuruan",
		},
		Book: books.Domain{
			Id:        1,
			Title:     "Java itu Mudah",
			Price:     100000,
			Author:    "Wawan",
			Publisher: "Aksara jawi",
		},
	}
	listWishlistDomain = append(listWishlistDomain, wishlistDomain)
}

func TestInsertWishlist(t *testing.T) {
	setup()
	wishlistRepository.On("InsertWishlist", mock.Anything, mock.Anything).Return(wishlistDomain, nil)
	t.Run("Test Case 1 | Success Insert Wishlist", func(t *testing.T) {
		wishlist, err := wishlistService.InsertWishlist(context.Background(), wishlists.Domain{
			Id:      1,
			User_Id: 1,
			Book_Id: 1,
			User: users.Domain{
				Id:      1,
				Name:    "wawan",
				Address: "pasuruan",
			},
			Book: books.Domain{
				Id:        1,
				Title:     "Java itu Mudah",
				Price:     100000,
				Author:    "Wawan",
				Publisher: "Aksara jawi",
			},
		})

		assert.Error(t, err)
		assert.NotEqual(t, wishlistDomain, wishlist)
	})
}

func TestGetListWishlist(t *testing.T) {
	t.Run("Test case 1 | Success GetListWishlists", func(t *testing.T) {
		setup()
		wishlistRepository.On("GetListWishlist", mock.Anything, mock.Anything).Return(listWishlistDomain, nil).Once()
		data, err := wishlistService.GetListWishlist(context.Background(), wishlistDomain.Book.Title)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listWishlistDomain))
	})

	t.Run("Test case 2 | Error GetListWishlists", func(t *testing.T) {
		setup()
		wishlistRepository.On("GetListWishlist", mock.Anything, mock.Anything).Return([]wishlists.Domain{}, errors.New("Wishlists Not Found")).Once()
		data, err := wishlistService.GetListWishlist(context.Background(), "")

		assert.Error(t, err)
		assert.Equal(t, data, []wishlists.Domain{})
	})
}

func TestGetListWishlistById(t *testing.T) {
	t.Run("Test case 1 | Success GetListWishlistById", func(t *testing.T) {
		setup()
		wishlistRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(wishlistDomain, nil).Once()
		data, err := wishlistService.GetById(context.Background(), wishlistDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error GetListWishlistById(wishlist Id = 0)", func(t *testing.T) {
		setup()
		wishlistDomain.Id = 0
		wishlistRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(wishlistDomain, nil).Once()
		data, err := wishlistService.GetById(context.Background(), wishlistDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, data, wishlists.Domain{})
	})

	t.Run("Test case 3 | Error GetListWishlistById", func(t *testing.T) {
		setup()
		wishlistRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(wishlists.Domain{}, nil).Once()
		data, err := wishlistService.GetById(context.Background(), 7)

		assert.Error(t, err)
		assert.Equal(t, data, wishlists.Domain{})
	})
}

func TestUpdateWishlist(t *testing.T) {
	t.Run("Test case 1 | Success Update Wishlist", func(t *testing.T) {
		setup()
		wishlistRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(wishlistDomain, nil).Once()
		data, err := wishlistService.Update(context.Background(), wishlistDomain, wishlistDomain.Id)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update Wishlist", func(t *testing.T) {
		setup()
		wishlistRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(wishlistDomain, errors.New("Wishlists Not Found")).Once()
		data, err := wishlistService.Update(context.Background(), wishlistDomain, wishlistDomain.Id)

		assert.Equal(t, data, wishlists.Domain{})
		assert.Error(t, err)
	})
}

func TestDeleteWishlist(t *testing.T) {
	t.Run("Test case 1 | Success Delete Wishlist", func(t *testing.T) {
		setup()
		wishlistRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		err := wishlistService.Delete(context.Background(), wishlistDomain.Id)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete Wishlist", func(t *testing.T) {
		setup()
		wishlistRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.New("Wishlists Not Found")).Once()
		err := wishlistService.Delete(context.Background(), wishlistDomain.Id)

		assert.NotEqual(t, err, nil)
		assert.Error(t, err)
	})
}
