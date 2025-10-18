package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// the login page
func Page(c echo.Context) error {
	if err := c.File("html/auth.html"); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
