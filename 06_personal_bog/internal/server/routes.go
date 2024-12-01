package server

import (
	"personal_blog/internal/handlers"
	"personal_blog/internal/middleware"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.app.Use(middleware.UserContext(s.userService))
	s.app.Get("/", handlers.NewArticleHandler(s.articleService).ListArticles)
	s.app.Get("/article/:slug", handlers.NewArticleHandler(s.articleService).ViewArticle)

	// Auth routes
	auth := s.app.Group("/auth")
	authHandler := handlers.NewAuthHandler(s.userService)
	auth.Get("/login", authHandler.LoginPage)
	auth.Post("/login", authHandler.Login)
	auth.Get("/register", authHandler.RegisterPage)
	auth.Post("/register", authHandler.Register)
	auth.Post("/logout", authHandler.Logout)

	//Protected routes
	//protected := s.app.Group("/", middleware.AuthMiddleware(s.userService))

	// Apply AuthMiddleware to the home page (optional, depending on your requirements)
	//protected.Get("/", handlers.NewArticleHandler(s.articleService).ListArticles)

	// API routes (protected)
	//api := protected.Group("/api")
	//articleHandler := handlers.NewArticleHandler(s.articleService)
	//api.Post("/article", articleHandler.CreateArticle)
	//api.Put("/article/:slug", articleHandler.UpdateArticle)
	//api.Delete("/article/:slug", articleHandler.DeleteArticle)
}
