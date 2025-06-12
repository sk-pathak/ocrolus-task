package routes

import (
	"ocrolus-task/internal/app/handler"
	"ocrolus-task/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterArticleRoutes(r *gin.Engine, articleHandler *handler.ArticleHandler, jwtSecret []byte) {
	r.GET("/articles", articleHandler.ListArticles)

	protectedGroup := r.Group("/articles")
	protectedGroup.Use(middlewares.AuthMiddleware(jwtSecret))

	protectedGroup.GET("/:id", articleHandler.GetArticle)
	protectedGroup.GET("/me", articleHandler.ListArticlesByAuthor)
	protectedGroup.POST("", articleHandler.CreateArticle)
	protectedGroup.PUT("/:id", articleHandler.UpdateArticle)
	protectedGroup.DELETE("/:id", articleHandler.DeleteArticle)
}
