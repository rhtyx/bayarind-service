package repository

import (
	"context"

	"github.com/rhtyx/bayarind-service.git/model"
	"github.com/rhtyx/bayarind-service.git/utils"

	"gorm.io/gorm"

	"github.com/sirupsen/logrus"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) model.AuthorRepository {
	return &AuthorRepository{db: db}
}

func (a AuthorRepository) Create(ctx context.Context, author *model.Author) (*model.Author, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("author", utils.Dump(author))

	author.ID = utils.GenerateID()
	err := a.db.WithContext(ctx).Create(author).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return author, nil
}

func (a AuthorRepository) FindByID(ctx context.Context, authorID int64) (*model.Author, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("authorID", authorID)

	author := &model.Author{}
	err := a.db.WithContext(ctx).Take(author, "id = ?", authorID).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return author, nil
}

func (a AuthorRepository) FindAll(ctx context.Context) ([]*model.Author, error) {
	logger := logrus.WithContext(ctx)

	authors := []*model.Author{}
	err := a.db.WithContext(ctx).Find(&authors).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return authors, nil
}

func (a AuthorRepository) Update(ctx context.Context, author *model.Author) (*model.Author, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("author", utils.Dump(author))

	err := a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Save(author).Error
		if err != nil {
			logger.Error(err)
			return err
		}

		err = tx.Take(author).Error
		if err != nil {
			logger.Error(err)
			return err
		}

		return nil
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return author, nil
}

func (a AuthorRepository) Delete(ctx context.Context, authorID int64) error {
	logger := logrus.
		WithContext(ctx).
		WithField("authorID", authorID)

	err := a.db.WithContext(ctx).Delete(&model.Author{}, "id = ?", authorID).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
