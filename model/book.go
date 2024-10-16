package model

import (
	"context"
	"time"
)

type Book struct {
	ID        int64      `json:"id" gorm:"primaryKey"`
	ISBN      string     `json:"isbn"`
	Title     string     `json:"title"`
	AuthorID  int64      `json:"author"`
	CreatedAt time.Time  `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"<-:update"`
}

type BookRepository interface {
	Create(ctx context.Context, book *Book) (*Book, error)
	FindByID(ctx context.Context, bookID int64) (*Book, error)
	FindByISBN(ctx context.Context, isbn string) (*Book, error)
	FindAll(ctx context.Context) ([]*Book, error)
	Update(ctx context.Context, book *Book) (*Book, error)
	Delete(ctx context.Context, bookID int64) error
}

type BookService interface {
	Create(ctx context.Context, book *Book) (*Book, error)
	FindByID(ctx context.Context, bookID int64) (*Book, error)
	FindByISBN(ctx context.Context, isbn string) (*Book, error)
	FindAll(ctx context.Context) ([]*Book, error)
	Update(ctx context.Context, book *Book) (*Book, error)
	Delete(ctx context.Context, bookID int64) error
}
