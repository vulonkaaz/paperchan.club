package router

import (
	"paperchan.club/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	app.Get("/", controllers.ThreadList)
	app.Get("/thread/:id", controllers.Thread)
	app.Post("/api/post", controllers.Publish)
	app.Delete("/api/post", controllers.Delete)
}
