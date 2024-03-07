package customer

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

// GetCustomer godoc
// @Summary List of Customer
// @Description List of Customer
// @Param page query int false "Page number start from zero"
// @Param size query int false "Size per page, default `0`"
// @Param sort query string false "Sort by field, adding dash (`-`) at the beginning means descending and vice versa"
// @Param fields query string false "Select specific fields with comma separated"
// @Param filters query string false "custom filters, see [more details](https://github.com/morkid/paginate#filter-format)"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Page{items=[]model.Customer} "List of Customer"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /customers [get]
// @Tags Customer
func GetCustomer(c *fiber.Ctx) error {
	db := services.DB
	pg := paginate.New()

	mod := db.Model(&model.Customer{})

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Customer{})

	return lib.OK(c, page)
}
