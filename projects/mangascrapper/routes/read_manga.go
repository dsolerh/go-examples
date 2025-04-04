package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ReadManga(state *state) fiber.Handler {
	return func(c *fiber.Ctx) error {
		slug := c.Params("slug")
		chap, err := c.ParamsInt("chap", 1)
		if err != nil {
			return c.Render("read-manga", fiber.Map{
				"MangaName": state.mangas[slug].Name,
				"Error":     "invalid chapter",
			})
		}

		manga, exist := state.mangas[slug]
		if !exist {
			return c.Render("read-manga", fiber.Map{
				"MangaName": "",
				"Error":     "invalid manga",
			})
		}

		if len(manga.Chapters) <= chap {
			return c.Render("read-manga", fiber.Map{
				"MangaName": manga.Name,
				"Error":     "invalid chapter",
			})
		}
		chapter := manga.Chapters[chap]
		isLast := len(manga.Chapters) == chap+1

		return c.Render("read-manga", fiber.Map{
			"MangaName": manga.Name,
			"Chapter":   chapter.ChapNo,
			"Images":    chapter.Images,
			"IsLast":    isLast,
			"NextUrl":   fmt.Sprintf("/read-manga/%s/%d", slug, chap+1),
		})
	}
}
