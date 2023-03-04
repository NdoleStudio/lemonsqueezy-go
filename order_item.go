package client

import "time"

// OrderItemAttributes is a line item for an order that includes product, variant and price information.
type OrderItemAttributes struct {
	OrderID     int       `json:"order_id"`
	ProductID   int       `json:"product_id"`
	VariantID   int       `json:"variant_id"`
	ProductName string    `json:"product_name"`
	VariantName string    `json:"variant_name"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ApiResponseRelationshipsOrderItem relationships of an order-item
type ApiResponseRelationshipsOrderItem struct {
	Order   ApiResponseLinks `json:"order"`
	Product ApiResponseLinks `json:"product"`
	Variant ApiResponseLinks `json:"variant"`
}

// OrderItemApiResponse is the api response for one order item
type OrderItemApiResponse = ApiResponse[OrderItemAttributes, ApiResponseRelationshipsOrderItem]

// OrderItemsApiResponse is the api response for a list of order items.
type OrderItemsApiResponse = ApiResponseList[OrderItemAttributes, ApiResponseRelationshipsOrderItem]
