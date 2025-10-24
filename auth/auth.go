package auth

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"

	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
)

func init() {
	// TODO inisialisasi dulu
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "", "profile", "email"),
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), ""),
	)
}

// middleware for checking is the user authenticated
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "missing header")
		}

		token = token[len("Bearer "):]
		if err := verify(token); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		return next(c)
	}
}
