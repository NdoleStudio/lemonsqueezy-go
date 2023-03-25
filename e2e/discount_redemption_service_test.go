package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscountRedemptionService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	_, response, err := client.DiscountRedemptions.Get(context.Background(), "9959")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusNotFound, response.HTTPResponse.StatusCode)
}

func TestDiscountRedemptionsService_List(t *testing.T) {
	// Setup
	t.Parallel()

	redemptions, response, err := client.DiscountRedemptions.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 0, len(redemptions.Data))
}
