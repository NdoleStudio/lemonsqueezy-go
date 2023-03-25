package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdersService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	_, response, err := client.Orders.Get(context.Background(), "325103")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
}

func TestOrdersService_List(t *testing.T) {
	// Setup
	t.Parallel()

	orders, response, err := client.Orders.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 7, len(orders.Data))
}
