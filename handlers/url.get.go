package handlers

import (
	"github.com/gofiber/fiber/v2"
	"rinku/db"
)

func (u *url) Get(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	if id == "" {
		return fiber.ErrNotFound
	}

	var stored db.ShortLink

	res := u.db.Where("short_id = ?", id).Find(&stored)

	if res.RowsAffected == 0 {
		return fiber.ErrNotFound
	}

	return ctx.Redirect(stored.OriginLink)
}
