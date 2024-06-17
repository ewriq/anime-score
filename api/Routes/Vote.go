package Routes

import (
	Handler "auth-api/Handler"
	"github.com/gofiber/fiber/v2"
)

func Vote(app fiber.Router) {
	app.Post("/list", Handler.ListVote)
	app.Post("/add", Handler.VoteAdd)
	app.Post("/del", Handler.VoteDel)
}