package main

import (
	"crud-api/routes"
	"crud-api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)



var db *gorm.DB

func main() {

	database, err := gorm.Open(sqlite.Open("data.sqlite3"), &gorm.Config{})
	if err != nil {
		panic("failed ...........")
	}
	db = database

	db.AutoMigrate(&models.UserRequest{})

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	routes.DefineRoutes(app, db)

	app.Listen(":3000")
}
