package controllers

import (
	"log"
	"paperchan.club/database"
	"paperchan.club/themagicpipe"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"strings"
	"strconv"
)

// data received by /api/post
type PostApi struct {
    Picture string `json:"picture" xml:"picture" form:"picture"`
    Thread string `json:"thread" xml:"thread" form:"thread"`
}

func Publish(c *fiber.Ctx) error {
	p := new(PostApi)
	if err := c.BodyParser(p); err != nil {
		return c.JSON(fiber.Map{
			"status": "error",
		})
   }
	picture := p.Picture
	var thread sql.NullInt32
	if parsed, err := strconv.ParseInt(p.Thread, 10, 32); err != nil {
		thread.Valid = false
	} else {
		thread.Int32 = int32(parsed)
		thread.Valid = true
	}
	ip := c.IP()+" "+strings.Join(c.IPs()," ")
	fixedPic, err := themagicpipe.DataURLConverter(picture)
	if err != nil {
		return c.JSON(fiber.Map{
			"status": "error",
		})
	}
	db := database.DB
	if _, err := db.Exec("INSERT INTO \"post\" (picture, ip_address, thread) VALUES ($1, $2, $3)", fixedPic, ip, thread); err == nil {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	} else {
		log.Println(err)
		return c.JSON(fiber.Map{
			"status": "database error",
		})
	}
}
