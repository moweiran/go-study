package http

import (
	"github.com/gofiber/fiber/v2"

	api "app/api/http/controller"
)

func InitRouter(app *fiber.App) {
	v1 := app.Group("/api/v1")

	helloRouter := v1.Group("/hello")
	{
		helloRouter.Get("/world", api.StudentsApi.HelloWorld)
	}

}
