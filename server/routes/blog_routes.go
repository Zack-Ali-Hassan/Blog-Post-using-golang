package routes

import (
	"BLOG-APP/server/controller"

	"github.com/gin-gonic/gin"
	// "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/api/post/", controller.GetAllPosts)
	router.GET("/api/post/:id", controller.GetPost)
	router.POST("/api/post/", controller.CreatePost)
	router.PATCH("/api/post/:id", controller.UpdatePost)
	router.DELETE("/api/post/:id", controller.DeletePost)
}
