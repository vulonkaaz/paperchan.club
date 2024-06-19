package main

import (
	"log"
	"os"
	"paperchan.club/database"
	"paperchan.club/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	database.DBConnect(os.Getenv("DBSTRING"))

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./static")
	router.SetRoutes(app)

	app.Listen(":3000")
}
