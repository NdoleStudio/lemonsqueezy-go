package stubs

// StoreGetResponse returns the response for the GET /v1/stores/:id endpoint
func StoreGetResponse() []byte {
	return []byte(`
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1"
  },
  "data": {
    "type": "stores",
    "id": "1",
    "attributes": {
      "name": "My Store",
      "slug": "my-store",
      "domain": "my-store.lemonsqueezy.com",
      "url": "https:\/\/my-store.lemonsqueezy.com",
      "avatar_url": "https:\/\/app.lemonsqueezy.com\/storage\/avatars\/stores\/1\/czTkMkDm4Vfb8PZehb5c29XFCm9JZyJx0AjEZP7s.png",
      "plan": "fresh",
      "country": "US",
      "country_nicename": "United States",
      "currency": "USD",
      "total_sales": 1,
      "total_revenue": 999,
      "thirty_day_sales": 0,
      "thirty_day_revenue": 0,
      "created_at": "2021-05-24T14:15:06.000000Z",
      "updated_at": "2021-06-15T10:03:14.000000Z"
    },
    "relationships": {
      "products": {
        "links": {
          "related": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/products",
          "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/relationships\/products"
        }
      },
      "orders": {
        "links": {
          "related": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/orders",
          "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/relationships\/orders"
        }
      },
      "subscriptions": {
        "links": {
          "related": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/subscriptions",
          "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/relationships\/subscriptions"
        }
      },
      "discounts": {
        "links": {
          "related": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/discounts",
          "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/relationships\/discounts"
        }
      },
      "license-keys": {
        "links": {
          "related": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/license-keys",
          "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/relationships\/license-keys"
        }
      }
    },
    "links": {
      "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1"
    }
  }
}
`)
}

// StoresListResponse is a dummy response for the GET /v1/stores endpoint
func StoresListResponse() []byte {
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
    "first": "https:\/\/api.lemonsqueezy.com\/v1\/stores?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=name",
    "last": "https:\/\/api.lemonsqueezy.com\/v1\/stores?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=name"
  },
  "data": [
    {
      "type": "stores",
      "id": "1",
      "attributes": {
        "name": "My Store",
        "slug": "my-store",
        "domain": "my-store.lemonsqueezy.com",
        "url": "https:\/\/my-store.lemonsqueezy.com",
        "avatar_url": "https:\/\/app.lemonsqueezy.com\/storage\/avatars\/stores\/1\/czTkMkDm4Vfb8PZehb5c29XFCm9JZyJx0AjEZP7s.png",
        "plan": "fresh",
        "country": "US",
        "country_nicename": "United States",
        "currency": "USD",
        "total_sales": 1,
        "total_revenue": 999,
        "thirty_day_sales": 0,
        "thirty_day_revenue": 0,
        "created_at": "2021-05-24T14:15:06.000000Z",
        "updated_at": "2021-06-15T10:03:14.000000Z"
      },
      "relationships": {
        "products": {
          "links": {
            "related": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/products",
            "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/relationships\/products"
          }
        },
        "orders": {
          "links": {
            "related": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/orders",
            "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/relationships\/orders"
          }
        },
        "subscriptions": {
          "links": {
            "related": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/subscriptions",
            "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/relationships\/subscriptions"
          }
        },
        "discounts": {
          "links": {
            "related": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/discounts",
            "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/relationships\/discounts"
          }
        },
        "license-keys": {
          "links": {
            "related": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/license-keys",
            "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/relationships\/license-keys"
          }
        }
      },
      "links": {
        "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1"
      }
    },
	{
      "type": "stores",
      "id": "2",
      "attributes": {
        "name": "My Store 2",
        "slug": "my-store-2",
        "domain": "my-store2.lemonsqueezy.com",
        "url": "https:\/\/my-store2.lemonsqueezy.com",
        "avatar_url": "https:\/\/app.lemonsqueezy.com\/storage\/avatars\/stores\/1\/czTkMkDm4Vfb8PZehb5c29XFCm9JZyJx0AjEZP7s.png",
        "plan": "fresh 2",
        "country": "US",
        "country_nicename": "United States",
        "currency": "USD",
        "total_sales": 2,
        "total_revenue": 99,
        "thirty_day_sales": 0,
        "thirty_day_revenue": 0,
        "created_at": "2021-05-24T14:15:06.000000Z",
        "updated_at": "2021-06-15T10:03:14.000000Z"
      },
      "relationships": {
        "products": {
          "links": {
            "related": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/products",
            "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/relationships\/products"
          }
        },
        "orders": {
          "links": {
            "related": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/orders",
            "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/relationships\/orders"
          }
        },
        "subscriptions": {
          "links": {
            "related": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/subscriptions",
            "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/relationships\/subscriptions"
          }
        },
        "discounts": {
          "links": {
            "related": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/discounts",
            "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/relationships\/discounts"
          }
        },
        "license-keys": {
          "links": {
            "related": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/license-keys",
            "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1\/relationships\/license-keys"
          }
        }
      },
      "links": {
        "self": "https:\/\/api.lemonsqueezy.com\/v1\/stores\/1"
      }
    }
  ]
}

`)
}
