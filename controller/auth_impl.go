package controller

import (
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/rhtyx/bayarind-service.git/dto"
	"github.com/rhtyx/bayarind-service.git/token"
	"github.com/rhtyx/bayarind-service.git/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (c Controller) Login(e echo.Context) error {
	ctx := e.Request().Context()
	logger := logrus.WithContext(ctx)

	body := &dto.LoginRequest{}
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

	session, err := c.sessionService.Create(ctx, body.Username, body.Password)
	if err != nil {
		logger.WithField("body", utils.Dump(body)).Error(err)
		return parseError(e, err)
	}

	hmacSecretKey := hex.EncodeToString(token.Hmac.SecretKey)
	response := &dto.TokenResponse{
		RefreshToken:  session.RefreshToken,
		AccessToken:   session.AccessToken,
		HMACSecretKey: hmacSecretKey,
	}

	return e.JSON(http.StatusOK, response)
}

func (c Controller) Logout(e echo.Context) error {
	ctx := e.Request().Context()
	logger := logrus.WithContext(ctx)

	body := &dto.LogoutRequest{}
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

	err = c.sessionService.DeleteByRefreshToken(ctx, body.RefreshToken)
	if err != nil {
		logger.WithField("body", utils.Dump(body)).Error(err)
		return parseError(e, err)
	}

	return e.JSON(http.StatusOK, "Logged out successful")
}

func (c Controller) RefreshAccessToken(e echo.Context) error {
	ctx := e.Request().Context()
	logger := logrus.WithContext(ctx)

	body := &dto.RefreshTokenRequest{}
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

	session, err := c.sessionService.RefreshAccessToken(ctx, body.RefreshToken)
	if err != nil {
		logger.WithField("body", utils.Dump(body)).Error(err)
		return parseError(e, err)
	}

	return e.JSON(
		http.StatusOK,
		dto.TokenResponse{
			RefreshToken: session.RefreshToken,
			AccessToken:  session.AccessToken,
		})
}
