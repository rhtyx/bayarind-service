package repository

import (
	"context"

	"github.com/rhtyx/bayarind-service.git/model"
	"github.com/rhtyx/bayarind-service.git/utils"

	"gorm.io/gorm"

	"github.com/sirupsen/logrus"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) model.BookRepository {
	return &BookRepository{db: db}
}

func (b BookRepository) Create(ctx context.Context, book *model.Book) (*model.Book, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("book", utils.Dump(book))

	book.ID = utils.GenerateID()
	err := b.db.WithContext(ctx).Create(book).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return book, nil
}

func (b BookRepository) FindByID(ctx context.Context, bookID int64) (*model.Book, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("bookID", bookID)

	book := &model.Book{}
	err := b.db.WithContext(ctx).Take(book, "id = ?", bookID).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return book, nil
}

func (b BookRepository) FindByISBN(ctx context.Context, isbn string) (*model.Book, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("isbn", isbn)

	book := &model.Book{}
	err := b.db.WithContext(ctx).Take(book, "isbn = ?", isbn).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return book, nil
}

func (b BookRepository) FindAll(ctx context.Context) ([]*model.Book, error) {
	logger := logrus.
		WithContext(ctx)

	book := []*model.Book{}
	err := b.db.WithContext(ctx).Find(&book).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return book, nil
}

func (b BookRepository) Update(ctx context.Context, book *model.Book) (*model.Book, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("book", utils.Dump(book))

	err := b.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Save(book).Error
		if err != nil {
			logger.Error(err)
			tx.Rollback()
			return err
		}

		err = tx.Take(book).Error
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

	return book, nil
}

func (b BookRepository) Delete(ctx context.Context, bookID int64) error {
	logger := logrus.
		WithContext(ctx).
		WithField("bookID", bookID)

	err := b.db.WithContext(ctx).Delete(&model.Book{}, "id = ?", bookID).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
