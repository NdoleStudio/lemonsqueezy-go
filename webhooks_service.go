package client

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// WebhooksService is the used to verify the signature in webhook requests
type WebhooksService service

// Verify the signature in webhook requests
// https://docs.lemonsqueezy.com/api/webhooks#signing-requests
func (service *WebhooksService) Verify(_ context.Context, signature string, body []byte) bool {
	key := []byte(service.client.signingSecret)
	h := hmac.New(sha256.New, key)
	h.Write(body)
	return hex.EncodeToString(h.Sum(nil)) == signature
}
