package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"ocrolus-task/internal/app/service"
	"ocrolus-task/internal/db"
)

type UserHandler struct {
	userService *service.UserService
	articleService *service.ArticleService
}

func NewUserHandler(userService *service.UserService, articleService *service.ArticleService) *UserHandler {
	return &UserHandler{userService: userService, articleService: articleService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user db.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	ctx := c.Request.Context()

	if err := h.userService.CreateUser(ctx, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := h.userService.GetAllUsers(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	ctx := c.Request.Context()

	user, err := h.userService.GetUser(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) ListArticlesByAuthor(c *gin.Context) {
	userID := c.MustGet("user_id").(int64)
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	ctx := c.Request.Context()

	articles, err := h.userService.ListArticlesByAuthor(ctx, userID, int32(limit), int32(offset))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalCount, err := h.articleService.CountArticlesByAuthor(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
		"total":    totalCount,
		"limit":    limit,
		"offset":   offset,
	})
}
