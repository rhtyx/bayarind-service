package service

import (
	"context"

	"github.com/rhtyx/bayarind-service.git/model"
	"github.com/rhtyx/bayarind-service.git/utils"

	"github.com/sirupsen/logrus"
)

type AuthorService struct {
	authorRepository model.AuthorRepository
}

func NewAuthorService(authoRepository model.AuthorRepository) model.AuthorService {
	return &AuthorService{authorRepository: authoRepository}
}

func (a AuthorService) Create(ctx context.Context, author *model.Author) (*model.Author, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("author", utils.Dump(author))

	author, err := a.authorRepository.Create(ctx, author)
	if err != nil {
		logger.Error(err)
		return nil, parseError(err, "author")
	}

	return author, nil
}

func (a AuthorService) FindByID(ctx context.Context, authorID int64) (*model.Author, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("authorID", authorID)

	author, err := a.authorRepository.FindByID(ctx, authorID)
	if err != nil {
		logger.Error(err)
		return nil, parseError(err, "id")
	}

	return author, nil
}

func (a AuthorService) FindAll(ctx context.Context) ([]*model.Author, error) {
	logger := logrus.WithContext(ctx)

	authors, err := a.authorRepository.FindAll(ctx)
	if err != nil {
		logger.Error(err)
		return nil, parseError(err, "author")
	}

	return authors, nil
}

func (a AuthorService) Update(ctx context.Context, author *model.Author) (*model.Author, error) {
	logger := logrus.
		WithContext(ctx).
		WithField("author", utils.Dump(author))

	author, err := a.authorRepository.Update(ctx, author)
	if err != nil {
		logger.Error(err)
		return nil, parseError(err, "author")
	}

	return author, nil
}

func (a AuthorService) Delete(ctx context.Context, authorID int64) error {
	logger := logrus.
		WithContext(ctx).
		WithField("authorID", authorID)

	err := a.authorRepository.Delete(ctx, authorID)
	if err != nil {
		logger.Error(err)
		return parseError(err, "author")
	}

	return nil
}
