package api

import (
	"github.com/labstack/echo/v4"
)

//Resource resource api operations
type Resource interface {
	Store(context echo.Context) error
	Get(context echo.Context) error
	Delete(context echo.Context) error
}
