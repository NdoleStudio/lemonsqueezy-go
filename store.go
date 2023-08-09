package lemonsqueezy

import "time"

// StoreAttributes Everything in Lemon Squeezy belongs to a store.
type StoreAttributes struct {
	Name             string    `json:"name"`
	Slug             string    `json:"slug"`
	Domain           string    `json:"domain"`
	Url              string    `json:"url"`
	AvatarUrl        string    `json:"avatar_url"`
	Plan             string    `json:"plan"`
	Country          string    `json:"country"`
	CountryNiceName  string    `json:"country_nicename"`
	Currency         string    `json:"currency"`
	TotalSales       int       `json:"total_sales"`
	TotalRevenue     int       `json:"total_revenue"`
	ThirtyDaySales   int       `json:"thirty_day_sales"`
	ThirtyDayRevenue int       `json:"thirty_day_revenue"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// ApiResponseRelationshipsStore relationships of a store object
type ApiResponseRelationshipsStore struct {
	Subscriptions ApiResponseLinks `json:"subscriptions"`
	Orders        ApiResponseLinks `json:"orders"`
	Products      ApiResponseLinks `json:"products"`
	LicenseKeys   ApiResponseLinks `json:"license-keys"`
	Discounts     ApiResponseLinks `json:"discounts"`
}

// StoreApiResponse represents a store api response
type StoreApiResponse = ApiResponse[StoreAttributes, ApiResponseRelationshipsStore]

// StoresApiResponse represents a list of store api responses.
type StoresApiResponse = ApiResponseList[StoreAttributes, ApiResponseRelationshipsStore]
