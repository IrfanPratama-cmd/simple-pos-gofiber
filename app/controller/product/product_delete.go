package product

import (
	"api/app/model"
	"api/app/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// DeleteProduct godoc
// @Summary Delete Product by id
// @Description Delete Product by id
// @Param id path string true "Product ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Response
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /products/{id} [delete]
// @Tags Product
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := services.DB

	var product model.Product
	result := db.Model(&product).Where(`id = ?`, id).First(&product)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "not found",
		})
	}

	result.Delete(&product)

	message := fmt.Sprintf(`Product with id %s has been deleted`, id)
	return c.JSON(fiber.Map{
		"message": message,
	})
}
