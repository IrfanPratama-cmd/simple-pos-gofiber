package model

type Payment struct {
	Base
	PaymentAPI
}

type PaymentAPI struct {
	CheckoutLink  string  `json:"checkout_link,omitempty" `
	ExternalID    string  `json:"external_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	PaymentStatus string  `json:"payment_status,omitempty" gorm:"type:varchar(191);not null" example:"due" `
	Amount        float64 `json:"amount,omitempty"`
}

type PaymentRequest struct {
	Amount float64 `json:"amount,omitempty"`
}

type PaymentLinkResponse struct {
	CheckoutLink string `json:"checkout_link,omitempty" example:"https://checkout-staging.xendit.co/latest/65a0eaaa63346762f641c28a" `
}
