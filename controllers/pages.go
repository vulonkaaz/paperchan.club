package controllers

import (
	"paperchan.club/database"
	"paperchan.club/models"
	"github.com/gofiber/fiber/v2"
	"log"
)

func ThreadList(c *fiber.Ctx) error {
	var threads []models.Thread
	db := database.DB
	if err := db.Select(&threads, "SELECT a.*, (SELECT COUNT(*) FROM \"post\" AS b WHERE b.thread = a.id) AS replies FROM \"post\" AS \"a\" WHERE \"thread\" IS NULL ORDER BY (SELECT c.created_at FROM \"post\" AS c WHERE c.thread = a.id OR c.id = a.id ORDER BY c.created_at DESC LIMIT 1) DESC"); err != nil {
		log.Println(err)
		return c.Status(500).SendString("ERROR")
	}
	return c.Render("index", fiber.Map{
		"posts": threads,
	})
}

func Thread(c *fiber.Ctx) error {
	id := c.Params("id")
	var posts []models.Post
	db := database.DB
	if err := db.Select(&posts, "SELECT * FROM \"post\" WHERE id = $1 OR thread = $1 ORDER BY created_at ASC", id); err != nil {
		log.Println(err)
		return c.Status(500).SendString("ERROR")
	}
	return c.Render("thread", fiber.Map{
		"threadId": id,
		"posts": posts,
	})
}
