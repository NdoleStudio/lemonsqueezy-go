package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	store, response, err := client.Stores.Get(context.Background(), 11559)

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, "11559", store.Data.ID)
}

func TestStoresService_List(t *testing.T) {
	// Setup
	t.Parallel()

	stores, response, err := client.Stores.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 2, len(stores.Data))
}
