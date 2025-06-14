package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	config "ocrolus-task/configs"
	handler "ocrolus-task/internal/app/handler"
	repo "ocrolus-task/internal/app/repository"
	routes "ocrolus-task/internal/app/routes"
	service "ocrolus-task/internal/app/service"
	db "ocrolus-task/internal/db"
	middlewares "ocrolus-task/internal/middlewares"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	connStr := cfg.DBDriver + "://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName

	dbPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer dbPool.Close()

	if err := dbPool.Ping(context.Background()); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	queries := db.New(dbPool)

	userRepo := repo.NewUserRepository(queries)
	articleRepo := repo.NewArticleRepository(queries)

	userService := service.NewUserService(userRepo)
	articleService := service.NewArticleService(articleRepo)
	authService := service.NewAuthService([]byte(cfg.JWTSecret), userService, userRepo)

	userHandler := handler.NewUserHandler(userService,articleService)
	articleHandler := handler.NewArticleHandler(articleService)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()
	r.RedirectTrailingSlash = true
	r.Use(middlewares.SetupCORS())
	r.Use(middlewares.RateLimiterMiddleware())

	routes.RegisterAuthRoutes(r, authHandler, []byte(cfg.JWTSecret))
	routes.RegisterUserRoutes(r, userHandler, []byte(cfg.JWTSecret))
	routes.RegisterArticleRoutes(r, articleHandler, []byte(cfg.JWTSecret))

	log.Printf("Server is running on port %s", cfg.Port)
	log.Fatal(r.Run(":" + cfg.Port))
}
