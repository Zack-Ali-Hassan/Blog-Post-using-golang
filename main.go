package main

import (
	"BLOG-APP/server/controller"
	"BLOG-APP/server/database"
	"BLOG-APP/server/routes"

	"log"
	"os"

	// "github.com/gofiber/fiber/v2"
	"github.com/gin-gonic/gin"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if os.Getenv("ENV") != "production" {
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	Db := database.ConnectDb()
	controller.SetCollection(Db.Collection("Blogs"))
	app := gin.Default()
	app.RedirectTrailingSlash = false
	// Use Gin CORS middleware with the desired config
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},            // Allow specific origin
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},   // Allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"}, // Allowed headers
		AllowCredentials: true,                                         // Allow credentials like cookies, etc.
	}))
	app.Use(gin.Logger())

	routes.SetupRoutes(app)
	if os.Getenv("ENV") == "production" {
		app.Static("/", "./client/dist")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "4411"
	}
	log.Fatal(app.Run(":" + port))
}
