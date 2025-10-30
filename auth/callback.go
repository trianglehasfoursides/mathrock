package auth

import (
	"github.com/Masterminds/squirrel"
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	auth "github.com/shareed2k/goth_fiber"
	"github.com/trianglehasfoursides/mathrock/db"
)

func callback(ctx *fiber.Ctx) (err error) {
	user, err := auth.CompleteUserAuth(ctx)
	if err != nil {
		return
	}

	query, args, err := squirrel.Insert("users").
		Columns("id", "name", "email", "provider").
		Values(uuid.NewString(), user.Name, user.Email, "").
		ToSql()

	if err != nil {
		log.Error(err.Error())
		return
	}

	if _, err = db.DB.Exec(query, args); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "",
		})
	}

	return
}
