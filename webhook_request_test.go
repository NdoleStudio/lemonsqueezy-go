package lemonsqueezy

import (
	"encoding/json"
	"testing"

	"github.com/NdoleStudio/lemonsqueezy-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestWebhookRequest_Order(t *testing.T) {
	// Setup
	t.Parallel()

	var request WebhookRequestOrder
	err := json.Unmarshal(stubs.WebhookRequestOrderCreated(), &request)
	assert.Nil(t, err)

	assert.Equal(t, WebhookRequestMeta{
		EventName: "order_created",
		TestMode:  false,
		CustomData: map[string]any{
			"customer_id": float64(25),
		},
	}, request.Meta)

	assert.Equal(t, "89b36d62-4f5c-4353-853f-0c769d0535c8", request.Data.Attributes.Identifier)
}
