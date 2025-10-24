package storage

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	sg "github.com/trianglehasfoursides/mathrock/db/model/storage"
	"github.com/trianglehasfoursides/mathrock/internal"
	"github.com/trianglehasfoursides/mathrock/rock"
	"github.com/trianglehasfoursides/mathrock/valid"
)

func New(c echo.Context) error {
	if contype := c.Request().Header.Get(""); contype != "" {
		return echo.NewHTTPError(http.StatusBadRequest, "")
	}

	strg := new(sg.Storage)
	if err := c.Bind(strg); err != nil {
		return internal.Deferr(http.StatusBadRequest, "", err)
	}

	if err := valid.Valid.Struct(strg); err != nil {
		return internal.Deferr(http.StatusBadRequest, "", err)
	}

	if _, exist := rock.Rock.Get("storage:" + strg.Name); exist {
		return echo.NewHTTPError(http.StatusConflict, "")
	}

	strg.UserID = auth.UserId(c)
	if err := db.DB.Create(strg); err != nil {
		return internal.Deferr(http.StatusInternalServerError, "", err.Error)
	}

	_ = rock.Rock.Set("storage:"+strg.Name, "", 0)
	return nil
}
