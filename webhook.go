package lemonsqueezy

import "time"

// WebhookCreateParams are parameters for creating a webhook
type WebhookCreateParams struct {
	URL    string   `json:"url"`
	Events []string `json:"events"`
	Secret string   `json:"secret"`
}

// WebhookUpdateParams are parameters for updating a webhook
type WebhookUpdateParams struct {
	ID     string   `json:"id"`
	Secret string   `json:"secret"`
	Events []string `json:"events"`
}

// WebhookAttributes contains information about a webhook.
type WebhookAttributes struct {
	StoreID    int        `json:"store_id"`
	URL        string     `json:"url"`
	Events     []string   `json:"events"`
	LastSentAt *time.Time `json:"last_sent_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	TestMode   bool       `json:"test_mode"`
}

// ApiResponseRelationshipWebhook relationships of a webhook
type ApiResponseRelationshipWebhook struct {
	Store ApiResponseLinks `json:"store"`
}

// WebhookApiResponse is the api response for one webhook
type WebhookApiResponse = ApiResponse[WebhookAttributes, ApiResponseRelationshipWebhook]

// WebhooksApiResponse is the api response for a list of webhooks.
type WebhooksApiResponse = ApiResponseList[WebhookAttributes, ApiResponseRelationshipWebhook]
