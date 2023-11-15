package lemonsqueezy

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/NdoleStudio/lemonsqueezy-go/internal/helpers"
	"github.com/NdoleStudio/lemonsqueezy-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestSubscriptionsService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.SubscriptionGetResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	subscription, response, err := client.Subscriptions.Get(context.Background(), "1")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.SubscriptionGetResponse(), *response.Body)

	assert.Equal(t, &SubscriptionApiResponse{
		Jsonapi: ApiResponseJSONAPI{
			Version: "1.0",
		},
		Links: ApiResponseSelfLink{
			Self: "https://api.lemonsqueezy.com/v1/subscriptions/1",
		},
		Data: ApiResponseData[Subscription, ApiResponseRelationshipsSubscription]{
			Type: "subscriptions",
			ID:   "1",
			Attributes: Subscription{
				StoreID:         1,
				CustomerID:      1,
				OrderID:         1,
				OrderItemID:     1,
				ProductID:       1,
				VariantID:       1,
				ProductName:     "Example Product",
				VariantName:     "Example Variant",
				UserName:        "Darlene Daugherty",
				UserEmail:       "gernser@yahoo.com",
				Status:          "active",
				StatusFormatted: "Active",
				CardBrand:       "visa",
				CardLastFour:    "42424",
				Pause:           nil,
				Cancelled:       false,
				TrialEndsAt:     nil,
				BillingAnchor:   12,
				FirstSubscriptionItem: &SubscriptionFirstSubscriptionItem{
					ID: 1,
					SubscriptionItem: SubscriptionItem{
						SubscriptionID: 1,
						PriceID:        1,
						Quantity:       5,
						CreatedAt:      time.Date(2021, time.August, 11, 13, 47, 28, 0, time.UTC),
						UpdatedAt:      time.Date(2021, time.August, 11, 13, 47, 28, 0, time.UTC),
					},
				},
				Urls: SubscriptionURLs{
					UpdatePaymentMethod: "https://my-store.lemonsqueezy.com/subscription/1/payment-details?expires=1666869343&signature=9985e3bf9007840aeb3951412be475abc17439c449c1af3e56e08e45e1345413",
					CustomerPortal:      "https://my-store.lemonsqueezy.com/billing?expires=1666869343&signature=82ae290ceac8edd4190c82825dd73a8743346d894a8ddbc4898b97eb96d105a5",
				},
				RenewsAt:  time.Date(2022, time.November, 12, 0, 0, 0, 0, time.UTC),
				EndsAt:    nil,
				CreatedAt: time.Date(2021, time.August, 11, 13, 47, 27, 0, time.UTC),
				UpdatedAt: time.Date(2021, time.August, 11, 13, 54, 19, 0, time.UTC),
				TestMode:  false,
			},
			Relationships: ApiResponseRelationshipsSubscription{
				Store: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/store",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/store",
					},
				},
				Customer: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/customer",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/customer",
					},
				},
				Order: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/order",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/order",
					},
				},
				OrderItem: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/order-item",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/order-item",
					},
				},
				Product: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/product",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/product",
					},
				},
				Variant: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/variant",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/variant",
					},
				},
				SubscriptionItems: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/subscription-items",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/subscription-items",
					},
				},
				SubscriptionInvoices: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/subscription-invoices",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/subscription-invoices",
					},
				},
			},
			Links: ApiResponseSelfLink{
				Self: "https://api.lemonsqueezy.com/v1/subscriptions/1",
			},
		},
	}, subscription)

	// Teardown
	server.Close()
}

