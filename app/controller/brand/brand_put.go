package brand

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PutBrand godoc
// @Summary Update Brand by id
// @Description Update Brand by id
// @Param id path string true "Brand ID"
// @Param data body model.BrandAPI true "Brand data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Brand "Brand data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /brands/{id} [put]
// @Tags Brand
func PutBrand(c *fiber.Ctx) error {
	api := new(model.BrandAPI)

	if err := c.BodyParser(&api); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	id := c.Params("id")
	db := services.DB

	var data model.Brand
	result := db.Model(&data).Where("id = ?", &id).Take(&data)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Brand Not Found",
		})
	}

	lib.Merge(api, &data)

	db.Model(&data).Updates(&data)

	return lib.OK(c, data)
}
