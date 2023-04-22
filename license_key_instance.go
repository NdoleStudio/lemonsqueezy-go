package client

import "time"

// LicenseKeyInstanceAttributes represents a single instance (or activation) of a license key that has been issued to a customer.
type LicenseKeyInstanceAttributes struct {
	LicenseKeyID int       `json:"license_key_id"`
	Identifier   string    `json:"identifier"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// ApiResponseRelationshipsLicenseKeyInstance relationships of a license key
type ApiResponseRelationshipsLicenseKeyInstance struct {
	LicenseKey ApiResponseLinks `json:"license-key"`
}

// LicenseKeyInstanceApiResponse is the api response for one subscription invoice
type LicenseKeyInstanceApiResponse = ApiResponse[LicenseKeyInstanceAttributes, ApiResponseRelationshipsLicenseKeyInstance]

// LicenseKeyInstancesApiResponse is the api response for a list of subscription invoices.
type LicenseKeyInstancesApiResponse = ApiResponseList[LicenseKeyInstanceAttributes, ApiResponseRelationshipsLicenseKeyInstance]
