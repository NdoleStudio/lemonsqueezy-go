# lemonsqueezy-go

[![Build](https://github.com/NdoleStudio/lemonsqueezy-go/actions/workflows/main.yml/badge.svg)](https://github.com/NdoleStudio/lemonsqueezy-go/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/NdoleStudio/lemonsqueezy-go/branch/main/graph/badge.svg)](https://codecov.io/gh/NdoleStudio/lemonsqueezy-go)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/NdoleStudio/lemonsqueezy-go/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/NdoleStudio/lemonsqueezy-go/?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/NdoleStudio/lemonsqueezy-go)](https://goreportcard.com/report/github.com/NdoleStudio/lemonsqueezy-go)
[![GitHub contributors](https://img.shields.io/github/contributors/NdoleStudio/lemonsqueezy-go)](https://github.com/NdoleStudio/lemonsqueezy-go/graphs/contributors)
[![GitHub license](https://img.shields.io/github/license/NdoleStudio/lemonsqueezy-go?color=brightgreen)](https://github.com/NdoleStudio/lemonsqueezy-go/blob/master/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NdoleStudio/lemonsqueezy-go)](https://pkg.go.dev/github.com/NdoleStudio/lemonsqueezy-go)

This package provides a go API client for the lemonsqueezy API

## Installation

`lemonsqueezy-go` is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/NdoleStudio/lemonsqueezy-go
```

Alternatively the same can be achieved if you use `import` in a package:

```go
import "github.com/NdoleStudio/lemonsqueezy-go"
```

## Implemented

- **Users**
  - `GET /v1/users/me`: Retrieves the currently authenticated user.
- **Stores**
  - `GET /v1/stores/:id`: Retrieve a store
  - `GET /v1/stores`: List all stores
- **Customers**
  - `GET /v1/customers/:id`: Retrieve a customer
  - `GET /v1/customers`: List all customers
- **Products**
  - `GET /v1/products/:id`: Retrieve a product
  - `GET /v1/products`: List all products
- **Variants**
  - `GET /v1/variants/:id`: Retrieve a variant
  - `GET /v1/variants`: List all variants
- **Prices**
  - `GET /v1/prices/:id`: Retrieve a price
  - `GET /v1/prices`: List all prices
- **Files**
  - `GET /v1/files/:id`: Retrieve a file
  - `GET /v1/files`: List all files
- **Orders**
  - `GET /v1/orders/:id`: Retrieve an order
  - `GET /v1/orders`: List all orders
- **Order Items**
  - `GET /v1/order-items/:id`: Retrieve an order item
  - `GET /v1/order-items`: List all order items
- **Subscriptions**
  - `PATCH /v1/subscriptions/:id`: Update a subscription
  - `GET /v1/subscriptions/:id`: Retrieve a subscription
  - `GET /v1/subscriptions`: List all subscriptions
  - `DELETE /v1/subscriptions/{id}`: Cancel an active subscription
- **Subscription Invoices**
  - `GET /v1/subscription-invoices/:id`: Retrieve a subscription invoice
  - `GET /v1/subscription-invoices`: List all subscription invoices
- **Subscription Items**
  - `GET /v1/subscription-items/:id`: Retrieve a subscription item
  - `PATCH /v1/subscription-items/:id`: Update a subscription item
  - `GET /v1/subscription-items`: List all subscription items
  - `GET /v1/subscription-items/:id/current-usage`: Retrieve a subscription item's current usage
- **Discounts**
  - `POST /v1/discounts`: Create a discount
  - `GET /v1/discounts/:id`: Retrieve a discount
  - `DELETE /v1/discounts/:id`: Delete a discount
  - `GET /v1/discounts`: List all discounts
- **Discount Redemptions**
  - `GET /v1/discount-redemptions/:id`: Retrieve a discount redemption
  - `GET /v1/discount-redemptions`: List all discount redemptions
- **License Keys**
  - `GET /v1/license-keys/:id`: Retrieve a license key
  - `GET /v1/license-keys`: List all license keys
- **License Key Instances**
  - `GET /v1/license-key-instances/:id`: Retrieve a license key instance
  - `GET /v1/license-key-instances`: List all license keys instance
- **Licenses**
  - `POST /v1/licenses/validate`: Validate a license
  - `POST /v1/licenses/activate`: Activate a license
  - `POST /v1/licenses/deactivate`: Deactivate a license
- **Checkouts**
  - `POST /v1/checkouts`: Create a checkout
  - `GET /v1/checkouts/:id`: Retrieve a checkout
  - `GET /v1/checkouts`: List all checkouts
- **Webhooks**
  - `PATCH /v1/webhooks/:id`: Update a webhook
  - `GET /v1/webhooks/:id`: Retrieve a webhook
  - `GET /v1/webhooks`: List all webhooks
  - `DELETE /v1/webhooks/{id}`: Update a webhook
  - `Verify`: Verify that webhook requests are coming from Lemon Squeezy

## Usage

### Initializing the Client

An instance of the client can be created using `New()`.

```go
package main

import (
    "github.com/NdoleStudio/lemonsqueezy-go"
)

func main() {
    client := lemonsqueezy.New(lemonsqueezy.WithAPIKey(""))
}
```

### Error handling

All API calls return an `error` as the last return object. All successful calls will return a `nil` error.

```go
subscription, response, err := client.Subscriptions.Get(context.Background(), "1")
if err != nil {
    //handle error
}
```

### WebHooks

Webhooks allow Lemon Squeezy to send new data to your application when certain events occur inside your store.
You can use the sample code below as inspiration for a basic `http.HandlerFunc` which processes webhook events on your server.

```go
func WebhookHandler(_ http.ResponseWriter, req *http.Request) {

	// 1. Authenticate the webhook request from Lemon Squeezy using the `X-Signature` header

	// 2. Process the payload if the request is authenticated
	eventName := req.Header.Get("X-Event-Name")
	payload, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	switch eventName {
	case lemonsqueezy.WebhookEventSubscriptionCreated:
		var request lemonsqueezy.WebhookRequestSubscription
		if err = json.Unmarshal(payload, &request); err != nil {
			log.Fatal(err)
		}
		// handle subscription_created request
	case lemonsqueezy.WebhookEventOrderCreated:
		var request lemonsqueezy.WebhookRequestOrder
		if err = json.Unmarshal(payload, &request); err != nil {
			log.Fatal(err)
		}
		// handle order_created request
	default:
		log.Fatal(fmt.Sprintf("invalid event [%s] received with request [%s]", eventName, string(payload)))
	}
}
```


## Testing

You can run the unit tests for this client from the root directory using the command below:

```bash
go test -v
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
