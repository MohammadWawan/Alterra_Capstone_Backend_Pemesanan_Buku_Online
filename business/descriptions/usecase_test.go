package descriptions_test

import (
	"alterra/business/descriptions"
	"alterra/business/descriptions/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var descriptionRepository = mocks.Repository{Mock: mock.Mock{}}
var descriptionService descriptions.Usecase
var descriptionDomain descriptions.Domain
var listDescriptionDomain []descriptions.Domain

func setup() {
	descriptionService = descriptions.NewDescriptionUsecase(&descriptionRepository, time.Hour*10)
	descriptionDomain = descriptions.Domain{
		Id:          1,
		Description: "ini adalah buku tentang bahasa java",
	}
	listDescriptionDomain = append(listDescriptionDomain, descriptionDomain)
}

func TestInsertDescription(t *testing.T) {
	setup()
	descriptionRepository.On("InsertDescription", mock.Anything, mock.Anything).Return(descriptionDomain, nil)
	t.Run("Test Case 1 | Success Insert Description", func(t *testing.T) {
		description, err := descriptionService.InsertDescription(context.Background(), descriptions.Domain{
			Id:          1,
			Description: "ini adalah buku tentang bahasa java",
		})

		assert.NoError(t, err)
		assert.Equal(t, descriptionDomain, description)
	})
}

func TestSearchDescription(t *testing.T) {
	t.Run("Test case 1 | Success SearchDescriptions", func(t *testing.T) {
		setup()
		descriptionRepository.On("GetListDescription", mock.Anything, mock.Anything).Return(listDescriptionDomain, nil).Once()
		data, err := descriptionService.GetListDescription(context.Background(), descriptionDomain.Description)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listDescriptionDomain))
	})

	t.Run("Test case 2 | Error SearchDescriptions(search empty)", func(t *testing.T) {
		setup()
		descriptionRepository.On("GetListDescription", mock.Anything, mock.Anything).Return([]descriptions.Domain{}, errors.New("Descriptions Not Found")).Once()
		data, err := descriptionService.GetListDescription(context.Background(), "")

		assert.Error(t, err)
		assert.Equal(t, data, []descriptions.Domain{})
	})
}

func TestSearchDescriptionById(t *testing.T) {
	t.Run("Test case 1 | Success SearchDescriptionById", func(t *testing.T) {
		setup()
		descriptionRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(descriptionDomain, nil).Once()
		data, err := descriptionService.GetById(context.Background(), descriptionDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error SearchDescriptionById(description Id = 0)", func(t *testing.T) {
		setup()
		descriptionDomain.Id = 0
		descriptionRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(descriptionDomain, nil).Once()
		data, err := descriptionService.GetById(context.Background(), descriptionDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, data, descriptions.Domain{})
	})

	t.Run("Test case 3 | Error SearchDescriptionById", func(t *testing.T) {
		setup()
		descriptionRepository.On("GetById", mock.Anything, mock.AnythingOfType("uint")).Return(descriptions.Domain{}, nil).Once()
		data, err := descriptionService.GetById(context.Background(), 7)

		assert.Error(t, err)
		assert.Equal(t, data, descriptions.Domain{})
	})
}

func TestUpdateDescription(t *testing.T) {
	t.Run("Test case 1 | Success Update Description", func(t *testing.T) {
		setup()
		descriptionRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(descriptionDomain, nil).Once()
		data, err := descriptionService.Update(context.Background(), descriptionDomain, descriptionDomain.Id)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update Description", func(t *testing.T) {
		setup()
		descriptionRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(descriptionDomain, errors.New("Descriptions Not Found")).Once()
		data, err := descriptionService.Update(context.Background(), descriptionDomain, descriptionDomain.Id)

		assert.Equal(t, data, descriptions.Domain{})
		assert.Error(t, err)
	})
}

func TestDeleteDescription(t *testing.T) {
	t.Run("Test case 1 | Success Delete Description", func(t *testing.T) {
		setup()
		descriptionRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		err := descriptionService.Delete(context.Background(), descriptionDomain.Id)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete Description", func(t *testing.T) {
		setup()
		descriptionRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.New("Descriptions  Not Found")).Once()
		err := descriptionService.Delete(context.Background(), descriptionDomain.Id)

		assert.NotEqual(t, err, errors.New("Descriptions Not Found"))
		assert.Error(t, err)
	})
}
