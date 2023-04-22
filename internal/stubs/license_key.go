package stubs

// LicenseKeyGetResponse is a dummy response to the GET /v1/license-key/:id endpoint
func LicenseKeyGetResponse() []byte {
	return []byte(`
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "self": "https://api.lemonsqueezy.com/v1/license-keys/1"
  },
  "data": {
    "type": "license-keys",
    "id": "1",
    "attributes": {
      "store_id": 1,
      "customer_id": 1,
      "order_id": 1,
      "order_item_id": 1,
      "product_id": 1,
      "user_name": "Darlene Daugherty",
      "user_email": "gernser@yahoo.com",
      "key": "80e15db5-c796-436b-850c-8f9c98a48abe",
      "key_short": "XXXX-8f9c98a48abe",
      "activation_limit": 5,
      "instances_count": 0,
      "disabled": 0,
      "status": "inactive",
      "status_formatted": "Inactive",
      "expires_at": null,
      "created_at": "2021-05-24T14:15:07.000000Z",
      "updated_at": "2021-05-24T14:15:07.000000Z"
    },
    "relationships": {
      "store": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/license-keys/1/store",
          "self": "https://api.lemonsqueezy.com/v1/license-keys/1/relationships/store"
        }
      },
      "customer": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/license-keys/1/customer",
          "self": "https://api.lemonsqueezy.com/v1/license-keys/1/relationships/customer"
        }
      },
      "order": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/license-keys/1/order",
          "self": "https://api.lemonsqueezy.com/v1/license-keys/1/relationships/order"
        }
      },
      "order-item": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/license-keys/1/order-item",
          "self": "https://api.lemonsqueezy.com/v1/license-keys/1/relationships/order-item"
        }
      },
      "product": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/license-keys/1/product",
          "self": "https://api.lemonsqueezy.com/v1/license-keys/1/relationships/product"
        }
      },
      "license-key-instances": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/license-keys/1/license-key-instances",
          "self": "https://api.lemonsqueezy.com/v1/license-keys/1/relationships/license-key-instances"
        }
      }
    },
    "links": {
      "self": "https://api.lemonsqueezy.com/v1/license-keys/1"
    }
  }
}
`)
}

// LicenseKeysListResponse is a dummy response to GET /v1/license-key
func LicenseKeysListResponse() []byte {
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
    "first": "https://api.lemonsqueezy.com/v1/license-keys?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=id",
    "last": "https://api.lemonsqueezy.com/v1/license-keys?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=id"
  },
  "data": [
    {
      "type": "license-keys",
      "id": "1",
      "attributes": {
        "store_id": 1,
        "customer_id": 1,
        "order_id": 1,
        "order_item_id": 1,
        "product_id": 1,
        "user_name": "Darlene Daugherty",
        "user_email": "gernser@yahoo.com",
        "key": "80e15db5-c796-436b-850c-8f9c98a48abe",
        "key_short": "XXXX-8f9c98a48abe",
        "activation_limit": 5,
        "instances_count": 0,
        "disabled": 0,
        "status": "inactive",
        "status_formatted": "Inactive",
        "expires_at": null,
        "created_at": "2021-05-24T14:15:07.000000Z",
        "updated_at": "2021-05-24T14:15:07.000000Z"
      },
      "relationships": {
        "store": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/license-keys/1/store",
            "self": "https://api.lemonsqueezy.com/v1/license-keys/1/relationships/store"
          }
        },
        "customer": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/license-keys/1/customer",
            "self": "https://api.lemonsqueezy.com/v1/license-keys/1/relationships/customer"
          }
        },
        "order": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/license-keys/1/order",
            "self": "https://api.lemonsqueezy.com/v1/license-keys/1/relationships/order"
          }
        },
        "order-item": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/license-keys/1/order-item",
            "self": "https://api.lemonsqueezy.com/v1/license-keys/1/relationships/order-item"
          }
        },
        "product": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/license-keys/1/product",
            "self": "https://api.lemonsqueezy.com/v1/license-keys/1/relationships/product"
          }
        },
        "license-key-instances": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/license-keys/1/license-key-instances",
            "self": "https://api.lemonsqueezy.com/v1/license-keys/1/relationships/license-key-instances"
          }
        }
      },
      "links": {
        "self": "https://api.lemonsqueezy.com/v1/license-keys/1"
      }
    }
  ]
}
`)
}
