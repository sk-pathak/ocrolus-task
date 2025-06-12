package repository

import (
	"context"
	"ocrolus-task/internal/db"

	"github.com/jackc/pgx/v5/pgtype"
)

type ArticleRepository struct {
	queries *db.Queries
}

func NewArticleRepository(queries *db.Queries) *ArticleRepository {
	return &ArticleRepository{queries: queries}
}

func (r *ArticleRepository) Create(ctx context.Context, title, content string, authorID int64) (db.Article, error) {
	return r.queries.CreateArticle(ctx, db.CreateArticleParams{
		Title:    title,
		Content:  content,
		AuthorID: pgtype.Int8{Int64: authorID, Valid: true},
	})
}

func (r *ArticleRepository) Get(ctx context.Context, id int64) (db.Article, error) {
	return r.queries.GetArticle(ctx, id)
}

func (r *ArticleRepository) List(ctx context.Context, limit, offset int32) ([]db.Article, error) {
	return r.queries.ListArticles(ctx, db.ListArticlesParams{
		Limit:  limit,
		Offset: offset,
	})
}

func (r *ArticleRepository) Update(ctx context.Context, id int64, title, content string) (db.Article, error) {
	return r.queries.UpdateArticle(ctx, db.UpdateArticleParams{
		ID:      id,
		Title:   title,
		Content: content,
	})
}

func (r *ArticleRepository) Delete(ctx context.Context, id int64) error {
	return r.queries.DeleteArticle(ctx, id)
}
