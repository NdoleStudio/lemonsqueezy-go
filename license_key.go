package lemonsqueezy

import "time"

// LicenseKeyAttributes contains information about a license key which can be used to externally verify that customer has access to a product.
type LicenseKeyAttributes struct {
	StoreID         int         `json:"store_id"`
	CustomerID      int         `json:"customer_id"`
	OrderID         int         `json:"order_id"`
	OrderItemID     int         `json:"order_item_id"`
	ProductID       int         `json:"product_id"`
	UserName        string      `json:"user_name"`
	UserEmail       string      `json:"user_email"`
	Key             string      `json:"key"`
	KeyShort        string      `json:"key_short"`
	ActivationLimit int         `json:"activation_limit"`
	InstancesCount  int         `json:"instances_count"`
	Disabled        bool        `json:"disabled"`
	Status          string      `json:"status"`
	StatusFormatted string      `json:"status_formatted"`
	ExpiresAt       interface{} `json:"expires_at"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

// ApiResponseRelationshipsLicenseKey relationships of a license key
type ApiResponseRelationshipsLicenseKey struct {
	Store               ApiResponseLinks `json:"store"`
	Customer            ApiResponseLinks `json:"customer"`
	Order               ApiResponseLinks `json:"order"`
	OrderItem           ApiResponseLinks `json:"order-item"`
	Product             ApiResponseLinks `json:"product"`
	LicenseKeyInstances ApiResponseLinks `json:"license-key-instances"`
}

// LicenseKeyApiResponse is the api response for one subscription invoice
type LicenseKeyApiResponse = ApiResponse[LicenseKeyAttributes, ApiResponseRelationshipsLicenseKey]

// LicenseKeysApiResponse is the api response for a list of subscription invoices.
type LicenseKeysApiResponse = ApiResponseList[LicenseKeyAttributes, ApiResponseRelationshipsLicenseKey]
