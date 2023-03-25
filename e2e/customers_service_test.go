package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomersService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	_, response, err := client.Customers.Get(context.Background(), "361810")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusNotFound, response.HTTPResponse.StatusCode)
}

func TestCustomersService_List(t *testing.T) {
	// Setup
	t.Parallel()

	customers, response, err := client.Customers.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 0, len(customers.Data))
}
