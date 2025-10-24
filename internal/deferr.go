package internal

import (
	"os"

	"github.com/labstack/echo/v4"
)

func Deferr(code int, msg string, err error) *echo.HTTPError {
	if os.Getenv("STATUS") == "DEBUG" {
		return echo.NewHTTPError(
			code,
			err.Error(),
		)
	}

	return echo.NewHTTPError(
		code,
		msg,
	)
}
