package controller

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrBadRequest     = errors.New("bad request")
	ErrInternalServer = errors.New("internal server error")
	ErrNotFound       = errors.New("id not found")
	ErrDuplicate      = errors.New("duplicate entry")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrCredentials    = errors.New("wrong username or password")
)

func parseError(e echo.Context, err error) error {
	switch {
	case errors.Is(err, ErrBadRequest):
		return e.JSON(http.StatusBadRequest, err.Error())
	case errors.Is(err, ErrInternalServer):
		return e.JSON(http.StatusInternalServerError, err.Error())
	case errors.Is(err, ErrNotFound):
		return e.JSON(http.StatusNotFound, err.Error())
	case errors.Is(err, ErrDuplicate):
		return e.JSON(http.StatusConflict, err.Error())
	case errors.Is(err, ErrCredentials):
		return e.JSON(http.StatusUnauthorized, err.Error())
	default:
		return e.JSON(http.StatusInternalServerError, ErrInternalServer)
	}
}
