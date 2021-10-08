package books_test

import (
	"alterra/business/books"
	"alterra/business/books/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var bookRepository = mocks.Repository{Mock: mock.Mock{}}
var bookService books.Usecase
var bookDomain books.Domain
var listBookDomain []books.Domain

func setup() {
	bookService = books.NewBookUsecase(&bookRepository, time.Hour*10)
	bookDomain = books.Domain{
		Id:        1,
		Title:     "Java itu Mudah",
		Price:     100000,
		Author:    "Wawan",
		Publisher: "Aksara jawi",
	}
	listBookDomain = append(listBookDomain, bookDomain)
}

func TestInsertBook(t *testing.T) {
	setup()
	bookRepository.On("InsertBook", mock.Anything, mock.Anything).Return(bookDomain, nil)
	t.Run("Test Case 1 | Success Insert Book", func(t *testing.T) {
		category, err := bookService.InsertBook(context.Background(), books.Domain{
			Id:        1,
			Title:     "Java itu Mudah",
			Price:     100000,
			Author:    "Wawan",
			Publisher: "Aksara jawi",
		})

		assert.NoError(t, err)
		assert.Equal(t, bookDomain, category)
	})
}

func TestSearchBook(t *testing.T) {
	t.Run("Test case 1 | Success SearchBook(by title)", func(t *testing.T) {
		setup()
		bookRepository.On("GetListBook", mock.Anything, mock.Anything).Return(listBookDomain, nil).Once()
		data, err := bookService.GetListBook(context.Background(), bookDomain.Title)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listBookDomain))
	})

	t.Run("Test case 2 | Success SearchBook(by author)", func(t *testing.T) {
		setup()
		bookRepository.On("GetListBook", mock.Anything, mock.Anything).Return(listBookDomain, nil).Once()
		data, err := bookService.GetListBook(context.Background(), bookDomain.Author)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listBookDomain))
	})

	t.Run("Test case 3 | Success SearchBook(by publisher)", func(t *testing.T) {
		setup()
		bookRepository.On("GetListBook", mock.Anything, mock.Anything).Return(listBookDomain, nil).Once()
		data, err := bookService.GetListBook(context.Background(), bookDomain.Publisher)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listBookDomain))
	})

	t.Run("Test case 4 | Error SearchCategories(search empty)", func(t *testing.T) {
		setup()
		bookRepository.On("GetListBook", mock.Anything, mock.Anything).Return([]books.Domain{}, errors.New("Books Not Found")).Once()
		data, err := bookService.GetListBook(context.Background(), "")

		assert.Error(t, err)
		assert.Equal(t, data, []books.Domain{})
	})
}

func TestSearchCategoryById(t *testing.T) {
	t.Run("Test case 1 | Success SearchBookById", func(t *testing.T) {
		setup()
		bookRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(bookDomain, nil).Once()
		data, err := bookService.GetById(context.Background(), bookDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error SearchBookById(book Id = 0)", func(t *testing.T) {
		setup()
		bookDomain.Id = 0
		bookRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(bookDomain, nil).Once()
		data, err := bookService.GetById(context.Background(), bookDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, data, books.Domain{})
	})

	t.Run("Test case 3 | Error SearchBookById", func(t *testing.T) {
		setup()
		bookRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(books.Domain{}, nil).Once()
		data, err := bookService.GetById(context.Background(), 7)

		assert.Error(t, err)
		assert.Equal(t, data, books.Domain{})
	})
}

func TestUpdateBook(t *testing.T) {
	t.Run("Test case 1 | Success Update Book", func(t *testing.T) {
		setup()
		bookRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(bookDomain, nil).Once()
		data, err := bookService.Update(context.Background(), bookDomain, bookDomain.Id)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update Book", func(t *testing.T) {
		setup()
		bookRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(bookDomain, errors.New("Books Not Found")).Once()
		data, err := bookService.Update(context.Background(), bookDomain, bookDomain.Id)

		assert.Equal(t, data, books.Domain{})
		assert.Error(t, err)
	})
}

func TestDeleteBook(t *testing.T) {
	t.Run("Test case 1 | Success Delete Book", func(t *testing.T) {
		setup()
		bookRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		err := bookService.Delete(context.Background(), bookDomain.Id)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete Category", func(t *testing.T) {
		setup()
		bookRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.New("Books  Not Found")).Once()
		err := bookService.Delete(context.Background(), bookDomain.Id)

		assert.NotEqual(t, err, errors.New("Books Not Found"))
		assert.Error(t, err)
	})
}
