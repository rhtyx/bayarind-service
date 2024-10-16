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

func TestUserCreate(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, false, false, false, false, 2),
		}

		userRepository := mock.NewMockUserRepository(ctrl)
		userRepository.EXPECT().
			FindByUsername(ctx, user.Username).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		userRepository.EXPECT().
			Create(ctx, user).
			Times(1).
			Return(user, nil)

		userService := service.NewUserService(userRepository)
		resUser, err := userService.Create(ctx, user)
		assert.Nil(t, err)
		assert.NotNil(t, resUser)
		assert.ObjectsAreEqualValues(user, resUser)
	})

	t.Run("error: duplicate username", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, false, false, false, false, 2),
		}

		userRepository := mock.NewMockUserRepository(ctrl)
		userRepository.EXPECT().
			FindByUsername(ctx, user.Username).
			Times(1).
			Return(user, nil)

		userService := service.NewUserService(userRepository)
		resUser, err := userService.Create(ctx, user)
		assert.Nil(t, resUser)
		assert.Error(t, err)
		assert.EqualError(t, err, "duplicate entry\n: username")
	})

	t.Run("error: create", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, false, false, false, false, 2),
		}

		userRepository := mock.NewMockUserRepository(ctrl)
		userRepository.EXPECT().
			FindByUsername(ctx, user.Username).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		userRepository.EXPECT().
			Create(ctx, user).
			Times(1).
			Return(nil, gorm.ErrInvalidDB)

		userService := service.NewUserService(userRepository)
		resUser, err := userService.Create(ctx, user)
		assert.Nil(t, resUser)
		assert.Error(t, err)
		assert.EqualError(t, err, controller.ErrInternalServer.Error())
	})
}

func TestUserFindByID(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, false, false, false, false, 2),
		}

		userRepository := mock.NewMockUserRepository(ctrl)
		userRepository.EXPECT().
			FindByID(ctx, user.ID).
			Times(1).
			Return(user, nil)

		userService := service.NewUserService(userRepository)
		resUser, err := userService.FindByID(ctx, user.ID)
		assert.Nil(t, err)
		assert.NotNil(t, resUser)
		assert.ObjectsAreEqualValues(user, resUser)
	})

	t.Run("error: id not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, false, false, false, false, 2),
		}

		userRepository := mock.NewMockUserRepository(ctrl)
		userRepository.EXPECT().
			FindByID(ctx, user.ID).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		userService := service.NewUserService(userRepository)
		resUser, err := userService.FindByID(ctx, user.ID)
		assert.Nil(t, resUser)
		assert.Error(t, err)
		assert.EqualError(t, err, "id not found\n: user")
	})
}

func TestUserFindByUsername(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, false, false, false, false, 2),
		}

		userRepository := mock.NewMockUserRepository(ctrl)
		userRepository.EXPECT().
			FindByUsername(ctx, user.Username).
			Times(1).
			Return(user, nil)

		userService := service.NewUserService(userRepository)
		resUser, err := userService.FindByUsername(ctx, user.Username)
		assert.Nil(t, err)
		assert.NotNil(t, resUser)
		assert.ObjectsAreEqualValues(user, resUser)
	})

	t.Run("error: username not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, false, false, false, false, 2),
		}

		userRepository := mock.NewMockUserRepository(ctrl)
		userRepository.EXPECT().
			FindByUsername(ctx, user.Username).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		userService := service.NewUserService(userRepository)
		resUser, err := userService.FindByUsername(ctx, user.Username)
		assert.Nil(t, resUser)
		assert.Error(t, err)
		assert.EqualError(t, err, "id not found\n: username")
	})
}

