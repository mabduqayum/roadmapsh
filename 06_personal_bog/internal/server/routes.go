package server

import (
	"personal_blog/internal/handlers"
)

func (s *FiberServer) RegisterFiberRoutes() {
	// Routes
	s.app.Get("/", handlers.NewArticleHandler(s.articleService).ListArticles)
	s.app.Get("/article/:slug", handlers.NewArticleHandler(s.articleService).ViewArticle)

	//api := s.app.Group("/api/v1", middleware.AuthMiddleware(s.clientService))
	//api := s.app.Group("/api/v1")

	//articleHandler := handlers.NewArticleHandler(s.articleService)
	//
	//article := api.Group("/article")
	//article.Get("/:slug", articleHandler.GetArticleBySlug)
	//article.Get("/", articleHandler.GetAllArticles)
	//article.Get("/search")

	//article.Post("/")
	//article.Put("/:id")
	//article.Delete("/:id")

	//auth := api.Group("/auth")
	//auth.Post("/sign-up")
	//auth.Post("/sign-in")
	//auth.Get("/status")
}
