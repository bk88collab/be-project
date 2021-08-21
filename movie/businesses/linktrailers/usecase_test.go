package linktrailers_test

import (
	"context"
	"movie/businesses"
	link "movie/businesses/linktrailers"
	linkMocks "movie/businesses/linktrailers/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	linkRepository linkMocks.Repository
	linkUsecase    link.UseCase
)

func setup() {
	linkUsecase = link.NewLinkUseCase(&linkRepository)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestCreate(t *testing.T) {
	t.Run("test case 1, valid id update", func(t *testing.T) {
		domain := link.Domain{
			Id_trailer: 1,
			Url:        "abcde",
		}
		linkRepository.On("GetLinkbyId", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()
		result := linkUsecase.Update(context.Background(), 1, &link.Domain{})
		assert.Nil(t, result)
		assert.Equal(t, result, nil)
	})

	t.Run("test case 2, valid get-all-link create", func(t *testing.T) {
		domain := link.Domain{
			Id_trailer: 1,
			Url:        "abcde",
		}
		linkRepository.On("GetAllLinkRepository", mock.Anything, mock.Anything).Return(domain, nil).Once()
		result, err := linkUsecase.Create(context.Background(), &link.Domain{})
		assert.Equal(t, result, link.Domain{})
		assert.Equal(t, err, businesses.ErrLinkExisted)
	})

	t.Run("test case 3, invalid get-all-link create", func(t *testing.T) {
		domain := link.Domain{
			Id_trailer: 1,
			Url:        "abcde",
		}
		linkRepository.On("GetAllLinkRepository", mock.Anything, mock.Anything).Return(domain, nil).Once()
		result, err := linkUsecase.Create(context.Background(), &link.Domain{})
		assert.Equal(t, result, link.Domain{})
		assert.NotEqual(t, err, &gorm.ErrRecordNotFound)
	})

}
