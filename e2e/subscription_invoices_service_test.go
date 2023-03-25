package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubscriptionInvoicesService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	_, response, err := client.SubscriptionInvoices.Get(context.Background(), "9959")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
}

func TestSubscriptionInvoicesService_List(t *testing.T) {
	// Setup
	t.Parallel()

	invoices, response, err := client.SubscriptionInvoices.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 10, len(invoices.Data))
}
