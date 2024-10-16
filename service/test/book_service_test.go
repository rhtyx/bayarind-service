package test

import (
	"context"
	"testing"

	"go.uber.org/mock/gomock"
	"gorm.io/gorm"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/rhtyx/bayarind-service.git/controller"
	"github.com/rhtyx/bayarind-service.git/model"
	"github.com/rhtyx/bayarind-service.git/model/mock"
	"github.com/rhtyx/bayarind-service.git/service"
	"github.com/rhtyx/bayarind-service.git/utils"
	"github.com/stretchr/testify/assert"
)

func TestBookCreate(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		book := &model.Book{
			ID:       utils.GenerateID(),
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		author := &model.Author{
			ID:        utils.GenerateID(),
			Name:      gofakeit.Name(),
			BirthDate: gofakeit.Date(),
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindByISBN(ctx, book.ISBN).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		authorRepository.EXPECT().
			FindByID(ctx, book.AuthorID).
			Times(1).
			Return(author, nil)

		bookRepository.EXPECT().
			Create(ctx, book).
			Times(1).
			Return(book, nil)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBook, err := bookService.Create(ctx, book)
		assert.Nil(t, err)
		assert.NotNil(t, resBook)
		assert.ObjectsAreEqualValues(book, resBook)
	})

	t.Run("error: find isbn", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		book := &model.Book{
			ID:       utils.GenerateID(),
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindByISBN(ctx, book.ISBN).
			Times(1).
			Return(nil, gorm.ErrInvalidDB)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBook, err := bookService.Create(ctx, book)
		assert.Nil(t, resBook)
		assert.Error(t, err)
		assert.EqualError(t, err, controller.ErrInternalServer.Error())
	})

	t.Run("error: isbn duplicated", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		book := &model.Book{
			ID:       utils.GenerateID(),
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		bookDuplicate := &model.Book{
			ID:       utils.GenerateID(),
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindByISBN(ctx, book.ISBN).
			Times(1).
			Return(bookDuplicate, nil)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBook, err := bookService.Create(ctx, book)
		assert.Nil(t, resBook)
		assert.Error(t, err)
		assert.EqualError(t, err, "duplicate entry\n: isbn")
	})

	t.Run("error: find author", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		book := &model.Book{
			ID:       utils.GenerateID(),
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindByISBN(ctx, book.ISBN).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		authorRepository.EXPECT().
			FindByID(ctx, book.AuthorID).
			Times(1).
			Return(nil, gorm.ErrInvalidDB)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBook, err := bookService.Create(ctx, book)
		assert.Nil(t, resBook)
		assert.Error(t, err)
		assert.EqualError(t, err, controller.ErrInternalServer.Error())
	})

	t.Run("error: author not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		book := &model.Book{
			ID:       utils.GenerateID(),
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindByISBN(ctx, book.ISBN).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		authorRepository.EXPECT().
			FindByID(ctx, book.AuthorID).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBook, err := bookService.Create(ctx, book)
		assert.Nil(t, resBook)
		assert.Error(t, err)
		assert.EqualError(t, err, "id not found\n: author")
	})

	t.Run("error: create book", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		book := &model.Book{
			ID:       utils.GenerateID(),
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		author := &model.Author{
			ID:        utils.GenerateID(),
			Name:      gofakeit.Name(),
			BirthDate: gofakeit.Date(),
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindByISBN(ctx, book.ISBN).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		authorRepository.EXPECT().
			FindByID(ctx, book.AuthorID).
			Times(1).
			Return(author, nil)

		bookRepository.EXPECT().
			Create(ctx, book).
			Times(1).
			Return(nil, gorm.ErrInvalidDB)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBook, err := bookService.Create(ctx, book)
		assert.Nil(t, resBook)
		assert.Error(t, err)
		assert.EqualError(t, err, controller.ErrInternalServer.Error())
	})
}

func TestBookFindByID(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		book := &model.Book{
			ID:       utils.GenerateID(),
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindByID(ctx, book.ID).
			Times(1).
			Return(book, nil)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBook, err := bookService.FindByID(ctx, book.ID)
		assert.Nil(t, err)
		assert.NotNil(t, resBook)
		assert.ObjectsAreEqualValues(book, resBook)
	})

	t.Run("error: id not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		book := &model.Book{
			ID:       utils.GenerateID(),
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindByID(ctx, book.ID).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBook, err := bookService.FindByID(ctx, book.ID)
		assert.Nil(t, resBook)
		assert.Error(t, err)
		assert.EqualError(t, err, "id not found\n: book")
	})
}

