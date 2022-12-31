package client

import (
	"context"
	"testing"

	"github.com/NdoleStudio/lemonsqueezy-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestWebhooksService_Verify(t *testing.T) {
	// Arrange
	signingSecret := "signing-key"
	client := New(WithSigningSecret(signingSecret))

	// Act
	isValid := client.Webhooks.Verify(context.Background(), "c1a1893f05aaa601cb02ee36f191f70422aa472fbdfcfae6229853d5a875c609", stubs.WebhookRequest())

	// Assert
	assert.True(t, isValid)
}
