package product

import (
	"api/app/model"
	"api/app/services"
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PostProduct godoc
// @Summary Create new Product
// @Description Create new Product
// @Param data body model.ProductRequest true "Product data"
// @Accept  application/json
// @Produce application/json
// @Success 201 {object} model.Product "Product data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /products [post]
// @Tags Product
func PostProduct(c *fiber.Ctx) error {
	var productAPI model.ProductRequest

	db := services.DB

	if err := c.BodyParser(&productAPI); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(productAPI)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var productAsset model.ProductAsset

	var product model.Product
	product.ProductName = productAPI.ProductName
	product.Description = productAPI.Description
	product.Price = productAPI.Price
	product.Quantity = productAPI.Quantity
	product.CategoryID = productAPI.CategoryID
	product.BrandID = productAPI.BrandID
	db.Create(&product)

	// Handler File
	fileName := c.Locals("thumbnail")

	if fileName == nil {
		return c.Status(422).JSON(fiber.Map{
			"message": "thumbnail is required",
		})
	} else {
		thumbnail := fmt.Sprintf("%v", fileName)
		// errSaveFile := c.SaveFile(fileName, fmt.Sprintf("./public/product-asset/%s", thumbnail))
		filePath := fmt.Sprintf("./public/thumbnail/%s", thumbnail)

		db.Model(&productAsset).Create(map[string]interface{}{
			"id":         uuid.New(),
			"created_at": time.Now(),
			"updated_at": time.Now(),
			"product_id": product.ID,
			"file_name":  thumbnail,
			"file_path":  filePath,
			"is_primary": true,
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["product_asset"]

	for _, file := range files {
		errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/product-asset/%s", file.Filename))
		filePath := fmt.Sprintf("./public/product-asset/%s", file.Filename)
		if errSaveFile != nil {
			log.Println("Upload Failed ")
		}
		db.Model(&productAsset).Create(map[string]interface{}{
			"id":         uuid.New(),
			"created_at": time.Now(),
			"updated_at": time.Now(),
			"product_id": product.ID,
			"file_name":  file.Filename,
			"file_path":  filePath,
			"is_primary": false,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    product,
	})
}
