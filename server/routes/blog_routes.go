package routes

import (
	"BLOG-APP/server/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/post/", controller.GetAllPosts)
	app.Get("/api/post/:id", controller.GetPost)
	app.Post("/api/post/", controller.CreatePost)
	app.Patch("/api/post/:id", controller.UpdatePost)
	app.Delete("/api/post/:id", controller.DeletePost)
}
