package e2e

import (
	"context"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/NdoleStudio/lemonsqueezy-go"
)

func TestCheckoutsService_Create(t *testing.T) {
	storeID := 11559
	variantId := 36096
	expiresAt := time.Now().UTC().Add(time.Hour * 24)
	customPrice := 5000

	// Act
	checkout, response, err := client.Checkouts.Create(context.Background(), variantId, storeID, &lemonsqueezy.CheckoutCreateAttributes{
		CustomPrice: &customPrice,
		ProductOptions: lemonsqueezy.CheckoutCreateProductOptions{
			EnabledVariants: []int{variantId},
		},
		CheckoutOptions: lemonsqueezy.CheckoutCreateOptions{
			ButtonColor: "#2DD272",
		},
		CheckoutData: lemonsqueezy.CheckoutCreateData{
			Custom: map[string]any{"user_id": "123"},
		},
		ExpiresAt: &expiresAt,
	})

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusCreated, response.HTTPResponse.StatusCode)
	assert.Equal(t, storeID, strconv.Itoa(checkout.Data.Attributes.StoreID))
}

func TestCheckoutsService_Get(t *testing.T) {
	// Arrange
	checkoutID := "b75461df-3832-4ec3-b7d3-f371a07a4eaa"

	// Act
	checkout, response, err := client.Checkouts.Get(context.Background(), checkoutID)

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, checkoutID, checkout.Data.ID)
}

func TestCheckoutsService_List(t *testing.T) {
	// Act
	checkouts, response, err := client.Checkouts.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.GreaterOrEqual(t, 10, len(checkouts.Data))
}
