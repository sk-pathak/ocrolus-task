package service

import (
	"context"
	"errors"
	"fmt"

	"ocrolus-task/internal/app/repository"
	"ocrolus-task/internal/db"
	"ocrolus-task/internal/utils"
)

type AuthService struct {
	JWTSecret   []byte
	UserService *UserService
	UserRepo    *repository.UserRepository
}

func NewAuthService(secret []byte, userService *UserService) *AuthService {
	return &AuthService{
		JWTSecret:   secret,
		UserService: userService,
	}
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *AuthService) RegisterUser(ctx context.Context, req RegisterRequest) error {
	hashedPassword := utils.HashPassword(req.Password)

	return s.UserService.CreateUser(ctx, &db.User{
		Name:     req.Name,
		Email:    req.Email,
		Username: req.Username,
		Password: hashedPassword,
	})
}

func (s *AuthService) Login(ctx context.Context, req LoginRequest) (string, error) {
	user, err := s.UserRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateJWT(fmt.Sprint(user.ID), s.JWTSecret)
}
