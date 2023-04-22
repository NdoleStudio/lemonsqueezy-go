package client

import "time"

// DiscountRedemptionAttributes is a record of a discount being applied to an order.
type DiscountRedemptionAttributes struct {
	DiscountID         int       `json:"discount_id"`
	OrderID            int       `json:"order_id"`
	DiscountName       string    `json:"discount_name"`
	DiscountCode       string    `json:"discount_code"`
	DiscountAmount     int       `json:"discount_amount"`
	DiscountAmountType string    `json:"discount_amount_type"`
	Amount             int       `json:"amount"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// ApiResponseRelationshipsDiscountRedemption relationships of a discount redemption
type ApiResponseRelationshipsDiscountRedemption struct {
	Discount ApiResponseLinks `json:"discount"`
	Order    ApiResponseLinks `json:"order"`
}

// DiscountRedemptionApiResponse is the api response for one discount redemption
type DiscountRedemptionApiResponse = ApiResponse[DiscountRedemptionAttributes, ApiResponseRelationshipsDiscountRedemption]

// DiscountRedemptionsApiResponse is the api response for a list of discount redemptions.
type DiscountRedemptionsApiResponse = ApiResponseList[DiscountRedemptionAttributes, ApiResponseRelationshipsDiscountRedemption]
