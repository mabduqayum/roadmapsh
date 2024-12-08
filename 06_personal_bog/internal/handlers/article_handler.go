package handlers

import (
	"personal_blog/internal/services"

	"github.com/gofiber/fiber/v2"
)

type ArticleHandler struct {
	articleService *services.ArticleService
}

func NewArticleHandler(articleService *services.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		articleService: articleService,
	}
}

func (h *ArticleHandler) ArticlesPage(c *fiber.Ctx) error {
	articles, err := h.articleService.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"Error": "Failed to fetch articles",
		})
	}

	return c.Render("articles", fiber.Map{
		"Articles": articles,
		"User":     c.Locals("user"),
	})
}

func (h *ArticleHandler) ArticlePage(c *fiber.Ctx) error {
	slug := c.Params("slug")
	article, err := h.articleService.GetBySlug(c.Context(), slug)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"Error": "Failed to fetch article",
		})
	}

	if article == nil {
		return c.Status(fiber.StatusNotFound).Render("error", fiber.Map{
			"Error": "Article not found",
		})
	}

	return c.Render("article", fiber.Map{
		"Article": article,
		"User":    c.Locals("user"),
	})
}
