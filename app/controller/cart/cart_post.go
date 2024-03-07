package cart

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PostCart godoc
// @Summary Create new Cart
// @Description Create new Cart
// @Param data body model.CartRequest true "Cart data"
// @Accept  application/json
// @Produce application/json
// @Success 201 {object} model.Cart "Cart data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /carts [post]
// @Tags Cart
func PostCart(c *fiber.Ctx) error {
	var cartAPI model.CartRequest

	db := services.DB

	if err := c.BodyParser(&cartAPI); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	var cart model.Cart
	res := db.Model(&cart).Where("product_id = ?", cartAPI.ProductID).First(&cart)
	if res.RowsAffected != 0 {
		qty := *cart.Qty
		qty += *cartAPI.Qty
		cart.Qty = &qty
		db.Model(&cart).Where("product_id = ?", cartAPI.ProductID).Updates(&cart)
	} else {
		cart.ProductID = cartAPI.ProductID
		cart.Qty = cartAPI.Qty
		cart.CustomerID = cartAPI.CustomerID
		db.Create(&cart)
	}

	return lib.Created(c, cart)
}
