package lemonsqueezy

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/lemonsqueezy-go/internal/helpers"
	"github.com/NdoleStudio/lemonsqueezy-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestSubscriptionInvoicesService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.SubscriptionInvoiceGetResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	invoice, response, err := client.SubscriptionInvoices.Get(context.Background(), "1")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.SubscriptionInvoiceGetResponse(), *response.Body)
	assert.Equal(t, "1", invoice.Data.ID)

	// Teardown
	server.Close()
}

func TestSubscriptionInvoicesService_GetWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.SubscriptionInvoices.Get(context.Background(), "1")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestSubscriptionInvoicesService_List(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.SubscriptionInvoicesListResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	invoices, response, err := client.SubscriptionInvoices.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.SubscriptionInvoicesListResponse(), *response.Body)
	assert.Equal(t, 1, len(invoices.Data))

	// Teardown
	server.Close()
}

func TestSubscriptionInvoicesService_ListWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.SubscriptionInvoices.List(context.Background())

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
