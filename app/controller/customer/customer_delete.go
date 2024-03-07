package customer

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// DeleteCustomer godoc
// @Summary Delete Customer by id
// @Description Delete Customer by id
// @Param id path string true "Customer ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Response
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /Customers/{id} [delete]
// @Tags Customer
func DeleteCustomer(c *fiber.Ctx) error {
	db := services.DB
	id := c.Params("id")

	var customer model.Customer
	result := db.Model(&customer).Where("id = ?", id).First(&customer)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Customer not found",
		})
	}

	db.Delete(&customer)

	return lib.OK(c)
}
