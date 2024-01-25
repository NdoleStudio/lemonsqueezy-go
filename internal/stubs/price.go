package stubs

// PriceGetResponse is a dummy response to the GET /v1/prices/:id endpoint
func PriceGetResponse() []byte {
	return []byte(`
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "self": "https://api.lemonsqueezy.com/v1/prices/1"
  },
  "data": {
    "type": "prices",
    "id": "1",
    "attributes": {
      "variant_id": 1,
      "category": "subscription",
      "scheme": "graduated",
      "usage_aggregation": null,
      "unit_price": 999,
      "unit_price_decimal": null,
      "setup_fee_enabled": false,
      "setup_fee": null,
      "package_size": 1,
      "tiers": [
        {
          "last_unit": 2,
          "unit_price": 10000,
          "unit_price_decimal": null,
          "fixed_fee": 1000
        },
        {
          "last_unit": "inf",
          "unit_price": 1000,
          "unit_price_decimal": null,
          "fixed_fee": 1000
        }
      ],
      "renewal_interval_unit": "year",
      "renewal_interval_quantity": 1,
      "trial_interval_unit": "day",
      "trial_interval_quantity": 30,
      "min_price": null,
      "suggested_price": null,
      "tax_code": "eservice",
      "created_at": "2023-05-24T14:15:06.000000Z",
      "updated_at": "2023-06-24T14:44:38.000000Z"
    },
    "relationships": {
      "variant": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/prices/1/variant",
          "self": "https://api.lemonsqueezy.com/v1/prices/1/relationships/variant"
        }
      }
    },
    "links": {
      "self": "https://api.lemonsqueezy.com/v1/prices/1"
    }
  }
}
`)
}

// PriceListResponse is a dummy response to GET /v1/prices
func PriceListResponse() []byte {
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
    "first": "https://api.lemonsqueezy.com/v1/prices?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=-created_at",
    "last": "https://api.lemonsqueezy.com/v1/prices?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=-created_at"
  },
  "data": [
    {
      "type": "prices",
      "id": "1",
      "attributes": {
        "variant_id": 1,
        "category": "subscription",
        "scheme": "graduated",
        "usage_aggregation": null,
        "unit_price": 999,
        "unit_price_decimal": null,
        "setup_fee_enabled": false,
        "setup_fee": null,
        "package_size": 1,
        "tiers": [
          {
            "last_unit": 2,
            "unit_price": 10000,
            "unit_price_decimal": null,
            "fixed_fee": 1000
          },
          {
            "last_unit": "inf",
            "unit_price": 1000,
            "unit_price_decimal": null,
            "fixed_fee": 1000
          }
        ],
        "renewal_interval_unit": "year",
        "renewal_interval_quantity": 1,
        "trial_interval_unit": "day",
        "trial_interval_quantity": 30,
        "min_price": null,
        "suggested_price": null,
        "tax_code": "eservice",
        "created_at": "2023-05-24T14:15:06.000000Z",
        "updated_at": "2023-06-24T14:44:38.000000Z"
      },
      "relationships": {
        "variant": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/prices/1/variant",
            "self": "https://api.lemonsqueezy.com/v1/prices/1/relationships/variant"
          }
        }
      },
      "links": {
        "self": "https://api.lemonsqueezy.com/v1/prices/1"
      }
    },
	{
      "type": "prices",
      "id": "2",
      "attributes": {
        "variant_id": 1,
        "category": "subscription",
        "scheme": "graduated",
        "usage_aggregation": null,
        "unit_price": 999,
        "unit_price_decimal": null,
        "setup_fee_enabled": false,
        "setup_fee": null,
        "package_size": 1,
        "tiers": [
          {
            "last_unit": 2,
            "unit_price": 10000,
            "unit_price_decimal": null,
            "fixed_fee": 1000
          },
          {
            "last_unit": "inf",
            "unit_price": 1000,
            "unit_price_decimal": null,
            "fixed_fee": 1000
          }
        ],
        "renewal_interval_unit": "year",
        "renewal_interval_quantity": 1,
        "trial_interval_unit": "day",
        "trial_interval_quantity": 30,
        "min_price": null,
        "suggested_price": null,
        "tax_code": "eservice",
        "created_at": "2023-05-24T14:15:06.000000Z",
        "updated_at": "2023-06-24T14:44:38.000000Z"
      },
      "relationships": {
        "variant": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/prices/1/variant",
            "self": "https://api.lemonsqueezy.com/v1/prices/1/relationships/variant"
          }
        }
      },
      "links": {
        "self": "https://api.lemonsqueezy.com/v1/prices/1"
      }
    }
  ]
}
`)
}
