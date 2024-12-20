package repository

import (
	"context"
	"errors"
	"personal_blog/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ArticleRepository interface {
	Create(ctx context.Context, article *models.Article) error
	GetBySlug(ctx context.Context, slug string) (*models.Article, error)
	GetAll(ctx context.Context) ([]models.Article, error)
	Update(ctx context.Context, article *models.Article) error
	Delete(ctx context.Context, slug string) error
}

type PostgresArticleRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresArticleRepository(pool *pgxpool.Pool) *PostgresArticleRepository {
	return &PostgresArticleRepository{pool}
}

func (r *PostgresArticleRepository) Create(ctx context.Context, article *models.Article) error {
	query := `
        INSERT INTO articles (author_id, title, content, slug, published_at, status, view_count)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id, created_at, updated_at`

	err := r.pool.QueryRow(ctx, query,
		article.AuthorId, article.Title, article.Content, article.Slug,
		article.PublishedAt, article.Status, article.ViewCount,
	).Scan(&article.ID, &article.CreatedAt, &article.UpdatedAt)

	return err
}

func (r *PostgresArticleRepository) GetBySlug(ctx context.Context, slug string) (*models.Article, error) {
	query := `
        SELECT id, author_id, title, content, slug, published_at, created_at, updated_at, status, view_count
        FROM articles
        WHERE slug = $1`

	var article models.Article
	err := r.pool.QueryRow(ctx, query, slug).Scan(
		&article.ID, &article.AuthorId, &article.Title, &article.Content, &article.Slug,
		&article.PublishedAt, &article.CreatedAt, &article.UpdatedAt,
		&article.Status, &article.ViewCount,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &article, nil
}

func (r *PostgresArticleRepository) GetAll(ctx context.Context) ([]models.Article, error) {
	query := `
        SELECT id, author_id, title, content, slug, published_at, created_at, updated_at, status, view_count
        FROM articles
        ORDER BY created_at DESC`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article
		err := rows.Scan(
			&article.ID, &article.AuthorId, &article.Title, &article.Content, &article.Slug,
			&article.PublishedAt, &article.CreatedAt, &article.UpdatedAt,
			&article.Status, &article.ViewCount,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (r *PostgresArticleRepository) Update(ctx context.Context, article *models.Article) error {
	query := `
        UPDATE articles
        SET author_id = $1, title = $2, content = $3, published_at = $4,
            status = $5, view_count = $6
        WHERE id = $7
        RETURNING updated_at`

	err := r.pool.QueryRow(ctx, query,
		article.AuthorId, article.Title, article.Content, article.PublishedAt,
		article.Status, article.ViewCount, article.ID,
	).Scan(&article.UpdatedAt)

	return err
}

func (r *PostgresArticleRepository) Delete(ctx context.Context, slug string) error {
	query := `DELETE FROM articles WHERE slug = $1`

	_, err := r.pool.Exec(ctx, query, slug)
	return err
}
