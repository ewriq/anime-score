package Routes

import (
	Handler "auth-api/Handler"
	"github.com/gofiber/fiber/v2"
)

func Anime(app fiber.Router) {
	app.Get("/list", Handler.AnimeList)
	app.Post("/add", Handler.AnimeAdd)
	app.Post("/del", Handler.AnimeDel)
	app.Post("/", Handler.Anime)
	app.Post("/search", Handler.Search)
}