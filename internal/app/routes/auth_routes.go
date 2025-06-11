package routes

import (
	"github.com/gin-gonic/gin"
	"ocrolus-task/internal/app/handler"
)

func RegisterAuthRoutes(r *gin.Engine, authHandler *handler.AuthHandler, jwtSecret []byte) {
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
}
