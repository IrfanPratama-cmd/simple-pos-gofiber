package model

import "github.com/google/uuid"

type TransactionDetail struct {
	Base
	TransactionDetailAPI
	Customer    *Customer    `json:"customer,omitempty" gorm:"foreignKey:CustomerID;references:ID"`
	Transaction *Transaction `json:"transaction,omitempty" gorm:"foreignKey:TransactionID;references:ID"`
	Product     *Product     `json:"product,omitempty" gorm:"foreignKey:ProductID;references:ID"`
}

type TransactionDetailAPI struct {
	CustomerID    *uuid.UUID `json:"customer_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	ProductID     *uuid.UUID `json:"product_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	TransactionID *uuid.UUID `json:"transaction_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	Qty           *int       `json:"qty,omitempty"`
	Amount        *float64   `json:"amount,omitempty"`
	TotalAmount   *float64   `json:"total_amount,omitempty"`
}

type TransactionDetailResponse struct {
	Base
	TransactionDetailAPI
	Customer    *Customer        `json:"customer,omitempty" gorm:"foreignKey:CustomerID;references:ID"`
	Transaction *Transaction     `json:"transaction,omitempty" gorm:"foreignKey:TransactionID;references:ID"`
	Product     *ProductCheckout `json:"product,omitempty" gorm:"foreignKey:ProductID;references:ID"`
}
