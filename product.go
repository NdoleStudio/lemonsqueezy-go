package lemonsqueezy

import "time"

// ProductAttributes are the attributes of a lemonsqueezy product
type ProductAttributes struct {
	StoreID            int       `json:"store_id"`
	Name               string    `json:"name"`
	Slug               string    `json:"slug"`
	Description        string    `json:"description"`
	Status             string    `json:"status"`
	StatusFormatted    string    `json:"status_formatted"`
	ThumbURL           string    `json:"thumb_url"`
	LargeThumbURL      string    `json:"large_thumb_url"`
	Price              any       `json:"price"`
	PayWhatYouWant     bool      `json:"pay_what_you_want"`
	FromPrice          *int      `json:"from_price"`
	FromPriceFormatted *string   `json:"from_price_formatted"`
	ToPrice            *int      `json:"to_price"`
	ToPriceFormatted   *string   `json:"to_price_formatted"`
	BuyNowURL          string    `json:"buy_now_url"`
	PriceFormatted     string    `json:"price_formatted"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// ApiResponseRelationshipsProduct relationships of a product
type ApiResponseRelationshipsProduct struct {
	Store    ApiResponseLinks `json:"store"`
	Variants ApiResponseLinks `json:"variants"`
}

// ProductApiResponse represents a product api response
type ProductApiResponse = ApiResponse[ProductAttributes, ApiResponseRelationshipsProduct]

// ProductsApiResponse represents a list of products api responses.
type ProductsApiResponse = ApiResponseList[ProductAttributes, ApiResponseRelationshipsProduct]
