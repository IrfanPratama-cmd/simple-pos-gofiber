package product

import (
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

// GetProduct godoc
// @Summary List of Product
// @Description List of Product
// @Param page query int false "Page number start from zero"
// @Param size query int false "Size per page, default `0`"
// @Param sort query string false "Sort by field, adding dash (`-`) at the beginning means descending and vice versa"
// @Param fields query string false "Select specific fields with comma separated"
// @Param filters query string false "custom filters, see [more details](https://github.com/morkid/paginate#filter-format)"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Page{items=[]model.Product} "List of Product"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /products [get]
// @Tags Product
func GetProduct(c *fiber.Ctx) error {
	db := services.DB
	pg := paginate.New()

	var product model.Product

	mod := db.Model(&product).
		Select(`products.id, products.created_at, products.updated_at, 
				products.product_name, products.description, products.price, products.quantity,
				c.id "Category__id",
				c.category_name "Category__category_name",
				c.category_code "Category__category_code",
				b.id "Brand__id",
				b.brand_name "Brand__brand_name",
				b.brand_code "Brand__brand_code",
				pa.id "ProductAsset__id",
				pa.file_name "ProductAsset__file_name",
				pa.file_path "ProductAsset__file_path",
				pa.is_primary "ProductAsset__is_primary"
				`).
		Joins(`LEFT JOIN brands b on b.id = products.brand_id`).
		Joins(`LEFT JOIN categories c on c.id = products.category_id`).
		Joins(`LEFT JOIN product_assets pa on pa.product_id = products.id`).
		Where(`pa.is_primary = ?`, true)

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Product{})
	return c.Status(200).JSON(page)
}
