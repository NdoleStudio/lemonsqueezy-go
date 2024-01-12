package lemonsqueezy

import (
	"context"
	"encoding/json"
	"net/http"
)

// LicensesService is the API client for the `/v1/licenses` endpoint
type LicensesService service

// Activate a license key.
//
// https://docs.lemonsqueezy.com/help/licensing/license-api
func (service *LicensesService) Activate(ctx context.Context, licenseKey, instanceName string) (*LicenseActivateApiResponse, *Response, error) {
	payload := map[string]any{
		"license_key":   licenseKey,
		"instance_name": instanceName,
	}

	response, err := service.client.do(ctx, http.MethodPost, "/v1/licenses/activate", payload)
	if err != nil {
		return nil, response, err
	}

	activation := new(LicenseActivateApiResponse)
	if err = json.Unmarshal(*response.Body, activation); err != nil {
		return nil, response, err
	}

	return activation, response, nil
}

// Validate a license key.
//
// https://docs.lemonsqueezy.com/help/licensing/license-api
func (service *LicensesService) Validate(ctx context.Context, licenseKey, instanceID string) (*LicenseValidateApiResponse, *Response, error) {
	payload := map[string]any{
		"license_key": licenseKey,
		"instance_id": instanceID,
	}

	response, err := service.client.do(ctx, http.MethodPost, "/v1/licenses/validate", payload)
	if err != nil {
		return nil, response, err
	}

	validation := new(LicenseValidateApiResponse)
	if err = json.Unmarshal(*response.Body, validation); err != nil {
		return nil, response, err
	}

	return validation, response, nil
}

// Deactivate a license key.
//
// https://docs.lemonsqueezy.com/help/licensing/license-api
func (service *LicensesService) Deactivate(ctx context.Context, licenseKey, instanceID string) (*LicenseDeactivateApiResponse, *Response, error) {
	payload := map[string]any{
		"license_key": licenseKey,
		"instance_id": instanceID,
	}

	response, err := service.client.do(ctx, http.MethodPost, "/v1/licenses/deactivate", payload)
	if err != nil {
		return nil, response, err
	}

	deactivation := new(LicenseDeactivateApiResponse)
	if err = json.Unmarshal(*response.Body, deactivation); err != nil {
		return nil, response, err
	}

	return deactivation, response, nil
}
