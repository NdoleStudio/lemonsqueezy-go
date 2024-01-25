package lemonsqueezy

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/NdoleStudio/lemonsqueezy-go/internal/helpers"
	"github.com/NdoleStudio/lemonsqueezy-go/internal/stubs"
)

func TestPricesService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.PriceGetResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	price, response, err := client.Prices.Get(context.Background(), 1)

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.PriceGetResponse(), *response.Body)
	assert.Equal(t, "1", price.Data.ID)

	// Teardown
	server.Close()
}

func TestPricesService_GetWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Prices.Get(context.Background(), 1)

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestPricesService_List(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.PriceListResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	prices, response, err := client.Prices.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.PriceListResponse(), *response.Body)
	assert.Equal(t, 2, len(prices.Data))

	// Teardown
	server.Close()
}

func TestPricesService_ListWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Prices.List(context.Background())

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
