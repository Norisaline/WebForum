package router

import (
	"github.com/Norisaline/WebForum.git/controller"
	"github.com/gofiber/fiber/v2"
)

// Настройка машрутов для запросов
func SetupRoutes(app *fiber.App) {

	/*list => get
	  add => post
	  update => put
	  delete => delete
	*/
	app.Get("/", controller.BlogList)
	app.Post("/", controller.BlogCreate)
	app.Put("/:id", controller.BlogUpdate)
	app.Delete("/", controller.BlogDelete)

}
