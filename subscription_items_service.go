package lemonsqueezy

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// SubscriptionItemsService is the API client for the `/subscription-items` endpoint
type SubscriptionItemsService service

// Update a subscription item
//
// https://docs.lemonsqueezy.com/api/subscription-items#update-a-subscription-item
func (service *SubscriptionItemsService) Update(ctx context.Context, params *SubscriptionItemUpdateParams) (*SubscriptionItemApiResponse, *Response, error) {
	payload := map[string]any{
		"data": map[string]any{
			"type":       "subscription-items",
			"id":         params.ID,
			"attributes": params.Attributes,
		},
	}

	response, err := service.client.do(ctx, http.MethodPatch, "/v1/subscriptions-items/"+params.ID, payload)
	if err != nil {
		return nil, response, err
	}

	subscriptionItem := new(SubscriptionItemApiResponse)
	if err = json.Unmarshal(*response.Body, subscriptionItem); err != nil {
		return nil, response, err
	}

	return subscriptionItem, response, nil
}

// List returns a paginated list of subscription items you can add extra query params to your request
//
// https://docs.lemonsqueezy.com/api/subscriptions#list-all-subscriptions
func (service *SubscriptionItemsService) List(ctx context.Context, queryParams map[string]string) (*SubscriptionItemsApiResponse, *Response, error) {
	basePath := "/v1/subscription-items"
	parsedURL, err := url.Parse(basePath)
	if err != nil {
		return nil, nil, err
	}

	if queryParams != nil {
		query := parsedURL.Query()
		for key, val := range queryParams {
			query.Add(key, val)
		}
		parsedURL.RawQuery = query.Encode()
	}

	response, err := service.client.do(ctx, http.MethodGet, parsedURL.String())
	if err != nil {
		return nil, response, err
	}

	subscriptionItems := new(SubscriptionItemsApiResponse)
	if err = json.Unmarshal(*response.Body, subscriptionItems); err != nil {
		return nil, response, err
	}

	return subscriptionItems, response, nil
}

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

// Current usage returns a subscription item's current usage with the given ID.
//
// https://docs.lemonsqueezy.com/api/subscription-items#retrieve-a-subscription-item-s-current-usage
func (service *SubscriptionItemsService) CurrentUsage(ctx context.Context, subscriptionItemID string) (*SubscriptionItemCurrentUsageApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/subscription-items/"+subscriptionItemID+"/current-usage")
	if err != nil {
		return nil, response, err
	}

	subscriptionItemCurrentUsage := new(SubscriptionItemCurrentUsageApiResponse)
	if err = json.Unmarshal(*response.Body, subscriptionItemCurrentUsage); err != nil {
		return nil, response, err
	}

	return subscriptionItemCurrentUsage, response, nil
}
