package service

import (
	"context"
	"errors"
	"ocrolus-task/internal/app/repository"
	"ocrolus-task/internal/db"
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

func (s *ArticleService) GetArticle(ctx context.Context, articleId int64) (db.Article, error) {
	article, err := s.articleRepo.Get(ctx, articleId)
	if err != nil {
		return db.Article{}, err
	}

	return article, nil
}

func (s *ArticleService) ListArticles(ctx context.Context, limit, offset int32) ([]db.Article, error) {
	return s.articleRepo.List(ctx, limit, offset)
}

func (s *ArticleService) UpdateArticle(ctx context.Context, userId, articleId int64, title, content string) (db.Article, error) {
	article, err := s.articleRepo.Get(ctx, articleId)
	if err != nil {
		return db.Article{}, err
	}

	if article.AuthorID.Int64 != userId {
		return db.Article{}, ErrUnauthorized
	}
	return s.articleRepo.Update(ctx, articleId, title, content)
}

func (s *ArticleService) DeleteArticle(ctx context.Context, userId, articleId int64) error {
	article, err := s.articleRepo.Get(ctx, articleId)
	if err != nil {
		return err
	}

	if article.AuthorID.Int64 != userId {
		return ErrUnauthorized
	}
	return s.articleRepo.Delete(ctx, articleId)
}

func (s *ArticleService) CountArticlesByAuthor(ctx context.Context, authorID int64) (int64, error) {
	return s.articleRepo.CountArticlesByAuthor(ctx, authorID)
}

func (s *ArticleService) CountArticles(ctx context.Context) (int64, error) {
	return s.articleRepo.CountArticles(ctx)
}
