package routes

import (
	"mangascrapper/mangas"

	"github.com/gofiber/fiber/v2"
)

type state struct {
	mangas       map[string]mangas.Manga
	lastReadChap map[string]int
}

func MountRoutes(app *fiber.App, mangas map[string]mangas.Manga) {
	state := &state{
		mangas:       mangas,
		lastReadChap: map[string]int{},
	}
	app.Get("/", HomePage(state))
	app.Get("/read-manga/:slug/:chap", ReadManga(state))
}
