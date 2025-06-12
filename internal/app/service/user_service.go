package service

import (
	"context"
	"errors"

	"ocrolus-task/internal/app/repository"
	"ocrolus-task/internal/db"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(ctx context.Context, user *db.User) error {
	_, err := s.userRepo.GetUserByID(ctx, user.ID)
	if err == nil {
		return errors.New("user already exists")
	}

	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		return errors.New("failed to create user in repository: " + err.Error())
	}
	return nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]db.User, error) {
	users, err := s.userRepo.GetAllUsers(ctx)
	if err != nil {
		return nil, errors.New("failed to retrieve users from repository: " + err.Error())
	}
	return users, nil
}

func (s *UserService) GetUser(ctx context.Context, id int64) (db.User, error) {
	user, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return db.User{}, errors.New("failed to retrieve user from repository: " + err.Error())
	}
	return *user, nil
}
