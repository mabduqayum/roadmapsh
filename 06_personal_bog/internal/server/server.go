package server

import (
	"personal_blog/internal/config"
	"personal_blog/internal/database"
	"personal_blog/internal/repository"
	"personal_blog/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type FiberServer struct {
	app *fiber.App
	db  database.Service
	cfg *config.ServerConfig

	articleService *services.ArticleService
}

func New(cfg *config.ServerConfig, db database.Service) *FiberServer {
	articleRepository := repository.NewPostgresArticleRepository(db.GetPool())
	articleService := services.NewArticleService(articleRepository)

	server := &FiberServer{
		app: fiber.New(fiber.Config{
			ServerHeader: "personal_blog",
			AppName:      "personal_blog v" + cfg.Version,
		}),
		db:  db,
		cfg: cfg,

		articleService: articleService,
	}

	// Add recover middleware
	server.app.Use(recover.New())

	return server
}

func (s *FiberServer) Listen() error {
	return s.app.Listen(s.cfg.Address())
}

func (s *FiberServer) Shutdown() error {
	return s.app.Shutdown()
}
