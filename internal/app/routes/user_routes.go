package routes

import (
	"github.com/gin-gonic/gin"
	"ocrolus-task/internal/app/handler"
	"ocrolus-task/internal/middlewares"
)

func RegisterUserRoutes(r *gin.Engine, userHandler *handler.UserHandler, jwtSecret []byte) {
	protectedGroup := r.Group("/")
	protectedGroup.Use(middlewares.AuthMiddleware(jwtSecret))

	protectedGroup.GET("/users", userHandler.GetUsers)
	protectedGroup.GET("/users/:id", userHandler.GetUser)
	protectedGroup.GET("/users/me/articles", userHandler.ListArticlesByAuthor)
}
