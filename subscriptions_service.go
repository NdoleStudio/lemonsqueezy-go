package lemonsqueezy

import (
	"context"
	"encoding/json"
	"net/http"
)

// SubscriptionsService is the API client for the `/subscriptions` endpoint
type SubscriptionsService service

// Update an existing subscription to specific parameter
//
// https://docs.lemonsqueezy.com/api/subscriptions#update-a-subscription
func (service *SubscriptionsService) Update(ctx context.Context, params *SubscriptionUpdateParams) (*Resource[Subscription], *Response, error) {
	response, err := service.client.do(ctx, http.MethodPatch, "/v1/subscriptions/"+params.ID, params)
	if err != nil {
		return nil, response, err
	}

	subscription := new(Resource[Subscription])
	if err = json.Unmarshal(*response.Body, subscription); err != nil {
		return nil, response, err
	}

	return subscription, response, nil
}

// List returns a paginated list of subscriptions ordered by created_at (descending)
//
// https://docs.lemonsqueezy.com/api/subscriptions#list-all-subscriptions
func (service *SubscriptionsService) List(ctx context.Context) (*ApiResponseSubscriptionList, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/subscriptions")
	if err != nil {
		return nil, response, err
	}

	subscriptions := new(ApiResponseSubscriptionList)
	if err = json.Unmarshal(*response.Body, subscriptions); err != nil {
		return nil, response, err
	}

	return subscriptions, response, nil
}

// Get returns the subscription with the given ID.
//
// https://docs.lemonsqueezy.com/api/subscriptions#retrieve-a-subscription
func (service *SubscriptionsService) Get(ctx context.Context, subscriptionID string) (*ApiResponseSubscription, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/subscriptions/"+subscriptionID)
	if err != nil {
		return nil, response, err
	}

	subscription := new(ApiResponseSubscription)
	if err = json.Unmarshal(*response.Body, subscription); err != nil {
		return nil, response, err
	}

	return subscription, response, nil
}

// Cancel an active subscription the given ID.
//
// https://docs.lemonsqueezy.com/api/subscriptions#retrieve-a-subscription
func (service *SubscriptionsService) Cancel(ctx context.Context, subscriptionID string) (*ApiResponseSubscription, *Response, error) {
	response, err := service.client.do(ctx, http.MethodDelete, "/v1/subscriptions/"+subscriptionID)
	if err != nil {
		return nil, response, err
	}

	subscription := new(ApiResponseSubscription)
	if err = json.Unmarshal(*response.Body, subscription); err != nil {
		return nil, response, err
	}

	return subscription, response, nil
}
