package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/lemonsqueezy-go/internal/helpers"

	"github.com/NdoleStudio/lemonsqueezy-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestWebhooksService_Verify(t *testing.T) {
	// Arrange
	signingSecret := "signing-secret"
	client := New(WithSigningSecret(signingSecret))

	// Act
	isValid := client.Webhooks.Verify(context.Background(), "cde3c645ad532b305247f0668e1882aa3186d3d74d009f23f0e544632ccf6edd", stubs.WebhookRequest())

	// Assert
	assert.True(t, isValid)
}

func TestWebhooksService_Create(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusCreated, stubs.WebhookGetResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	webhook, response, err := client.Webhooks.Create(context.Background(), &WebhookCreateParams{
		URL:     "https://mysite.com/webhooks/",
		Events:  []string{"order_created", "subscription_created"},
		Secret:  "SIGNING_SECRET",
		StoreID: "1",
	})

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusCreated, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.WebhookGetResponse(), *response.Body)
	assert.Equal(t, "1", webhook.Data.ID)

	// Teardown
	server.Close()
}

func TestWebhooksService_CreateWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Webhooks.Create(context.Background(), &WebhookCreateParams{
		URL:     "https://mysite.com/webhooks/",
		Events:  []string{"order_created", "subscription_created"},
		Secret:  "SIGNING_SECRET",
		StoreID: "1",
	})

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestWebhooksService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.WebhookGetResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	webhook, response, err := client.Webhooks.Get(context.Background(), "1")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.WebhookGetResponse(), *response.Body)
	assert.Equal(t, "1", webhook.Data.ID)

	// Teardown
	server.Close()
}

func TestWebhooksService_GetWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Webhooks.Get(context.Background(), "1")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestWebhooksService_Delete(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusNoContent, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	response, err := client.Webhooks.Delete(context.Background(), "1")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNoContent, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestWebhooksService_DeleteWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	response, err := client.Webhooks.Delete(context.Background(), "1")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestWebhooksService_List(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.WebhooksListResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	webhooks, response, err := client.Webhooks.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.WebhooksListResponse(), *response.Body)
	assert.Equal(t, 1, len(webhooks.Data))

	// Teardown
	server.Close()
}

func TestWebhooksService_ListWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Webhooks.List(context.Background())

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestWebhooksService_Update(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.WebhookGetResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	webhook, response, err := client.Webhooks.Update(context.Background(), &WebhookUpdateParams{
		ID:     "1",
		Events: []string{"order_created", "subscription_created"},
		Secret: "SIGNING_SECRET",
	})

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.WebhookGetResponse(), *response.Body)
	assert.Equal(t, "1", webhook.Data.ID)

	// Teardown
	server.Close()
}

func TestWebhooksService_UpdateWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Webhooks.Update(context.Background(), &WebhookUpdateParams{})

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
