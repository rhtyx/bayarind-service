package test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/rhtyx/bayarind-service.git/controller"
	"github.com/rhtyx/bayarind-service.git/model"
	"github.com/rhtyx/bayarind-service.git/model/mock"
	"github.com/rhtyx/bayarind-service.git/service"
	"github.com/rhtyx/bayarind-service.git/token"
	"github.com/rhtyx/bayarind-service.git/utils"
	"github.com/stretchr/testify/assert"
)

func init() {
	privateKey, _ := os.ReadFile("./../../cert/id_rsa.pri")
	publicKey, _ := os.ReadFile("./../../cert/id_rsa.pub")
	token.Jwt = &token.JWT{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}
}

func TestSessionCreate(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		now := time.Now()

		password := gofakeit.Password(true, false, false, false, false, 2)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: string(hashedPassword),
		}

		refreshToken, _ := token.Jwt.CreateToken(user.ID, now, 24*time.Hour)
		accessToken, _ := token.Jwt.CreateToken(user.ID, now, 5*time.Minute)
		session := &model.Session{
			UserID:                user.ID,
			RefreshToken:          refreshToken,
			RefreshTokenExpiredAt: now.Add(24 * time.Hour),
		}

		sessionRepository := mock.NewMockSessionRepository(ctrl)
		userRepository := mock.NewMockUserRepository(ctrl)
		jwtService := mock.NewMockJWTService(ctrl)

		userRepository.EXPECT().
			FindByUsername(ctx, user.Username).
			Times(1).
			Return(user, nil)

		jwtService.EXPECT().
			CreateToken(user.ID, gomock.Any(), 24*time.Hour).
			Times(1).
			Return(session.RefreshToken, nil)

		jwtService.EXPECT().
			CreateToken(user.ID, gomock.Any(), 5*time.Minute).
			Times(1).
			Return(accessToken, nil)

		sessionRepository.EXPECT().
			Create(ctx, gomock.Any()).
			Times(1).
			Return(session, nil)

		sessionService := service.NewSessionService(sessionRepository, userRepository, jwtService)
		resSession, err := sessionService.Create(ctx, user.Username, password)
		assert.Nil(t, err)
		assert.NotNil(t, resSession)
		assert.ObjectsAreEqualValues(session, resSession)
	})

	t.Run("error: username not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		password := gofakeit.Password(true, false, false, false, false, 2)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: string(hashedPassword),
		}

		sessionRepository := mock.NewMockSessionRepository(ctrl)
		userRepository := mock.NewMockUserRepository(ctrl)
		jwtService := mock.NewMockJWTService(ctrl)

		userRepository.EXPECT().
			FindByUsername(ctx, user.Username).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		sessionService := service.NewSessionService(sessionRepository, userRepository, jwtService)
		resSession, err := sessionService.Create(ctx, user.Username, password)
		assert.Nil(t, resSession)
		assert.Error(t, err)
		assert.EqualError(t, err, "id not found\n: username")
	})

	t.Run("error: wrong password", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()

		password := gofakeit.Password(true, false, false, false, false, 2)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: string(hashedPassword),
		}

		sessionRepository := mock.NewMockSessionRepository(ctrl)
		userRepository := mock.NewMockUserRepository(ctrl)
		jwtService := mock.NewMockJWTService(ctrl)

		userRepository.EXPECT().
			FindByUsername(ctx, user.Username).
			Times(1).
			Return(user, nil)

		sessionService := service.NewSessionService(sessionRepository, userRepository, jwtService)
		resSession, err := sessionService.Create(ctx, user.Username, "wrong"+password)
		assert.Nil(t, resSession)
		assert.Error(t, err)
		assert.EqualError(t, err, controller.ErrCredentials.Error())
	})

	t.Run("error: create refresh token", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()

		password := gofakeit.Password(true, false, false, false, false, 2)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: string(hashedPassword),
		}

		sessionRepository := mock.NewMockSessionRepository(ctrl)
		userRepository := mock.NewMockUserRepository(ctrl)
		jwtService := mock.NewMockJWTService(ctrl)

		userRepository.EXPECT().
			FindByUsername(ctx, user.Username).
			Times(1).
			Return(user, nil)

		jwtService.EXPECT().
			CreateToken(user.ID, gomock.Any(), 24*time.Hour).
			Times(1).
			Return("", errors.New("error create refresh token"))

		sessionService := service.NewSessionService(sessionRepository, userRepository, jwtService)
		resSession, err := sessionService.Create(ctx, user.Username, password)
		assert.Nil(t, resSession)
		assert.Error(t, err)
		assert.EqualError(t, err, controller.ErrInternalServer.Error())
	})

	t.Run("error: create access token", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		now := time.Now()

		password := gofakeit.Password(true, false, false, false, false, 2)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: string(hashedPassword),
		}

		refreshToken, _ := token.Jwt.CreateToken(user.ID, now, 24*time.Hour)
		session := &model.Session{
			UserID:                user.ID,
			RefreshToken:          refreshToken,
			RefreshTokenExpiredAt: now.Add(24 * time.Hour),
		}

		sessionRepository := mock.NewMockSessionRepository(ctrl)
		userRepository := mock.NewMockUserRepository(ctrl)
		jwtService := mock.NewMockJWTService(ctrl)

		userRepository.EXPECT().
			FindByUsername(ctx, user.Username).
			Times(1).
			Return(user, nil)

		jwtService.EXPECT().
			CreateToken(user.ID, gomock.Any(), 24*time.Hour).
			Times(1).
			Return(session.RefreshToken, nil)

		jwtService.EXPECT().
			CreateToken(user.ID, gomock.Any(), 5*time.Minute).
			Times(1).
			Return("", errors.New("error creating access token"))

		sessionService := service.NewSessionService(sessionRepository, userRepository, jwtService)
		resSession, err := sessionService.Create(ctx, user.Username, password)
		assert.Nil(t, resSession)
		assert.Error(t, err)
		assert.EqualError(t, err, controller.ErrInternalServer.Error())
	})

	t.Run("error: duplicated session", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		now := time.Now()

		password := gofakeit.Password(true, false, false, false, false, 2)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: string(hashedPassword),
		}

		refreshToken, _ := token.Jwt.CreateToken(user.ID, now, 24*time.Hour)
		accessToken, _ := token.Jwt.CreateToken(user.ID, now, 5*time.Minute)
		session := &model.Session{
			UserID:                user.ID,
			RefreshToken:          refreshToken,
			RefreshTokenExpiredAt: now.Add(24 * time.Hour),
		}

		sessionRepository := mock.NewMockSessionRepository(ctrl)
		userRepository := mock.NewMockUserRepository(ctrl)
		jwtService := mock.NewMockJWTService(ctrl)

		userRepository.EXPECT().
			FindByUsername(ctx, user.Username).
			Times(1).
			Return(user, nil)

		jwtService.EXPECT().
			CreateToken(user.ID, gomock.Any(), 24*time.Hour).
			Times(1).
			Return(session.RefreshToken, nil)

		jwtService.EXPECT().
			CreateToken(user.ID, gomock.Any(), 5*time.Minute).
			Times(1).
			Return(accessToken, nil)

		sessionRepository.EXPECT().
			Create(ctx, gomock.Any()).
			Times(1).
			Return(nil, gorm.ErrDuplicatedKey)

		sessionService := service.NewSessionService(sessionRepository, userRepository, jwtService)
		resSession, err := sessionService.Create(ctx, user.Username, password)
		assert.Nil(t, resSession)
		assert.Error(t, err)
		assert.EqualError(t, err, "duplicate entry\n: session")
	})
}

