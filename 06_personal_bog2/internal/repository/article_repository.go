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
	GetByID(ctx context.Context, id int64) (*models.Article, error)
	GetBySlug(ctx context.Context, slug string) (*models.Article, error)
	GetAll(ctx context.Context) ([]models.Article, error)
	GetAllByAuthorID(ctx context.Context, authorID int64) ([]models.Article, error)
	Update(ctx context.Context, article *models.Article) error
	Delete(ctx context.Context, id int64) error
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

func (r *PostgresArticleRepository) GetByID(ctx context.Context, id int64) (*models.Article, error) {
	query := `
        SELECT id, author_id, title, content, slug, published_at, created_at, updated_at, status, view_count
        FROM articles
        WHERE id = $1`

	var article models.Article
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&article.ID, &article.AuthorId, &article.Title, &article.Content, &article.Slug,
		&article.PublishedAt, &article.CreatedAt, &article.UpdatedAt, &article.Status, &article.ViewCount,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &article, nil
}

func (r *PostgresArticleRepository) GetBySlug(ctx context.Context, slug string) (*models.Article, error) {
	query := `
        SELECT a.id, a.author_id, a.title, a.content, a.slug, a.published_at, a.created_at, a.updated_at, a.status, a.view_count,
               u.id, u.username, u.email, u.full_name, u.created_at, u.updated_at, u.is_admin, u.last_login
        FROM articles a
        LEFT JOIN users u ON a.author_id = u.id
        WHERE a.slug = $1`

	var article models.Article
	var user models.User
	err := r.pool.QueryRow(ctx, query, slug).Scan(
		&article.ID, &article.AuthorId, &article.Title, &article.Content, &article.Slug,
		&article.PublishedAt, &article.CreatedAt, &article.UpdatedAt, &article.Status, &article.ViewCount,
		&user.ID, &user.Username, &user.Email, &user.FullName, &user.CreatedAt, &user.UpdatedAt, &user.IsAdmin, &user.LastLogin,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	article.Author = &user
	return &article, nil
}

func (r *PostgresArticleRepository) GetAll(ctx context.Context) ([]models.Article, error) {
	query := `
        SELECT a.id, a.author_id, a.title, a.content, a.slug, a.published_at, a.created_at, a.updated_at, a.status, a.view_count,
               u.id, u.username, u.email, u.full_name, u.created_at, u.updated_at, u.is_admin, u.last_login
        FROM articles a
        LEFT JOIN users u ON a.author_id = u.id
        WHERE A.published_at IS NOT NULL
        ORDER BY a.view_count,
                 a.created_at DESC`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article
		var user models.User
		err := rows.Scan(
			&article.ID, &article.AuthorId, &article.Title, &article.Content, &article.Slug,
			&article.PublishedAt, &article.CreatedAt, &article.UpdatedAt, &article.Status, &article.ViewCount,
			&user.ID, &user.Username, &user.Email, &user.FullName, &user.CreatedAt, &user.UpdatedAt, &user.IsAdmin, &user.LastLogin,
		)
		if err != nil {
			return nil, err
		}
		article.Author = &user
		articles = append(articles, article)
	}

	return articles, nil
}

func (r *PostgresArticleRepository) GetAllByAuthorID(ctx context.Context, authorID int64) ([]models.Article, error) {
	query := `
        SELECT a.id, a.author_id, a.title, a.content, a.slug, a.published_at, a.created_at, a.updated_at, a.status, a.view_count,
               u.id, u.username, u.email, u.full_name, u.created_at, u.updated_at, u.is_admin, u.last_login
        FROM articles a
        LEFT JOIN users u ON a.author_id = u.id
        WHERE a.author_id = $1
        ORDER BY a.created_at DESC`

	rows, err := r.pool.Query(ctx, query, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article
		var user models.User
		err := rows.Scan(
			&article.ID, &article.AuthorId, &article.Title, &article.Content, &article.Slug,
			&article.PublishedAt, &article.CreatedAt, &article.UpdatedAt, &article.Status, &article.ViewCount,
			&user.ID, &user.Username, &user.Email, &user.FullName, &user.CreatedAt, &user.UpdatedAt, &user.IsAdmin, &user.LastLogin,
		)
		if err != nil {
			return nil, err
		}
		article.Author = &user
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

func (r *PostgresArticleRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM articles WHERE id = $1`

	_, err := r.pool.Exec(ctx, query, id)
	return err
}
