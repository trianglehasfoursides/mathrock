package storage

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/trianglehasfoursides/mathrock/rock"
)

func Upload(c echo.Context) error {
	name := c.Request().
		Header.
		Get("X-Name")

	if name == "" {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"",
		)
	}

	if _, exist := rock.Rock.Get(name); exist {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"",
		)
	}

	buf := make([]byte, 512)
	n, _ := c.Request().Body.Read(buf)
	detected := http.DetectContentType(buf[:n])
}
