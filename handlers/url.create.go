package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"rinku/db"
	"rinku/utils"
)

type createLinkRequest struct {
	Link string
}

type createLinkResponse struct {
	Id string
}

func (u *url) Create(ctx *fiber.Ctx) error {
	r := new(createLinkRequest)

	if err := ctx.BodyParser(r); err != nil {
		return fiber.ErrBadRequest
	}

	var stored db.ShortLink

	res := u.db.Where("origin_link = ?", r.Link).Find(&stored)

	if res.RowsAffected > 0 {
		return ctx.JSON(createLinkResponse{Id: stored.ShortId})
	}

	shortId := utils.GenerateShortId()

	stored = db.ShortLink{
		ShortId:    shortId,
		OriginLink: r.Link,
		Expires:    nil,
	}

	u.db.Create(&stored)
	fmt.Println(stored)
	return ctx.JSON(createLinkResponse{Id: shortId})
}
