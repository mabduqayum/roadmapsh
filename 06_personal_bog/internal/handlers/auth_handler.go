package handlers

import (
	"personal_blog/internal/jwt"
	"personal_blog/internal/models"
	"personal_blog/internal/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userService *services.UserService
}

func NewAuthHandler(userService *services.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

func (h *AuthHandler) LoginPage(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := h.userService.GetByUsername(c.Context(), username)
	if err != nil {
		return c.Render("login", fiber.Map{
			"Error": "Invalid credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return c.Render("login", fiber.Map{
			"Error": "Invalid credentials",
		})
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return c.Render("login", fiber.Map{
			"Error": "Failed to generate token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		HTTPOnly: true,
	})

	return c.Redirect("/admin")
}

func (h *AuthHandler) RegisterPage(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Title": "Register!",
	})
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.Render("register", fiber.Map{
			"Error": "Failed to hash password",
		})
	}

	user := &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	if err := h.userService.Create(c.Context(), user); err != nil {
		return c.Render("register", fiber.Map{
			"Error": "Failed to create user",
		})
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return c.Render("register", fiber.Map{
			"Error": "Failed to generate token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		HTTPOnly: true,
	})

	return c.Redirect("/admin")
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Time{},
		HTTPOnly: true,
		SameSite: "Lax",
		Path:     "/",
	})
	return c.Redirect("/")
}
