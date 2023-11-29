package db

import (
	"gorm.io/gorm"
	"time"
)

type ShortLink struct {
	gorm.Model
	ShortId    string
	OriginLink string
	Expires    *time.Time
}
