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

func TestSubscriptionItemsService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.SubscriptionItemGetResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	subscriptionItem, response, err := client.SubscriptionItems.Get(context.Background(), "1")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.SubscriptionItemGetResponse(), *response.Body)

	assert.Equal(t, &SubscriptionItemApiResponse{
		Jsonapi: ApiResponseJSONAPI{
			Version: "1.0",
		},
		Links: ApiResponseSelfLink{
			Self: "https://api.lemonsqueezy.com/v1/subscription-item/1",
		},
		Data: ApiResponseData[SubscriptionItem, ApiResponseRelationshipsSubscriptionItem]{
			Type: "subscription-items",
			ID:   "1",
			Attributes: SubscriptionItem{
				SubscriptionID: 1,
				PriceID:        1,
				Quantity:       1,
				IsUsageBased:   false,
				CreatedAt:      time.Date(2023, time.July, 18, 12, 16, 24, 0, time.UTC),
				UpdatedAt:      time.Date(2023, time.July, 18, 12, 16, 24, 0, time.UTC),
			},
			Relationships: ApiResponseRelationshipsSubscriptionItem{
				Subscription: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscription-items/1/subscription",
						Self:    "https://api.lemonsqueezy.com/v1/subscription-items/1/relationships/subscription",
					},
				},
				Price: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscription-items/1/price",
						Self:    "https://api.lemonsqueezy.com/v1/subscription-items/1/relationships/price",
					},
				},
				UsageRecords: ApiResponseLinks{
					Links: ApiResponseLink{
						Related: "https://api.lemonsqueezy.com/v1/subscription-items/1/usage-records",
						Self:    "https://api.lemonsqueezy.com/v1/subscription-items/1/relationships/usage-records",
					},
				},
			},
			Links: ApiResponseSelfLink{
				Self: "https://api.lemonsqueezy.com/v1/subscription-items/1",
			},
		},
	}, subscriptionItem)

	// Teardown
	server.Close()
}

func TestSubscriptionItemsService_GetWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.SubscriptionItems.Get(context.Background(), "1")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestSubscriptionItemsService_Update(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.SubscriptionItemUpdateResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	subscriptionItem, response, err := client.SubscriptionItems.Update(context.Background(), &SubscriptionItemUpdateParams{
		ID: "1",
		Attributes: SubscriptionItemUpdateParamsAttributes{
			Quantity: 10,
		},
	})

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.SubscriptionItemUpdateResponse(), *response.Body)
	assert.Equal(t, "1", subscriptionItem.Data.ID)
	assert.Equal(t, 10, subscriptionItem.Data.Attributes.Quantity)

	// Teardown
	server.Close()
}

func TestSubscriptionItemsService_UpdateWithError(t *testing.T) {
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

func TestSubscriptionItemsService_List(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.SubscriptionItemsListResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	subscriptionItems, response, err := client.SubscriptionItems.List(context.Background(), nil)

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.SubscriptionItemsListResponse(), *response.Body)
	assert.Equal(t, 1, len(subscriptionItems.Data))
	assert.Equal(t, "1", subscriptionItems.Data[0].ID)

	// Teardown
	server.Close()
}

func TestSubscriptionItemsService_ListWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.SubscriptionItems.List(context.Background(), nil)

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestSubscriptionItemsService_CurrentUsage(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.SubscriptionItemCurrentUsageResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	subscriptionItemCurrentUsage, response, err := client.SubscriptionItems.CurrentUsage(context.Background(), "1")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.SubscriptionItemCurrentUsageResponse(), *response.Body)

	assert.Equal(t, &SubscriptionItemCurrentUsageApiResponse{
		Jsonapi: ApiResponseJSONAPI{
			Version: "1.0",
		},
		Meta: ApiResponseMetaSubscriptionItemCurrentUsage{
			PeriodStart:      time.Date(2023, time.August, 10, 13, 8, 16, 0, time.UTC),
			PeriodEnd:        time.Date(2023, time.September, 10, 13, 3, 16, 0, time.UTC),
			Quantity:         5,
			IntervalUnit:     "month",
			IntervalQuantity: 1,
		},
	}, subscriptionItemCurrentUsage)

	// Teardown
	server.Close()
}

func TestSubscriptionItemsService_CurrentUsageWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.SubscriptionItems.CurrentUsage(context.Background(), "1")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
