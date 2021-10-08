package payment_methods_test

import (
	"alterra/business/payment_methods"
	"alterra/business/payment_methods/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var payment_methodRepository = mocks.Repository{Mock: mock.Mock{}}
var payment_methodService payment_methods.Usecase
var payment_methodDomain payment_methods.Domain
var listPayment_MethodDomain []payment_methods.Domain

func setup() {
	payment_methodService = payment_methods.NewPayment_MethodUsecase(&payment_methodRepository, time.Hour*10)
	payment_methodDomain = payment_methods.Domain{
		Id:   1,
		Type: "Debit BRI: 1247 0100 5834 534",
	}
	listPayment_MethodDomain = append(listPayment_MethodDomain, payment_methodDomain)
}

func TestInsertPayment_Method(t *testing.T) {
	setup()
	payment_methodRepository.On("InsertPayment_Method", mock.Anything, mock.Anything).Return(payment_methodDomain, nil)
	t.Run("Test Case 1 | Success Insert Payment_Method", func(t *testing.T) {
		payment_method, err := payment_methodService.InsertPayment_Method(context.Background(), payment_methods.Domain{
			Id:   1,
			Type: "Debit BRI: 1247 0100 5834 534",
		})

		assert.NoError(t, err)
		assert.Equal(t, payment_methodDomain, payment_method)
	})
}

func TestSearchPayment_Method(t *testing.T) {
	t.Run("Test case 1 | Success SearchPayment_Methods", func(t *testing.T) {
		setup()
		payment_methodRepository.On("GetListPayment_Method", mock.Anything, mock.Anything).Return(listPayment_MethodDomain, nil).Once()
		data, err := payment_methodService.GetListPayment_Method(context.Background(), payment_methodDomain.Type)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listPayment_MethodDomain))
	})

	t.Run("Test case 2 | Error SearchPayment_Methods(search empty)", func(t *testing.T) {
		setup()
		payment_methodRepository.On("GetListPayment_Method", mock.Anything, mock.Anything).Return([]payment_methods.Domain{}, errors.New("Payment_Methods Not Found")).Once()
		data, err := payment_methodService.GetListPayment_Method(context.Background(), "")

		assert.Error(t, err)
		assert.Equal(t, data, []payment_methods.Domain{})
	})
}

func TestSearchPayment_MethodById(t *testing.T) {
	t.Run("Test case 1 | Success SearchPayment_MethodById", func(t *testing.T) {
		setup()
		payment_methodRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(payment_methodDomain, nil).Once()
		data, err := payment_methodService.GetById(context.Background(), payment_methodDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error SearchPayment_MethodById(payment_method Id = 0)", func(t *testing.T) {
		setup()
		payment_methodDomain.Id = 0
		payment_methodRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(payment_methodDomain, nil).Once()
		data, err := payment_methodService.GetById(context.Background(), payment_methodDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, data, payment_methods.Domain{})
	})

	t.Run("Test case 3 | Error SearchPayment_MethodById", func(t *testing.T) {
		setup()
		payment_methodRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(payment_methods.Domain{}, nil).Once()
		data, err := payment_methodService.GetById(context.Background(), 7)

		assert.Error(t, err)
		assert.Equal(t, data, payment_methods.Domain{})
	})
}

func TestUpdatePayment_Method(t *testing.T) {
	t.Run("Test case 1 | Success Update Payment_Method", func(t *testing.T) {
		setup()
		payment_methodRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(payment_methodDomain, nil).Once()
		data, err := payment_methodService.Update(context.Background(), payment_methodDomain, payment_methodDomain.Id)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update Payment_Method", func(t *testing.T) {
		setup()
		payment_methodRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(payment_methodDomain, errors.New("Payment_Methods Not Found")).Once()
		data, err := payment_methodService.Update(context.Background(), payment_methodDomain, payment_methodDomain.Id)

		assert.Equal(t, data, payment_methods.Domain{})
		assert.Error(t, err)
	})
}

func TestDeletePayment_Method(t *testing.T) {
	t.Run("Test case 1 | Success Delete Payment_Method", func(t *testing.T) {
		setup()
		payment_methodRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		err := payment_methodService.Delete(context.Background(), payment_methodDomain.Id)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete Payment_Method", func(t *testing.T) {
		setup()
		payment_methodRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.New("Payment_Methods  Not Found")).Once()
		err := payment_methodService.Delete(context.Background(), payment_methodDomain.Id)

		assert.NotEqual(t, err, errors.New("Payment_Methods Not Found"))
		assert.Error(t, err)
	})
}
