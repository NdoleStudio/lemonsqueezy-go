package lemonsqueezy

import "time"

// In Lemon Squeezy A subscription item is an object that links a price to a subscription and also contains quantity information.
// https://docs.lemonsqueezy.com/api/subscription-items#the-subscription-item-object
type SubscriptionItem struct {
	SubscriptionID int       `json:"subscription_id"`
	PriceID        int       `json:"price_id"`
	Quantity       int       `json:"quantity"`
	IsUsageBased   bool      `json:"is_usage_based"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// SubscriptionUpdateParams are parameters for updating a subscription
type SubscriptionItemUpdateParams struct {
	ID         string                                 `json:"id"`
	Attributes SubscriptionItemUpdateParamsAttributes `json:"attributes"`
}

// SubscriptionItemListParams are parameters for filtering list responses
type SubscriptionItemListParams struct {
	SubscriptionID string
	PriceID        string
}

// SubscriptionUpdateParamsAttributes are subscription update attributes
type SubscriptionItemUpdateParamsAttributes struct {
	Quantity int `json:"quantity,omitempty"`
}

type ApiResponseMetaSubscriptionItemCurrentUsage struct {
	PeriodStart      time.Time `json:"period_start"`
	PeriodEnd        time.Time `json:"period_end"`
	Quantity         int       `json:"quantity"`
	IntervalUnit     string    `json:"interval_unit"`
	IntervalQuantity int       `json:"interval_quantity"`
}

// ApiResponseRelationshipsSubscription relationships of a subscription item object
type ApiResponseRelationshipsSubscriptionItem struct {
	Subscription ApiResponseLinks `json:"subscription"`
	Price        ApiResponseLinks `json:"price"`
	UsageRecords ApiResponseLinks `json:"usage-records"`
}

// SubscriptionItemApiResponse represents a subscription item api response
type SubscriptionItemApiResponse = ApiResponse[SubscriptionItem, ApiResponseRelationshipsSubscriptionItem]

// SubscriptionItemsApiResponse represents a list of subscription items api responses
type SubscriptionItemsApiResponse = ApiResponseList[SubscriptionItem, ApiResponseRelationshipsSubscriptionItem]

// SubscriptionItemsCurrentUsageApiResponse represents the subscription item's current usage api response
type SubscriptionItemCurrentUsageApiResponse struct {
	Jsonapi ApiResponseJSONAPI                          `json:"jsonapi"`
	Meta    ApiResponseMetaSubscriptionItemCurrentUsage `json:"meta"`
}
