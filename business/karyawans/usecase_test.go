package karyawans_test

import (
	"alterra/app/middlewares"
	"alterra/business/karyawans"
	"alterra/business/karyawans/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var karyawanRepository = mocks.Repository{Mock: mock.Mock{}}
var karyawanService karyawans.Usecase
var token string
var karyawanDomain karyawans.Domain
var listKaryawanDomain []karyawans.Domain

func setup() {
	karyawanService = karyawans.NewKaryawanUsecase(&karyawanRepository, time.Hour*10, &middlewares.ConfigJWT{})
	karyawanDomain = karyawans.Domain{
		Id:       1,
		Name:     "hasim",
		Email:    "Hasim@gmail.com",
		Password: "hasim1",
	}
	token = "token"
	listKaryawanDomain = append(listKaryawanDomain, karyawanDomain)
}

func TestRegister(t *testing.T) {
	setup()
	karyawanRepository.On("Register", mock.Anything, mock.Anything).Return(karyawanDomain, nil)
	karyawanRepository.On("GetByEmail", mock.Anything, mock.Anything).Return(karyawans.Domain{}, nil)
	t.Run("Test Case 1 | Success Register", func(t *testing.T) {
		karyawan, err := karyawanService.Register(context.Background(), karyawans.Domain{
			Id:       1,
			Name:     "hasim",
			Email:    "Hasim@gmail.com",
			Password: "hasim1",
		})

		assert.NoError(t, err)
		assert.Equal(t, karyawanDomain, karyawan)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Test Case 1 | Success Login", func(t *testing.T) {
		setup()
		karyawanRepository.On("GetByEmail",
			mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(karyawans.Domain{}, nil).Once()
		karyawan, token, err := karyawanService.Login(context.Background(), karyawanDomain.Email, "123")

		assert.NotNil(t, token)
		assert.NoError(t, err)
		assert.Equal(t, karyawan, karyawans.Domain{})
	})

	t.Run("Test Case 2 | Cannot Login (Password Not Found)", func(t *testing.T) {
		data, token, err := karyawanService.Login(context.Background(), karyawanDomain.Email, "")

		assert.Equal(t, karyawans.Domain{}, data)
		assert.Error(t, err)
		assert.Equal(t, token, "")
	})

	t.Run("Test Case 3 | Cannot Login (Wrong Auth)", func(t *testing.T) {
		setup()
		karyawanRepository.On("GetByEmail",
			mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(karyawans.Domain{}, errors.New("Users Not Found")).Once()
		data, token, err := karyawanService.Login(context.Background(), karyawanDomain.Email, "1234")

		assert.Equal(t, karyawans.Domain{}, data)
		assert.NoError(t, err)
		assert.NotEqual(t, token, "")
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Test case 1 | Success GetAllUsers", func(t *testing.T) {
		setup()
		karyawanRepository.On("GetAll", mock.Anything, mock.Anything).Return(listKaryawanDomain, nil).Once()
		data, err := karyawanService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listKaryawanDomain))
	})

	t.Run("Test case 2 | Error GetAllUsers", func(t *testing.T) {
		setup()
		karyawanRepository.On("GetAll", mock.Anything, mock.Anything).Return([]karyawans.Domain{}, errors.New("Users Not Found")).Once()
		data, err := karyawanService.GetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, data, []karyawans.Domain{})
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test case 1 | Success GetAllById", func(t *testing.T) {
		setup()
		karyawanRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(karyawanDomain, nil).Once()
		data, err := karyawanService.GetById(context.Background(), karyawanDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error GetAllById(karyawan Id = 0)", func(t *testing.T) {
		setup()
		karyawanDomain.Id = 0
		karyawanRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(karyawanDomain, nil).Once()
		data, err := karyawanService.GetById(context.Background(), karyawanDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, data, karyawans.Domain{})
	})

	t.Run("Test case 3 | Error GetAllById", func(t *testing.T) {
		setup()
		karyawanRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(karyawans.Domain{}, nil).Once()
		data, err := karyawanService.GetById(context.Background(), 7)

		assert.Error(t, err)
		assert.Equal(t, data, karyawans.Domain{})
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test case 1 | Success Update", func(t *testing.T) {
		setup()
		karyawanRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(karyawanDomain, nil).Once()
		data, err := karyawanService.Update(context.Background(), karyawanDomain, karyawanDomain.Id)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update", func(t *testing.T) {
		setup()
		karyawanRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(karyawanDomain, errors.New("Users Not Found")).Once()
		data, err := karyawanService.Update(context.Background(), karyawanDomain, karyawanDomain.Id)

		assert.Equal(t, data, karyawans.Domain{})
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test case 1 | Success Delete", func(t *testing.T) {
		setup()
		karyawanRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		err := karyawanService.Delete(context.Background(), karyawanDomain.Id)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete", func(t *testing.T) {
		setup()
		karyawanRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.New("Users Not Found")).Once()
		err := karyawanService.Delete(context.Background(), karyawanDomain.Id)

		assert.Equal(t, err, errors.New("Users Not Found"))
		assert.Error(t, err)
	})
}
