package category

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PutCategory godoc
// @Summary Update Category by id
// @Description Update Category by id
// @Param id path string true "Category ID"
// @Param data body model.CategoryAPI true "Category data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Category "Category data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /categories/{id} [put]
// @Tags Category
func PutCategory(c *fiber.Ctx) error {
	api := new(model.CategoryAPI)

	if err := c.BodyParser(&api); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	id := c.Params("id")
	db := services.DB

	var data model.Category
	result := db.Model(&data).Where("id = ?", &id).Take(&data)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Category Not Found",
		})
	}

	lib.Merge(api, &data)

	db.Model(&data).Updates(&data)

	return lib.OK(c, data)
}
