package lemonsqueezy

import (
	"context"
	"encoding/json"
	"net/http"
)

// OrdersService is the API client for the `/v1/orders` endpoint
type OrdersService service

// Get returns the order with the given ID.
//
// https://docs.lemonsqueezy.com/api/orders#retrieve-an-order
func (service *OrdersService) Get(ctx context.Context, orderID string) (*OrderApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/orders/"+orderID)
	if err != nil {
		return nil, response, err
	}

	order := new(OrderApiResponse)
	if err = json.Unmarshal(*response.Body, order); err != nil {
		return nil, response, err
	}

	return order, response, nil
}

// List returns a paginated list of orders.
//
// https://docs.lemonsqueezy.com/api/orders#list-all-orders
func (service *OrdersService) List(ctx context.Context) (*OrdersApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/orders")
	if err != nil {
		return nil, response, err
	}

	orders := new(OrdersApiResponse)
	if err = json.Unmarshal(*response.Body, orders); err != nil {
		return nil, response, err
	}

	return orders, response, nil
}
