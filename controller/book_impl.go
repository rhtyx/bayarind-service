package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rhtyx/bayarind-service.git/dto"
	"github.com/rhtyx/bayarind-service.git/model"
	"github.com/rhtyx/bayarind-service.git/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (c Controller) CreateBook(e echo.Context) error {
	ctx := e.Request().Context()
	logger := logrus.WithContext(ctx)

	body := &dto.BookRequest{}
	err := json.NewDecoder(e.Request().Body).Decode(body)
	if err != nil {
		logger.Error(err)
		return e.JSON(http.StatusBadRequest, ErrBadRequest.Error())
	}

	validate := validator.New()
	err = validate.Struct(body)
	if err != nil {
		logger.WithField("body", utils.Dump(body)).Error(err)
		return e.JSON(http.StatusBadRequest, utils.ParseValidationError(err))
	}

	book := &model.Book{
		ISBN:     body.ISBN,
		Title:    body.Title,
		AuthorID: body.AuthorID,
	}
	book, err = c.bookService.Create(ctx, book)
	if err != nil {
		logger.Error(err)
		return parseError(e, err)
	}

	return e.JSON(http.StatusCreated, book)
}

func (c Controller) FindBookByID(e echo.Context) error {
	ctx := e.Request().Context()
	logger := logrus.WithContext(ctx)

	bookID, err := strconv.ParseInt(e.Param("id"), 10, 64)
	if err != nil {
		logger.WithField("bookID", e.Param("id")).Error(err)
		return e.JSON(http.StatusBadRequest, fmt.Errorf("%s: invalid param id", ErrBadRequest.Error()))
	}

	book, err := c.bookService.FindByID(ctx, bookID)
	if err != nil {
		logger.WithField("bookID", bookID).Error(err)
		return parseError(e, err)
	}

	return e.JSON(http.StatusOK, book)
}

func (c Controller) FindAllBooks(e echo.Context) error {
	ctx := e.Request().Context()
	logger := logrus.WithContext(ctx)

	books, err := c.bookService.FindAll(ctx)
	if err != nil {
		logger.Error(err)
		return parseError(e, err)
	}

	return e.JSON(http.StatusOK, books)
}

func (c Controller) UpdateBook(e echo.Context) error {
	ctx := e.Request().Context()
	logger := logrus.WithContext(ctx)

	bookID, err := strconv.ParseInt(e.Param("id"), 10, 64)
	if err != nil {
		logger.WithField("bookID", e.Param("id")).Error(err)
		return e.JSON(http.StatusBadRequest, fmt.Errorf("%s: invalid param id", ErrBadRequest.Error()))
	}

	body := &dto.BookRequest{}
	err = json.NewDecoder(e.Request().Body).Decode(body)
	if err != nil {
		logger.Error(err)
		return e.JSON(http.StatusBadRequest, ErrBadRequest.Error())
	}

	validate := validator.New()
	err = validate.Struct(body)
	if err != nil {
		logger.WithField("body", utils.Dump(body)).Error(err)
		return e.JSON(http.StatusBadRequest, utils.ParseValidationError(err))
	}

	book := &model.Book{
		ID:       bookID,
		ISBN:     body.ISBN,
		Title:    body.Title,
		AuthorID: body.AuthorID,
	}
	book, err = c.bookService.Update(ctx, book)
	if err != nil {
		logger.WithField("book", utils.Dump(book)).Error(err)
		return parseError(e, err)
	}

	return e.JSON(http.StatusOK, book)
}

func (c Controller) DeleteBook(e echo.Context) error {
	ctx := e.Request().Context()
	logger := logrus.WithContext(ctx)

	bookID, err := strconv.ParseInt(e.Param("id"), 10, 64)
	if err != nil {
		logger.WithField("bookID", e.Param("id")).Error(err)
		return e.JSON(http.StatusBadRequest, fmt.Errorf("%s: invalid param id", ErrBadRequest.Error()))
	}

	err = c.bookService.Delete(ctx, bookID)
	if err != nil {
		logger.WithField("bookID", bookID).Error(err)
		return parseError(e, err)
	}

	return e.JSON(http.StatusOK, "Book deleted")
}
