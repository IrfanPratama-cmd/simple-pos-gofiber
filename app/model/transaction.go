package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type Transaction struct {
	Base
	TransactionAPI
	Customer          *Customer          `json:"customer,omitempty" gorm:"foreignKey:CustomerID;references:ID"`
	TransactionDetail *TransactionDetail `json:"transaction_detail,omitempty" gorm:"foreignKey:ID;references:TransactionID"`
}

type TransactionAPI struct {
	CustomerID        *uuid.UUID      `json:"customer_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	InvoiceNo         *string         `json:"invoice_no,omitempty"  gorm:"type:varchar(191)" example:"INV-000000000000072270623"`
	TransactionDate   strfmt.DateTime `json:"transaction_date,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz;not null"`
	TransactionStatus *string         `json:"transaction_status,omitempty" gorm:"type:varchar(191);not null" example:"pending" `
	TransactionType   *string         `json:"transaction_type,omitempty" gorm:"type:varchar(191);not null" example:"cash" `
	TotalAmount       *float64        `json:"total_amount,omitempty" example:"127000"`
}

type TransactionRequest struct {
	TransactionType *string `json:"transaction_type,omitempty" gorm:"type:varchar(191);not null" example:"cash" `
}

type TransactionResponse struct {
	Base
	TransactionAPI
	Customer          *Customer          `json:"customer,omitempty" gorm:"foreignKey:CustomerID;references:ID"`
	TransactionDetail *TransactionDetail `json:"TransactionDetail,omitempty" gorm:"foreignKey:ID;references:TransactionID"`
}
