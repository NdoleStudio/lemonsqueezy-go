package stubs

// OrderGetResponse is a dummy response to the GET /v1/orders/:id endpoint
func OrderGetResponse() []byte {
	return []byte(`
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "self": "https://api.lemonsqueezy.com/v1/orders/1"
  },
  "data": {
    "type": "orders",
    "id": "1",
    "attributes": {
      "store_id": 1,
      "identifier": "104e18a2-d755-4d4b-80c4-a6c1dcbe1c10",
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
      "refunded": 0,
      "refunded_at": null,
      "subtotal_formatted": "$9.99",
      "discount_total_formatted": "$0.00",
      "tax_formatted": "$2.00",
      "total_formatted": "$11.99",
      "first_order_item": {
        "id": 1,
        "order_id": 1,
        "product_id": 1,
        "variant_id": 1,
        "product_name": "Test Limited Licencse for 2 years",
        "variant_name": "Default",
        "price": 1199,
        "created_at": "2021-08-17T09:45:53.000000Z",
        "updated_at": "2021-08-17T09:45:53.000000Z",
        "deleted_at": null,
        "test_mode": false
      },
      "created_at": "2021-08-17T09:45:53.000000Z",
      "updated_at": "2021-08-17T09:45:53.000000Z"
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

// OrdersListResponse is a dummy response to GET /v1/orders
func OrdersListResponse() []byte {
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
    "first": "https://api.lemonsqueezy.com/v1/orders?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=-createdAt",
    "last": "https://api.lemonsqueezy.com/v1/orders?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=-createdAt"
  },
  "data": [
    {
      "type": "orders",
      "id": "1",
      "attributes": {
        "store_id": 1,
        "identifier": "104e18a2-d755-4d4b-80c4-a6c1dcbe1c10",
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
        "refunded": 0,
        "refunded_at": null,
        "subtotal_formatted": "$9.99",
        "discount_total_formatted": "$0.00",
        "tax_formatted": "$2.00",
        "total_formatted": "$11.99",
        "first_order_item": {
          "id": 1,
          "order_id": 1,
          "product_id": 1,
          "variant_id": 1,
          "product_name": "Test Limited Licencse for 2 years",
          "variant_name": "Default",
          "price": 1199,
          "created_at": "2021-08-17T09:45:53.000000Z",
          "updated_at": "2021-08-17T09:45:53.000000Z",
          "deleted_at": null,
          "test_mode": false
        },
        "created_at": "2021-08-17T09:45:53.000000Z",
        "updated_at": "2021-08-17T09:45:53.000000Z"
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
  ]
}
`)
}
