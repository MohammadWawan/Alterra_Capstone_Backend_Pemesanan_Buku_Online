package transaction_details_test

import (
	"alterra/business/books"
	"alterra/business/karyawans"
	"alterra/business/payment_methods"
	"alterra/business/transaction_details"
	"alterra/business/transaction_details/mocks"
	"alterra/business/transactions"
	"alterra/business/users"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var transaction_detailRepository = mocks.Repository{Mock: mock.Mock{}}
var transaction_detailService transaction_details.Usecase
var transaction_detailDomain transaction_details.Domain
var listTransaction_DetailDomain []transaction_details.Domain

func setup() {
	transaction_detailService = transaction_details.NewTransaction_DetailUsecase(&transaction_detailRepository, time.Hour*10)
	transaction_detailDomain = transaction_details.Domain{
		Id: 1,
		Book: books.Domain{
			Id:        1,
			Title:     "Java itu Mudah",
			Price:     100000,
			Author:    "Wawan",
			Publisher: "Aksara jawi",
		},
		Transaction: transactions.Domain{
			Id:                1,
			User_Id:           1,
			Method_Payment_Id: 1,
			Karyawan_Id:       1,
			User: users.Domain{
				Id:      1,
				Name:    "wawan",
				Address: "malang",
			},
			Payment_Method: payment_methods.Domain{
				Id:   1,
				Type: "Debit BRI: 1247 0100 5834 534",
			},
			Karyawan: karyawans.Domain{
				Id:   1,
				Name: "jamal",
			},
			Total_Qty:   1,
			Total_Price: 100000,
		},
		Qty:   1,
		Price: 100000,
	}
	listTransaction_DetailDomain = append(listTransaction_DetailDomain, transaction_detailDomain)
}

func TestInsertTransaction_Detail(t *testing.T) {
	setup()
	transaction_detailRepository.On("InsertTransaction_Details", mock.Anything, mock.Anything).Return(transaction_detailDomain, nil)
	t.Run("Test Case 1 | Success Insert Transaction_Detail", func(t *testing.T) {
		transaction_detail, err := transaction_detailService.InsertTransaction_Details(context.Background(), transaction_details.Domain{
			Id: 1,
			Book: books.Domain{
				Id:        1,
				Title:     "Java itu Mudah",
				Price:     100000,
				Author:    "Wawan",
				Publisher: "Aksara jawi",
			},
			Transaction: transactions.Domain{
				Id:                1,
				User_Id:           1,
				Method_Payment_Id: 1,
				Karyawan_Id:       1,
				User: users.Domain{
					Id:      1,
					Name:    "wawan",
					Address: "malang",
				},
				Payment_Method: payment_methods.Domain{
					Id:   1,
					Type: "Debit BRI: 1247 0100 5834 534",
				},
				Karyawan: karyawans.Domain{
					Id:   1,
					Name: "jamal",
				},
				Total_Qty:   1,
				Total_Price: 100000,
			},
			Qty:   1,
			Price: 100000,
		})

		assert.NoError(t, err)
		assert.Equal(t, transaction_detailDomain, transaction_detail)
	})
}

func TestGetListTransaction_Details(t *testing.T) {
	t.Run("Test case 1 | Success GetListTransaction_Details", func(t *testing.T) {
		setup()
		transaction_detailRepository.On("GetListTransaction_Details", mock.Anything, mock.Anything).Return(listTransaction_DetailDomain, nil).Once()
		data, err := transaction_detailService.GetListTransaction_Details(context.Background(), transaction_detailDomain.Book.Title)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listTransaction_DetailDomain))
	})

	t.Run("Test case 2 | Error GetListTransaction_Details", func(t *testing.T) {
		setup()
		transaction_detailRepository.On("GetListTransaction_Details", mock.Anything, mock.Anything).Return([]transaction_details.Domain{}, errors.New("Transaction_Details Not Found")).Once()
		data, err := transaction_detailService.GetListTransaction_Details(context.Background(), "")

		assert.Error(t, err)
		assert.Equal(t, data, []transaction_details.Domain{})
	})
}

func TestSearchTransaction_DetailById(t *testing.T) {
	t.Run("Test case 1 | Success GetListTransaction_DetailById", func(t *testing.T) {
		setup()
		transaction_detailRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(transaction_detailDomain, nil).Once()
		data, err := transaction_detailService.GetById(context.Background(), transaction_detailDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error GetListTransaction_DetailById(transaction Id = 0)", func(t *testing.T) {
		setup()
		transaction_detailDomain.Id = 0
		transaction_detailRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(transaction_detailDomain, nil).Once()
		data, err := transaction_detailService.GetById(context.Background(), transaction_detailDomain.Id)

		assert.Error(t, err)
		assert.NotEqual(t, data, transactions.Domain{})
	})

	t.Run("Test case 3 | Error GetListTransaction_DetailById", func(t *testing.T) {
		setup()
		transaction_detailRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(transaction_details.Domain{}, nil).Once()
		data, err := transaction_detailService.GetById(context.Background(), 7)

		assert.Error(t, err)
		assert.Equal(t, data, transaction_details.Domain{})
	})
}

func TestUpdateTransaction_Detail(t *testing.T) {
	t.Run("Test case 1 | Success Update Transaction_Detail", func(t *testing.T) {
		setup()
		transaction_detailRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(transaction_detailDomain, nil).Once()
		data, err := transaction_detailService.Update(context.Background(), transaction_detailDomain, transaction_detailDomain.Id)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update Transaction_Detail", func(t *testing.T) {
		setup()
		transaction_detailRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(transaction_detailDomain, errors.New("transaction_details Not Found")).Once()
		data, err := transaction_detailService.Update(context.Background(), transaction_detailDomain, transaction_detailDomain.Id)

		assert.NotEqual(t, data, transactions.Domain{})
		assert.Error(t, err)
	})
}

func TestDeleteTransaction_Detail(t *testing.T) {
	t.Run("Test case 1 | Success Delete Transaction_Detail", func(t *testing.T) {
		setup()
		transaction_detailRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		err := transaction_detailService.Delete(context.Background(), transaction_detailDomain.Id)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete transaction_detail", func(t *testing.T) {
		setup()
		transaction_detailRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.New("Transaction_details  Not Found")).Once()
		err := transaction_detailService.Delete(context.Background(), transaction_detailDomain.Id)

		assert.NotEqual(t, err, errors.New("Transaction_details Not Found"))
		assert.Error(t, err)
	})
}
