package brand

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetBrandID godoc
// @Summary Get a Brand by id
// @Description Get a Brand by id
// @Param id path string true "Brand ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Brand "Brand data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /brands/{id} [get]
// @Tags Brand
func GetBrandID(c *fiber.Ctx) error {
	db := services.DB
	id := c.Params("id")

	var brand model.Brand
	result := db.Model(&brand).Where("id = ?", id).First(&brand)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Brand not found",
		})
	}

	return lib.OK(c, brand)
}
