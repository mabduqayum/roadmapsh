package server

import (
	"context"
	"fmt"
	"log"
	"personal_blog/internal/handlers"
	"time"

	//"personal_blog/internal/middleware"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.app.Get("/", s.HelloWorldHandler)
	s.app.Get("/health", s.healthHandler)
	s.app.Get("/websocket", websocket.New(s.websocketHandler))

	//api := s.app.Group("/api/v1", middleware.AuthMiddleware(s.clientService))
	api := s.app.Group("/api/v1")

	articleHandler := handlers.NewArticleHandler(s.articleService)

	article := api.Group("/article")
	article.Get("/", articleHandler.GetAllArticles)
	article.Get("/:slug", articleHandler.GetArticleBySlug)
	//article.Get("/search")

	//article.Post("/")
	//article.Put("/:id")
	//article.Delete("/:id")

	//auth := api.Group("/auth")
	//auth.Post("/sign-up")
	//auth.Post("/sign-in")
	//auth.Get("/status")
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}

func (s *FiberServer) websocketHandler(con *websocket.Conn) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			_, _, err := con.ReadMessage()
			if err != nil {
				cancel()
				log.Println("Receiver Closing", err)
				break
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
			if err := con.WriteMessage(websocket.TextMessage, []byte(payload)); err != nil {
				log.Printf("could not write to socket: %v", err)
				return
			}
			time.Sleep(time.Second * 2)
		}
	}
}
