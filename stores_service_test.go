package lemonsqueezy

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/NdoleStudio/lemonsqueezy-go/internal/helpers"
	"github.com/NdoleStudio/lemonsqueezy-go/internal/stubs"
)

func TestStoreService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.StoreGetResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	stores, response, err := client.Stores.Get(context.Background(), 1)

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.StoreGetResponse(), *response.Body)
	assert.Equal(t, "1", stores.Data.ID)

	// Teardown
	server.Close()
}

func TestStoreService_GetWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Stores.Get(context.Background(), 1)

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestStoresService_List(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.StoresListResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	stores, response, err := client.Stores.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.StoresListResponse(), *response.Body)
	assert.Equal(t, 2, len(stores.Data))

	// Teardown
	server.Close()
}

func TestStoresService_ListWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Stores.List(context.Background())

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
