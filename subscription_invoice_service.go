package lemonsqueezy

import (
	"context"
	"encoding/json"
	"net/http"
)

// SubscriptionInvoicesService is the API client for the `/v1/subscription-invoices` endpoint
type SubscriptionInvoicesService service

// Get returns the subscription invoice with the given ID.
//
// https://docs.lemonsqueezy.com/api/subscription-invoices#retrieve-a-subscription-invoice
func (service *SubscriptionInvoicesService) Get(ctx context.Context, invoiceID string) (*SubscriptionInvoiceApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/subscription-invoices/"+invoiceID)
	if err != nil {
		return nil, response, err
	}

	invoice := new(SubscriptionInvoiceApiResponse)
	if err = json.Unmarshal(*response.Body, invoice); err != nil {
		return nil, response, err
	}

	return invoice, response, nil
}

// List a paginated list of subscription invoices.
//
// https://docs.lemonsqueezy.com/api/subscription-invoices#list-all-subscription-invoices
func (service *SubscriptionInvoicesService) List(ctx context.Context) (*SubscriptionInvoicesApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/subscription-invoices")
	if err != nil {
		return nil, response, err
	}

	invoices := new(SubscriptionInvoicesApiResponse)
	if err = json.Unmarshal(*response.Body, invoices); err != nil {
		return nil, response, err
	}

	return invoices, response, nil
}
