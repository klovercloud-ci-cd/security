package api

import (
	v1 "github.com/klovercloud-ci-cd/security/api/v1"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"net/http"
)

// Routes base router
func Routes(e *echo.Echo) {
	// Index Page
	e.GET("/", index)

	// Health Page
	e.GET("/health", health)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	v1.Router(e.Group("/api/v1"))
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "This is KloverCloud CI Security Service")
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "I am live!")
}
