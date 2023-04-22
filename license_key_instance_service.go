package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// LicenseKeyInstancesService is the API client for the `/v1/license-key-instances` endpoint
type LicenseKeyInstancesService service

// Get retrieves the license key instance with the given ID.
//
// https://docs.lemonsqueezy.com/api/license-key-instances#retrieve-a-license-key-instance
func (service *LicenseKeyInstancesService) Get(ctx context.Context, licenseKeyID string) (*LicenseKeyInstanceApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/license-key-instances/"+licenseKeyID)
	if err != nil {
		return nil, response, err
	}

	licenseKeyInstance := new(LicenseKeyInstanceApiResponse)
	if err = json.Unmarshal(*response.Body, licenseKeyInstance); err != nil {
		return nil, response, err
	}

	return licenseKeyInstance, response, nil
}

// List a paginated list of license key instances.
//
// https://docs.lemonsqueezy.com/api/license-key-instances#list-all-license-key-instances
func (service *LicenseKeyInstancesService) List(ctx context.Context) (*LicenseKeyInstancesApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/license-key-instances")
	if err != nil {
		return nil, response, err
	}

	licenseKeyInstances := new(LicenseKeyInstancesApiResponse)
	if err = json.Unmarshal(*response.Body, licenseKeyInstances); err != nil {
		return nil, response, err
	}

	return licenseKeyInstances, response, nil
}
