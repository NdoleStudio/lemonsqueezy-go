package lemonsqueezy

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/lemonsqueezy-go/internal/helpers"
	"github.com/NdoleStudio/lemonsqueezy-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestLicenseKeyInstancesService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.LicenseKeyInstanceGetResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	licenseKeyInstance, response, err := client.LicenseKeyInstances.Get(context.Background(), "1")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.LicenseKeyInstanceGetResponse(), *response.Body)
	assert.Equal(t, "1", licenseKeyInstance.Data.ID)

	// Teardown
	server.Close()
}

func TestLicenseKeyInstancesService_GetWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.LicenseKeyInstances.Get(context.Background(), "1")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestLicenseKeyInstancesService_List(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.LicenseKeyInstancesListResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	licenseKeyInstances, response, err := client.LicenseKeyInstances.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.LicenseKeyInstancesListResponse(), *response.Body)
	assert.Equal(t, 1, len(licenseKeyInstances.Data))

	// Teardown
	server.Close()
}

func TestLicenseKeyInstancesService_ListWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.LicenseKeyInstances.List(context.Background())

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
