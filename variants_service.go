package lemonsqueezy

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// VariantsService is the API client for the `/v1/variants` endpoint
type VariantsService service

// Get returns the variant with the given ID.
//
// https://docs.lemonsqueezy.com/api/variants#retrieve-a-variant
func (service *VariantsService) Get(ctx context.Context, variantID int) (*VariantAPIResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/variants/"+strconv.Itoa(variantID))
	if err != nil {
		return nil, response, err
	}

	variant := new(VariantAPIResponse)
	if err = json.Unmarshal(*response.Body, variant); err != nil {
		return nil, response, err
	}

	return variant, response, nil
}

// List returns a paginated list of variants.
//
// https://docs.lemonsqueezy.com/api/variants#list-all-variants
func (service *VariantsService) List(ctx context.Context) (*VariantsAPIResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/variants")
	if err != nil {
		return nil, response, err
	}

	variants := new(VariantsAPIResponse)
	if err = json.Unmarshal(*response.Body, variants); err != nil {
		return nil, response, err
	}

	return variants, response, nil
}
