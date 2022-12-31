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
	isValid := client.Webhooks.Verify(context.Background(), "waGJPwWqpgHLAu428ZH3BCKqRy+9/PrmIphT1ah1xgk=", stubs.WebhookRequest())

	// Assert
	assert.True(t, isValid)
}
