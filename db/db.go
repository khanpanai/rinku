package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewDatabase(dbName string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&ShortLink{})

	return db
}
