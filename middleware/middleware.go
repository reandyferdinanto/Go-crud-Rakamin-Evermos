package middleware

import (
	"tryFiberGo/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	token := c.Get("x-token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message": "unauthenticated",
		})
	}

	// _, err := utils.VerifyToken(token)
	claims, err := utils.DecodeToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message" : "Unauthenticated",
	})
}

role := claims["role"].(string)
if role != "admin" {
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"Message" : "forbidden access",
	})
}

c.Locals("userInfo", claims)


	return c.Next()
}

// func PermissionCreate(c *fiber.Ctx) error {
// 	return c.Next()
// }
