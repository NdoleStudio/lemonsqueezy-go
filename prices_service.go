package lemonsqueezy

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// PricesService is the API client for the `/v1/prices` endpoint
type PricesService service

// Get returns the price with the given ID.
//
// https://docs.lemonsqueezy.com/api/prices#retrieve-a-price
func (service *PricesService) Get(ctx context.Context, priceID int) (*PriceAPIResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/prices/"+strconv.Itoa(priceID))
	if err != nil {
		return nil, response, err
	}

	price := new(PriceAPIResponse)
	if err = json.Unmarshal(*response.Body, price); err != nil {
		return nil, response, err
	}

	return price, response, nil
}

// List returns a paginated list of prices.
//
// https://docs.lemonsqueezy.com/api/prices#list-all-prices
func (service *PricesService) List(ctx context.Context) (*PricesAPIResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/prices")
	if err != nil {
		return nil, response, err
	}

	prices := new(PricesAPIResponse)
	if err = json.Unmarshal(*response.Body, prices); err != nil {
		return nil, response, err
	}

	return prices, response, nil
}
