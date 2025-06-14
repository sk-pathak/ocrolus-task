package service

import (
	"context"
	"errors"
	"log"
	"ocrolus-task/internal/app/repository"
	"ocrolus-task/internal/db"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

var ErrUnauthorized = errors.New("unauthorized")

type ArticleService struct {
	articleRepo *repository.ArticleRepository
}

func NewArticleService(articleRepo *repository.ArticleRepository) *ArticleService {
	return &ArticleService{
		articleRepo: articleRepo,
	}
}

func (s *ArticleService) CreateArticle(ctx context.Context, title, content string, authorID int64) (db.Article, error) {
	return s.articleRepo.Create(ctx, title, content, authorID)
}

func (s *ArticleService) GetArticle(ctx context.Context, userID, articleID int64) (db.Article, error) {
	article, err := s.articleRepo.Get(ctx, articleID)
	if err != nil {
		return db.Article{}, err
	}

	if err := s.articleRepo.UpsertArticleView(ctx, db.UpsertArticleViewParams{
		UserID:    userID,
		ArticleID: articleID,
		ViewedAt:  pgtype.Timestamp{Time: time.Now(), Valid: true},
	}); err != nil {
		log.Printf("failed to upsert article view: %v", err)
	}

	const maxRecentViews = 15
	if err := s.articleRepo.DeleteOldArticleViews(ctx, db.DeleteOldArticleViewsParams{
		UserID: userID,
		Limit:  int32(maxRecentViews),
	}); err != nil {
		log.Printf("failed to delete old article views: %v", err)
	}

	return article, nil
}

func (s *ArticleService) ListArticles(ctx context.Context, limit, offset int32) ([]db.Article, error) {
	return s.articleRepo.List(ctx, limit, offset)
}

func (s *ArticleService) UpdateArticle(ctx context.Context, userID, articleID int64, title, content string) (db.Article, error) {
	article, err := s.articleRepo.Get(ctx, articleID)
	if err != nil {
		return db.Article{}, err
	}

	if article.AuthorID.Int64 != userID {
		return db.Article{}, ErrUnauthorized
	}
	return s.articleRepo.Update(ctx, articleID, title, content)
}

func (s *ArticleService) DeleteArticle(ctx context.Context, userID, articleID int64) error {
	article, err := s.articleRepo.Get(ctx, articleID)
	if err != nil {
		return err
	}

	if article.AuthorID.Int64 != userID {
		return ErrUnauthorized
	}
	return s.articleRepo.Delete(ctx, articleID)
}

func (s *ArticleService) CountArticlesByAuthor(ctx context.Context, authorID int64) (int64, error) {
	return s.articleRepo.CountArticlesByAuthor(ctx, authorID)
}

func (s *ArticleService) CountArticles(ctx context.Context) (int64, error) {
	return s.articleRepo.CountArticles(ctx)
}

func (s *ArticleService) GetRecentlyViewedArticles(ctx context.Context, userID int64) ([]db.Article, error) {
	articles, err := s.articleRepo.GetRecentlyViewedArticles(ctx, userID)
	if err != nil {
		return nil, errors.New("failed to retrieve articles from repository: " + err.Error())
	}
	return articles, nil
}
