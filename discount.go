package lemonsqueezy

import "time"

// DiscountAttributes contains information about a percentage or amount discount that can be applied to an order at checkout via a code.
type DiscountAttributes struct {
	StoreID              int        `json:"store_id"`
	Name                 string     `json:"name"`
	Code                 string     `json:"code"`
	Amount               int        `json:"amount"`
	AmountType           string     `json:"amount_type"`
	IsLimitedToProducts  bool       `json:"is_limited_to_products"`
	IsLimitedRedemptions bool       `json:"is_limited_redemptions"`
	MaxRedemptions       int        `json:"max_redemptions"`
	StartsAt             *time.Time `json:"starts_at"`
	ExpiresAt            *time.Time `json:"expires_at"`
	Duration             string     `json:"duration"`
	DurationInMonths     int        `json:"duration_in_months"`
	Status               string     `json:"status"`
	StatusFormatted      string     `json:"status_formatted"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}

// ApiResponseRelationshipsDiscount relationships of a discount
type ApiResponseRelationshipsDiscount struct {
	Store               ApiResponseLinks `json:"store"`
	DiscountRedemptions ApiResponseLinks `json:"discount-redemptions"`
	Variants            ApiResponseLinks `json:"variants"`
}

// DiscountApiResponse is the api response for one discount
type DiscountApiResponse = ApiResponse[DiscountAttributes, ApiResponseRelationshipsDiscount]

// DiscountsApiResponse is the api response for a list of discounts.
type DiscountsApiResponse = ApiResponseList[DiscountAttributes, ApiResponseRelationshipsDiscount]

// DiscountCreateParams are parameters for creating a discount
type DiscountCreateParams struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	Amount     int    `json:"amount"`
	AmountType string `json:"amountType"`
	StoreID    int    `json:"storeID"`
}
