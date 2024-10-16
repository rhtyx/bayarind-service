package controller

import (
	"encoding/hex"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rhtyx/bayarind-service.git/token"
)

func HmacMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		hmacString := c.Request().Header.Get("X-HMAC")
		if hmacString == "" {
			return c.JSON(http.StatusBadRequest, ErrBadRequest)
		}

		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrInternalServer)
		}

		digest, err := hex.DecodeString(hmacString)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrInternalServer)
		}

		if !token.Hmac.ValidMAC(body, digest) {
			return c.JSON(http.StatusBadRequest, ErrBadRequest)
		}

		return next(c)
	}
}
