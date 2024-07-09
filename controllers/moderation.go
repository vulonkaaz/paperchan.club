package controllers

import (
	"log"
	"os"
	"paperchan.club/database"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

// data received by DELETE /api/post
type DeleteApi struct {
    Id string `json:"id" xml:"id" form:"id"`
    Pass string `json:"pass" xml:"pass" form:"pass"` //moderation password defined in .env
}
func Delete(c *fiber.Ctx) error {
	d := new(DeleteApi)
	if err := c.BodyParser(d); err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{
			"status": "error",
		})
   }
   id, err := strconv.ParseInt(d.Id, 10, 32)
   if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{
			"status": "error",
		})
   }
   if bcrypt.CompareHashAndPassword([]byte(os.Getenv("MODPASS")),[]byte(d.Pass)) != nil {
		log.Println(os.Getenv("MODPASS"))
		log.Println(err)
		return c.JSON(fiber.Map{
			"status": "forbidden",
		})
	} else {
		db := database.DB
		if _, err := db.Exec("DELETE FROM \"post\" WHERE id = $1", id); err == nil {
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
}
