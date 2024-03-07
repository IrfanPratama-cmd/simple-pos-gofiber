package model

import "github.com/google/uuid"

type Cart struct {
	Base
	CartAPI
	Customer *Customer `json:"customer,omitempty" gorm:"foreignKey:CustomerID;references:ID"`
	Product  *Product  `json:"product,omitempty" gorm:"foreignKey:ProductID;references:ID"`
}

type CartAPI struct {
	CustomerID *uuid.UUID `json:"customer_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	ProductID  *uuid.UUID `json:"product_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	Qty        *int       `json:"qty,omitempty"`
}

type CartRequest struct {
	CustomerID *uuid.UUID `json:"customer_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	ProductID  *uuid.UUID `json:"product_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	Qty        *int       `json:"qty,omitempty"`
}

type CartUpdate struct {
	Qty *int `json:"qty,omitempty"`
}
