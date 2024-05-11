package lemonsqueezy

import "time"

// PriceAttributes represents a price added to a variant.
type PriceAttributes struct {
	VariantID               int                    `json:"variant_id"`
	Category                string                 `json:"category"`
	Scheme                  string                 `json:"scheme"`
	UsageAggregation        any                    `json:"usage_aggregation"`
	UnitPrice               int                    `json:"unit_price"`
	UnitPriceDecimal        any                    `json:"unit_price_decimal"`
	SetupFeeEnabled         bool                   `json:"setup_fee_enabled"`
	SetupFee                any                    `json:"setup_fee"`
	PackageSize             int                    `json:"package_size"`
	Tiers                   []PriceAttributesTiers `json:"tiers"`
	RenewalIntervalUnit     string                 `json:"renewal_interval_unit"`
	RenewalIntervalQuantity int                    `json:"renewal_interval_quantity"`
	TrialIntervalUnit       string                 `json:"trial_interval_unit"`
	TrialIntervalQuantity   int                    `json:"trial_interval_quantity"`
	MinPrice                any                    `json:"min_price"`
	SuggestedPrice          any                    `json:"suggested_price"`
	TaxCode                 string                 `json:"tax_code"`
	CreatedAt               time.Time              `json:"created_at"`
	UpdatedAt               time.Time              `json:"updated_at"`
}

// PriceAttributesTiers is list of pricing tier objects when using volume and graduated pricing.
type PriceAttributesTiers struct {
	LastUnit         any     `json:"last_unit"`
	UnitPrice        int     `json:"unit_price"`
	UnitPriceDecimal *string `json:"unit_price_decimal"`
	FixedFee         int     `json:"fixed_fee"`
}

// APIResponseRelationshipsPrice relationships of a variant
type APIResponseRelationshipsPrice struct {
	Variant ApiResponseLinks `json:"variant"`
}

// PriceAPIResponse is the api response for one variant
type PriceAPIResponse = ApiResponse[PriceAttributes, APIResponseRelationshipsPrice]

// PricesAPIResponse is the api response for a list of variants.
type PricesAPIResponse = ApiResponseList[PriceAttributes, APIResponseRelationshipsPrice]
