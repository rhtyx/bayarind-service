package repository

import (
	"context"

	"github.com/rhtyx/bayarind-service.git/model"
	"github.com/rhtyx/bayarind-service.git/utils"

	"gorm.io/gorm"

	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("user", utils.Dump(user))

	user.ID = utils.GenerateID()
	err := u.db.WithContext(ctx).Create(user).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return user, nil
}

func (u UserRepository) FindByID(ctx context.Context, userID int64) (*model.User, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("userID", userID)

	user := &model.User{}
	err := u.db.WithContext(ctx).Take(user, "id = ?", userID).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return user, nil
}

func (u UserRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("username", username)

	user := &model.User{}
	err := u.db.WithContext(ctx).Take(user, "username = ?", username).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return user, nil
}

func (u UserRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("user", utils.Dump(user))

	err := u.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Save(user).Error
		if err != nil {
			logger.Error(err)
			tx.Rollback()
			return err
		}

		err = tx.Take(user).Error
		if err != nil {
			logger.Error(err)
			tx.Rollback()
			return err
		}

		return nil
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return user, nil
}

func (u UserRepository) Delete(ctx context.Context, userID int64) error {
	logger := logrus.
		WithContext(ctx).
		WithField("userID", userID)

	err := u.db.WithContext(ctx).Delete(&model.User{}, "id = ?", userID).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
