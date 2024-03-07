package brand

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// DeleteBrand godoc
// @Summary Delete Brand by id
// @Description Delete Brand by id
// @Param id path string true "Brand ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Response
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /brands/{id} [delete]
// @Tags Brand
func DeleteBrand(c *fiber.Ctx) error {
	db := services.DB
	id := c.Params("id")

	var brand model.Brand
	result := db.Model(&brand).Where("id = ?", id).First(&brand)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Brand not found",
		})
	}

	db.Delete(&brand)

	return lib.OK(c)
}
