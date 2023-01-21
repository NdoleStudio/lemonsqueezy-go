package stubs

// CustomerGetResponse is a dummy response to the GET /v1/customers/:id endpoint
func CustomerGetResponse() []byte {
	return []byte(`
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "self": "https://api.lemonsqueezy.com/v1/customers/1"
  },
  "data": {
    "type": "customers",
    "id": "1",
    "attributes": {
      "store_id": 1,
      "name": "Luke Skywalker",
      "email": "lukeskywalker@example.com",
      "status": "subscribed",
      "city": null,
      "region": null,
      "country": "US",
      "total_revenue_currency": 84332,
      "mrr": 1999,
      "status_formatted": "Subscribed",
      "country_formatted": "United States",
      "total_revenue_currency_formatted": "$843.32",
      "mrr_formatted": "$19.99",
      "created_at": "2022-12-01T13:01:07.000000Z",
      "updated_at": "2022-12-09T09:05:21.000000Z",
      "test_mode": false
    },
    "relationships": {
      "store": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/customers/1/store",
          "self": "https://api.lemonsqueezy.com/v1/customers/1/relationships/store"
        }
      },
      "orders": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/customers/1/orders",
          "self": "https://api.lemonsqueezy.com/v1/customers/1/relationships/orders"
        }
      },
      "subscriptions": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/customers/1/subscriptions",
          "self": "https://api.lemonsqueezy.com/v1/customers/1/relationships/subscriptions"
        }
      },
      "license-keys": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/customers/1/license-keys",
          "self": "https://api.lemonsqueezy.com/v1/customers/1/relationships/license-keys"
        }
      }
    },
    "links": {
      "self": "https://api.lemonsqueezy.com/v1/customers/1"
    }
  }
}
`)
}

// CustomerListResponse is a dummy response to GET /v1/customers
func CustomerListResponse() []byte {
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
    "first": "https://api.lemonsqueezy.com/v1/customers?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=-createdAt",
    "last": "https://api.lemonsqueezy.com/v1/customers?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=-createdAt"
  },
  "data": [
    {
      "type": "customers",
      "id": "1",
      "attributes": {
        "store_id": 1,
        "name": "Luke Skywalker",
        "email": "lukeskywalker@example.com",
        "status": "subscribed",
        "city": null,
        "region": null,
        "country": "US",
        "total_revenue_currency": 84332,
        "mrr": 1999,
        "status_formatted": "Subscribed",
        "country_formatted": "United States",
        "total_revenue_currency_formatted": "$843.32",
        "mrr_formatted": "$19.99",
        "created_at": "2022-12-01T13:01:07.000000Z",
        "updated_at": "2022-12-09T09:05:21.000000Z",
        "test_mode": false
      },
      "relationships": {
        "store": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/customers/1/store",
            "self": "https://api.lemonsqueezy.com/v1/customers/1/relationships/store"
          }
        },
        "orders": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/customers/1/orders",
            "self": "https://api.lemonsqueezy.com/v1/customers/1/relationships/orders"
          }
        },
        "subscriptions": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/customers/1/subscriptions",
            "self": "https://api.lemonsqueezy.com/v1/customers/1/relationships/subscriptions"
          }
        },
        "license-keys": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/customers/1/license-keys",
            "self": "https://api.lemonsqueezy.com/v1/customers/1/relationships/license-keys"
          }
        }
      },
      "links": {
        "self": "https://api.lemonsqueezy.com/v1/customers/1"
      }
    },
	{
      "type": "customers",
      "id": "2",
      "attributes": {
        "store_id": 1,
        "name": "Luke Skywalker 1",
        "email": "lukeskywalker1@example.com",
        "status": "subscribed",
        "city": null,
        "region": null,
        "country": "US",
        "total_revenue_currency": 84332,
        "mrr": 1999,
        "status_formatted": "Subscribed",
        "country_formatted": "United States",
        "total_revenue_currency_formatted": "$843.32",
        "mrr_formatted": "$19.99",
        "created_at": "2022-12-01T13:01:07.000000Z",
        "updated_at": "2022-12-09T09:05:21.000000Z",
        "test_mode": false
      },
      "relationships": {
        "store": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/customers/1/store",
            "self": "https://api.lemonsqueezy.com/v1/customers/1/relationships/store"
          }
        },
        "orders": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/customers/1/orders",
            "self": "https://api.lemonsqueezy.com/v1/customers/1/relationships/orders"
          }
        },
        "subscriptions": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/customers/1/subscriptions",
            "self": "https://api.lemonsqueezy.com/v1/customers/1/relationships/subscriptions"
          }
        },
        "license-keys": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/customers/1/license-keys",
            "self": "https://api.lemonsqueezy.com/v1/customers/1/relationships/license-keys"
          }
        }
      },
      "links": {
        "self": "https://api.lemonsqueezy.com/v1/customers/1"
      }
    }
  ]
}
`)
}
