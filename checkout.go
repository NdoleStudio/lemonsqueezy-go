package lemonsqueezy

import "time"

// CheckoutAttributes contains information about a percentage or amount discount that can be applied to an order at checkout via a code.
type CheckoutAttributes struct {
	StoreID         int                    `json:"store_id"`
	VariantID       int                    `json:"variant_id"`
	CustomPrice     interface{}            `json:"custom_price"`
	ProductOptions  CheckoutProductOptions `json:"product_options"`
	CheckoutOptions CheckoutOptions        `json:"checkout_options"`
	CheckoutData    CheckoutData           `json:"checkout_data"`
	ExpiresAt       *time.Time             `json:"expires_at"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
	TestMode        bool                   `json:"test_mode"`
	URL             string                 `json:"url"`
}

// CheckoutProductOptions are options for a checkout product
type CheckoutProductOptions struct {
	Name                string `json:"name"`
	Description         string `json:"description"`
	Media               []any  `json:"media"`
	RedirectURL         string `json:"redirect_url"`
	ReceiptButtonText   string `json:"receipt_button_text"`
	ReceiptLinkURL      string `json:"receipt_link_url"`
	ReceiptThankYouNote string `json:"receipt_thank_you_note"`
	EnabledVariants     []int  `json:"enabled_variants"`
}

// CheckoutOptions are options for a checkout
type CheckoutOptions struct {
	Embed               bool   `json:"embed"`
	Media               bool   `json:"media"`
	Logo                bool   `json:"logo"`
	Desc                bool   `json:"desc"`
	Discount            bool   `json:"discount"`
	Dark                bool   `json:"dark"`
	SubscriptionPreview bool   `json:"subscription_preview"`
	ButtonColor         string `json:"button_color"`
}

// BillingAddress contains the checkout billing address
type BillingAddress struct {
	Country string `json:"country"`
	Zip     string `json:"zip"`
}

// CheckoutData contains information about a checkout
type CheckoutData struct {
	Email          string           `json:"email"`
	Name           string           `json:"name"`
	BillingAddress []BillingAddress `json:"billing_address"`
	TaxNumber      string           `json:"tax_number"`
	DiscountCode   string           `json:"discount_code"`
	Custom         map[string]any   `json:"custom"`
}

// CheckoutPreview contains information about a percentage or amount discount that can be applied to an order at checkout via a code.
type CheckoutPreview struct {
	Currency               string `json:"currency"`
	CurrencyRate           int    `json:"currency_rate"`
	Subtotal               int    `json:"subtotal"`
	DiscountTotal          int    `json:"discount_total"`
	Tax                    int    `json:"tax"`
	Total                  int    `json:"total"`
	SubtotalUsd            int    `json:"subtotal_usd"`
	DiscountTotalUsd       int    `json:"discount_total_usd"`
	TaxUsd                 int    `json:"tax_usd"`
	TotalUsd               int    `json:"total_usd"`
	SubtotalFormatted      string `json:"subtotal_formatted"`
	DiscountTotalFormatted string `json:"discount_total_formatted"`
	TaxFormatted           string `json:"tax_formatted"`
	TotalFormatted         string `json:"total_formatted"`
}

// ApiResponseRelationshipsCheckout relationships of a checkout
type ApiResponseRelationshipsCheckout struct {
	Store   ApiResponseLinks `json:"store"`
	Variant ApiResponseLinks `json:"variant"`
}

// CheckoutApiResponse is the api response for one checkout
type CheckoutApiResponse = ApiResponse[CheckoutAttributes, ApiResponseRelationshipsDiscount]

// CheckoutsApiResponse is the api response for a list of checkout.
type CheckoutsApiResponse = ApiResponseList[CheckoutAttributes, ApiResponseRelationshipsDiscount]

// Quantities represents variant quantities when creating checkout
type Quantities struct {
	VariantId int `json:"variant_id"`
	Quantity  int `json:"quantity"`
}

// CheckoutCreateData represents the data options for creating a checkout.
type CheckoutCreateData struct {
	Email                 string         `json:"email,omitempty"`
	Name                  string         `json:"name,omitempty"`
	BillingAddressCountry string         `json:"billing_address.country,omitempty"`
	BillingAddressZip     string         `json:"billing_address.zip,omitempty"`
	TaxNumber             string         `json:"tax_number,omitempty"`
	DiscountCode          string         `json:"discount_code,omitempty"`
	Custom                map[string]any `json:"custom,omitempty"`
	VariantQuantities     []Quantities   `json:"variant_quantities,omitempty"`
}

// CheckoutCreateOptions represents the checkout options for creating a checkout.
//
// Note: We use pointers for the boolean fields as otherwise, setting them to "false" would omit them, which would
// break some of the boolean checks in the API. See: https://docs.lemonsqueezy.com/api/checkouts#create-a-checkout
type CheckoutCreateOptions struct {
	Embed               *bool  `json:"embed,omitempty"`
	Media               *bool  `json:"media,omitempty"`
	Logo                *bool  `json:"logo,omitempty"`
	Desc                *bool  `json:"desc,omitempty"`
	Discount            *bool  `json:"discount,omitempty"`
	Dark                *bool  `json:"dark,omitempty"`
	SubscriptionPreview *bool  `json:"subscription_preview,omitempty"`
	ButtonColor         string `json:"button_color,omitempty"`
}

// CheckoutCreateProductOptions represents product options for creating a checkout.
type CheckoutCreateProductOptions struct {
	Name                string   `json:"name,omitempty"`
	Description         string   `json:"description,omitempty"`
	Media               []string `json:"media,omitempty"`
	RedirectUrl         string   `json:"redirect_url,omitempty"`
	ReceiptButtonText   string   `json:"receipt_button_text,omitempty"`
	ReceiptLinkUrl      string   `json:"receipt_link_url,omitempty"`
	ReceiptThankYouNote string   `json:"receipt_thank_you_note,omitempty"`
	EnabledVariants     []int    `json:"enabled_variants,omitempty"`
}

// CheckoutCreateAttributes represents individual parameters for creating a checkout.
type CheckoutCreateAttributes struct {
	CustomPrice     *int                         `json:"custom_price,omitempty"`
	ProductOptions  CheckoutCreateProductOptions `json:"product_options,omitempty"`
	CheckoutOptions CheckoutCreateOptions        `json:"checkout_options,omitempty"`
	CheckoutData    CheckoutCreateData           `json:"checkout_data,omitempty"`
	Preview         *bool                        `json:"preview,omitempty"`
	TestMode        *bool                        `json:"test_mode,omitempty"`
	ExpiresAt       *time.Time                   `json:"expires_at,omitempty"`
}
