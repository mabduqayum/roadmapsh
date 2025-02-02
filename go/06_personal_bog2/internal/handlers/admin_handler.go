package handlers

import (
	"personal_blog/internal/models"
	"personal_blog/internal/services"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct {
	articleService *services.ArticleService
}

func NewAdminHandler(articleService *services.ArticleService) *AdminHandler {
	return &AdminHandler{
		articleService: articleService,
	}
}

func (h *AdminHandler) DashboardPage(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	articles, err := h.articleService.GetAllByAuthorID(c.Context(), user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching articles")
	}

	return c.Render("admin/dashboard", fiber.Map{
		"Articles": articles,
		"User":     user,
	})
}

func (h *AdminHandler) NewArticlePage(c *fiber.Ctx) error {
	return c.Render("admin/new_article", fiber.Map{
		"User": c.Locals("user"),
	})
}

func (h *AdminHandler) CreateArticle(c *fiber.Ctx) error {
	article := &models.Article{
		Title:    c.FormValue("title"),
		Content:  c.FormValue("content"),
		AuthorId: c.Locals("user").(*models.User).ID,
		Status:   models.ArticleStatus(c.FormValue("status")),
	}

	if err := h.articleService.Create(c.Context(), article); err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"Error": "Failed to create article",
		})
	}

	return c.Redirect("/admin")
}

func (h *AdminHandler) EditArticlePage(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Render("error", fiber.Map{
			"Error": "Invalid article ID",
		})
	}

	article, err := h.articleService.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"Error": "Failed to fetch article",
		})
	}

	return c.Render("admin/edit_article", fiber.Map{
		"Article": article,
		"User":    c.Locals("user"),
	})
}

func (h *AdminHandler) UpdateArticle(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Render("error", fiber.Map{
			"Error": "Invalid article ID",
		})
	}

	article := &models.Article{
		ID:       id,
		Title:    c.FormValue("title"),
		Content:  c.FormValue("content"),
		Status:   models.ArticleStatus(c.FormValue("status")),
		AuthorId: c.Locals("user").(*models.User).ID,
	}
	now := time.Now()
	if article.Status == models.StatusPublished {
		article.PublishedAt = &now
	} else {
		article.PublishedAt = nil
	}

	if err := h.articleService.Update(c.Context(), article); err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"Error": "Failed to update article",
		})
	}

	return c.Redirect("/admin")
}

func (h *AdminHandler) DeleteArticle(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Render("error", fiber.Map{
			"Error": "Invalid article ID",
		})
	}

	if err := h.articleService.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"Error": "Failed to delete article",
		})
	}

	return c.Redirect("/admin")
}
