package main

import (
	"BLOG-APP/server/controller"
	"BLOG-APP/server/database"
	"BLOG-APP/server/routes"

	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Blog struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title"`
	Content string             `json:"content"`
}

var collection *mongo.Collection

func main() {
	Db := database.ConnectDb()
	controller.SetCollection(Db.Collection("Blogs"))
	app := fiber.New()
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "http://localhost:5173/",
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// }))
	routes.SetupRoutes(app)
	if os.Getenv("ENV") == "production" {
		app.Static("/", "./client/dist")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "4411"
	}
	log.Fatal(app.Listen(":" + port))
}
