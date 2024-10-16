package service

import (
	"context"
	"time"

	"github.com/rhtyx/bayarind-service.git/config"
	"github.com/rhtyx/bayarind-service.git/controller"
	"github.com/rhtyx/bayarind-service.git/model"
	"github.com/rhtyx/bayarind-service.git/token"
	"github.com/rhtyx/bayarind-service.git/utils"

	"github.com/sirupsen/logrus"
)

type SessionService struct {
	sessionRepository model.SessionRepository
	userRepository    model.UserRepository
	jwtService        token.JWTService
}

func NewSessionService(sessionRepository model.SessionRepository, userRepository model.UserRepository, jwtService token.JWTService) model.SessionService {
	return SessionService{
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
		jwtService:        jwtService,
	}
}

func (s SessionService) Create(ctx context.Context, username, password string) (*model.Session, error) {
	logger := logrus.
		WithContext(ctx).
		WithFields(logrus.Fields{
			"username": username,
			"password": password,
		})

	user, err := s.userRepository.FindByUsername(ctx, username)
	if err != nil {
		logger.Error(err)
		return nil, parseError(err, "username")
	}

	if !utils.IsPasswordCorrect(password, user.Password) {
		return nil, controller.ErrCredentials
	}

	now := time.Now()
	refreshToken, err := s.jwtService.CreateToken(user.ID, now, config.RefreshTokenDuration())
	if err != nil {
		logger.WithField("user", utils.Dump(user)).Error(err)
		return nil, parseError(err, "refreshToken")
	}

	accessToken, err := s.jwtService.CreateToken(user.ID, now, config.AccessTokenDuration())
	if err != nil {
		logger.WithField("user", utils.Dump(user)).Error(err)
		return nil, parseError(err, "accessToken")
	}

	session := &model.Session{
		UserID:                user.ID,
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: now.Add(config.RefreshTokenDuration()),
	}
	session, err = s.sessionRepository.Create(ctx, session)
	if err != nil {
		logger.Error(err)
		return nil, parseError(err, "session")
	}

	session.AccessToken = accessToken
	session.AccessTokenExpiredAt = now.Add(config.AccessTokenDuration())
	return session, nil
}

func (s SessionService) FindByRefreshToken(ctx context.Context, refreshToken string) (*model.Session, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("refreshToken", refreshToken)

	session, err := s.sessionRepository.FindByRefreshToken(ctx, refreshToken)
	if err != nil {
		logger.Error(err)
		return nil, parseError(err, "refreshToken")
	}

	return session, nil
}

func (s SessionService) DeleteByRefreshToken(ctx context.Context, refreshToken string) error {
	logger := logrus.
		WithContext(ctx).
		WithField("refreshToken", refreshToken)

	err := s.sessionRepository.DeleteByRefreshToken(ctx, refreshToken)
	if err != nil {
		logger.Error(err)
		return parseError(err, "refreshToken")
	}

	return nil
}

func (s SessionService) RefreshAccessToken(ctx context.Context, refreshToken string) (*model.Session, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("refreshToken", refreshToken)

	session, err := s.sessionRepository.FindByRefreshToken(ctx, refreshToken)
	if err != nil {
		logger.Error(err)
		return nil, parseError(err, "session")
	}

	now := time.Now()
	accessToken, err := s.jwtService.CreateToken(session.UserID, now, config.AccessTokenDuration())
	if err != nil {
		logger.WithField("session", utils.Dump(session)).Error(err)
		return nil, parseError(err, "accessToken")
	}

	session.AccessToken = accessToken
	session.AccessTokenExpiredAt = now.Add(config.AccessTokenDuration())
	return session, nil
}
