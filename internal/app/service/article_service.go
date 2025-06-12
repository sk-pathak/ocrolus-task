package service

import (
	"context"
	"ocrolus-task/internal/app/repository"
	"ocrolus-task/internal/db"
)

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

func (s *ArticleService) GetArticle(ctx context.Context, userID, id int64) (db.Article, error) {
	article, err := s.articleRepo.Get(ctx, id)
	if err != nil {
		return db.Article{}, err
	}

	return article, nil
}

func (s *ArticleService) ListArticles(ctx context.Context, limit, offset int32) ([]db.Article, error) {
	return s.articleRepo.List(ctx, limit, offset)
}

func (s *ArticleService) UpdateArticle(ctx context.Context, id int64, title, content string) (db.Article, error) {
	return s.articleRepo.Update(ctx, id, title, content)
}

func (s *ArticleService) DeleteArticle(ctx context.Context, id int64) error {
	return s.articleRepo.Delete(ctx, id)
}
