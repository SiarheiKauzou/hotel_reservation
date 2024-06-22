package main

import (
	"flag"

	"github.com/SiarheiKauzou/hotel_reservation/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8080", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandlerGetUsers)
	apiv1.Get("/user/:id", api.HandlerGetUser)
	app.Listen(*listenAddr)
}
