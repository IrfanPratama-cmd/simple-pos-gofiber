package brand

import (
	"api/app/model"
	"api/app/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// PostBrand godoc
// @Summary Create new Brand
// @Description Create new Brand
// @Param data body model.BrandAPI true "Brand data"
// @Accept  application/json
// @Produce application/json
// @Success 201 {object} model.Brand "Brand data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /brands [post]
// @Tags Brand
func PostBrand(c *fiber.Ctx) error {
	var brandAPI model.BrandAPI

	db := services.DB

	if err := c.BodyParser(&brandAPI); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(brandAPI)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}
	brand := &model.Brand{BrandAPI: brandAPI}
	db.Model(&model.Brand{}).Create(brand)

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    brand,
	})
}
