package lemonsqueezy

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// SubscriptionInvoicesService is the API client for the `/v1/subscription-invoices` endpoint
type SubscriptionInvoicesService service

// Get returns the subscription invoice with the given ID.
//
// https://docs.lemonsqueezy.com/api/subscription-invoices#retrieve-a-subscription-invoice
func (service *SubscriptionInvoicesService) Get(ctx context.Context, invoiceID string) (*SubscriptionInvoiceAPIResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/subscription-invoices/"+invoiceID)
	if err != nil {
		return nil, response, err
	}

	invoice := new(SubscriptionInvoiceAPIResponse)
	if err = json.Unmarshal(*response.Body, invoice); err != nil {
		return nil, response, err
	}

	return invoice, response, nil
}

// List a paginated list of subscription invoices.
//
// https://docs.lemonsqueezy.com/api/subscription-invoices#list-all-subscription-invoices
func (service *SubscriptionInvoicesService) List(ctx context.Context, filters map[string]string) (*SubscriptionInvoicesAPIResponse, *Response, error) {
	endpoint, err := url.Parse("/v1/subscription-invoices")
	if err != nil {
		return nil, nil, err
	}

	query := endpoint.Query()
	for key, val := range filters {
		query.Add(key, val)
	}
	endpoint.RawQuery = query.Encode()

	response, err := service.client.do(ctx, http.MethodGet, endpoint.String())
	if err != nil {
		return nil, response, err
	}

	invoices := new(SubscriptionInvoicesAPIResponse)
	if err = json.Unmarshal(*response.Body, invoices); err != nil {
		return nil, response, err
	}

	return invoices, response, nil
}

// Generate a Subscription Invoice.
//
// https://docs.lemonsqueezy.com/api/subscription-invoices/generate-subscription-invoice
func (service *SubscriptionInvoicesService) Generate(ctx context.Context, subscriptionID string, params map[string]string) (*SubscriptionInvoiceGenerateResponse, *Response, error) {
	endpoint, err := url.Parse(fmt.Sprintf("/v1/subscription-invoices/%s/generate-invoice", subscriptionID))
	if err != nil {
		return nil, nil, err
	}

	query := endpoint.Query()
	for key, val := range params {
		query.Add(key, val)
	}
	endpoint.RawQuery = query.Encode()

	response, err := service.client.do(ctx, http.MethodPost, endpoint.String())
	if err != nil {
		return nil, response, err
	}

	invoice := new(SubscriptionInvoiceGenerateResponse)
	if err = json.Unmarshal(*response.Body, invoice); err != nil {
		return nil, response, err
	}

	return invoice, response, nil
}
