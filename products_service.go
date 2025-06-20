package lemonsqueezy

import (
	"context"
	"encoding/json"
	"net/http"
)

// ProductsService is the API client for the `/v1/products` endpoint
type ProductsService service

// Get returns the product with the given ID.
//
// https://docs.lemonsqueezy.com/api/products/retrieve-product
func (service *ProductsService) Get(ctx context.Context, productID string) (*ProductApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/products/"+productID)
	if err != nil {
		return nil, response, err
	}

	product := new(ProductApiResponse)
	if err = json.Unmarshal(*response.Body, product); err != nil {
		return nil, response, err
	}

	return product, response, nil
}

// List returns a paginated list of products.
//
// https://docs.lemonsqueezy.com/api/products/list-all-products
func (service *ProductsService) List(ctx context.Context, filtersKV ...string) (*ProductsApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, pathWithFilters("/v1/products", filtersKV...))
	if err != nil {
		return nil, response, err
	}

	products := new(ProductsApiResponse)
	if err = json.Unmarshal(*response.Body, products); err != nil {
		return nil, response, err
	}

	return products, response, nil
}

// ListWithVariants returns a paginated list of products with variants.
//
// https://docs.lemonsqueezy.com/api/products/list-all-products
func (service *ProductsService) ListWithVariants(ctx context.Context, filtersKV ...string) (*ProductsWithVariantsApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, pathWithFilters("/v1/products?include=variants", filtersKV...))
	if err != nil {
		return nil, response, err
	}

	products := new(ProductsWithVariantsApiResponse)
	if err = json.Unmarshal(*response.Body, products); err != nil {
		return nil, response, err
	}

	return products, response, nil
}