func TestBookFindByISBN(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		book := &model.Book{
			ID:       utils.GenerateID(),
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindByISBN(ctx, book.ISBN).
			Times(1).
			Return(book, nil)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBook, err := bookService.FindByISBN(ctx, book.ISBN)
		assert.Nil(t, err)
		assert.NotNil(t, resBook)
		assert.ObjectsAreEqualValues(book, resBook)
	})

	t.Run("error: isbn not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		book := &model.Book{
			ID:       utils.GenerateID(),
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindByISBN(ctx, book.ISBN).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBook, err := bookService.FindByISBN(ctx, book.ISBN)
		assert.Nil(t, resBook)
		assert.Error(t, err)
		assert.EqualError(t, err, "id not found\n: book")
	})
}

func TestBookFindAll(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		books := []*model.Book{
			{
				ID:       utils.GenerateID(),
				ISBN:     "9789295055025",
				Title:    gofakeit.BookTitle(),
				AuthorID: utils.GenerateID(),
			},
			{
				ID:       utils.GenerateID(),
				ISBN:     "9780323776714",
				Title:    gofakeit.BookTitle(),
				AuthorID: utils.GenerateID(),
			},
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindAll(ctx).
			Times(1).
			Return(books, nil)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBooks, err := bookService.FindAll(ctx)
		assert.Nil(t, err)
		assert.NotNil(t, resBooks)
		assert.ObjectsAreEqualValues(books, resBooks)
	})

	t.Run("error: find all", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindAll(ctx).
			Times(1).
			Return(nil, gorm.ErrInvalidDB)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBooks, err := bookService.FindAll(ctx)
		assert.Nil(t, resBooks)
		assert.Error(t, err)
		assert.EqualError(t, err, controller.ErrInternalServer.Error())
	})
}

