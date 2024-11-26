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

func (h *ArticleHandler) ListArticles(c *fiber.Ctx) error {
	articles, err := h.articleService.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"Error": "Failed to fetch articles",
		})
	}

	return c.Render("articles", fiber.Map{
		"Articles": articles,
	})
}

func (h *ArticleHandler) ViewArticle(c *fiber.Ctx) error {
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
	})
}

//func (h *ArticleHandler) GetAllArticles(c *fiber.Ctx) error {
//	articles, err := h.articleService.GetAll(c.Context())
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"error": "Failed to fetch articles",
//		})
//	}
//
//	return c.JSON(articles)
//}
//
//func (h *ArticleHandler) GetArticleBySlug(c *fiber.Ctx) error {
//	slug := c.Params("slug")
//	article, err := h.articleService.GetBySlug(c.Context(), slug)
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"error": "Failed to fetch article",
//		})
//	}
//
//	if article == nil {
//		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
//			"error": "Article not found",
//		})
//	}
//
//	return c.JSON(article)
//}
//
//func (h *ArticleHandler) CreateArticle(c *fiber.Ctx) error {
//	var article models.Article
//	if err := c.BodyParser(&article); err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"error": "Invalid request body",
//		})
//	}
//
//	// Set PublishedAt if article is published
//	if article.Status == models.StatusPublished {
//		now := time.Now()
//		article.PublishedAt = &now
//		article.Status = models.StatusPublished
//	}
//
//	if err := h.articleService.Create(c.Context(), &article); err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"error": "Failed to create article",
//		})
//	}
//
//	return c.Status(fiber.StatusCreated).JSON(article)
//}
//
//func (h *ArticleHandler) UpdateArticle(c *fiber.Ctx) error {
//	slug := c.Params("slug")
//
//	existingArticle, err := h.articleService.GetBySlug(c.Context(), slug)
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"error": "Failed to fetch article",
//		})
//	}
//
//	if existingArticle == nil {
//		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
//			"error": "Article not found",
//		})
//	}
//
//	// Parse the update request
//	var updateArticle models.Article
//	if err := c.BodyParser(&updateArticle); err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"error": "Invalid request body",
//		})
//	}
//
//	// Update the existing article with new values
//	updateArticle.ID = existingArticle.ID
//	updateArticle.Slug = existingArticle.Slug // Preserve the original slug
//
//	if existingArticle.Status != models.StatusPublished && updateArticle.Status == models.StatusPublished {
//		now := time.Now()
//		updateArticle.PublishedAt = &now
//		updateArticle.Status = models.StatusPublished
//	}
//
//	if err := h.articleService.Update(c.Context(), &updateArticle); err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"error": "Failed to update article",
//		})
//	}
//
//	return c.JSON(updateArticle)
//}
//
//// DeleteArticle handles DELETE /api/v1/article/:id
//func (h *ArticleHandler) DeleteArticle(c *fiber.Ctx) error {
//	slug := c.Params("slug")
//
//	if err := h.articleService.Delete(c.Context(), slug); err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"error": "Failed to delete article",
//		})
//	}
//
//	return c.Status(fiber.StatusNoContent).Send(nil)
//}
