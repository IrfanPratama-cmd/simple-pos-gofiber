package transaction

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gofiber/fiber/v2"
)

// PostTransaction godoc
// @Summary Create new Transaction
// @Description Create new Transaction
// @Accept  application/json
// @Produce application/json
// @Success 201 {object} model.Transaction "Transaction data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /transactions [post]
// @Tags Transaction
func PostTransaction(c *fiber.Ctx) error {
	var api model.TransactionRequest

	if err := c.BodyParser(&api); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	db := services.DB

	id := c.Params("customer_id")

	var customer model.Customer
	result := db.Model(&customer).Where("id", id).First(&customer)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Customer not found",
		})
	}

	var cart []model.Cart
	db.Model(&cart).Where("customer_id", customer.ID).Find(&cart)

	transactionID := lib.GenUUID()
	currentTime := time.Now()

	var totalAmount float64

	for _, c := range cart {
		var product model.Product
		db.Model(&product).Where("id", c.ProductID).First(&product)

		var data model.TransactionDetail
		data.CustomerID = customer.ID
		data.ProductID = c.ProductID
		data.TransactionID = transactionID
		data.Qty = c.Qty
		data.Amount = &product.Price
		data.TotalAmount = lib.Float64ptr(float64(*c.Qty) * product.Price)
		db.Create(&data)

		totalAmount += *data.TotalAmount
	}

	invoiceNo := lib.RandomNumber(6)

	var transaction model.Transaction
	transaction.ID = transactionID
	transaction.TotalAmount = &totalAmount
	transaction.TransactionDate = strfmt.DateTime(currentTime)
	transaction.TransactionStatus = lib.Strptr("pending")
	transaction.InvoiceNo = &invoiceNo
	transaction.CustomerID = customer.ID
	transaction.TransactionType = api.TransactionType
	db.Create(&transaction)

	db.Delete(&cart)

	return lib.Created(c, transaction)
}
