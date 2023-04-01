package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/stretchr/testify/assert"
)

func TestVariantsService_Get(t *testing.T) {
	// Act
	variant, response, err := client.Variants.Get(context.Background(), "36095")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, "36095", variant.Data.ID)
}

func TestVariantsService_List(t *testing.T) {
	// Act
	variants, response, err := client.Variants.List(context.Background())

	spew.Dump(variants)

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 6, len(variants.Data))
}
