package category

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetCategoryID godoc
// @Summary Get a Category by id
// @Description Get a Category by id
// @Param id path string true "Category ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Category "Category data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /categories/{id} [get]
// @Tags Category
func GetCategoryID(c *fiber.Ctx) error {
	db := services.DB
	id := c.Params("id")

	var category model.Category
	result := db.Model(&category).Where("id = ?", id).First(&category)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Category not found",
		})
	}

	return lib.OK(c, category)
}
