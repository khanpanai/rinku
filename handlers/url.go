package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UrlHandler interface {
	Get(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}

type url struct {
	db *gorm.DB
}

var _ UrlHandler = (*url)(nil)

func NewUrlHandler(db *gorm.DB) UrlHandler {
	return &url{
		db: db,
	}
}
