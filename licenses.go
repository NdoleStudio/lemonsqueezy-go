package lemonsqueezy

import "time"

type LicenseAttributes struct {
	Error      string          `json:"error"`
	LicenseKey LicenseKey      `json:"license_key"`
	Instance   LicenseInstance `json:"instance"`
	Meta       LicenseMeta     `json:"meta"`
}

type LicenseKey struct {
	ID              int        `json:"id"`
	Status          string     `json:"status"`
	Key             string     `json:"key"`
	ActivationLimit int        `json:"activation_limit"`
	ActivationUsage int        `json:"activation_usage"`
	CreatedAt       time.Time  `json:"created_at"`
	ExpiresAt       *time.Time `json:"expires_at"`
	TestMode        bool       `json:"test_mode"`
}

type LicenseInstance struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type LicenseMeta struct {
	StoreID       int    `json:"store_id"`
	OrderID       int    `json:"order_id"`
	OrderItemID   int    `json:"order_item_id"`
	VariantID     int    `json:"variant_id"`
	VariantName   string `json:"variant_name"`
	ProductID     int    `json:"product_id"`
	ProductName   string `json:"product_name"`
	CustomerID    int    `json:"customer_id"`
	CustomerName  string `json:"customer_name"`
	CustomerEmail string `json:"customer_email"`
}

type LicenseActivateApiResponse struct {
	Activated bool `json:"activated"`
	LicenseAttributes
}

type LicenseValidateApiResponse struct {
	Valid bool `json:"valid"`
	LicenseAttributes
}
type LicenseDeactivateApiResponse struct {
	Deactivated bool `json:"deactivated"`
	LicenseAttributes
}
