package lemonsqueezy

import (
	"context"
	"encoding/json"
	"net/http"
)

// LicenseKeysService is the API client for the `/v1/license-keys` endpoint
type LicenseKeysService service

// Get retrieves the license key with the given ID.
//
// https://docs.lemonsqueezy.com/api/license-keys#retrieve-a-license-key
func (service *LicenseKeysService) Get(ctx context.Context, licenseKeyID string) (*LicenseKeyApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/license-keys/"+licenseKeyID)
	if err != nil {
		return nil, response, err
	}

	licenseKey := new(LicenseKeyApiResponse)
	if err = json.Unmarshal(*response.Body, licenseKey); err != nil {
		return nil, response, err
	}

	return licenseKey, response, nil
}

// List a paginated list of discount redemptions.
//
// https://docs.lemonsqueezy.com/api/license-keys#list-all-license-keys
func (service *LicenseKeysService) List(ctx context.Context) (*LicenseKeysApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/license-keys")
	if err != nil {
		return nil, response, err
	}

	licenseKeys := new(LicenseKeysApiResponse)
	if err = json.Unmarshal(*response.Body, licenseKeys); err != nil {
		return nil, response, err
	}

	return licenseKeys, response, nil
}
