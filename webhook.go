package client

import "time"

// WebHookRequestSubscription is the webhook request for the `subscription_created` event
type WebHookRequestSubscription webhookRequest[SubscriptionCreatedWebhookRequestAttributes, SubscriptionCreatedWebhookRequestRelationships]

// WebhookRequest is a generic webhook request
type WebhookRequest webhookRequest[any, any]

// WebhookRequest represents a webhook request
type webhookRequest[T, R any] struct {
	Meta WebhookRequestMeta       `json:"meta"`
	Data WebhookRequestData[T, R] `json:"data"`
}

// WebhookRequestMeta contains meta data about the request
type WebhookRequestMeta struct {
	EventName  string         `json:"event_name"`
	TestMode   bool           `json:"test_mode"`
	CustomData map[string]any `json:"custom_data"`
}

// WebhookRequestAttributes attributes of a webhook request
type OrderCreatedWebhookRequestAttributes struct {
	StoreID          int        `json:"store_id"`
	Identifier       string     `json:"identifier"`
	OrderNumber      int        `json:"order_number"`
	UserName         string     `json:"user_name"`
	UserEmail        string     `json:"user_email"`
	Currency         string     `json:"currency"`
	CurrencyRate     string     `json:"currency_rate"`
	Subtotal         int        `json:"subtotal"`
	DiscountTotal    int        `json:"discount_total"`
	Tax              int        `json:"tax"`
	Total            int        `json:"total"`
	SubtotalUsd      int        `json:"subtotal_usd"`
	DiscountTotalUsd int        `json:"discount_total_usd"`
	TaxUsd           int        `json:"tax_usd"`
	TotalUsd         int        `json:"total_usd"`
	TaxName          string     `json:"tax_name"`
	TaxRate          string     `json:"tax_rate"`
	Status           string     `json:"status"`
	StatusFormatted  string     `json:"status_formatted"`
	Refunded         bool       `json:"refunded"`
	RefundedAt       *time.Time `json:"refunded_at"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

// SubscriptionCreatedWebhookRequestAttributes are attributes for the subscription created event
type SubscriptionCreatedWebhookRequestAttributes struct {
	Urls            SubscriptionURLs   `json:"urls"`
	Pause           *SubscriptionPause `json:"pause"`
	Status          string             `json:"status"`
	EndsAt          *time.Time         `json:"ends_at"`
	OrderID         int                `json:"order_id"`
	StoreID         int                `json:"store_id"`
	Cancelled       bool               `json:"cancelled"`
	RenewsAt        time.Time          `json:"renews_at"`
	TestMode        bool               `json:"test_mode"`
	UserName        string             `json:"user_name"`
	CreatedAt       time.Time          `json:"created_at"`
	ProductID       int                `json:"product_id"`
	UpdatedAt       time.Time          `json:"updated_at"`
	UserEmail       string             `json:"user_email"`
	VariantID       int                `json:"variant_id"`
	ProductName     string             `json:"product_name"`
	VariantName     string             `json:"variant_name"`
	OrderItemID     int                `json:"order_item_id"`
	TrialEndsAt     time.Time          `json:"trial_ends_at"`
	BillingAnchor   int                `json:"billing_anchor"`
	StatusFormatted string             `json:"status_formatted"`
}

// WebhookRequestRelationships webhook request relationships
type OrderCreatedWebhookRequestRelationships struct {
	Store         ApiResponseLinks `json:"store"`
	OrderItem     ApiResponseLinks `json:"order-items"`
	Subscriptions ApiResponseLinks `json:"subscriptions"`
	LicenseKeys   ApiResponseLinks `json:"license-keys"`
}

// SubscriptionCreatedWebhookRequestRelationships are relationships for the subsciption created event
type SubscriptionCreatedWebhookRequestRelationships struct {
	Order     ApiResponseLinks `json:"order"`
	Store     ApiResponseLinks `json:"store"`
	Product   ApiResponseLinks `json:"product"`
	Variant   ApiResponseLinks `json:"variant"`
	OrderItem ApiResponseLinks `json:"order-item"`
}

// WebhookRequestData webhook request data
type WebhookRequestData[T, R any] struct {
	Type          string          `json:"type"`
	ID            string          `json:"id"`
	Attributes    T               `json:"attributes"`
	Relationships R               `json:"relationships"`
	Links         ApiResponseLink `json:"links"`
}
