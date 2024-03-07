package customer

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetCustomerID godoc
// @Summary Get a Customer by id
// @Description Get a Customer by id
// @Param id path string true "Customer ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Customer "Customer data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /customers/{id} [get]
// @Tags Customer
func GetCustomerID(c *fiber.Ctx) error {
	db := services.DB
	id := c.Params("id")

	var customer model.Customer
	result := db.Model(&customer).Where("id = ?", id).First(&customer)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Customer not found",
		})
	}

	return lib.OK(c, customer)
}
