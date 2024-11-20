package services

import (
	"context"
	"personal_blog/internal/models"
	"personal_blog/internal/repository"
)

type ArticleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) *ArticleService {
	return &ArticleService{
		repo: repo,
	}
}

func (s *ArticleService) GetAll(ctx context.Context) ([]models.Article, error) {
	return s.repo.GetAll(ctx)
}

func (s *ArticleService) GetBySlug(ctx context.Context, slug string) (*models.Article, error) {
	return s.repo.GetBySlug(ctx, slug)
}

func (s *ArticleService) Create(ctx context.Context, article *models.Article) error {
	return s.repo.Create(ctx, article)
}

func (s *ArticleService) Update(ctx context.Context, article *models.Article) error {
	return s.repo.Update(ctx, article)
}

func (s *ArticleService) Delete(ctx context.Context, slug string) error {
	return s.repo.Delete(ctx, slug)
}
