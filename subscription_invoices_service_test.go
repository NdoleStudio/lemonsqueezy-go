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
	invoices, response, err := client.SubscriptionInvoices.List(context.Background(), nil)

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
	_, response, err := client.SubscriptionInvoices.List(context.Background(), nil)

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestSubscriptionInvoicesService_Generate(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.SubscriptionInvoicesGenerateResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	invoice, response, err := client.SubscriptionInvoices.Generate(context.Background(), "1234", map[string]string{
		"address":  "123 Main St",
		"city":     "Anytown",
		"country":  "US",
		"name":     "John Doe",
		"notes":    "Thank you for your business",
		"state":    "CA",
		"zip_code": "12345",
	})

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.SubscriptionInvoicesGenerateResponse(), *response.Body)
	assert.Equal(
		t,
		"https://app.lemonsqueezy.com/my-orders/c1e4de31-b4cf-4668-a6fe-019d071d41ab/invoice/download?address=123%20Main%20St&city=Anytown&country=US&name=John%20Doe&notes=Thank%20you%20for%20your%20business&state=CA&zip_code=12345&signature=c21a17a13d9deeacc99ff52144a06b49df141af37dd668cc70a76bcc8022888e",
		invoice.Meta.Urls.DownloadInvoice,
	)

	// Teardown
	server.Close()
}

func TestSubscriptionInvoicesService_GenerateWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.SubscriptionInvoices.Generate(context.Background(), "1234", map[string]string{
		"address":  "123 Main St",
		"city":     "Anytown",
		"country":  "US",
		"name":     "John Doe",
		"notes":    "Thank you for your business",
		"state":    "CA",
		"zip_code": "12345",
	})

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
