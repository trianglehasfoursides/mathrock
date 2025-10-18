package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

// redirect the user to the provider
func Redirect(c echo.Context) error {
	gothic.BeginAuthHandler(c.Response(), c.Request())
	return nil
}
