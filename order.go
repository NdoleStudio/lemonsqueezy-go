package client

import "time"

// OrderAttributes for an order which is created when a customer purchases a product.
type OrderAttributes struct {
	StoreID                int        `json:"store_id"`
	Identifier             string     `json:"identifier"`
	OrderNumber            int        `json:"order_number"`
	UserName               string     `json:"user_name"`
	UserEmail              string     `json:"user_email"`
	Currency               string     `json:"currency"`
	CurrencyRate           string     `json:"currency_rate"`
	Subtotal               int        `json:"subtotal"`
	DiscountTotal          int        `json:"discount_total"`
	Tax                    int        `json:"tax"`
	Total                  int        `json:"total"`
	SubtotalUsd            int        `json:"subtotal_usd"`
	DiscountTotalUsd       int        `json:"discount_total_usd"`
	TaxUsd                 int        `json:"tax_usd"`
	TotalUsd               int        `json:"total_usd"`
	TaxName                string     `json:"tax_name"`
	TaxRate                string     `json:"tax_rate"`
	Status                 string     `json:"status"`
	StatusFormatted        string     `json:"status_formatted"`
	Refunded               int        `json:"refunded"`
	RefundedAt             *time.Time `json:"refunded_at"`
	SubtotalFormatted      string     `json:"subtotal_formatted"`
	DiscountTotalFormatted string     `json:"discount_total_formatted"`
	TaxFormatted           string     `json:"tax_formatted"`
	TotalFormatted         string     `json:"total_formatted"`
	FirstOrderItem         struct {
		ID          int       `json:"id"`
		OrderID     int       `json:"order_id"`
		ProductID   int       `json:"product_id"`
		VariantID   int       `json:"variant_id"`
		ProductName string    `json:"product_name"`
		VariantName string    `json:"variant_name"`
		Price       int       `json:"price"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		TestMode    bool      `json:"test_mode"`
	} `json:"first_order_item"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ApiResponseRelationshipsOrder relationships of an order
type ApiResponseRelationshipsOrder struct {
	Store         ApiResponseLinks `json:"store"`
	OrderItems    ApiResponseLinks `json:"order-items"`
	Subscriptions ApiResponseLinks `json:"subscriptions"`
	LicenseKeys   ApiResponseLinks `json:"license-keys"`
}

// OrderApiResponse is the api response for one order
type OrderApiResponse = ApiResponse[OrderAttributes, ApiResponseRelationshipsOrder]

// OrdersApiResponse is the api response for a list of orders.
type OrdersApiResponse = ApiResponseList[OrderAttributes, ApiResponseRelationshipsOrder]