func TestUserUpdate(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		userID := utils.GenerateID()
		req := &model.User{
			ID:       userID,
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, false, false, false, false, 2),
		}

		user := &model.User{
			ID:       userID,
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, false, false, false, false, 2),
		}

		userRepository := mock.NewMockUserRepository(ctrl)
		userRepository.EXPECT().
			FindByID(ctx, req.ID).
			Times(1).
			Return(user, nil)

		userRepository.EXPECT().
			FindByUsername(ctx, req.Username).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		userRepository.EXPECT().
			Update(ctx, req).
			Times(1).
			Return(req, nil)

		userService := service.NewUserService(userRepository)
		resUser, err := userService.Update(ctx, req)
		assert.Nil(t, err)
		assert.NotNil(t, resUser)
		assert.ObjectsAreEqualValues(req, resUser)
	})

	t.Run("error: id not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		userID := utils.GenerateID()
		req := &model.User{
			ID:       userID,
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, false, false, false, false, 2),
		}

		userRepository := mock.NewMockUserRepository(ctrl)
		userRepository.EXPECT().
			FindByID(ctx, req.ID).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		userService := service.NewUserService(userRepository)
		resUser, err := userService.Update(ctx, req)
		assert.Nil(t, resUser)
		assert.Error(t, err)
		assert.EqualError(t, err, "id not found\n: user")
	})

	t.Run("error: find username", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		userID := utils.GenerateID()
		req := &model.User{
			ID:       userID,
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, false, false, false, false, 2),
		}

		user := &model.User{
			ID:       userID,
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, false, false, false, false, 2),
		}

		userRepository := mock.NewMockUserRepository(ctrl)
		userRepository.EXPECT().
			FindByID(ctx, req.ID).
			Times(1).
			Return(user, nil)

		userRepository.EXPECT().
			FindByUsername(ctx, req.Username).
			Times(1).
			Return(nil, gorm.ErrInvalidDB)

		userService := service.NewUserService(userRepository)
		resUser, err := userService.Update(ctx, req)
		assert.Nil(t, resUser)
		assert.Error(t, err)
		assert.EqualError(t, err, controller.ErrInternalServer.Error())
	})

	t.Run("error: duplicate username", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		userID := utils.GenerateID()
		req := &model.User{
			ID:       userID,
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, false, false, false, false, 2),
		}

		user := &model.User{
			ID:       userID,
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, false, false, false, false, 2),
		}

		userRepository := mock.NewMockUserRepository(ctrl)
		userRepository.EXPECT().
			FindByID(ctx, req.ID).
			Times(1).
			Return(user, nil)

		userRepository.EXPECT().
			FindByUsername(ctx, req.Username).
			Times(1).
			Return(user, nil)

		userService := service.NewUserService(userRepository)
		resUser, err := userService.Update(ctx, req)
		assert.Nil(t, resUser)
		assert.Error(t, err)
		assert.EqualError(t, err, "duplicate entry\n: username")
	})

	t.Run("error: update", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		userID := utils.GenerateID()
		req := &model.User{
			ID:       userID,
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, false, false, false, false, 2),
		}

		user := &model.User{
			ID:       userID,
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, false, false, false, false, 2),
		}

		userRepository := mock.NewMockUserRepository(ctrl)
		userRepository.EXPECT().
			FindByID(ctx, req.ID).
			Times(1).
			Return(user, nil)

		userRepository.EXPECT().
			FindByUsername(ctx, req.Username).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		userRepository.EXPECT().
			Update(ctx, req).
			Times(1).
			Return(nil, gorm.ErrInvalidDB)

		userService := service.NewUserService(userRepository)
		resUser, err := userService.Update(ctx, req)
		assert.Nil(t, resUser)
		assert.Error(t, err)
		assert.EqualError(t, err, controller.ErrInternalServer.Error())
	})
}

func TestUserDelete(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, false, false, false, false, 2),
		}

		userRepository := mock.NewMockUserRepository(ctrl)
		userRepository.EXPECT().
			Delete(ctx, user.ID).
			Times(1).
			Return(nil)

		userService := service.NewUserService(userRepository)
		err := userService.Delete(ctx, user.ID)
		assert.Nil(t, err)
	})

	t.Run("error: id not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, false, false, false, false, 2),
		}

		userRepository := mock.NewMockUserRepository(ctrl)
		userRepository.EXPECT().
			Delete(ctx, user.ID).
			Times(1).
			Return(gorm.ErrRecordNotFound)

		userService := service.NewUserService(userRepository)
		err := userService.Delete(ctx, user.ID)
		assert.Error(t, err)
		assert.EqualError(t, err, "id not found\n: user")
	})
}
