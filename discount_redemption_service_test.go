package lemonsqueezy

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/lemonsqueezy-go/internal/helpers"
	"github.com/NdoleStudio/lemonsqueezy-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestDiscountRedemptionsService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.DiscountRedemptionGetResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	redemption, response, err := client.DiscountRedemptions.Get(context.Background(), "1")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.DiscountRedemptionGetResponse(), *response.Body)
	assert.Equal(t, "1", redemption.Data.ID)

	// Teardown
	server.Close()
}

func TestDiscountRedemptionsService_GetWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.DiscountRedemptions.Get(context.Background(), "1")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestDiscountRedemptionsService_List(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.DiscountRedemptionsListResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	redemptions, response, err := client.DiscountRedemptions.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.DiscountRedemptionsListResponse(), *response.Body)
	assert.Equal(t, 1, len(redemptions.Data))

	// Teardown
	server.Close()
}

func TestDiscountRedemptionsService_ListWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.DiscountRedemptions.List(context.Background())

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
