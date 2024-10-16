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

type BookService struct {
	bookRepository   model.BookRepository
	authorRepository model.AuthorRepository
}

func NewBookService(bookRepository model.BookRepository, authorRepository model.AuthorRepository) model.BookService {
	return &BookService{
		bookRepository:   bookRepository,
		authorRepository: authorRepository,
	}
}

func (b BookService) Create(ctx context.Context, book *model.Book) (*model.Book, error) {
	logger := logrus.WithContext(ctx)

	currBook, err := b.bookRepository.FindByISBN(ctx, book.ISBN)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.WithField("bookID", book.ID).Error(err)
		return nil, parseError(err, "book")
	}

	if currBook != nil {
		return nil, errors.Join(controller.ErrDuplicate, errors.New(": isbn"))
	}

	_, err = b.authorRepository.FindByID(ctx, book.AuthorID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Join(controller.ErrNotFound, errors.New(": author"))
		}

		logger.WithField("authorID", book.AuthorID).Error(err)
		return nil, parseError(err, "author")
	}

	book, err = b.bookRepository.Create(ctx, book)
	if err != nil {
		logger.WithField("book", utils.Dump(book)).Error(err)
		return nil, parseError(err, "book")
	}

	return book, nil
}

func (b BookService) FindByID(ctx context.Context, bookID int64) (*model.Book, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("bookID", bookID)

	book, err := b.bookRepository.FindByID(ctx, bookID)
	if err != nil {
		logger.Error(err)
		return nil, parseError(err, "book")
	}

	return book, nil
}

func (b BookService) FindByISBN(ctx context.Context, isbn string) (*model.Book, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("isbn", isbn)

	book, err := b.bookRepository.FindByISBN(ctx, isbn)
	if err != nil {
		logger.Error(err)
		return nil, parseError(err, "book")
	}

	return book, nil
}

func (b BookService) FindAll(ctx context.Context) ([]*model.Book, error) {
	logger := logrus.
		WithContext(ctx)

	books, err := b.bookRepository.FindAll(ctx)
	if err != nil {
		logger.Error(err)
		return nil, parseError(err, "book")
	}

	return books, nil
}

func (b BookService) Update(ctx context.Context, book *model.Book) (*model.Book, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("book", utils.Dump(book))

	currBook, err := b.bookRepository.FindByID(ctx, book.ID)
	if err != nil {
		logger.WithField("bookID", book.ID).Error(err)
		return nil, parseError(err, "book")
	}

	if currBook.ISBN != book.ISBN {
		currBook, err = b.bookRepository.FindByISBN(ctx, book.ISBN)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error(err)
			return nil, parseError(err, "isbn")
		}

		if currBook != nil {
			return nil, errors.Join(controller.ErrDuplicate, errors.New(": isbn"))
		}
	}

	_, err = b.authorRepository.FindByID(ctx, book.AuthorID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Join(controller.ErrNotFound, errors.New(": author"))
		}

		logger.WithField("authorID", book.AuthorID).Error(err)
		return nil, parseError(err, "author")
	}

	book, err = b.bookRepository.Update(ctx, book)
	if err != nil {
		logger.Error(err)
		return nil, parseError(err, "book")
	}

	return book, nil
}

func (b BookService) Delete(ctx context.Context, bookID int64) error {
	logger := logrus.
		WithContext(ctx)

	err := b.bookRepository.Delete(ctx, bookID)
	if err != nil {
		logger.Error(err)
		return parseError(err, "book")
	}

	return nil
}
