package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rinku/db"
	"rinku/handlers"
	ui "rinku/ui"
)

func main() {
	app := fiber.New()

	sqliteDb := db.NewDatabase("url.db")

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(ui.EmbedDirStatic),
		PathPrefix: "dist",
	}))

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
