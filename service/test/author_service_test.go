package test

import (
	"context"
	"testing"

	"go.uber.org/mock/gomock"
	"gorm.io/gorm"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/rhtyx/bayarind-service.git/controller"
	"github.com/rhtyx/bayarind-service.git/model"
	"github.com/rhtyx/bayarind-service.git/model/mock"
	"github.com/rhtyx/bayarind-service.git/service"
	"github.com/rhtyx/bayarind-service.git/utils"
	"github.com/stretchr/testify/assert"
)

func TestAuthorCreate(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		author := &model.Author{
			ID:        utils.GenerateID(),
			Name:      gofakeit.Name(),
			BirthDate: gofakeit.Date(),
		}

		authorRepository := mock.NewMockAuthorRepository(ctrl)
		authorRepository.EXPECT().
			Create(ctx, author).
			Times(1).
			Return(author, nil)

		authorService := service.NewAuthorService(authorRepository)
		resAuthor, err := authorService.Create(ctx, author)
		assert.Nil(t, err)
		assert.NotNil(t, resAuthor)
		assert.ObjectsAreEqualValues(author, resAuthor)
	})

	t.Run("error: id duplicate", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		author := &model.Author{
			ID:        utils.GenerateID(),
			Name:      gofakeit.Name(),
			BirthDate: gofakeit.Date(),
		}

		authorRepository := mock.NewMockAuthorRepository(ctrl)
		authorRepository.EXPECT().
			Create(ctx, author).
			Times(1).
			Return(nil, gorm.ErrDuplicatedKey)

		authorService := service.NewAuthorService(authorRepository)
		resAuthor, err := authorService.Create(ctx, author)
		assert.Nil(t, resAuthor)
		assert.Error(t, err)
		assert.EqualError(t, err, "duplicate entry\n: author")
	})
}

func TestAuthorFindByID(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		author := &model.Author{
			ID:        utils.GenerateID(),
			Name:      gofakeit.Name(),
			BirthDate: gofakeit.Date(),
		}

		authorRepository := mock.NewMockAuthorRepository(ctrl)
		authorRepository.EXPECT().
			FindByID(ctx, author.ID).
			Times(1).
			Return(author, nil)

		authorService := service.NewAuthorService(authorRepository)
		resAuthor, err := authorService.FindByID(ctx, author.ID)
		assert.Nil(t, err)
		assert.NotNil(t, resAuthor)
		assert.ObjectsAreEqualValues(author, resAuthor)
	})

	t.Run("error: id not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		author := &model.Author{
			ID:        utils.GenerateID(),
			Name:      gofakeit.Name(),
			BirthDate: gofakeit.Date(),
		}

		authorRepository := mock.NewMockAuthorRepository(ctrl)
		authorRepository.EXPECT().
			FindByID(ctx, author.ID).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		authorService := service.NewAuthorService(authorRepository)
		resAuthor, err := authorService.FindByID(ctx, author.ID)
		assert.Nil(t, resAuthor)
		assert.Error(t, err)
		assert.EqualError(t, err, "id not found\n: id")
	})
}

func TestAuthorFindAll(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		author := []*model.Author{
			{
				ID:        utils.GenerateID(),
				Name:      gofakeit.Name(),
				BirthDate: gofakeit.Date(),
			},
			{
				ID:        utils.GenerateID(),
				Name:      gofakeit.Name(),
				BirthDate: gofakeit.Date(),
			},
		}

		authorRepository := mock.NewMockAuthorRepository(ctrl)
		authorRepository.EXPECT().
			FindAll(ctx).
			Times(1).
			Return(author, nil)

		authorService := service.NewAuthorService(authorRepository)
		resAuthor, err := authorService.FindAll(ctx)
		assert.Nil(t, err)
		assert.NotNil(t, resAuthor)
		assert.ObjectsAreEqualValues(author, resAuthor)
	})

	t.Run("error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()

		authorRepository := mock.NewMockAuthorRepository(ctrl)
		authorRepository.EXPECT().
			FindAll(ctx).
			Times(1).
			Return(nil, gorm.ErrInvalidDB)

		authorService := service.NewAuthorService(authorRepository)
		resAuthor, err := authorService.FindAll(ctx)
		assert.Nil(t, resAuthor)
		assert.Error(t, err)
		assert.EqualError(t, err, controller.ErrInternalServer.Error())
	})
}

func TestAuthorUpdate(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		author := &model.Author{
			ID:        utils.GenerateID(),
			Name:      gofakeit.Name(),
			BirthDate: gofakeit.Date(),
		}

		authorRepository := mock.NewMockAuthorRepository(ctrl)
		authorRepository.EXPECT().
			Update(ctx, author).
			Times(1).
			Return(author, nil)

		authorService := service.NewAuthorService(authorRepository)
		resAuthor, err := authorService.Update(ctx, author)
		assert.Nil(t, err)
		assert.NotNil(t, resAuthor)
		assert.ObjectsAreEqualValues(author, resAuthor)
	})

	t.Run("error: id not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		author := &model.Author{
			ID:        utils.GenerateID(),
			Name:      gofakeit.Name(),
			BirthDate: gofakeit.Date(),
		}

		authorRepository := mock.NewMockAuthorRepository(ctrl)
		authorRepository.EXPECT().
			Update(ctx, author).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		authorService := service.NewAuthorService(authorRepository)
		resAuthor, err := authorService.Update(ctx, author)
		assert.Nil(t, resAuthor)
		assert.Error(t, err)
		assert.EqualError(t, err, "id not found\n: author")
	})
}

func TestAuthorDelete(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		authorID := utils.GenerateID()

		authorRepository := mock.NewMockAuthorRepository(ctrl)
		authorRepository.EXPECT().
			Delete(ctx, authorID).
			Times(1).
			Return(nil)

		authorService := service.NewAuthorService(authorRepository)
		err := authorService.Delete(ctx, authorID)
		assert.Nil(t, err)
	})

	t.Run("error: id not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		authorID := utils.GenerateID()

		authorRepository := mock.NewMockAuthorRepository(ctrl)
		authorRepository.EXPECT().
			Delete(ctx, authorID).
			Times(1).
			Return(gorm.ErrRecordNotFound)

		authorService := service.NewAuthorService(authorRepository)
		err := authorService.Delete(ctx, authorID)
		assert.Error(t, err)
		assert.EqualError(t, err, "id not found\n: author")
	})
}
