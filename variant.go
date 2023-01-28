package client

import "time"

// VariantAttributes represents a different option that is presented to the customer at checkout.
type VariantAttributes struct {
	ProductID                int       `json:"product_id"`
	Name                     string    `json:"name"`
	Slug                     string    `json:"slug"`
	Description              string    `json:"description"`
	Price                    int       `json:"price"`
	IsSubscription           bool      `json:"is_subscription"`
	Interval                 *string   `json:"interval"`
	IntervalCount            *int      `json:"interval_count"`
	HasFreeTrial             bool      `json:"has_free_trial"`
	TrialInterval            string    `json:"trial_interval"`
	TrialIntervalCount       int       `json:"trial_interval_count"`
	PayWhatYouWant           bool      `json:"pay_what_you_want"`
	MinPrice                 int       `json:"min_price"`
	SuggestedPrice           int       `json:"suggested_price"`
	HasLicenseKeys           bool      `json:"has_license_keys"`
	LicenseActivationLimit   int       `json:"license_activation_limit"`
	IsLicenseLimitUnlimited  bool      `json:"is_license_limit_unlimited"`
	LicenseLengthValue       int       `json:"license_length_value"`
	LicenseLengthUnit        string    `json:"license_length_unit"`
	IsLicenseLengthUnlimited bool      `json:"is_license_length_unlimited"`
	Sort                     int       `json:"sort"`
	Status                   string    `json:"status"`
	StatusFormatted          string    `json:"status_formatted"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

// ApiResponseRelationshipsVariant relationships of a variant
type ApiResponseRelationshipsVariant struct {
	Product ApiResponseLinks `json:"product"`
}

// VariantApiResponse is the api response for one variant
type VariantApiResponse = ApiResponse[VariantAttributes, ApiResponseRelationshipsVariant]

// VariantsApiResponse is the api response for a list of variants.
type VariantsApiResponse = ApiResponseList[VariantAttributes, ApiResponseRelationshipsVariant]
