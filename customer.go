package client

import "time"

// CustomerAttributes are the attributes of a lemonsqueezy customer
type CustomerAttributes struct {
	StoreID                       int       `json:"store_id"`
	Name                          string    `json:"name"`
	Email                         string    `json:"email"`
	Status                        string    `json:"status"`
	City                          *string   `json:"city"`
	Region                        *string   `json:"region"`
	Country                       string    `json:"country"`
	TotalRevenueCurrency          int       `json:"total_revenue_currency"`
	Mrr                           int       `json:"mrr"`
	StatusFormatted               string    `json:"status_formatted"`
	CountryFormatted              string    `json:"country_formatted"`
	TotalRevenueCurrencyFormatted string    `json:"total_revenue_currency_formatted"`
	MrrFormatted                  string    `json:"mrr_formatted"`
	CreatedAt                     time.Time `json:"created_at"`
	UpdatedAt                     time.Time `json:"updated_at"`
	TestMode                      bool      `json:"test_mode"`
}

// ApiResponseRelationshipsCustomer relationships of a customer
type ApiResponseRelationshipsCustomer struct {
	Store         ApiResponseLinks `json:"store"`
	Subscriptions ApiResponseLinks `json:"subscriptions"`
	Orders        ApiResponseLinks `json:"orders"`
	LicenseKeys   ApiResponseLinks `json:"license-keys"`
}

// CustomerApiResponse represents a customer api response
type CustomerApiResponse = ApiResponse[CustomerAttributes, ApiResponseRelationshipsCustomer]

// CustomersApiResponse represents a list of customers api responses.
type CustomersApiResponse = ApiResponseList[CustomerAttributes, ApiResponseRelationshipsCustomer]
