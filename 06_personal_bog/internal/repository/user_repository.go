package repository

import (
	"context"
	"personal_blog/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(ctx context.Context, article *models.User) error
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	Update(ctx context.Context, article *models.User) error
	Delete(ctx context.Context, slug string) error
}

type PostgresUserRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresUserRepository(pool *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{pool}
}

//func (r *PostgresUserRepository) Create(ctx context.Context, article *models.User) error {
//}
//
//func (r *PostgresUserRepository) GetBySlug(ctx context.Context, slug string) (*models.User, error) {
//}
//
//func (r *PostgresUserRepository) GetAll(ctx context.Context) ([]models.User, error) {
//}
//
//func (r *PostgresUserRepository) Update(ctx context.Context, article *models.User) error {
//}
//
//func (r *PostgresUserRepository) Delete(ctx context.Context, slug string) error {
//}
