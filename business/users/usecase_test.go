package users_test

import (
	"alterra/app/middlewares"
	"alterra/business/users"
	"alterra/business/users/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = mocks.Repository{Mock: mock.Mock{}}
var userService users.Usecase
var token string
var userDomain users.Domain
var listUserDomain []users.Domain

func setup() {
	userService = users.NewUserUsecase(&userRepository, time.Hour*10, &middlewares.ConfigJWT{})
	userDomain = users.Domain{
		Id:       1,
		Name:     "wawan",
		Email:    "Wawan@gmail.com",
		Address:  "pasuruan",
		Password: "wawen1",
	}
	token = "token"
	listUserDomain = append(listUserDomain, userDomain)
}

func TestRegister(t *testing.T) {
	setup()
	userRepository.On("Register", mock.Anything, mock.Anything).Return(userDomain, nil)
	userRepository.On("GetByEmail", mock.Anything, mock.Anything).Return(users.Domain{}, nil)
	t.Run("Test Case 1 | Success Register", func(t *testing.T) {
		user, err := userService.Register(context.Background(), users.Domain{
			Id:       1,
			Name:     "wawan",
			Email:    "Wawan@gmail.com",
			Address:  "pasuruan",
			Password: "wawen1",
		})

		assert.NoError(t, err)
		assert.Equal(t, userDomain, user)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Test Case 1 | Success Login", func(t *testing.T) {
		setup()
		userRepository.On("GetByEmail",
			mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(users.Domain{}, nil).Once()
		user, token, err := userService.Login(context.Background(), userDomain.Email, "123")

		assert.NotNil(t, token)
		assert.NoError(t, err)
		assert.Equal(t, user, users.Domain{})
	})

	t.Run("Test Case 2 | Cannot Login (Password Not Found)", func(t *testing.T) {
		data, token, err := userService.Login(context.Background(), userDomain.Email, "")

		assert.Equal(t, users.Domain{}, data)
		assert.Error(t, err)
		assert.Equal(t, token, "")
	})

	t.Run("Test Case 3 | Cannot Login (Wrong Auth)", func(t *testing.T) {
		setup()
		userRepository.On("GetByEmail",
			mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(users.Domain{}, errors.New("Users Not Found")).Once()
		data, token, err := userService.Login(context.Background(), userDomain.Email, "1234")

		assert.Equal(t, users.Domain{}, data)
		assert.NoError(t, err)
		assert.NotEqual(t, token, "")
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Test case 1 | Success GetAllUsers", func(t *testing.T) {
		setup()
		userRepository.On("GetAll", mock.Anything, mock.Anything).Return(listUserDomain, nil).Once()
		data, err := userService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listUserDomain))
	})

	t.Run("Test case 2 | Error GetAllUsers", func(t *testing.T) {
		setup()
		userRepository.On("GetAll", mock.Anything, mock.Anything).Return([]users.Domain{}, errors.New("Users Not Found")).Once()
		data, err := userService.GetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, data, []users.Domain{})
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test case 1 | Success GetAllById", func(t *testing.T) {
		setup()
		userRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(userDomain, nil).Once()
		data, err := userService.GetById(context.Background(), userDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error GetAllById(user Id = 0)", func(t *testing.T) {
		setup()
		userDomain.Id = 0
		userRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(userDomain, nil).Once()
		data, err := userService.GetById(context.Background(), userDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, data, users.Domain{})
	})

	t.Run("Test case 3 | Error GetAllById", func(t *testing.T) {
		setup()
		userRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(users.Domain{}, nil).Once()
		data, err := userService.GetById(context.Background(), 7)

		assert.Error(t, err)
		assert.Equal(t, data, users.Domain{})
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test case 1 | Success Update", func(t *testing.T) {
		setup()
		userRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(userDomain, nil).Once()
		data, err := userService.Update(context.Background(), userDomain, userDomain.Id)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update", func(t *testing.T) {
		setup()
		userRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(userDomain, errors.New("Users Not Found")).Once()
		data, err := userService.Update(context.Background(), userDomain, userDomain.Id)

		assert.Equal(t, data, users.Domain{})
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test case 1 | Success Delete", func(t *testing.T) {
		setup()
		userRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		err := userService.Delete(context.Background(), userDomain.Id)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete", func(t *testing.T) {
		setup()
		userRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.New("Users Not Found")).Once()
		err := userService.Delete(context.Background(), userDomain.Id)

		assert.Equal(t, err, errors.New("Users Not Found"))
		assert.Error(t, err)
	})
}
