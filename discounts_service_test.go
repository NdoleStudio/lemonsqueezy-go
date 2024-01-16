package lemonsqueezy

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/NdoleStudio/lemonsqueezy-go/internal/helpers"
	"github.com/NdoleStudio/lemonsqueezy-go/internal/stubs"
)

func TestDiscountsService_Create(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusCreated, stubs.DiscountGetResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	discount, response, err := client.Discounts.Create(context.Background(), &DiscountCreateParams{
		Name:       "10% Off",
		Code:       "10PERCENT",
		Amount:     10,
		AmountType: "percent",
		StoreID:    1,
	})

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusCreated, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.DiscountGetResponse(), *response.Body)
	assert.Equal(t, "1", discount.Data.ID)

	// Teardown
	server.Close()
}

func TestDiscountsService_CreateWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Discounts.Create(context.Background(), &DiscountCreateParams{
		Name:       "10% Off",
		Code:       "10PERCENT",
		Amount:     10,
		AmountType: "percent",
		StoreID:    1,
	})

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestDiscountsService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.DiscountGetResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	discount, response, err := client.Discounts.Get(context.Background(), "1")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.DiscountGetResponse(), *response.Body)
	assert.Equal(t, "1", discount.Data.ID)

	// Teardown
	server.Close()
}

func TestDiscountsService_GetWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Discounts.Get(context.Background(), "1")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestDiscountsService_Delete(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusNoContent, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	response, err := client.Discounts.Delete(context.Background(), "1")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNoContent, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestDiscountsService_DeleteWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	response, err := client.Discounts.Delete(context.Background(), "1")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestDiscountsService_List(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.DiscountsListResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	discounts, response, err := client.Discounts.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.DiscountsListResponse(), *response.Body)
	assert.Equal(t, 1, len(discounts.Data))

	// Teardown
	server.Close()
}

func TestDiscountsService_ListWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Discounts.List(context.Background())

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
