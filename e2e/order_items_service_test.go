package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderItemsService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	_, response, err := client.OrderItems.Get(context.Background(), "382956")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
}

func TestOrderItemsService_List(t *testing.T) {
	// Setup
	t.Parallel()

	orderItems, response, err := client.OrderItems.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 7, len(orderItems.Data))
}
