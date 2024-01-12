package lemonsqueezy

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/lemonsqueezy-go/internal/helpers"
	"github.com/NdoleStudio/lemonsqueezy-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestLicensesService_Activate(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.LicenseActivateResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	licenseActivation, response, err := client.Licenses.Activate(context.Background(), "1234567890", "test")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.LicenseActivateResponse(), *response.Body)
	assert.Equal(t, true, licenseActivation.Activated)

	// Teardown
	server.Close()
}

func TestLicensesService_ActivateWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Licenses.Activate(context.Background(), "1234567890", "test")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestLicensesService_Validate(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.LicenseValidateResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	licenseValidation, response, err := client.Licenses.Validate(context.Background(), "1234567890", "")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.LicenseValidateResponse(), *response.Body)
	assert.Equal(t, true, licenseValidation.Valid)

	// Teardown
	server.Close()
}

func TestLicensesService_ValidateWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Licenses.Validate(context.Background(), "1234567890", "")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestLicensesService_Deactivate(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.LicenseDeactivateResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	licenseDeactivation, response, err := client.Licenses.Deactivate(context.Background(), "1234567890", "abc123")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.LicenseDeactivateResponse(), *response.Body)
	assert.Equal(t, true, licenseDeactivation.Deactivated)

	// Teardown
	server.Close()
}

func TestLicensesService_DeactivateWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Licenses.Deactivate(context.Background(), "1234567890", "abc123")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
