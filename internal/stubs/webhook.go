package stubs

// WebhookRequestOrderCreated is a dummy webhook request
func WebhookRequestOrderCreated() []byte {
	return []byte(`
{
  "meta": {
    "event_name": "order_created",
	"custom_data": {
      "customer_id": 25
    }
  },
  "data": {
    "type": "orders",
    "id": "1",
    "attributes": {
      "store_id": 1,
      "customer_id": 1,
      "identifier": "89b36d62-4f5c-4353-853f-0c769d0535c8",
      "order_number": 1,
      "user_name": "Dan R",
      "user_email": "dan@lemonsqueezy.com",
      "currency": "EUR",
      "currency_rate": "1.08405",
      "subtotal": 1499,
      "discount_total": 0,
      "tax": 359,
      "total": 1859,
      "subtotal_usd": 1626,
      "discount_total_usd": 0,
      "tax_usd": 390,
      "total_usd": 2016,
      "tax_name": "ALV",
      "tax_rate": "24.00",
      "status": "paid",
      "status_formatted": "Paid",
      "refunded": null,
      "refunded_at": null,
      "subtotal_formatted": "€14.99",
      "discount_total_formatted": "€0.00",
      "tax_formatted": "€3.59",
      "total_formatted": "€18.59",
      "first_order_item": {
        "order_id": 1,
        "product_id": 1,
        "variant_id": 1,
        "product_name": "Product One",
        "variant_name": "Default",
        "price": 1500,
        "created_at": "2023-01-17T12:26:23.000000Z",
        "updated_at": "2023-01-17T12:26:23.000000Z",
        "test_mode": false
      },
      "urls": {
        "receipt": "https://app.lemonsqueezy.com/my-orders/89b36d62-4f5c-4353-853f-0c769d0535c8?signature=8847fff02e1bfb0c7c43ff1cdf1b1657a8eed2029413692663b86859208c9f42"
      },
      "created_at": "2023-01-17T12:26:23.000000Z",
      "updated_at": "2023-01-17T12:26:23.000000Z",
      "test_mode": false
    },
    "relationships": {
      "store": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/orders/1/store",
          "self": "https://api.lemonsqueezy.com/v1/orders/1/relationships/store"
        }
      },
      "customer": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/orders/1/customer",
          "self": "https://api.lemonsqueezy.com/v1/orders/1/relationships/customer"
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
      },
      "discount-redemptions": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/orders/1/discount-redemptions",
          "self": "https://api.lemonsqueezy.com/v1/orders/1/relationships/discount-redemptions"
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
