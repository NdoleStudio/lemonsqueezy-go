package client

// WebhookRequestOrder is the payload for order related events e.g. order_created, order_refunded
type WebhookRequestOrder WebhookRequest[OrderAttributes, ApiResponseRelationshipsOrder]

// WebhookRequestSubscription is the payload for subscription related events e.g. subscription_created, subscription_updated
type WebhookRequestSubscription WebhookRequest[Subscription, ApiResponseRelationshipsSubscription]

// WebhookRequestSubscriptionInvoice is the payload for subscription invoice related events e.g. subscription_payment_success, subscription_payment_failed
type WebhookRequestSubscriptionInvoice WebhookRequest[SubscriptionInvoiceAttributes, ApiResponseRelationshipsSubscriptionInvoice]

// WebhookRequestLicenseKey is the payload for license key related events e.g. license_key_created
type WebhookRequestLicenseKey WebhookRequest[LicenseKeyAttributes, ApiResponseRelationshipsLicenseKey]

// WebhookRequest represents a webhook request
type WebhookRequest[T, R any] struct {
	Meta WebhookRequestMeta       `json:"meta"`
	Data WebhookRequestData[T, R] `json:"data"`
}

// WebhookRequestMeta contains meta data about the request
type WebhookRequestMeta struct {
	EventName  string         `json:"event_name"`
	TestMode   bool           `json:"test_mode"`
	CustomData map[string]any `json:"custom_data"`
}

// WebhookRequestData webhook request data
type WebhookRequestData[Attributes, Relationships any] struct {
	Type          string              `json:"type"`
	ID            string              `json:"id"`
	Attributes    Attributes          `json:"attributes"`
	Relationships Relationships       `json:"relationships"`
	Links         ApiResponseSelfLink `json:"links"`
}
