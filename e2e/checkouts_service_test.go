package e2e

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	lemonsqueezy "github.com/NdoleStudio/lemonsqueezy-go"
	"github.com/stretchr/testify/assert"
)

func TestCheckoutsService_Create(t *testing.T) {
	// Act
	storeID := "11559"
	checkout, response, err := client.Checkouts.Create(context.Background(), &lemonsqueezy.CheckoutCreateParams{
		CustomPrice:     5000,
		EnabledVariants: []int{36096},
		ButtonColor:     "#2DD272",
		CustomData:      map[string]string{"user_id": "123"},
		ExpiresAt:       time.Now().UTC().Add(time.Hour * 24),
		StoreID:         storeID,
		VariantID:       "36096",
	})

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusCreated, response.HTTPResponse.StatusCode)
	assert.Equal(t, storeID, fmt.Sprintf("%d", checkout.Data.Attributes.StoreID))
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
