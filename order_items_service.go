package lemonsqueezy

import (
	"context"
	"encoding/json"
	"net/http"
)

// OrderItemsService is the API client for the `/v1/order-items` endpoint
type OrderItemsService service

// Get returns the order-item with the given ID.
//
// https://docs.lemonsqueezy.com/api/order-items#retrieve-an-order-item
func (service *OrderItemsService) Get(ctx context.Context, orderItemID string) (*OrderItemApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/order-items/"+orderItemID)
	if err != nil {
		return nil, response, err
	}

	orderItem := new(OrderItemApiResponse)
	if err = json.Unmarshal(*response.Body, orderItem); err != nil {
		return nil, response, err
	}

	return orderItem, response, nil
}

// List returns a paginated list of order-items.
//
// https://docs.lemonsqueezy.com/api/order-items#list-all-order-items
func (service *OrderItemsService) List(ctx context.Context) (*OrderItemsApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/order-items")
	if err != nil {
		return nil, response, err
	}

	orderItems := new(OrderItemsApiResponse)
	if err = json.Unmarshal(*response.Body, orderItems); err != nil {
		return nil, response, err
	}

	return orderItems, response, nil
}
