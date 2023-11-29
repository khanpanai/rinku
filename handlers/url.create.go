package handlers

import (
	"github.com/gofiber/fiber/v2"
	"rinku/db"
	"rinku/utils"
)

type createLinkRequest struct {
	Link string
}

type createLinkResponse struct {
	Id string `json:"id"`
}

func (u *url) Create(ctx *fiber.Ctx) error {
	r := new(createLinkRequest)

	if err := ctx.BodyParser(r); err != nil {
		ctx.Status(400)
		return ctx.JSON("Bad request")
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
	return ctx.JSON(createLinkResponse{Id: shortId})
}
