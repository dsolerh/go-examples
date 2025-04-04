package main

import (
	"mangascrapper/mangas"
	"mangascrapper/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

// var beginning = []byte(`decoding="async" data-src="`)
// var ends = []byte(`.webp"`)
// url := "https://thebeginningaftertheendmanga.com/manga/the-beginning-after-the-end-ch-24/"

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./pages", ".html"),
	})
	app.Use(cors.New())

	mangas, err := mangas.LoadMangas()
	if err != nil {
		log.Fatal(err)
	}

	routes.MountRoutes(app, mangas)

	log.Fatal(app.Listen(":8081"))
}
