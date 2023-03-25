package stubs

// DiscountGetResponse is a dummy response to the GET /v1/discounts/:id endpoint
func DiscountGetResponse() []byte {
	return []byte(`
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "self": "https://api.lemonsqueezy.com/v1/discounts/1"
  },
  "data": {
    "type": "discounts",
    "id": "1",
    "attributes": {
      "store_id": 1,
      "name": "10% Off",
      "code": "10PERCENT",
      "amount": 10,
      "amount_type": "percent",
      "is_limited_to_products": false,
      "is_limited_redemptions": false,
      "max_redemptions": 0,
      "starts_at": null,
      "expires_at": null,
      "duration": "once",
      "duration_in_months": 1,
      "status": "published",
      "status_formatted": "Published",
      "created_at": "2021-05-24T14:15:06.000000Z",
      "updated_at": "2021-05-24T14:15:06.000000Z"
    },
    "relationships": {
      "store": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/discounts/1/store",
          "self": "https://api.lemonsqueezy.com/v1/discounts/1/relationships/store"
        }
      },
      "variants": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/discounts/1/variants",
          "self": "https://api.lemonsqueezy.com/v1/discounts/1/relationships/variants"
        }
      },
      "discount-redemptions": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/discounts/1/discount-redemptions",
          "self": "https://api.lemonsqueezy.com/v1/discounts/1/relationships/discount-redemptions"
        }
      }
    },
    "links": {
      "self": "https://api.lemonsqueezy.com/v1/discounts/1"
    }
  }
}
`)
}

// DiscountsListResponse is a dummy response to GET /v1/discounts
func DiscountsListResponse() []byte {
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
    "first": "https://api.lemonsqueezy.com/v1/discounts?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=createdAt",
    "last": "https://api.lemonsqueezy.com/v1/discounts?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=createdAt"
  },
  "data": [
    {
      "type": "discounts",
      "id": "1",
      "attributes": {
        "store_id": 1,
        "name": "10%",
        "code": "10PERC",
        "amount": 10,
        "amount_type": "percent",
        "is_limited_to_products": false,
        "is_limited_redemptions": false,
        "max_redemptions": 0,
        "starts_at": null,
        "expires_at": null,
        "duration": "once",
        "duration_in_months": 1,
        "status": "published",
        "status_formatted": "Published",
        "created_at": "2021-05-24T14:15:06.000000Z",
        "updated_at": "2021-05-24T14:15:06.000000Z"
      },
      "relationships": {
        "store": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/discounts/1/store",
            "self": "https://api.lemonsqueezy.com/v1/discounts/1/relationships/store"
          }
        },
        "variants": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/discounts/1/variants",
            "self": "https://api.lemonsqueezy.com/v1/discounts/1/relationships/variants"
          }
        },
        "discount-redemptions": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/discounts/1/discount-redemptions",
            "self": "https://api.lemonsqueezy.com/v1/discounts/1/relationships/discount-redemptions"
          }
        }
      },
      "links": {
        "self": "https://api.lemonsqueezy.com/v1/discounts/1"
      }
    }
  ]
}
`)
}
