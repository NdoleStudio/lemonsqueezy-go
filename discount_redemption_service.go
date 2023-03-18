package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// DiscountRedemptionsService is the API client for the `/v1/discount-redemptions` endpoint
type DiscountRedemptionsService service

// Get returns the discount redemption with the given ID.
//
// https://docs.lemonsqueezy.com/api/discount-redemptions#retrieve-a-discount-redemption
func (service *DiscountRedemptionsService) Get(ctx context.Context, invoiceID string) (*DiscountRedemptionApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/discount-redemptions/"+invoiceID)
	if err != nil {
		return nil, response, err
	}

	redemption := new(DiscountRedemptionApiResponse)
	if err = json.Unmarshal(*response.Body, redemption); err != nil {
		return nil, response, err
	}

	return redemption, response, nil
}

// List a paginated list of discount redemptions.
//
// https://docs.lemonsqueezy.com/api/discount-redemptions#list-all-discount-redemptions
func (service *DiscountRedemptionsService) List(ctx context.Context) (*DiscountRedemptionsApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/discount-redemptions")
	if err != nil {
		return nil, response, err
	}

	redemptions := new(DiscountRedemptionsApiResponse)
	if err = json.Unmarshal(*response.Body, redemptions); err != nil {
		return nil, response, err
	}

	return redemptions, response, nil
}