func TestSubscriptionsService_GetWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Subscriptions.Get(context.Background(), "1")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestSubscriptionsService_Cancel(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.SubscriptionCancelResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	subscription, response, err := client.Subscriptions.Cancel(context.Background(), "1")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.SubscriptionCancelResponse(), *response.Body)

	assert.Equal(t, &SubscriptionApiResponse{
		Jsonapi: ApiResponseJSONAPI{
			Version: "1.0",
		},
		Links: ApiResponseSelfLink{
			Self: "https://api.lemonsqueezy.com/v1/subscriptions/1",
		},
		Data: ApiResponseData[Subscription, ApiResponseRelationshipsSubscription]{
			Type: "subscriptions",
			ID:   "1",
			Attributes: Subscription{
				StoreID:         1,
				CustomerID:      1,
				OrderID:         1,
				OrderItemID:     1,
				ProductID:       1,
				VariantID:       1,
				ProductName:     "Example Product",
				VariantName:     "Example Variant",
				UserName:        "Darlene Daugherty",
				UserEmail:       "gernser@yahoo.com",
				Status:          "cancelled",
				StatusFormatted: "Cancelled",
				Pause:           nil,
				Cancelled:       true,
				TrialEndsAt:     nil,
				BillingAnchor:   12,
				FirstSubscriptionItem: &SubscriptionFirstSubscriptionItem{
					ID: 1,
					SubscriptionItem: SubscriptionItem{
						SubscriptionID: 1,
						PriceID:        1,
						Quantity:       5,
						CreatedAt:      time.Date(2021, time.August, 11, 13, 47, 28, 0, time.UTC),
						UpdatedAt:      time.Date(2021, time.August, 11, 13, 47, 28, 0, time.UTC),
					},
				},
				Urls: SubscriptionURLs{
					UpdatePaymentMethod: "https://my-store.lemonsqueezy.com/subscription/1/payment-details?expires=1666869343&signature=9985e3bf9007840aeb3951412be475abc17439c449c1af3e56e08e45e1345413",
					CustomerPortal:      "https://my-store.lemonsqueezy.com/billing?expires=1666869343&signature=82ae290ceac8edd4190c82825dd73a8743346d894a8ddbc4898b97eb96d105a5",
				},
				RenewsAt: time.Date(2022, time.November, 12, 0, 0, 0, 0, time.UTC),
				EndsAt: (func() *time.Time {
					val := time.Date(2022, time.November, 12, 0, 0, 0, 0, time.UTC)
					return &val
				})(),
				CreatedAt: time.Date(2021, time.August, 11, 13, 47, 27, 0, time.UTC),
				UpdatedAt: time.Date(2021, time.August, 11, 13, 54, 19, 0, time.UTC),
				TestMode:  false,
			},
			Relationships: ApiResponseRelationshipsSubscription{
				Store: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/store",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/store",
					},
				},
				Customer: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/customer",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/customer",
					},
				},
				Order: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/order",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/order",
					},
				},
				OrderItem: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/order-item",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/order-item",
					},
				},
				Product: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/product",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/product",
					},
				},
				Variant: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/variant",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/variant",
					},
				},
				SubscriptionItems: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/subscription-items",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/subscription-items",
					},
				},
				SubscriptionInvoices: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscriptions/1/subscription-invoices",
						Self:    "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/subscription-invoices",
					},
				},
			},
			Links: ApiResponseSelfLink{
				Self: "https://api.lemonsqueezy.com/v1/subscriptions/1",
			},
		},
	}, subscription)

	// Teardown
	server.Close()
}

func TestSubscriptionsService_CancelWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Subscriptions.Cancel(context.Background(), "1")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestSubscriptionsService_Update(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.SubscriptionUpdateResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	subscription, response, err := client.Subscriptions.Update(context.Background(), &SubscriptionUpdateParams{
		ID: "1",
		Attributes: SubscriptionUpdateParamsAttributes{
			ProductID:     9,
			VariantID:     11,
			BillingAnchor: 29,
		},
	})

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.SubscriptionUpdateResponse(), *response.Body)
	assert.Equal(t, "1", subscription.Data.ID)

	// Teardown
	server.Close()
}

func TestSubscriptionsService_UpdateWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Subscriptions.Update(context.Background(), &SubscriptionUpdateParams{})

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestSubscriptionsService_List(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.SubscriptionsListResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	subscriptions, response, err := client.Subscriptions.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.SubscriptionsListResponse(), *response.Body)
	assert.Equal(t, 2, len(subscriptions.Data))
	assert.Equal(t, "2", subscriptions.Data[1].ID)

	// Teardown
	server.Close()
}

func TestSubscriptionsService_ListWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Subscriptions.List(context.Background())

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
