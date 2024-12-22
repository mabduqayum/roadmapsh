package middleware

import (
	"personal_blog/internal/jwt"
	"personal_blog/internal/services"

	"github.com/gofiber/fiber/v2"
)

func UserContext(userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("jwt")
		if token == "" {
			// No token found, continue without setting user context
			return c.Next()
		}

		userID, err := jwt.ValidateToken(token)
		if err != nil {
			// Token is invalid, clear the cookie and continue
			c.ClearCookie("jwt")
			return c.Next()
		}

		user, err := userService.GetByID(c.Context(), userID)
		if err != nil {
			// User not found or other database error
			// You might want to log this error
			// log.Printf("Error fetching user: %v", err)
			c.ClearCookie("jwt")
			return c.Next()
		}

		// Set user in context
		c.Locals("user", user)
		return c.Next()
	}
}
