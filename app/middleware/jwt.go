package middleware

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

func JwtMiddleware(c *fiber.Ctx) error {
	token := c.Get("x-token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims, err := lib.DecodeToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	// role := claims["role"].(string)
	// if role != "user" {
	// 	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
	// 		"message": "forbidden access",
	// 	})
	// }

	userID := claims["user_id"].(string)

	db := services.DB

	var user model.User
	result := db.Where("id = ?", userID).First(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if user.ActivatedAt == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User not verified",
		})
	}

	c.Locals("userInfo", claims)
	c.Locals("userID", userID)

	return c.Next()
}

func PermissionCreate(c *fiber.Ctx) error {
	return c.Next()
}
