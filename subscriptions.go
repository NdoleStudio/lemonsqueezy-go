package lemonsqueezy

import "time"

// Subscription is created when a subscription product is purchased and will bill the customer on a recurring basis.
// https://docs.lemonsqueezy.com/api/subscriptions#the-subscription-object
type Subscription struct {
	StoreID               int                                `json:"store_id"`
	CustomerID            int                                `json:"customer_id"`
	OrderID               int                                `json:"order_id"`
	OrderItemID           int                                `json:"order_item_id"`
	ProductID             int                                `json:"product_id"`
	VariantID             int                                `json:"variant_id"`
	ProductName           string                             `json:"product_name"`
	VariantName           string                             `json:"variant_name"`
	UserName              string                             `json:"user_name"`
	UserEmail             string                             `json:"user_email"`
	Status                string                             `json:"status"`
	StatusFormatted       string                             `json:"status_formatted"`
	CardBrand             string                             `json:"card_brand"`
	CardLastFour          string                             `json:"card_last_four"`
	Pause                 *SubscriptionPause                 `json:"pause"`
	Cancelled             bool                               `json:"cancelled"`
	TrialEndsAt           *time.Time                         `json:"trial_ends_at"`
	BillingAnchor         int                                `json:"billing_anchor"`
	FirstSubscriptionItem *SubscriptionFirstSubscriptionItem `json:"first_subscription_item"`
	Urls                  SubscriptionURLs                   `json:"urls"`
	RenewsAt              time.Time                          `json:"renews_at"`
	EndsAt                *time.Time                         `json:"ends_at"`
	CreatedAt             time.Time                          `json:"created_at"`
	UpdatedAt             time.Time                          `json:"updated_at"`
	TestMode              bool                               `json:"test_mode"`
}

// SubscriptionSubscriptionItem is an object representing the first subscription item belonging to this subscription.
type SubscriptionFirstSubscriptionItem struct {
	ID int `json:"id"`
	SubscriptionItem
}

// SubscriptionURLs is object of customer-facing URLs for managing the subscription.
type SubscriptionURLs struct {
	UpdatePaymentMethod string `json:"update_payment_method"`
	CustomerPortal      string `json:"customer_portal"`
}

// SubscriptionPause is object of customer-facing URLs for managing the subscription.
type SubscriptionPause struct {
	Mode      string    `json:"mode"`
	ResumesAt time.Time `json:"resumes_at"`
}

// SubscriptionUpdateParams are parameters for updating a subscription
type SubscriptionUpdateParams struct {
	ID         string                             `json:"id"`
	Attributes SubscriptionUpdateParamsAttributes `json:"attributes"`
}

// SubscriptionUpdateParamsAttributes are subscription update attributes
type SubscriptionUpdateParamsAttributes struct {
	ProductID          int                `json:"product_id,omitempty"`
	VariantID          int                `json:"variant_id,omitempty"`
	BillingAnchor      int                `json:"billing_anchor,omitempty"`
	Cancelled          bool               `json:"cancelled"`
	Pause              *SubscriptionPause `json:"pause,omitempty"`
	InvoiceImmediately bool               `json:"invoice_immediately"`
	DisableProrations  bool               `json:"disable_prorations"`
}

// ApiResponseRelationshipsSubscription relationships of a subscription object
type ApiResponseRelationshipsSubscription struct {
	Store                ApiResponseLinks `json:"store"`
	Customer             ApiResponseLinks `json:"customer"`
	Order                ApiResponseLinks `json:"order"`
	OrderItem            ApiResponseLinks `json:"order-item"`
	Product              ApiResponseLinks `json:"product"`
	Variant              ApiResponseLinks `json:"variant"`
	SubscriptionItems    ApiResponseLinks `json:"subscription-items"`
	SubscriptionInvoices ApiResponseLinks `json:"subscription-invoices"`
}

// SubscriptionApiResponse represents a subscription api response
type SubscriptionApiResponse = ApiResponse[Subscription, ApiResponseRelationshipsSubscription]

// SubscriptionsApiResponse represents a list of subscription api responses.
type SubscriptionsApiResponse = ApiResponseList[Subscription, ApiResponseRelationshipsSubscription]
