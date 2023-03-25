package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilesService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	_, response, err := client.Files.Get(context.Background(), "36095")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusNotFound, response.HTTPResponse.StatusCode)
}

func TestFilesService_List(t *testing.T) {
	// Setup
	t.Parallel()

	variants, response, err := client.Files.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 0, len(variants.Data))
}
