package category

import (
	"api/app/model"
	"api/app/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// PostCategory godoc
// @Summary Create new Category
// @Description Create new Category
// @Param data body model.CategoryAPI true "Category data"
// @Accept  application/json
// @Produce application/json
// @Success 201 {object} model.Category "Category data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /categories [post]
// @Tags Category
func PostCategory(c *fiber.Ctx) error {
	var categoryAPI model.CategoryAPI

	db := services.DB

	if err := c.BodyParser(&categoryAPI); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(categoryAPI)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}
	category := &model.Category{CategoryAPI: categoryAPI}
	db.Model(&model.Category{}).Create(category)

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    category,
	})
}
