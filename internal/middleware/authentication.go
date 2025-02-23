package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) Authentication(ctx *fiber.Ctx) error {
	authHeader := ctx.GetReqHeaders()["Authorization"]

	if authHeader == nil {
		ctx.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	bearerToken := authHeader[0]

	if bearerToken == "" {
		ctx.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
		return nil
	}

	token := strings.Split(bearerToken, " ")[1]

	id, err := m.jwt.ValidateToken(token)
	if err != nil {
		ctx.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
		return nil
	}

	ctx.Locals("userId", id)
	return ctx.Next()
}
