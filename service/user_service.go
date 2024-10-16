package service

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/rhtyx/bayarind-service.git/controller"
	"github.com/rhtyx/bayarind-service.git/model"
	"github.com/rhtyx/bayarind-service.git/utils"

	"github.com/sirupsen/logrus"
)

type UserService struct {
	userRepository model.UserRepository
}

func NewUserService(userRepository model.UserRepository) model.UserService {
	return &UserService{userRepository: userRepository}
}

func (u UserService) Create(ctx context.Context, user *model.User) (*model.User, error) {
	logger := logrus.WithContext(ctx)

	currUser, err := u.userRepository.FindByUsername(ctx, user.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.WithField("user", utils.Dump(user)).Error(err)
		return nil, parseError(err, "username")
	}

	if currUser != nil {
		return nil, errors.Join(controller.ErrDuplicate, errors.New(": username"))
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	user.Password = hashedPassword
	logger.WithField("user", utils.Dump(user))

	user, err = u.userRepository.Create(ctx, user)
	if err != nil {
		logger.Error(err)
		return nil, parseError(err, "user")
	}

	user.Password = ""
	return user, nil
}

func (u UserService) FindByID(ctx context.Context, userID int64) (*model.User, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("userID", userID)

	user, err := u.userRepository.FindByID(ctx, userID)
	if err != nil {
		logger.Error(err)
		return nil, parseError(err, "user")
	}

	user.Password = ""
	return user, nil
}

func (u UserService) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("username", username)

	user, err := u.userRepository.FindByUsername(ctx, username)
	if err != nil {
		logger.Error(err)
		return nil, parseError(err, "username")
	}

	user.Password = ""
	return user, nil
}

func (u UserService) Update(ctx context.Context, user *model.User) (*model.User, error) {
	logger := logrus.WithContext(ctx).WithField("user", utils.Dump(user))

	currUser, err := u.userRepository.FindByID(ctx, user.ID)
	if err != nil {
		logger.Error(err)
		return nil, parseError(err, "user")
	}

	if currUser.Username != user.Username {
		currUser, err = u.userRepository.FindByUsername(ctx, user.Username)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error(err)
			return nil, parseError(err, "username")
		}

		if currUser != nil {
			return nil, errors.Join(controller.ErrDuplicate, errors.New(": username"))
		}
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	user.Password = hashedPassword
	user, err = u.userRepository.Update(ctx, user)
	if err != nil {
		logger.Error(err)
		return nil, parseError(err, "user")
	}

	user.Password = ""
	return user, nil
}

func (u UserService) Delete(ctx context.Context, userID int64) error {
	logger := logrus.
		WithContext(ctx).
		WithField("userID", userID)

	err := u.userRepository.Delete(ctx, userID)
	if err != nil {
		logger.Error(err)
		return parseError(err, "user")
	}

	return nil
}
