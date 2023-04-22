package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLicenseKeyService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	_, response, err := client.LicenseKeys.Get(context.Background(), "9959")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusNotFound, response.HTTPResponse.StatusCode)
}

func TestLicenseKeysService_List(t *testing.T) {
	// Setup
	t.Parallel()

	licenseKeys, response, err := client.LicenseKeys.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 0, len(licenseKeys.Data))
}
