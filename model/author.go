package model

import (
	"context"
	"time"
)

type Author struct {
	ID        int64      `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	BirthDate time.Time  `json:"birth_date"`
	CreatedAt time.Time  `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"<-:update"`
}

type AuthorRepository interface {
	Create(ctx context.Context, author *Author) (*Author, error)
	FindByID(ctx context.Context, authorID int64) (*Author, error)
	FindAll(ctx context.Context) ([]*Author, error)
	Update(ctx context.Context, author *Author) (*Author, error)
	Delete(ctx context.Context, authorID int64) error
}

type AuthorService interface {
	Create(ctx context.Context, author *Author) (*Author, error)
	FindByID(ctx context.Context, authorID int64) (*Author, error)
	FindAll(ctx context.Context) ([]*Author, error)
	Update(ctx context.Context, author *Author) (*Author, error)
	Delete(ctx context.Context, authorID int64) error
}
