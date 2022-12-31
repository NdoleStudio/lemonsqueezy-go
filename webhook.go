package client

import "time"

// WebhookRequest represents a webhook request
type WebhookRequest struct {
	Meta WebhookRequestMeta `json:"meta"`
	Data WebhookRequestData `json:"data"`
}

// WebhookRequestMeta contains meta data about the request
type WebhookRequestMeta struct {
	EventName  string         `json:"event_name"`
	TestMode   bool           `json:"test_mode"`
	CustomData map[string]any `json:"custom_data"`
}

// WebhookRequestAttributes attributes of a webhook request
type WebhookRequestAttributes struct {
	StoreID          int         `json:"store_id"`
	Identifier       string      `json:"identifier"`
	OrderNumber      int         `json:"order_number"`
	UserName         string      `json:"user_name"`
	UserEmail        string      `json:"user_email"`
	Currency         string      `json:"currency"`
	CurrencyRate     string      `json:"currency_rate"`
	Subtotal         int         `json:"subtotal"`
	DiscountTotal    int         `json:"discount_total"`
	Tax              int         `json:"tax"`
	Total            int         `json:"total"`
	SubtotalUsd      int         `json:"subtotal_usd"`
	DiscountTotalUsd int         `json:"discount_total_usd"`
	TaxUsd           int         `json:"tax_usd"`
	TotalUsd         int         `json:"total_usd"`
	TaxName          string      `json:"tax_name"`
	TaxRate          string      `json:"tax_rate"`
	Status           string      `json:"status"`
	StatusFormatted  string      `json:"status_formatted"`
	Refunded         bool        `json:"refunded"`
	RefundedAt       interface{} `json:"refunded_at"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
}

// WebhookRequestRelationships webhook request relationships
type WebhookRequestRelationships struct {
	Store         ApiResponseLinks `json:"store"`
	OrderItem     ApiResponseLinks `json:"order-items"`
	Subscriptions ApiResponseLinks `json:"subscriptions"`
	LicenseKeys   ApiResponseLinks `json:"license-keys"`
}

// WebhookRequestData webhook request data
type WebhookRequestData struct {
	Type          string                      `json:"type"`
	ID            string                      `json:"id"`
	Attributes    WebhookRequestAttributes    `json:"attributes"`
	Relationships WebhookRequestRelationships `json:"relationships"`
	Links         ApiResponseLink             `json:"links"`
}
