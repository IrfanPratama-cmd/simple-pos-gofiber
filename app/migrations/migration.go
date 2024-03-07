package migrations

import "api/app/model"

// ModelMigrations models to automigrate
var ModelMigrations = []interface{}{
	&model.Asset{},
	&model.User{},
	&model.Customer{},
	&model.Category{},
	&model.Cart{},
	&model.Brand{},
	&model.Product{},
	&model.ProductAsset{},
	&model.TransactionDetail{},
	&model.Transaction{},
	&model.Payment{},
}
