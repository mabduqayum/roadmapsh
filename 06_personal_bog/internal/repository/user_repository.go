package repository

import (
	"context"
	"personal_blog/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(ctx context.Context, article *models.User) error
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	GetByID(ctx context.Context, id int64) (*models.User, error)
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

func (r *PostgresUserRepository) Create(ctx context.Context, user *models.User) error {
	query := `
        INSERT INTO users (username, email, password_hash, full_name)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, updated_at`

	err := r.pool.QueryRow(ctx, query,
		user.Username, user.Email, user.PasswordHash, user.FullName,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	return err
}

func (r *PostgresUserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	// Implement the logic to fetch a user by username from the database
	// For example:
	var user models.User
	err := r.pool.QueryRow(ctx, "SELECT id, username, email, password_hash FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	err := r.pool.QueryRow(ctx, "SELECT id, username, email, password_hash FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	return nil, nil
}

func (r *PostgresUserRepository) Update(ctx context.Context, article *models.User) error {
	return nil
}

func (r *PostgresUserRepository) Delete(ctx context.Context, slug string) error {
	return nil
}
