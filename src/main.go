package main

import (
	"app/server/http"
	"app/svc"
	"app/svc/factory"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	svc.NewServiceContext()
	svc.ServiceComm.Service = factory.NewServiceFactory(svc.ServiceComm)

	app := fiber.New()

	http.InitRouter(app)

	log.Fatal(app.Listen(":3000"))
}
