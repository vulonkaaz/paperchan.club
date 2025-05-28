package router

import (
	"paperchan.club/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
	"strings"
)

func SetRoutes(app *fiber.App) {
	app.Get("/", controllers.ThreadList)
	app.Get("/thread/:id", controllers.Thread)
	app.Post("/api/post", rateLimiter, controllers.Publish)
	app.Delete("/api/post", controllers.Delete)
}

var rateLimiter = limiter.New(limiter.Config{
	Max: 5,
	Expiration: 10*time.Minute,
	KeyGenerator: func(c *fiber.Ctx) string {
		return c.IP()+" "+strings.Join(c.IPs()," ")
	},
})
