package repository

import (
	"context"

	"github.com/rhtyx/bayarind-service.git/model"
	"github.com/rhtyx/bayarind-service.git/utils"

	"gorm.io/gorm"

	"github.com/sirupsen/logrus"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) model.SessionRepository {
	return &SessionRepository{db: db}
}

func (s SessionRepository) Create(ctx context.Context, session *model.Session) (*model.Session, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("session", utils.Dump(session))

	session.ID = utils.GenerateID()
	err := s.db.WithContext(ctx).Create(session).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return session, nil
}

func (s SessionRepository) FindByRefreshToken(ctx context.Context, refreshToken string) (*model.Session, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("refreshToken", refreshToken)

	session := &model.Session{}
	err := s.db.WithContext(ctx).Take(session, "refresh_token = ?", refreshToken).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return session, nil
}

func (s SessionRepository) DeleteByRefreshToken(ctx context.Context, refreshToken string) error {
	logger := logrus.
		WithContext(ctx).
		WithField("refreshToken", refreshToken)

	err := s.db.WithContext(ctx).Delete(&model.Session{}, "refresh_token = ?", refreshToken).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
