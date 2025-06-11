package routes

import (
	"github.com/gin-gonic/gin"
	"ocrolus-task/internal/app/handler"
	"ocrolus-task/internal/middlewares"
)

func RegisterUserRoutes(r *gin.Engine, userHandler *handler.UserHandler, jwtSecret []byte) {
	r.POST("/users", userHandler.CreateUser)

	protectedGroup := r.Group("/users")
	protectedGroup.Use(middlewares.AuthMiddleware(jwtSecret))
	{
		protectedGroup.GET("/", userHandler.GetUsers)
		protectedGroup.GET("/:id", userHandler.GetUser)
	}
}
