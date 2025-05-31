package controllers

import (
	"paperchan.club/database"
	"paperchan.club/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"math/rand"
)

const maxThreadsPerPage = 24
const maxBanner = 12 // dont forget to increment whenever we add more

func ThreadList(c *fiber.Ctx) error {
	db := database.DB
	var nthread []int
	if err := db.Select(&nthread, "SELECT COUNT(*) FROM \"post\" WHERE \"thread\" IS NULL"); err != nil {
		log.Println(err)
		return c.Status(500).SendString("ERROR")
	}
	npage := (nthread[0] + maxThreadsPerPage - 1) / maxThreadsPerPage
	var pages []int
	for i := 1; i <= npage; i++ {
		pages = append(pages, i)
	}

	page,_ := c.ParamsInt("page")
	if page < 1 {
		page = 1
	}
	offset := maxThreadsPerPage * (page - 1)
	var threads []models.Thread
	if err := db.Select(&threads, "SELECT a.*, (SELECT COUNT(*) FROM \"post\" AS b WHERE b.thread = a.id) AS replies FROM \"post\" AS \"a\" WHERE \"thread\" IS NULL ORDER BY (SELECT c.created_at FROM \"post\" AS c WHERE c.thread = a.id OR c.id = a.id ORDER BY c.created_at DESC LIMIT 1) DESC LIMIT $1 OFFSET $2", maxThreadsPerPage, offset); err != nil {
		log.Println(err)
		return c.Status(500).SendString("ERROR")
	}
	return c.Render("index", fiber.Map{
		"posts": threads,
		"pages": pages,
		"banner": rand.Intn(maxBanner)+1,
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
		"banner": rand.Intn(maxBanner)+1,
	})
}
