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
      "webhook_total": 0,
      "tax": 200,
      "total": 1199,
      "subtotal_usd": 999,
      "webhook_total_usd": 0,
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

// WebhookGetResponse is a dummy response to the GET /v1/webhooks/:id endpoint
func WebhookGetResponse() []byte {
	return []byte(`
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "self": "https://api.lemonsqueezy.com/v1/webhooks/1"
  },
  "data": {
    "type": "webhooks",
    "id": "1",
    "attributes": {
      "store_id": 1,
      "url": "https://mysite.com/webhooks/",
      "events": [
        "order_created",
        "order_refunded",
        "subscription_created",
        "subscription_updated",
        "subscription_expired"
      ],
      "last_sent_at": "2022-11-22T07:38:06.000000Z",
      "created_at": "2022-06-07T08:32:47.000000Z",
      "updated_at": "2022-06-07T08:41:37.000000Z",
      "test_mode": false
    },
    "relationships": {
      "store": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/webhooks/1/store",
          "self": "https://api.lemonsqueezy.com/v1/webhooks/1/relationships/store"
        }
      }
    },
    "links": {
      "self": "https://api.lemonsqueezy.com/v1/webhooks/1"
    }
  }
}
`)
}

// WebhooksListResponse is a dummy response to GET /v1/webhooks
func WebhooksListResponse() []byte {
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
    "first": "https://api.lemonsqueezy.com/v1/webhooks?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=createdAt",
    "last": "https://api.lemonsqueezy.com/v1/webhooks?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=createdAt"
  },
  "data": [
    {
      "type": "webhooks",
      "id": "1",
      "attributes": {
      	"store_id": 1,
        "url": "https://mysite.com/webhooks/",
        "events": [
          "order_created",
          "subscription_created",
          "subscription_updated",
          "subscription_expired"
        ],
        "last_sent_at": "2022-11-22T07:38:06.000000Z",
        "created_at": "2022-06-07T08:32:47.000000Z",
        "updated_at": "2022-06-07T08:41:37.000000Z",
        "test_mode": false
      },
      "relationships": {
        "store": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/webhooks/1/store",
            "self": "https://api.lemonsqueezy.com/v1/webhooks/1/relationships/store"
          }
        }
      },
      "links": {
        "self": "https://api.lemonsqueezy.com/v1/webhooks/1"
      }
    }
  ]
}
`)
}
