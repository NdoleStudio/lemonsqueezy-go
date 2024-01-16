package lemonsqueezy

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"
)

// WebhooksService is the used to verify the signature in webhook requests
type WebhooksService service

// Verify the signature in webhook requests
//
// https://docs.lemonsqueezy.com/api/webhooks#signing-requests
func (service *WebhooksService) Verify(_ context.Context, signature string, body []byte) bool {
	key := []byte(service.client.signingSecret)
	h := hmac.New(sha256.New, key)
	h.Write(body)
	return hex.EncodeToString(h.Sum(nil)) == signature
}

// Create a webhook.
//
// https://docs.lemonsqueezy.com/api/webhooks#create-a-webhook
func (service *WebhooksService) Create(ctx context.Context, storeId int, params *WebhookCreateParams) (*WebhookApiResponse, *Response, error) {
	payload := map[string]any{
		"data": map[string]any{
			"type": "webhooks",
			"attributes": map[string]any{
				"url":    params.URL,
				"events": params.Events,
				"secret": params.Secret,
			},
			"relationships": map[string]any{
				"store": map[string]any{
					"data": map[string]any{
						"type": "stores",
						"id":   strconv.Itoa(storeId),
					},
				},
			},
		},
	}

	response, err := service.client.do(ctx, http.MethodPost, "/v1/webhooks/", payload)
	if err != nil {
		return nil, response, err
	}

	webhook := new(WebhookApiResponse)
	if err = json.Unmarshal(*response.Body, webhook); err != nil {
		return nil, response, err
	}

	return webhook, response, nil
}

// Get the webhook with the given ID.
//
// https://docs.lemonsqueezy.com/api/webhooks#retrieve-a-webhook
func (service *WebhooksService) Get(ctx context.Context, webhookID string) (*WebhookApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/webhooks/"+webhookID)
	if err != nil {
		return nil, response, err
	}

	webhook := new(WebhookApiResponse)
	if err = json.Unmarshal(*response.Body, webhook); err != nil {
		return nil, response, err
	}

	return webhook, response, nil
}

// Update an existing webhook
//
// https://docs.lemonsqueezy.com/api/subscriptions#update-a-webhook
func (service *WebhooksService) Update(ctx context.Context, params *WebhookUpdateParams) (*WebhookApiResponse, *Response, error) {
	payload := map[string]any{
		"data": map[string]any{
			"type": "webhooks",
			"id":   params.ID,
			"attributes": map[string]any{
				"events": params.Events,
				"secret": params.Secret,
			},
		},
	}

	response, err := service.client.do(ctx, http.MethodPatch, "/v1/webhooks/"+params.ID, payload)
	if err != nil {
		return nil, response, err
	}

	subscription := new(WebhookApiResponse)
	if err = json.Unmarshal(*response.Body, subscription); err != nil {
		return nil, response, err
	}

	return subscription, response, nil
}

// Delete a webhook with the given ID.
//
// https://docs.lemonsqueezy.com/api/webhooks#delete-a-webhook
func (service *WebhooksService) Delete(ctx context.Context, webhookID string) (*Response, error) {
	return service.client.do(ctx, http.MethodDelete, "/v1/webhooks/"+webhookID)
}

// List returns a paginated list of webhooks.
//
// https://docs.lemonsqueezy.com/api/webhooks#list-all-webhooks
func (service *WebhooksService) List(ctx context.Context) (*WebhooksApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/webhooks")
	if err != nil {
		return nil, response, err
	}

	webhooks := new(WebhooksApiResponse)
	if err = json.Unmarshal(*response.Body, webhooks); err != nil {
		return nil, response, err
	}

	return webhooks, response, nil
}
