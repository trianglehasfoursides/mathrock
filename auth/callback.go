package auth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
	"github.com/trianglehasfoursides/mathrock/db"
)

func Callback(c echo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := db.DB.Create(&User{
		Name:  user.Name,
		Email: user.Email,
	}).Error; err != nil {
		// TODO check the error first
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	token, err := generate(user.Email)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	url := fmt.Sprintf("http://localhost:9000/callback/%s", token)
	if err := c.Redirect(http.StatusTemporaryRedirect, url); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
