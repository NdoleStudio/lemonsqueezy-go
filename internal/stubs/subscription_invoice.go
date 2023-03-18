package stubs

// SubscriptionInvoiceGetResponse is a dummy response to the GET /v1/subscription-invoice/:id endpoint
func SubscriptionInvoiceGetResponse() []byte {
	return []byte(`
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "self": "https://api.lemonsqueezy.com/v1/subscription-invoices/1"
  },
  "data": {
    "type": "subscription-invoices",
    "id": "1",
    "attributes": {
      "store_id": 1,
      "subscription_id": 1,
      "billing_reason": "initial",
      "card_brand": "visa",
      "card_last_four": "4242",
      "currency": "USD",
      "currency_rate": "1.00000000",
      "subtotal": 999,
      "discount_total": 0,
      "tax": 0,
      "total": 999,
      "subtotal_usd": 999,
      "discount_total_usd": 0,
      "tax_usd": 0,
      "total_usd": 999,
      "status": "paid",
      "status_formatted": "Paid",
      "refunded": 0,
      "refunded_at": null,
      "subtotal_formatted": "$9.99",
      "discount_total_formatted": "$0.00",
      "tax_formatted": "$0.00",
      "total_formatted": "$9.99",
      "urls": {
        "invoice_url": "https://app.lemonsqueezy.com/my-orders/.../subscription-invoice/..."
      },
      "created_at": "2023-01-18T12:16:24.000000Z",
      "updated_at": "2023-01-18T12:16:24.000000Z",
      "test_mode": false
    },
    "relationships": {
      "store": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/subscription-invoices/1/store",
          "self": "https://api.lemonsqueezy.com/v1/subscription-invoices/1/relationships/store"
        }
      },
      "subscription": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/subscription-invoices/1/subscription",
          "self": "https://api.lemonsqueezy.com/v1/subscription-invoices/1/relationships/subscription"
        }
      }
    },
    "links": {
      "self": "https://api.lemonsqueezy.com/v1/subscription-invoices/1"
    }
  }
}
`)
}

// SubscriptionInvoicesListResponse is a dummy response to GET /v1/subscription-invoice
func SubscriptionInvoicesListResponse() []byte {
	return []byte(`
{
  "meta": {
    "page": {
      "currentPage": 1,
      "from": 1,
      "lastPage": 1,
      "perPage": 10,
      "to": 10,
      "total": 10
    }
  },
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "first": "https://api.lemonsqueezy.com/v1/subscription-invoices?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=-createdAt",
    "last": "https://api.lemonsqueezy.com/v1/subscription-invoices?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=-createdAt"
  },
  "data": [
    {
      "type": "subscription-invoices",
      "id": "1",
      "attributes": {
        "store_id": 1,
        "subscription_id": 1,
        "billing_reason": "initial",
        "card_brand": "visa",
        "card_last_four": "4242",
        "currency": "USD",
        "currency_rate": "1.00000000",
        "subtotal": 999,
        "discount_total": 0,
        "tax": 0,
        "total": 999,
        "subtotal_usd": 999,
        "discount_total_usd": 0,
        "tax_usd": 0,
        "total_usd": 999,
        "status": "paid",
        "status_formatted": "Paid",
        "refunded": 0,
        "refunded_at": null,
        "subtotal_formatted": "$9.99",
        "discount_total_formatted": "$0.00",
        "tax_formatted": "$0.00",
        "total_formatted": "$9.99",
        "urls": {
          "invoice_url": "https://app.lemonsqueezy.com/my-orders/.../subscription-invoice/..."
        },
        "created_at": "2023-01-18T12:16:24.000000Z",
        "updated_at": "2023-01-18T12:16:24.000000Z",
        "test_mode": false
      },
      "relationships": {
        "store": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/subscription-invoices/1/store",
            "self": "https://api.lemonsqueezy.com/v1/subscription-invoices/1/relationships/store"
          }
        },
        "subscription": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/subscription-invoices/1/subscription",
            "self": "https://api.lemonsqueezy.com/v1/subscription-invoices/1/relationships/subscription"
          }
        }
      },
      "links": {
        "self": "https://api.lemonsqueezy.com/v1/subscription-invoices/1"
      }
    }
  ]
}
`)
}
