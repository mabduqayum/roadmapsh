package services

import (
	"context"
	"personal_blog/internal/models"
	"personal_blog/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(ctx context.Context, user *models.User) error {
	return s.repo.Create(ctx, user)
}

func (s *UserService) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	return s.repo.GetByUsername(ctx, username)
}

func (s *UserService) GetByID(ctx context.Context, id int64) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}