func TestSessionFindByRefreshToken(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		now := time.Now()

		password := gofakeit.Password(true, false, false, false, false, 2)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: string(hashedPassword),
		}

		refreshToken, _ := token.Jwt.CreateToken(user.ID, now, 24*time.Hour)
		session := &model.Session{
			UserID:                user.ID,
			RefreshToken:          refreshToken,
			RefreshTokenExpiredAt: now.Add(24 * time.Hour),
		}

		sessionRepository := mock.NewMockSessionRepository(ctrl)
		userRepository := mock.NewMockUserRepository(ctrl)
		jwtService := mock.NewMockJWTService(ctrl)

		sessionRepository.EXPECT().
			FindByRefreshToken(ctx, gomock.Any()).
			Times(1).
			Return(session, nil)

		sessionService := service.NewSessionService(sessionRepository, userRepository, jwtService)
		resSession, err := sessionService.FindByRefreshToken(ctx, refreshToken)
		assert.Nil(t, err)
		assert.NotNil(t, resSession)
		assert.ObjectsAreEqualValues(session, resSession)
	})

	t.Run("error: refresh token not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		now := time.Now()

		password := gofakeit.Password(true, false, false, false, false, 2)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: string(hashedPassword),
		}

		refreshToken, _ := token.Jwt.CreateToken(user.ID, now, 24*time.Hour)

		sessionRepository := mock.NewMockSessionRepository(ctrl)
		userRepository := mock.NewMockUserRepository(ctrl)
		jwtService := mock.NewMockJWTService(ctrl)

		sessionRepository.EXPECT().
			FindByRefreshToken(ctx, gomock.Any()).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		sessionService := service.NewSessionService(sessionRepository, userRepository, jwtService)
		resSession, err := sessionService.FindByRefreshToken(ctx, refreshToken)
		assert.Nil(t, resSession)
		assert.Error(t, err)
		assert.EqualError(t, err, "id not found\n: refreshToken")
	})
}

