package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// CustomersService is the API client for the `/v1/customers` endpoint
type CustomersService service

// Get returns the customer with the given ID.
//
// https://docs.lemonsqueezy.com/api/customers#retrieve-a-customer
func (service *CustomersService) Get(ctx context.Context, customerID string) (*CustomerApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/customers/"+customerID)
	if err != nil {
		return nil, response, err
	}

	customer := new(CustomerApiResponse)
	if err = json.Unmarshal(*response.Body, customer); err != nil {
		return nil, response, err
	}

	return customer, response, nil
}

// List returns a paginated list of customers.
//
// https://docs.lemonsqueezy.com/api/customers#list-all-customers
func (service *CustomersService) List(ctx context.Context) (*CustomersApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/customers")
	if err != nil {
		return nil, response, err
	}

	customers := new(CustomersApiResponse)
	if err = json.Unmarshal(*response.Body, customers); err != nil {
		return nil, response, err
	}

	return customers, response, nil
}
