package server

import (
	"personal_blog/internal/handlers"
	"personal_blog/internal/middleware"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.app.Use(middleware.UserContext(s.userService))
	s.app.Get("/", handlers.NewArticleHandler(s.articleService).ArticlesPage)
	s.app.Get("/article/:slug", handlers.NewArticleHandler(s.articleService).ArticlePage)

	// Auth routes
	auth := s.app.Group("/auth")
	authHandler := handlers.NewAuthHandler(s.userService)
	auth.Get("/login", authHandler.LoginPage)
	auth.Post("/login", authHandler.Login)
	auth.Get("/register", authHandler.RegisterPage)
	auth.Post("/register", authHandler.Register)
	auth.Post("/logout", authHandler.Logout)

	// Admin routes
	admin := s.app.Group("/admin")
	admin.Use(middleware.AuthMiddleware(s.userService))
	adminHandler := handlers.NewAdminHandler(s.articleService)
	admin.Get("/", adminHandler.DashboardPage)
	admin.Get("/articles/new", adminHandler.NewArticlePage)
	admin.Post("/articles", adminHandler.CreateArticle)
	admin.Get("/articles/:id/edit", adminHandler.EditArticlePage)
	admin.Post("/articles/:id", adminHandler.UpdateArticle)
	admin.Post("/articles/:id/delete", adminHandler.DeleteArticle)
}
