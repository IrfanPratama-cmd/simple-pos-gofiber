package model

import "github.com/google/uuid"

type Product struct {
	Base
	ProductAPI
	Category     *Category     `json:"category,omitempty" gorm:"foreignKey:CategoryID;references:ID"`
	Brand        *Brand        `json:"brand,omitempty" gorm:"foreignKey:BrandID;references:ID"`
	ProductAsset *ProductAsset `json:"product_asset,omitempty" gorm:"foreignKey:ProductID;references:ID"`
}

type ProductAPI struct {
	ProductName string     `json:"product_name,omitempty" example:"Samsung"  validate:"required"`
	Sku         *string    `json:"sku,omitempty" example:"PR6900000007" gorm:"type:varchar(191)" `
	Description string     `json:"description,omitempty" example:"Samsung Desc"  validate:"required"`
	Quantity    int        `json:"quantity,omitempty" example:"10" `
	Price       float64    `json:"price,omitempty" example:"10000" `
	CategoryID  *uuid.UUID `json:"category_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	BrandID     *uuid.UUID `json:"brand_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
}

type ProductRequest struct {
	ProductName  string         `json:"product_name,omitempty" example:"Samsung" validate:"required"`
	Description  string         `json:"description,omitempty" example:"Samsung Desc" validate:"required"`
	Quantity     int            `json:"quantity,omitempty" example:"10"`
	Price        float64        `json:"price,omitempty" example:"10000"`
	CategoryID   *uuid.UUID     `json:"category_id,omitempty" swaggertype:"string" gorm:"type:varchar(36);index" format:"uuid"`
	BrandID      *uuid.UUID     `json:"brand_id,omitempty" swaggertype:"string" gorm:"type:varchar(36);index" format:"uuid"`
	ProductAsset []ProductAsset `json:"product_asset,omitempty"`
}

type ProductCheckout struct {
	Base
	ProductName string     `json:"product_name,omitempty" example:"Samsung"  validate:"required"`
	Description string     `json:"description,omitempty" example:"Samsung Desc"  validate:"required"`
	CategoryID  *uuid.UUID `json:"category_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	BrandID     *uuid.UUID `json:"brand_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
}
