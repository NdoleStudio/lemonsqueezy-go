package lemonsqueezy

import (
	"context"
	"encoding/json"
	"net/http"
)

// VariantsService is the API client for the `/v1/variants` endpoint
type VariantsService service

// Get returns the variant with the given ID.
//
// https://docs.lemonsqueezy.com/api/variants#retrieve-a-variant
func (service *VariantsService) Get(ctx context.Context, variantID string) (*VariantApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/variants/"+variantID)
	if err != nil {
		return nil, response, err
	}

	variant := new(VariantApiResponse)
	if err = json.Unmarshal(*response.Body, variant); err != nil {
		return nil, response, err
	}

	return variant, response, nil
}

// List returns a paginated list of variants.
//
// https://docs.lemonsqueezy.com/api/variants#list-all-variants
func (service *VariantsService) List(ctx context.Context) (*VariantsApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/variants")
	if err != nil {
		return nil, response, err
	}

	variants := new(VariantsApiResponse)
	if err = json.Unmarshal(*response.Body, variants); err != nil {
		return nil, response, err
	}

	return variants, response, nil
}