func TestBookUpdate(t *testing.T) {
	t.Run("ok: change all", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		bookID := utils.GenerateID()
		authorID := utils.GenerateID()
		req := &model.Book{
			ID:       bookID,
			ISBN:     "9789353008956",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		book := &model.Book{
			ID:       bookID,
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: authorID,
		}

		author := &model.Author{
			ID:        authorID,
			Name:      gofakeit.Name(),
			BirthDate: gofakeit.Date(),
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindByID(ctx, req.ID).
			Times(1).
			Return(book, nil)

		bookRepository.EXPECT().
			FindByISBN(ctx, req.ISBN).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		authorRepository.EXPECT().
			FindByID(ctx, req.AuthorID).
			Times(1).
			Return(author, nil)

		bookRepository.EXPECT().
			Update(ctx, req).
			Times(1).
			Return(req, nil)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBook, err := bookService.Update(ctx, req)
		assert.Nil(t, err)
		assert.NotNil(t, resBook)
		assert.ObjectsAreEqualValues(req, resBook)
	})

	t.Run("error: find id", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		req := &model.Book{
			ID:       utils.GenerateID(),
			ISBN:     "9789353008956",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindByID(ctx, req.ID).
			Times(1).
			Return(nil, gorm.ErrInvalidDB)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBook, err := bookService.Update(ctx, req)
		assert.Nil(t, resBook)
		assert.Error(t, err)
		assert.EqualError(t, err, controller.ErrInternalServer.Error())
	})

	t.Run("error: find isbn", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		bookID := utils.GenerateID()
		authorID := utils.GenerateID()
		req := &model.Book{
			ID:       bookID,
			ISBN:     "9789353008956",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		book := &model.Book{
			ID:       bookID,
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: authorID,
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindByID(ctx, req.ID).
			Times(1).
			Return(book, nil)

		bookRepository.EXPECT().
			FindByISBN(ctx, req.ISBN).
			Times(1).
			Return(nil, gorm.ErrInvalidDB)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBook, err := bookService.Update(ctx, req)
		assert.Nil(t, resBook)
		assert.Error(t, err)
		assert.EqualError(t, err, controller.ErrInternalServer.Error())
	})

	t.Run("error: isbn duplicated", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		bookID := utils.GenerateID()
		authorID := utils.GenerateID()
		req := &model.Book{
			ID:       bookID,
			ISBN:     "9789353008956",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		book := &model.Book{
			ID:       bookID,
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: authorID,
		}

		bookByISBN := &model.Book{
			ID:       utils.GenerateID(),
			ISBN:     "9789353008956",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindByID(ctx, req.ID).
			Times(1).
			Return(book, nil)

		bookRepository.EXPECT().
			FindByISBN(ctx, req.ISBN).
			Times(1).
			Return(bookByISBN, nil)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBook, err := bookService.Update(ctx, req)
		assert.Nil(t, resBook)
		assert.Error(t, err)
		assert.EqualError(t, err, "duplicate entry\n: isbn")
	})

	t.Run("error: find author", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		bookID := utils.GenerateID()
		authorID := utils.GenerateID()
		req := &model.Book{
			ID:       bookID,
			ISBN:     "9789353008956",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		book := &model.Book{
			ID:       bookID,
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: authorID,
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindByID(ctx, req.ID).
			Times(1).
			Return(book, nil)

		bookRepository.EXPECT().
			FindByISBN(ctx, req.ISBN).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		authorRepository.EXPECT().
			FindByID(ctx, req.AuthorID).
			Times(1).
			Return(nil, gorm.ErrInvalidDB)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBook, err := bookService.Update(ctx, req)
		assert.Nil(t, resBook)
		assert.Error(t, err)
		assert.EqualError(t, err, controller.ErrInternalServer.Error())
	})

	t.Run("error: author not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		bookID := utils.GenerateID()
		authorID := utils.GenerateID()
		req := &model.Book{
			ID:       bookID,
			ISBN:     "9789353008956",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		book := &model.Book{
			ID:       bookID,
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: authorID,
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindByID(ctx, req.ID).
			Times(1).
			Return(book, nil)

		bookRepository.EXPECT().
			FindByISBN(ctx, req.ISBN).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		authorRepository.EXPECT().
			FindByID(ctx, req.AuthorID).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBook, err := bookService.Update(ctx, req)
		assert.Nil(t, resBook)
		assert.Error(t, err)
		assert.EqualError(t, err, "id not found\n: author")
	})

	t.Run("error: create book", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		bookID := utils.GenerateID()
		authorID := utils.GenerateID()
		req := &model.Book{
			ID:       bookID,
			ISBN:     "9789353008956",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		book := &model.Book{
			ID:       bookID,
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: authorID,
		}

		author := &model.Author{
			ID:        authorID,
			Name:      gofakeit.Name(),
			BirthDate: gofakeit.Date(),
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			FindByID(ctx, req.ID).
			Times(1).
			Return(book, nil)

		bookRepository.EXPECT().
			FindByISBN(ctx, req.ISBN).
			Times(1).
			Return(nil, gorm.ErrRecordNotFound)

		authorRepository.EXPECT().
			FindByID(ctx, req.AuthorID).
			Times(1).
			Return(author, nil)

		bookRepository.EXPECT().
			Update(ctx, req).
			Times(1).
			Return(nil, gorm.ErrInvalidDB)

		bookService := service.NewBookService(bookRepository, authorRepository)
		resBook, err := bookService.Update(ctx, req)
		assert.Nil(t, resBook)
		assert.Error(t, err)
		assert.EqualError(t, err, controller.ErrInternalServer.Error())
	})
}

func TestBookDelete(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		book := &model.Book{
			ID:       utils.GenerateID(),
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			Delete(ctx, book.ID).
			Times(1).
			Return(nil)

		bookService := service.NewBookService(bookRepository, authorRepository)
		err := bookService.Delete(ctx, book.ID)
		assert.Nil(t, err)
	})

	t.Run("error: id not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		ctx := context.TODO()
		book := &model.Book{
			ID:       utils.GenerateID(),
			ISBN:     "9789295055025",
			Title:    gofakeit.BookTitle(),
			AuthorID: utils.GenerateID(),
		}

		bookRepository := mock.NewMockBookRepository(ctrl)
		authorRepository := mock.NewMockAuthorRepository(ctrl)

		bookRepository.EXPECT().
			Delete(ctx, book.ID).
			Times(1).
			Return(gorm.ErrRecordNotFound)

		bookService := service.NewBookService(bookRepository, authorRepository)
		err := bookService.Delete(ctx, book.ID)
		assert.Error(t, err)
		assert.EqualError(t, err, "id not found\n: book")
	})
}
