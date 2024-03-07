package cart

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PutCart godoc
// @Summary Update Cart by id
// @Description Update Cart by id
// @Param id path string true "Cart ID"
// @Param data body model.CartUpdate true "Cart data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Cart "Cart data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /carts/{id} [put]
// @Tags Cart
func PutCart(c *fiber.Ctx) error {
	api := new(model.CartUpdate)

	if err := c.BodyParser(&api); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	id := c.Params("id")
	db := services.DB

	var data model.Cart
	result := db.Model(&data).Where("id = ?", &id).Take(&data)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Cart Not Found",
		})
	}

	lib.Merge(api, &data)

	db.Model(&data).Updates(&data)

	return lib.OK(c, data)
}
