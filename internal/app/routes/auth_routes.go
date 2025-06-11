package routes

import (
	"github.com/gin-gonic/gin"
	"ocrolus-task/internal/app/handler"
	"ocrolus-task/internal/middlewares"
)

func RegisterAuthRoutes(r *gin.Engine, authHandler *handler.AuthHandler, jwtSecret []byte) {
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware(jwtSecret))
	auth.GET("/me", authHandler.GetUserByID)
}
