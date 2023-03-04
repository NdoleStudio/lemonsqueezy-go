package stubs

// OrderItemGetResponse is a dummy response to the GET /v1/order-items/:id endpoint
func OrderItemGetResponse() []byte {
	return []byte(`
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "self": "https://api.lemonsqueezy.com/v1/order-items/1"
  },
  "data": {
    "type": "order-items",
    "id": "1",
    "attributes": {
      "order_id": 1,
      "product_id": 1,
      "variant_id": 1,
      "product_name": "Example Product",
      "variant_name": "Example Variant",
      "price": 999,
      "created_at": "2021-05-24T14:15:06.000000Z",
      "updated_at": "2021-05-24T14:15:06.000000Z"
    },
    "relationships": {
      "order": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/order-items/1/order",
          "self": "https://api.lemonsqueezy.com/v1/order-items/1/relationships/order"
        }
      },
      "product": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/order-items/1/product",
          "self": "https://api.lemonsqueezy.com/v1/order-items/1/relationships/product"
        }
      },
      "variant": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/order-items/1/variant",
          "self": "https://api.lemonsqueezy.com/v1/order-items/1/relationships/variant"
        }
      }
    },
    "links": {
      "self": "https://api.lemonsqueezy.com/v1/order-items/1"
    }
  }
}
`)
}

// OrderItemsListResponse is a dummy response to GET /v1/order-items
func OrderItemsListResponse() []byte {
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
    "first": "https://api.lemonsqueezy.com/v1/order-items?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=id",
    "last": "https://api.lemonsqueezy.com/v1/order-items?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=id"
  },
  "data": [
    {
      "type": "order-items",
      "id": "1",
      "attributes": {
        "order_id": 1,
        "product_id": 1,
        "variant_id": 1,
        "product_name": "Example Product",
        "variant_name": "Example Variant",
        "price": 999,
        "created_at": "2021-05-24T14:15:06.000000Z",
        "updated_at": "2021-05-24T14:15:06.000000Z"
      },
      "relationships": {
        "order": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/order-items/1/order",
            "self": "https://api.lemonsqueezy.com/v1/order-items/1/relationships/order"
          }
        },
        "product": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/order-items/1/product",
            "self": "https://api.lemonsqueezy.com/v1/order-items/1/relationships/product"
          }
        },
        "variant": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/order-items/1/variant",
            "self": "https://api.lemonsqueezy.com/v1/order-items/1/relationships/variant"
          }
        }
      },
      "links": {
        "self": "https://api.lemonsqueezy.com/v1/order-items/1"
      }
    }
  ]
}
`)
}
