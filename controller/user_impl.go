package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rhtyx/bayarind-service.git/dto"
	"github.com/rhtyx/bayarind-service.git/model"
	"github.com/rhtyx/bayarind-service.git/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (c Controller) CreateUser(e echo.Context) error {
	ctx := e.Request().Context()
	logger := logrus.WithContext(ctx)

	body := &dto.UserRequest{}
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

	user := &model.User{
		Username: body.Username,
		Password: body.Password,
	}
	user, err = c.userService.Create(ctx, user)
	if err != nil {
		logger.WithField("user", utils.Dump(user)).Error(err)
		return parseError(e, err)
	}

	return e.JSON(http.StatusCreated, user)
}

func (c Controller) FindUserByID(e echo.Context) error {
	ctx := e.Request().Context()
	logger := logrus.WithContext(ctx)

	userID, ok := e.Get("userID").(int64)
	if !ok {
		return e.JSON(http.StatusInternalServerError, ErrInternalServer.Error())
	}

	user, err := c.userService.FindByID(ctx, userID)
	if err != nil {
		logger.WithField("userID", userID).Error(err)
		return parseError(e, err)
	}

	return e.JSON(http.StatusOK, user)
}

func (c Controller) UpdateUser(e echo.Context) error {
	ctx := e.Request().Context()
	logger := logrus.WithContext(ctx)

	userID, ok := e.Get("userID").(int64)
	if !ok {
		return e.JSON(http.StatusInternalServerError, ErrInternalServer.Error())
	}

	body := &dto.UserRequest{}
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

	user := &model.User{
		ID:       userID,
		Username: body.Username,
		Password: body.Password,
	}
	user, err = c.userService.Update(ctx, user)
	if err != nil {
		logger.WithField("user", utils.Dump(user)).Error(err)
		return parseError(e, err)
	}

	return e.JSON(http.StatusOK, user)
}

func (c Controller) DeleteUser(e echo.Context) error {
	ctx := e.Request().Context()
	logger := logrus.WithContext(ctx)

	userID, ok := e.Get("userID").(int64)
	if !ok {
		return e.JSON(http.StatusInternalServerError, ErrInternalServer.Error())
	}

	err := c.userService.Delete(ctx, userID)
	if err != nil {
		logger.WithField("userID", userID).Error(err)
		return parseError(e, err)
	}

	return e.JSON(http.StatusOK, "User deleted")
}
