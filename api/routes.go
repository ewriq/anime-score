package main

import (
	Handler "auth-api/Handler"
	Middleware "auth-api/Middleware"
	"auth-api/Routes"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func Initalize(app *fiber.App){
	app.Use(Middleware.Cors)
	app.Use(Middleware.RateLimit)
	app.Use(helmet.New())
	app.Use(Middleware.Compress)

	auth := app.Group("/api/auth/")
	anime := app.Group("/api/anime/")
	vote := app.Group("/api/vote/")

	app.Get("/", Handler.Home)

	Routes.Auth(auth)
	Routes.Anime(anime)
	Routes.Vote(vote)
	
	app.Use(Middleware.Notfound)

}