func TestSessionRefreshAccessToken(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		now := time.Now()

		password := gofakeit.Password(true, false, false, false, false, 2)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: string(hashedPassword),
		}

		refreshToken, _ := token.Jwt.CreateToken(user.ID, now, 24*time.Hour)
		accessToken, _ := token.Jwt.CreateToken(user.ID, now, 5*time.Minute)
		session := &model.Session{
			UserID:                user.ID,
			RefreshToken:          refreshToken,
			RefreshTokenExpiredAt: now.Add(24 * time.Hour),
		}

		sessionRepository := mock.NewMockSessionRepository(ctrl)
		userRepository := mock.NewMockUserRepository(ctrl)
		jwtService := mock.NewMockJWTService(ctrl)

		sessionRepository.EXPECT().
			FindByRefreshToken(ctx, gomock.Any()).
			Times(1).
			Return(session, nil)

		jwtService.EXPECT().
			CreateToken(user.ID, gomock.Any(), 5*time.Minute).
			Times(1).
			Return(accessToken, nil)

		sessionService := service.NewSessionService(sessionRepository, userRepository, jwtService)
		resSession, err := sessionService.RefreshAccessToken(ctx, refreshToken)
		assert.Nil(t, err)
		assert.NotNil(t, resSession)
		assert.ObjectsAreEqualValues(session, resSession)
	})

	t.Run("error: refresh token not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		now := time.Now()

		password := gofakeit.Password(true, false, false, false, false, 2)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: string(hashedPassword),
		}

		refreshToken, _ := token.Jwt.CreateToken(user.ID, now, 24*time.Hour)

		sessionRepository := mock.NewMockSessionRepository(ctrl)
		userRepository := mock.NewMockUserRepository(ctrl)
		jwtService := mock.NewMockJWTService(ctrl)

		sessionRepository.EXPECT().
			FindByRefreshToken(ctx, gomock.Any()).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		sessionService := service.NewSessionService(sessionRepository, userRepository, jwtService)
		resSession, err := sessionService.RefreshAccessToken(ctx, refreshToken)
		assert.Nil(t, resSession)
		assert.Error(t, err)
		assert.EqualError(t, err, "id not found\n: session")
	})

	t.Run("error: create access token", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		now := time.Now()

		password := gofakeit.Password(true, false, false, false, false, 2)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: string(hashedPassword),
		}

		refreshToken, _ := token.Jwt.CreateToken(user.ID, now, 24*time.Hour)
		session := &model.Session{
			UserID:                user.ID,
			RefreshToken:          refreshToken,
			RefreshTokenExpiredAt: now.Add(24 * time.Hour),
		}

		sessionRepository := mock.NewMockSessionRepository(ctrl)
		userRepository := mock.NewMockUserRepository(ctrl)
		jwtService := mock.NewMockJWTService(ctrl)

		sessionRepository.EXPECT().
			FindByRefreshToken(ctx, gomock.Any()).
			Times(1).
			Return(session, nil)

		jwtService.EXPECT().
			CreateToken(user.ID, gomock.Any(), 5*time.Minute).
			Times(1).
			Return("", errors.New("error creating access token"))

		sessionService := service.NewSessionService(sessionRepository, userRepository, jwtService)
		resSession, err := sessionService.RefreshAccessToken(ctx, refreshToken)
		assert.Nil(t, resSession)
		assert.Error(t, err)
		assert.EqualError(t, err, controller.ErrInternalServer.Error())
	})
}

func TestSessionDeleteByRefreshToken(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		now := time.Now()

		password := gofakeit.Password(true, false, false, false, false, 2)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: string(hashedPassword),
		}

		refreshToken, _ := token.Jwt.CreateToken(user.ID, now, 24*time.Hour)

		sessionRepository := mock.NewMockSessionRepository(ctrl)
		userRepository := mock.NewMockUserRepository(ctrl)
		jwtService := mock.NewMockJWTService(ctrl)

		sessionRepository.EXPECT().
			DeleteByRefreshToken(ctx, gomock.Any()).
			Times(1).
			Return(nil)

		sessionService := service.NewSessionService(sessionRepository, userRepository, jwtService)
		err := sessionService.DeleteByRefreshToken(ctx, refreshToken)
		assert.Nil(t, err)
	})

	t.Run("error: refresh token not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		now := time.Now()

		password := gofakeit.Password(true, false, false, false, false, 2)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user := &model.User{
			ID:       utils.GenerateID(),
			Username: gofakeit.Username(),
			Password: string(hashedPassword),
		}

		refreshToken, _ := token.Jwt.CreateToken(user.ID, now, 24*time.Hour)

		sessionRepository := mock.NewMockSessionRepository(ctrl)
		userRepository := mock.NewMockUserRepository(ctrl)
		jwtService := mock.NewMockJWTService(ctrl)

		sessionRepository.EXPECT().
			DeleteByRefreshToken(ctx, gomock.Any()).
			Times(1).
			Return(gorm.ErrRecordNotFound)

		sessionService := service.NewSessionService(sessionRepository, userRepository, jwtService)
		err := sessionService.DeleteByRefreshToken(ctx, refreshToken)
		assert.Error(t, err)
		assert.EqualError(t, err, "id not found\n: refreshToken")
	})
}
