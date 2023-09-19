package routes

import (
	"crud-api/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DefineRoutes(app *fiber.App, db *gorm.DB) {

	app.Get("/", controllers.GetAllEndpoints)
	app.Get("/history", controllers.GetHistory)
	app.Get("/:num1/into/:num2", controllers.Multiplication)
    app.Get("/:num1/divide/:num2", controllers.Division)
    app.Get("/:num1/plus/:num2", controllers.Addition)
    app.Get("/:num1/minus/:num2", controllers.Subtraction)
	app.Get("/:operations", controllers.ComplexOperation)


}
