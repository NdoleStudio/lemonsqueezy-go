package stubs

// WebhookRequest is a dummy webhook request
func WebhookRequest() []byte {
	return []byte(`
{
  "meta": {
    "event_name": "order_created"
  },
  "data": {
    "type": "orders",
    "id": "1",
    "attributes": {
      "store_id": 1,
      "identifier": "636f855c-1fb9-4c07-b75c-3a10afef010a",
      "order_number": 1,
      "user_name": "Darlene Daugherty",
      "user_email": "gernser@yahoo.com",
      "currency": "USD",
      "currency_rate": "1.0000",
      "subtotal": 999,
      "discount_total": 0,
      "tax": 200,
      "total": 1199,
      "subtotal_usd": 999,
      "discount_total_usd": 0,
      "tax_usd": 200,
      "total_usd": 1199,
      "tax_name": "VAT",
      "tax_rate": "20.00",
      "status": "paid",
      "status_formatted": "Paid",
      "refunded": false,
      "refunded_at": null,
      "created_at": "2021-08-11T13:47:29.000000Z",
      "updated_at": "2021-08-11T13:54:54.000000Z"
    },
    "relationships": {
      "store": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/orders/1/store",
          "self": "https://api.lemonsqueezy.com/v1/orders/1/relationships/store"
        }
      },
      "order-items": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/orders/1/order-items",
          "self": "https://api.lemonsqueezy.com/v1/orders/1/relationships/order-items"
        }
      },
      "subscriptions": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/orders/1/subscriptions",
          "self": "https://api.lemonsqueezy.com/v1/orders/1/relationships/subscriptions"
        }
      },
      "license-keys": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/orders/1/license-keys",
          "self": "https://api.lemonsqueezy.com/v1/orders/1/relationships/license-keys"
        }
      }
    },
    "links": {
      "self": "https://api.lemonsqueezy.com/v1/orders/1"
    }
  }
}
`)
}
