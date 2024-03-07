package account

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"log"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// PostLogin godoc
// @Summary Generate token
// @Description Generate token
// @Param data body model.LoginAPI true "Payload"
// @Security TokenKey
// @Success 200 {object} model.LoginAPI "Logedin"
// @Failure 401 {object} lib.Response "Unauthorized"
// @Failure 400 {object} lib.Response "Bad Request"
// @Failure 204 {object} lib.Response "No Content"
// @Failure 404 {object} lib.Response "Not Found"
// @Failure 409 {object} lib.Response "Conflict"
// @Failure 500 {object} lib.Response "Internal Server Error"
// @Router /account/login [post]
// @Tags Account
func PostLoginAccount(c *fiber.Ctx) error {
	loginRequest := new(model.LoginAPI)
	db := services.DB

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var user model.User
	err := db.First(&user, "email = ?", loginRequest.Username).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	checkPassword := lib.CheckPassword(*loginRequest.Password, *user.Password)
	if !checkPassword {
		return c.Status(404).JSON(fiber.Map{
			"message": "Wrong credential",
		})
	}

	// Generate JWT
	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 10000).Unix()

	token, errGenerateToken := lib.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong credential",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
