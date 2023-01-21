package stubs

// ProductGetResponse is a dummy response to the GET /v1/products/:id endpoint
func ProductGetResponse() []byte {
	return []byte(`
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "self": "https://api.lemonsqueezy.com/v1/products/1"
  },
  "data": {
    "type": "products",
    "id": "1",
    "attributes": {
      "store_id": 1,
      "name": "Example Product",
      "slug": "example-product",
      "description": "<p>Lorem ipsum...</p>",
      "status": "published",
      "status_formatted": "Published",
      "thumb_url": "https://app.lemonsqueezy.com/storage/media/1/1c40f227-aed5-4321-9ffc-62f37a06c9a0.jpg",
      "large_thumb_url": "https://app.lemonsqueezy.com/storage/media/1/1c40f227-aed5-4321-9ffc-62f37a06c9a0.jpg",
      "price": 999,
      "pay_what_you_want": false,
      "from_price": null,
      "to_price": null,
      "buy_now_url": "https://my-store.lemonsqueezy.com/checkout/buy/0a841cf5-4cc2-4bbb-ae5d-9529d97deec6",
      "price_formatted": "$9.99",
      "created_at": "2021-05-27T12:54:47.000000Z",
      "updated_at": "2021-07-14T11:25:24.000000Z"
    },
    "relationships": {
      "store": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/products/1/store",
          "self": "https://api.lemonsqueezy.com/v1/products/1/relationships/store"
        }
      },
      "variants": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/products/1/variants",
          "self": "https://api.lemonsqueezy.com/v1/products/1/relationships/variants"
        }
      }
    },
    "links": {
      "self": "https://api.lemonsqueezy.com/v1/products/1"
    }
  }
}
`)
}

// ProductsListResponse is a dummy response to GET /v1/products
func ProductsListResponse() []byte {
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
    "first": "https://api.lemonsqueezy.com/v1/products?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=name",
    "last": "https://api.lemonsqueezy.com/v1/products?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=name"
  },
  "data": [
    {
      "type": "products",
      "id": "1",
      "attributes": {
        "store_id": 1,
        "name": "Example Product",
        "slug": "example-product",
        "description": "<p>Lorem ipsum...</p>",
        "status": "published",
        "status_formatted": "Published",
        "thumb_url": "https://app.lemonsqueezy.com/storage/media/1/1c40f227-aed5-4321-9ffc-62f37a06c9a0.jpg",
        "large_thumb_url": "https://app.lemonsqueezy.com/storage/media/1/1c40f227-aed5-4321-9ffc-62f37a06c9a0.jpg",
        "price": 999,
        "pay_what_you_want": false,
        "from_price": null,
        "to_price": null,
        "buy_now_url": "https://my-store.lemonsqueezy.com/checkout/buy/0a841cf5-4cc2-4bbb-ae5d-9529d97deec6",
        "price_formatted": "$9.99",
        "created_at": "2021-05-27T12:54:47.000000Z",
        "updated_at": "2021-07-14T11:25:24.000000Z"
      },
      "relationships": {
        "store": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/products/1/store",
            "self": "https://api.lemonsqueezy.com/v1/products/1/relationships/store"
          }
        },
        "variants": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/products/1/variants",
            "self": "https://api.lemonsqueezy.com/v1/products/1/relationships/variants"
          }
        }
      },
      "links": {
        "self": "https://api.lemonsqueezy.com/v1/products/1"
      }
    },
	{
      "type": "products",
      "id": "2",
      "attributes": {
        "store_id": 2,
        "name": "Example Product",
        "slug": "example-product",
        "description": "<p>Lorem ipsum...</p>",
        "status": "published",
        "status_formatted": "Published",
        "thumb_url": "https://app.lemonsqueezy.com/storage/media/1/1c40f227-aed5-4321-9ffc-62f37a06c9a0.jpg",
        "large_thumb_url": "https://app.lemonsqueezy.com/storage/media/1/1c40f227-aed5-4321-9ffc-62f37a06c9a0.jpg",
        "price": 999,
        "pay_what_you_want": false,
        "from_price": null,
        "to_price": null,
        "buy_now_url": "https://my-store.lemonsqueezy.com/checkout/buy/0a841cf5-4cc2-4bbb-ae5d-9529d97deec6",
        "price_formatted": "$9.99",
        "created_at": "2021-05-27T12:54:47.000000Z",
        "updated_at": "2021-07-14T11:25:24.000000Z"
      },
      "relationships": {
        "store": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/products/1/store",
            "self": "https://api.lemonsqueezy.com/v1/products/1/relationships/store"
          }
        },
        "variants": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/products/1/variants",
            "self": "https://api.lemonsqueezy.com/v1/products/1/relationships/variants"
          }
        }
      },
      "links": {
        "self": "https://api.lemonsqueezy.com/v1/products/1"
      }
    }
  ]
}
`)
}
