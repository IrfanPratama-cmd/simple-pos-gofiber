package customer

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PutCustomer godoc
// @Summary Update Customer by id
// @Description Update Customer by id
// @Param id path string true "Customer ID"
// @Param data body model.CustomerAPI true "Customer data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Customer "Customer data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /customers/{id} [put]
// @Tags Customer
func PutCustomer(c *fiber.Ctx) error {
	api := new(model.CustomerAPI)

	if err := c.BodyParser(&api); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	id := c.Params("id")
	db := services.DB

	var data model.Customer
	result := db.Model(&data).Where("id = ?", &id).Take(&data)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Customer Not Found",
		})
	}

	lib.Merge(api, &data)

	db.Model(&data).Updates(&data)

	return lib.OK(c, data)
}
