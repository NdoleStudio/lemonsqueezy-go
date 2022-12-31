package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// subscriptionsService is the API client for the `/subscriptions` endpoint
type subscriptionsService service

// Get returns the subscription with the given ID.
// https://docs.lemonsqueezy.com/api/subscriptions#retrieve-a-subscription
func (service *subscriptionsService) Get(ctx context.Context, subscriptionID string) (*ApiResponse[Subscription], *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodGet, "/v1/subscriptions/"+subscriptionID, nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	status := new(ApiResponse[Subscription])
	if err = json.Unmarshal(*response.Body, status); err != nil {
		return nil, response, err
	}

	return status, response, nil
}

// Cancel an active subscription the given ID.
// https://docs.lemonsqueezy.com/api/subscriptions#retrieve-a-subscription
func (service *subscriptionsService) Cancel(ctx context.Context, subscriptionID string) (*ApiResponse[Subscription], *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodDelete, "/v1/subscriptions/"+subscriptionID, nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	status := new(ApiResponse[Subscription])
	if err = json.Unmarshal(*response.Body, status); err != nil {
		return nil, response, err
	}

	return status, response, nil
}
