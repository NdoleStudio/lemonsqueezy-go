package client

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/NdoleStudio/lemonsqueezy-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestWebhook(t *testing.T) {
	// Setup
	t.Parallel()

	var request webhookRequest[OrderCreatedWebhookRequestAttributes, OrderCreatedWebhookRequestRelationships]
	err := json.Unmarshal(stubs.WebhookRequest(), &request)
	assert.Nil(t, err)

	assert.Equal(t, webhookRequest[OrderCreatedWebhookRequestAttributes, OrderCreatedWebhookRequestRelationships]{
		Meta: WebhookRequestMeta{
			EventName:  "order_created",
			TestMode:   false,
			CustomData: map[string]interface{}(nil),
		},
		Data: WebhookRequestData[OrderCreatedWebhookRequestAttributes, OrderCreatedWebhookRequestRelationships]{
			Type: "orders",
			ID:   "1",
			Attributes: OrderCreatedWebhookRequestAttributes{
				StoreID:          1,
				Identifier:       "636f855c-1fb9-4c07-b75c-3a10afef010a",
				OrderNumber:      1,
				UserName:         "Darlene Daugherty",
				UserEmail:        "gernser@yahoo.com",
				Currency:         "USD",
				CurrencyRate:     "1.0000",
				Subtotal:         999,
				DiscountTotal:    0,
				Tax:              200,
				Total:            1199,
				SubtotalUsd:      999,
				DiscountTotalUsd: 0,
				TaxUsd:           200,
				TotalUsd:         1199,
				TaxName:          "VAT",
				TaxRate:          "20.00",
				Status:           "paid",
				StatusFormatted:  "Paid",
				Refunded:         false,
				RefundedAt:       nil,
				CreatedAt:        time.Date(2021, time.August, 11, 13, 47, 29, 0, time.UTC),
				UpdatedAt:        time.Date(2021, time.August, 11, 13, 54, 54, 0, time.UTC),
			},
			Relationships: OrderCreatedWebhookRequestRelationships{
				Store: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/orders/1/store",
						Self:    "https://api.lemonsqueezy.com/v1/orders/1/relationships/store",
					},
				},
				OrderItem: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/orders/1/order-items",
						Self:    "https://api.lemonsqueezy.com/v1/orders/1/relationships/order-items",
					},
				},
				Subscriptions: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/orders/1/subscriptions",
						Self:    "https://api.lemonsqueezy.com/v1/orders/1/relationships/subscriptions",
					},
				},
				LicenseKeys: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/orders/1/license-keys",
						Self:    "https://api.lemonsqueezy.com/v1/orders/1/relationships/license-keys",
					},
				},
			},
			Links: ApiResponseLink{
				Self: "https://api.lemonsqueezy.com/v1/orders/1",
			},
		},
	}, request)
}
