package stubs

// DiscountRedemptionGetResponse is a dummy response to the GET /v1/discount-redemption/:id endpoint
func DiscountRedemptionGetResponse() []byte {
	return []byte(`
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "self": "https://api.lemonsqueezy.com/v1/discount-redemptions/1"
  },
  "data": {
    "type": "discount-redemptions",
    "id": "1",
    "attributes": {
      "discount_id": 1,
      "order_id": 1,
      "discount_name": "10%",
      "discount_code": "10PERC",
      "discount_amount": 10,
      "discount_amount_type": "percent",
      "amount": 999,
      "created_at": "2023-02-07T10:30:01.000000Z",
      "updated_at": "2023-02-07T10:30:01.000000Z"
    },
    "relationships": {
      "discount": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/discount-redemptions/1/discount",
          "self": "https://api.lemonsqueezy.com/v1/discount-redemptions/1/relationships/discount"
        }
      },
      "order": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/discount-redemptions/1/order",
          "self": "https://api.lemonsqueezy.com/v1/discount-redemptions/1/relationships/order"
        }
      }
    },
    "links": {
      "self": "https://api.lemonsqueezy.com/v1/discount-redemptions/1"
    }
  }
}
`)
}

// DiscountRedemptionsListResponse is a dummy response to GET /v1/discount-redemption
func DiscountRedemptionsListResponse() []byte {
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
    "first": "https://api.lemonsqueezy.com/v1/discount-redemptions?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=-createdAt",
    "last": "https://api.lemonsqueezy.com/v1/discount-redemptions?page%5Bnumber%5D=1339&page%5Bsize%5D=10&sort=-createdAt",
    "next": "https://api.lemonsqueezy.com/v1/discount-redemptions?page%5Bnumber%5D=2&page%5Bsize%5D=10&sort=-createdAt"
  },
  "data": [
    {
      "type": "discount-redemptions",
      "id": "1",
      "attributes": {
        "discount_id": 1,
        "order_id": 1,
        "discount_name": "10%",
        "discount_code": "10PERC",
        "discount_amount": 10,
        "discount_amount_type": "percent",
        "amount": 999,
        "created_at": "2023-02-07T10:30:01.000000Z",
        "updated_at": "2023-02-07T10:30:01.000000Z"
      },
      "relationships": {
        "discount": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/discount-redemptions/1/discount",
            "self": "https://api.lemonsqueezy.com/v1/discount-redemptions/1/relationships/discount"
          }
        },
        "order": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/discount-redemptions/1/order",
            "self": "https://api.lemonsqueezy.com/v1/discount-redemptions/1/relationships/order"
          }
        }
      },
      "links": {
        "self": "https://api.lemonsqueezy.com/v1/discount-redemptions/1"
      }
    }
  ]
}
`)
}
