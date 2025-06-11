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

func NewAuthService(secret []byte, UserService *UserService, UserRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		JWTSecret:   secret,
		UserService: UserService,
		UserRepo: UserRepo,
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

func (s *AuthService) RegisterUser(ctx context.Context, req RegisterRequest) (string, error) {
	hashedPassword := utils.HashPassword(req.Password)

	err := s.UserService.CreateUser(ctx, &db.User{
		Name:     req.Name,
		Email:    req.Email,
		Username: req.Username,
		Password: hashedPassword,
	})
	if err != nil {
		return "", err
	}

	user, err := s.UserRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", errors.New("failed to fetch newly created user")
	}

	print("supppp")

	return utils.GenerateJWT(fmt.Sprint(user.ID), s.JWTSecret)
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
