package category

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// DeleteCategory godoc
// @Summary Delete Category by id
// @Description Delete Category by id
// @Param id path string true "Category ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Response
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /categories/{id} [delete]
// @Tags Category
func DeleteCategory(c *fiber.Ctx) error {
	db := services.DB
	id := c.Params("id")

	var category model.Category
	result := db.Model(&category).Where("id = ?", id).First(&category)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Category not found",
		})
	}

	db.Delete(&category)

	return lib.OK(c)
}
