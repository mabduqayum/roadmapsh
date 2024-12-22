package middleware

import (
	"personal_blog/internal/jwt"
	"personal_blog/internal/services"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("jwt")
		if token == "" {
			return c.Redirect("/auth/login")
		}

		userID, err := jwt.ValidateToken(token)
		if err != nil {
			c.ClearCookie("jwt")
			return c.Redirect("/auth/login")
		}

		user, err := userService.GetByID(c.Context(), userID)
		if err != nil {
			c.ClearCookie("jwt")
			return c.Redirect("/auth/login")
		}

		c.Locals("user", user)
		return c.Next()
	}
}
