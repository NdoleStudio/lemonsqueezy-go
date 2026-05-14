package lemonsqueezy

import "time"

// SubscriptionItem is an object that links a price to a subscription and also contains quantity information.
// https://docs.lemonsqueezy.com/api/subscription-items#the-subscription-item-object
type SubscriptionItem struct {
	SubscriptionID int       `json:"subscription_id"`
	PriceID        int       `json:"price_id"`
	Quantity       int       `json:"quantity"`
	IsUsageBased   bool      `json:"is_usage_based"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// SubscriptionItemUpdateParams are parameters for updating a subscription item
type SubscriptionItemUpdateParams struct {
	ID         string                                 `json:"id"`
	Attributes SubscriptionItemUpdateParamsAttributes `json:"attributes"`
}

// SubscriptionItemListParams are parameters for filtering list responses
type SubscriptionItemListParams struct {
	SubscriptionID string
	PriceID        string
}

// SubscriptionItemUpdateParamsAttributes are subscription item update attributes
type SubscriptionItemUpdateParamsAttributes struct {
	Quantity int `json:"quantity,omitempty"`
}

// APIResponseMetaSubscriptionItemCurrentUsage contains usage metadata for a subscription item
type APIResponseMetaSubscriptionItemCurrentUsage struct {
	PeriodStart      time.Time `json:"period_start"`
	PeriodEnd        time.Time `json:"period_end"`
	Quantity         int       `json:"quantity"`
	IntervalUnit     string    `json:"interval_unit"`
	IntervalQuantity int       `json:"interval_quantity"`
}

// APIResponseRelationshipsSubscriptionItem contains the relationships of a subscription item object
type APIResponseRelationshipsSubscriptionItem struct {
	Subscription ApiResponseLinks `json:"subscription"`
	Price        ApiResponseLinks `json:"price"`
	UsageRecords ApiResponseLinks `json:"usage-records"`
}

// SubscriptionItemAPIResponse represents a subscription item api response
type SubscriptionItemAPIResponse = ApiResponse[SubscriptionItem, APIResponseRelationshipsSubscriptionItem]

// SubscriptionItemsAPIResponse represents a list of subscription items api responses
type SubscriptionItemsAPIResponse = ApiResponseList[SubscriptionItem, APIResponseRelationshipsSubscriptionItem]

// SubscriptionItemCurrentUsageAPIResponse represents the subscription item's current usage api response
type SubscriptionItemCurrentUsageAPIResponse struct {
	Jsonapi ApiResponseJSONAPI                          `json:"jsonapi"`
	Meta    APIResponseMetaSubscriptionItemCurrentUsage `json:"meta"`
}
