package client

import "time"

// SubscriptionInvoiceAttributes is the invoice for a subscription
type SubscriptionInvoiceAttributes struct {
	StoreID                int        `json:"store_id"`
	SubscriptionID         int        `json:"subscription_id"`
	BillingReason          string     `json:"billing_reason"`
	CardBrand              string     `json:"card_brand"`
	CardLastFour           string     `json:"card_last_four"`
	Currency               string     `json:"currency"`
	CurrencyRate           string     `json:"currency_rate"`
	Subtotal               int        `json:"subtotal"`
	DiscountTotal          int        `json:"discount_total"`
	Tax                    int        `json:"tax"`
	Total                  int        `json:"total"`
	SubtotalUsd            int        `json:"subtotal_usd"`
	DiscountTotalUsd       int        `json:"discount_total_usd"`
	TaxUsd                 int        `json:"tax_usd"`
	TotalUsd               int        `json:"total_usd"`
	Status                 string     `json:"status"`
	StatusFormatted        string     `json:"status_formatted"`
	Refunded               bool       `json:"refunded"`
	RefundedAt             *time.Time `json:"refunded_at"`
	SubtotalFormatted      string     `json:"subtotal_formatted"`
	DiscountTotalFormatted string     `json:"discount_total_formatted"`
	TaxFormatted           string     `json:"tax_formatted"`
	TotalFormatted         string     `json:"total_formatted"`
	Urls                   struct {
		InvoiceURL string `json:"invoice_url"`
	} `json:"urls"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	TestMode  bool      `json:"test_mode"`
}

// ApiResponseRelationshipsSubscriptionInvoice relationships of a subscription invoice
type ApiResponseRelationshipsSubscriptionInvoice struct {
	Store        ApiResponseLinks `json:"store"`
	Subscription ApiResponseLinks `json:"subscription"`
}

// SubscriptionInvoiceApiResponse is the api response for one subscription invoice
type SubscriptionInvoiceApiResponse = ApiResponse[SubscriptionInvoiceAttributes, ApiResponseRelationshipsSubscriptionInvoice]

// SubscriptionInvoicesApiResponse is the api response for a list of subscription invoices.
type SubscriptionInvoicesApiResponse = ApiResponseList[SubscriptionInvoiceAttributes, ApiResponseRelationshipsSubscriptionInvoice]
