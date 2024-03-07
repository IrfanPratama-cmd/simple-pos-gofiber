package customer

import (
	"api/app/model"
	"api/app/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// PostCustomer godoc
// @Summary Create new Customer
// @Description Create new Customer
// @Param data body model.CustomerAPI true "Customer data"
// @Accept  application/json
// @Produce application/json
// @Success 201 {object} model.Customer "Customer data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /customers [post]
// @Tags Customer
func PostCustomer(c *fiber.Ctx) error {
	var customerAPI model.CustomerAPI

	db := services.DB

	if err := c.BodyParser(&customerAPI); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(customerAPI)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}
	Customer := &model.Customer{CustomerAPI: customerAPI}
	db.Model(&model.Customer{}).Create(Customer)

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    Customer,
	})
}
