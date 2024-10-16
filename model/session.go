package model

import (
	"context"
	"time"
)

type Session struct {
	ID                    int64      `json:"id" gorm:"primaryKey"`
	UserID                int64      `json:"user_id"`
	RefreshToken          string     `json:"refresh_token"`
	RefreshTokenExpiredAt time.Time  `json:"refresh_token_expired_at"`
	CreatedAt             time.Time  `json:"created_at" gorm:"<-:create"`
	UpdatedAt             *time.Time `json:"updated_at" gorm:"<-:update"`

	AccessToken          string    `json:"access_token" gorm:"-"`
	AccessTokenExpiredAt time.Time `json:"access_token_expired_at" gorm:"-"`
}

type SessionRepository interface {
	Create(ctx context.Context, session *Session) (*Session, error)
	FindByRefreshToken(ctx context.Context, refreshToken string) (*Session, error)
	DeleteByRefreshToken(ctx context.Context, refreshToken string) error
}

type SessionService interface {
	Create(ctx context.Context, username, password string) (*Session, error)
	FindByRefreshToken(ctx context.Context, refreshToken string) (*Session, error)
	DeleteByRefreshToken(ctx context.Context, refreshToken string) error

	RefreshAccessToken(ctx context.Context, refreshToken string) (*Session, error)
}
