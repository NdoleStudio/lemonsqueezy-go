package client

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/davecgh/go-spew/spew"
)

// webhooksService is the used to verify the signature in webhook requests
type webhooksService service

// Verify the signature in webhook requests
// https://docs.lemonsqueezy.com/api/webhooks#signing-requests
func (service *webhooksService) Verify(_ context.Context, signature string, body []byte) bool {
	key := []byte(service.client.signingSecret)
	h := hmac.New(sha256.New, key)
	h.Write(body)
	spew.Dump(service.client.signingSecret)
	spew.Dump(hex.EncodeToString((h.Sum(nil))))
	return hex.EncodeToString(h.Sum(nil)) == signature
}
