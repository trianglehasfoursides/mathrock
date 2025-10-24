package auth

import (
	"github.com/labstack/echo/v4"
)

func UserId(c echo.Context) uint {
	return c.Get("user_id").(uint)
}
