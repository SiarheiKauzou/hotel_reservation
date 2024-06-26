package main

import (
	"context"
	"flag"
	"log"

	"github.com/SiarheiKauzou/hotel_reservation/api"
	"github.com/SiarheiKauzou/hotel_reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	listenAddr := flag.String("listenAddr", ":8080", "The listen address of the API server")
	flag.Parse()

	// handlers initialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client, db.DBNAME))

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	apiv1.Put("/user/:id", userHandler.HandlerUpdateUser)
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandlerGetUsers)
	apiv1.Get("/user/:id", userHandler.HandlerGetUser)
	apiv1.Delete("/user/:id", userHandler.HandlerDeleteUser)
	app.Listen(*listenAddr)
}
