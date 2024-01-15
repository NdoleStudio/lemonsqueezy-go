package lemonsqueezy

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// CheckoutsService is the API client for the `/v1/checkouts` endpoint
type CheckoutsService service

// Create a custom checkout.
//
// https://docs.lemonsqueezy.com/api/checkouts#create-a-checkout
func (service *CheckoutsService) Create(ctx context.Context, variantId int, storeId int, attributes *CheckoutCreateAttributes) (*CheckoutApiResponse, *Response, error) {
	payload := map[string]any{
		"data": map[string]any{
			"type":       "checkouts",
			"attributes": attributes,
			"relationships": map[string]any{
				"store": map[string]any{
					"data": map[string]any{
						"id":   strconv.Itoa(storeId),
						"type": "stores",
					},
				},
				"variant": map[string]any{
					"data": map[string]any{
						"id":   strconv.Itoa(variantId),
						"type": "variants",
					},
				},
			},
		},
	}

	response, err := service.client.do(ctx, http.MethodPost, "/v1/checkouts", payload)
	if err != nil {
		return nil, response, err
	}

	checkout := new(CheckoutApiResponse)
	if err = json.Unmarshal(*response.Body, checkout); err != nil {
		return nil, response, err
	}

	return checkout, response, nil
}

// Get the checkout with the given ID.
//
// https://docs.lemonsqueezy.com/api/checkouts#retrieve-a-checkout
func (service *CheckoutsService) Get(ctx context.Context, checkoutID string) (*CheckoutApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/checkouts/"+checkoutID)
	if err != nil {
		return nil, response, err
	}

	checkout := new(CheckoutApiResponse)
	if err = json.Unmarshal(*response.Body, checkout); err != nil {
		return nil, response, err
	}

	return checkout, response, nil
}

// List returns a paginated list of checkouts.
//
// https://docs.lemonsqueezy.com/api/checkouts#list-all-checkouts
func (service *CheckoutsService) List(ctx context.Context) (*CheckoutsApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/checkouts")
	if err != nil {
		return nil, response, err
	}

	checkouts := new(CheckoutsApiResponse)
	if err = json.Unmarshal(*response.Body, checkouts); err != nil {
		return nil, response, err
	}

	return checkouts, response, nil
}
