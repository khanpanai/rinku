package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"rinku/db"
	"rinku/handlers"
)

func main() {
	app := fiber.New()

	sqliteDb := db.NewDatabase("url.db")

	urlHandler := handlers.NewUrlHandler(sqliteDb)
	app.Get("/:id", urlHandler.Get)

	urlGroup := app.Group("url")
	urlGroup.Post("/create", urlHandler.Create)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen(":9000"); err != nil {
		log.Panic(err)
	}
}
