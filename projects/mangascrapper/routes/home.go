package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type mangaItem struct {
	Name string
	Url  string
}

func HomePage(state *state) fiber.Handler {
	return func(c *fiber.Ctx) error {
		items := make([]mangaItem, 0, len(state.mangas))
		for slug, manga := range state.mangas {
			mItem := mangaItem{
				Name: manga.Name,
				Url:  fmt.Sprintf("/read-manga/%s/%d", slug, state.lastReadChap[slug]),
			}
			items = append(items, mItem)
		}

		// Render index template
		return c.Render("index", fiber.Map{
			"Mangas": items,
		})
	}
}
