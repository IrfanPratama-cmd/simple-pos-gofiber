package cart

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// DeleteCart godoc
// @Summary Delete Cart by id
// @Description Delete Cart by id
// @Param id path string true "Cart ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Response
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /carts/{id} [delete]
// @Tags Cart
func DeleteCart(c *fiber.Ctx) error {
	db := services.DB
	id := c.Params("id")

	var cart model.Cart
	result := db.Model(&cart).Where("id = ?", id).First(&cart)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Cart not found",
		})
	}

	db.Delete(&cart)

	return lib.OK(c)
}
