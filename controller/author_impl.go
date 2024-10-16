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

func (c Controller) CreateAuthor(e echo.Context) error {
	ctx := e.Request().Context()
	ctx.Value("ini")
	logger := logrus.WithContext(ctx)

	body := &dto.AuthorRequest{}
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

	birthDate, err := utils.ParseDate(body.BirthDate)
	if err != nil {
		logger.WithField("body", utils.Dump(body)).Error(err)
		return e.JSON(http.StatusInternalServerError, ErrInternalServer.Error())
	}

	author := &model.Author{
		Name:      body.Name,
		BirthDate: *birthDate,
	}
	author, err = c.authorService.Create(ctx, author)
	if err != nil {
		logger.WithField("author", utils.Dump(author)).Error(err)
		return parseError(e, err)
	}

	return e.JSON(http.StatusCreated, author)
}

func (c Controller) FindAuthorByID(e echo.Context) error {
	ctx := e.Request().Context()
	logger := logrus.WithContext(ctx)

	authorID, err := strconv.ParseInt(e.Param("id"), 10, 64)
	if err != nil {
		logger.WithField("authorID", e.Param("id")).Error(err)
		return e.JSON(http.StatusBadRequest, fmt.Errorf("%s: invalid param id", ErrBadRequest.Error()))
	}

	author, err := c.authorService.FindByID(ctx, authorID)
	if err != nil {
		logger.WithField("authorID", authorID).Error(err)
		return parseError(e, err)
	}

	return e.JSON(http.StatusOK, author)
}

func (c *Controller) FindAllAuthors(e echo.Context) error {
	ctx := e.Request().Context()
	logger := logrus.WithContext(ctx)

	authors, err := c.authorService.FindAll(ctx)
	if err != nil {
		logger.Error(err)
		return parseError(e, err)
	}

	return e.JSON(http.StatusOK, authors)
}

func (c Controller) UpdateAuthor(e echo.Context) error {
	ctx := e.Request().Context()
	logger := logrus.WithContext(ctx)

	authorID, err := strconv.ParseInt(e.Param("id"), 10, 64)
	if err != nil {
		logger.WithField("authorID", e.Param("id")).Error(err)
		return e.JSON(http.StatusBadRequest, fmt.Errorf("%s: invalid param id", ErrBadRequest.Error()))
	}

	body := &dto.AuthorRequest{}
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

	birthDate, err := utils.ParseDate(body.BirthDate)
	if err != nil {
		logger.WithField("body", utils.Dump(body)).Error(err)
		return e.JSON(http.StatusInternalServerError, ErrInternalServer.Error())
	}

	author := &model.Author{
		ID:        authorID,
		Name:      body.Name,
		BirthDate: *birthDate,
	}
	author, err = c.authorService.Update(ctx, author)
	if err != nil {
		logger.WithField("author", utils.Dump(author)).Error(err)
		return parseError(e, err)
	}

	return e.JSON(http.StatusOK, author)
}

func (c Controller) DeleteAuthor(e echo.Context) error {
	ctx := e.Request().Context()
	logger := logrus.WithContext(ctx)

	authorID, err := strconv.ParseInt(e.Param("id"), 10, 64)
	if err != nil {
		logger.WithField("authorID", e.Param("id")).Error(err)
		return e.JSON(http.StatusBadRequest, fmt.Errorf("%s: invalid param id", ErrBadRequest.Error()))
	}

	err = c.authorService.Delete(ctx, authorID)
	if err != nil {
		logger.WithField("authorID", authorID).Error(err)
		return parseError(e, err)
	}

	return e.JSON(http.StatusOK, "Author deleted")
}
