package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLicenseKeyInstancesService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	_, response, err := client.LicenseKeyInstances.Get(context.Background(), "9959")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusNotFound, response.HTTPResponse.StatusCode)
}

func TestLicenseKeyInstancesService_List(t *testing.T) {
	// Setup
	t.Parallel()

	licenseKeys, response, err := client.LicenseKeyInstances.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 0, len(licenseKeys.Data))
}
