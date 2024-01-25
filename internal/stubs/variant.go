package stubs

// VariantGetResponse is a dummy response to the GET /v1/variants/:id endpoint
func VariantGetResponse() []byte {
	return []byte(`
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "self": "https://api.lemonsqueezy.com/v1/variants/1"
  },
  "data": {
    "type": "variants",
    "id": "1",
    "attributes": {
      "product_id": 1,
      "name": "Example Variant",
      "slug": "46beb127-a8a9-33e6-89b5-078505657239",
      "description": "<p>Lorem ipsum...</p>",
      "price": 999,
      "is_subscription": false,
      "interval": null,
      "interval_count": null,
      "has_free_trial": false,
      "trial_interval": "day",
      "trial_interval_count": 30,
      "pay_what_you_want": false,
      "min_price": 0,
      "suggested_price": 0,
      "has_license_keys": false,
      "license_activation_limit": 5,
      "is_license_limit_unlimited": false,
      "license_length_value": 1,
      "license_length_unit": "years",
      "is_license_length_unlimited": false,
      "sort": 1,
      "status": "published",
      "status_formatted": "Published",
      "created_at": "2021-05-24T14:15:06.000000Z",
      "updated_at": "2021-06-24T14:44:38.000000Z",
      "test_mode": false
    },
    "relationships": {
      "product": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/variants/1/product",
          "self": "https://api.lemonsqueezy.com/v1/variants/1/relationships/product"
        }
      }
    },
    "links": {
      "self": "https://api.lemonsqueezy.com/v1/variants/1"
    }
  }
}
`)
}

// VariantsListResponse is a dummy response to GET /v1/variants
func VariantsListResponse() []byte {
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
    "first": "https:\/\/api.lemonsqueezy.com\/v1\/variants?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=sort",
    "last": "https:\/\/api.lemonsqueezy.com\/v1\/variants?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=sort"
  },
  "data": [
    {
      "type": "variants",
      "id": "1",
      "attributes": {
        "product_id": 1,
        "name": "Example Variant",
        "slug": "46beb127-a8a9-33e6-89b5-078505657239",
        "description": "<p>Lorem ipsum...<\/p>",
        "price": 999,
        "is_subscription": false,
        "interval": null,
        "interval_count": null,
        "has_free_trial": false,
        "trial_interval": "day",
        "trial_interval_count": 30,
        "pay_what_you_want": false,
        "min_price": 0,
        "suggested_price": 0,
        "has_license_keys": false,
        "license_activation_limit": 5,
        "is_license_limit_unlimited": false,
        "license_length_value": 1,
        "license_length_unit": "years",
        "is_license_length_unlimited": false,
        "sort": 1,
        "status": "published",
        "status_formatted": "Published",
        "created_at": "2021-05-24T14:15:06.000000Z",
        "updated_at": "2021-06-24T14:44:38.000000Z",
        "test_mode": false
      },
      "relationships": {
        "product": {
          "links": {
            "related": "https:\/\/api.lemonsqueezy.com\/v1\/variants\/1\/product",
            "self": "https:\/\/api.lemonsqueezy.com\/v1\/variants\/1\/relationships\/product"
          }
        }
      },
      "links": {
        "self": "https:\/\/api.lemonsqueezy.com\/v1\/variants\/1"
      }
    },
	{
      "type": "variants",
      "id": "2",
      "attributes": {
        "product_id": 1,
        "name": "Example Variant",
        "slug": "46beb127-a8a9-33e6-89b5-078505657239",
        "description": "<p>Lorem ipsum...<\/p>",
        "price": 999,
        "is_subscription": false,
        "interval": null,
        "interval_count": null,
        "has_free_trial": false,
        "trial_interval": "day",
        "trial_interval_count": 30,
        "pay_what_you_want": false,
        "min_price": 0,
        "suggested_price": 0,
        "has_license_keys": false,
        "license_activation_limit": 5,
        "is_license_limit_unlimited": false,
        "license_length_value": 1,
        "license_length_unit": "years",
        "is_license_length_unlimited": false,
        "sort": 1,
        "status": "published",
        "status_formatted": "Published",
        "created_at": "2021-05-24T14:15:06.000000Z",
        "updated_at": "2021-06-24T14:44:38.000000Z",
        "test_mode": false
      },
      "relationships": {
        "product": {
          "links": {
            "related": "https:\/\/api.lemonsqueezy.com\/v1\/variants\/1\/product",
            "self": "https:\/\/api.lemonsqueezy.com\/v1\/variants\/1\/relationships\/product"
          }
        }
      },
      "links": {
        "self": "https:\/\/api.lemonsqueezy.com\/v1\/variants\/1"
      }
    }
  ]
}
`)
}
