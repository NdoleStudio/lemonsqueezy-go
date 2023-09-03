package lemonsqueezy

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// SubscriptionItemsService is the API client for the `/v1/subscription-items` endpoint
type SubscriptionItemsService service

// Get returns the subscription item with the given ID.
//
// https://docs.lemonsqueezy.com/api/subscription-items#retrieve-a-subscription-item
func (service *SubscriptionItemsService) Get(ctx context.Context, subscriptionItemID string) (*SubscriptionItemApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/subscription-items/"+subscriptionItemID)
	if err != nil {
		return nil, response, err
	}

	subscriptionItem := new(SubscriptionItemApiResponse)
	if err = json.Unmarshal(*response.Body, subscriptionItem); err != nil {
		return nil, response, err
	}

	return subscriptionItem, response, nil
}

// List a paginated list of subscription items.
//
// https://docs.lemonsqueezy.com/api/subscription-items#list-all-subscription-items
func (service *SubscriptionItemsService) List(ctx context.Context, params *SubscriptionItemsListParams) (*SubscriptionItemsApiResponse, *Response, error) {
	paramsURI := ""
	if params.SubscriptionID != 0 {
		paramsURI = fmt.Sprintf("?filter[subscription_id]=%v", params.SubscriptionID)
	}
	if params.PriceID != 0 {
		paramsURI += fmt.Sprintf("&[price_id]=%v", params.PriceID)
	}
	response, err := service.client.do(ctx, http.MethodGet, "/v1/subscription-items"+paramsURI)
	if err != nil {
		return nil, response, err
	}

	subscriptionItems := new(SubscriptionItemsApiResponse)
	if err = json.Unmarshal(*response.Body, subscriptionItems); err != nil {
		return nil, response, err
	}

	return subscriptionItems, response, nil
}
