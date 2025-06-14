package repository

import (
	"context"

	"ocrolus-task/internal/db"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(queries *db.Queries) *UserRepository {
	return &UserRepository{queries: queries}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *db.User) error {
	_, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	})
	return err
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]db.User, error) {
	users, err := r.queries.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	var result []db.User
	for _, u := range users {
		result = append(result, db.User{
			ID:       u.ID,
			Name:     u.Name,
			Email:    u.Email,
			Username: u.Username,
			Password: u.Password,
		})
	}

	return result, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (*db.User, error) {
	user, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return &db.User{}, err
	}

	result := db.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}

	return &result, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*db.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return &db.User{}, err
	}

	result := db.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}

	return &result, nil
}

func (r *UserRepository) ListArticlesByAuthor(ctx context.Context, authorID int64, limit, offset int32) ([]db.Article, error) {
	return r.queries.ListArticlesByAuthor(ctx, db.ListArticlesByAuthorParams{
		AuthorID: pgtype.Int8{Int64: authorID, Valid: true},
		Limit:    limit,
		Offset:   offset,
	})
}
