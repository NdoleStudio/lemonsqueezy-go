package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductsService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	product, response, err := client.Products.Get(context.Background(), "41443")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, "41443", product.Data.ID)
}

func TestProductsService_List(t *testing.T) {
	// Setup
	t.Parallel()

	products, response, err := client.Products.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 2, len(products.Data))
}
