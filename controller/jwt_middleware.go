package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rhtyx/bayarind-service.git/token"
)

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, ErrUnauthorized.Error())
		}

		claims, err := token.Jwt.ValidateToken(tokenString[7:])
		if err != nil {
			return c.JSON(http.StatusUnauthorized, ErrUnauthorized.Error())
		}

		c.Set("userID", claims.UserID)
		return next(c)
	}
}
