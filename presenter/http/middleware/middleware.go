package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewMiddleware middlewareの設定を行います.
func NewMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
}
