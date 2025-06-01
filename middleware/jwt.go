package middleware

import (
	"strings"
	"github.com/gofiber/fiber/v2"
	"your_project/utils"
)

func  JWTProtected()  fiber.Handler { 
	return func(c *fiber.Ctx) error {
		header := c.Get("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid token"})
		}
		tokenStr := strings.TrimPrefix(header, "Bearer ")
		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
		}
		c.Locals("user_id", claims.UserID)
		return c.Next()
	}
}
