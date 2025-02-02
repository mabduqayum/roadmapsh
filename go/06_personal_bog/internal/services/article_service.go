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
	// todo: add pagination
	return s.repo.GetAll(ctx)
}

func (s *ArticleService) GetAllByAuthorID(ctx context.Context, authorID int64) ([]models.Article, error) {
	return s.repo.GetAllByAuthorID(ctx, authorID)
}

func (s *ArticleService) GetByID(ctx context.Context, id int64) (*models.Article, error) {
	return s.repo.GetByID(ctx, id)
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

func (s *ArticleService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
