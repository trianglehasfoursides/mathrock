package storage

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	sg "github.com/trianglehasfoursides/mathrock/db/model/storage"
	"github.com/trianglehasfoursides/mathrock/internal"
	"github.com/trianglehasfoursides/mathrock/rock"
)

func Upload(c echo.Context) (err error) {
	storage, dir := c.FormValue("storage"), c.FormValue("dir")

	file, err := c.FormFile("file")
	if err != nil {
		return
	}

	if storage == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "")
	}

	if id, _ := rock.Rock.Get("storage:" + storage); id != strconv.Itoa(int(auth.UserId(c))) {
		return echo.NewHTTPError(http.StatusUnauthorized, "")
	}

	directory := new(sg.Directory)
	if err = db.DB.Where("name = ?", dir).First(directory).Error; err != nil {
	}

	if directory.UserID != auth.UserId(c) {
	}

	return
}
