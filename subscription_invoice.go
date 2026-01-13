package lemonsqueezy

import "time"

// SubscriptionInvoiceAttributes is the invoice for a subscription
// https://docs.lemonsqueezy.com/api/subscription-invoices#the-subscription-invoice-object
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

// APIResponseRelationshipsSubscriptionInvoice relationships of a subscription invoice
type APIResponseRelationshipsSubscriptionInvoice struct {
	Store        ApiResponseLinks `json:"store"`
	Subscription ApiResponseLinks `json:"subscription"`
}

// SubscriptionInvoiceAPIResponse is the api response for one subscription invoice
type SubscriptionInvoiceAPIResponse = ApiResponse[SubscriptionInvoiceAttributes, APIResponseRelationshipsSubscriptionInvoice]

// SubscriptionInvoicesAPIResponse is the api response for a list of subscription invoices.
type SubscriptionInvoicesAPIResponse = ApiResponseList[SubscriptionInvoiceAttributes, APIResponseRelationshipsSubscriptionInvoice]

// SubscriptionInvoiceGenerateResponse is the response when generating a subscription invoice
type SubscriptionInvoiceGenerateResponse struct {
	JSONAPI ApiResponseJSONAPI `json:"jsonapi"`
	Meta    struct {
		Urls struct {
			DownloadInvoice string `json:"download_invoice"`
		} `json:"urls"`
	} `json:"meta"`
}
