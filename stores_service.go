package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// StoresService is the API client for the `/v1/stores` endpoint
type StoresService service

// Get returns the store with the given ID.
// https://docs.lemonsqueezy.com/api/stores#retrieve-a-store
func (service *StoresService) Get(ctx context.Context, storeID string) (*StoreApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/stores/"+storeID)
	if err != nil {
		return nil, response, err
	}

	store := new(StoreApiResponse)
	if err = json.Unmarshal(*response.Body, store); err != nil {
		return nil, response, err
	}

	return store, response, nil
}

// List returns a paginated list of stores.
// https://docs.lemonsqueezy.com/api/stores#list-all-stores
func (service *StoresService) List(ctx context.Context) (*StoresApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/stores/")
	if err != nil {
		return nil, response, err
	}

	stores := new(StoresApiResponse)
	if err = json.Unmarshal(*response.Body, stores); err != nil {
		return nil, response, err
	}

	return stores, response, nil
}
