package stubs

// CheckoutGetResponse is a dummy response to the GET /v1/checkouts/:id endpoint
func CheckoutGetResponse() []byte {
	return []byte(`
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "self": "https://api.lemonsqueezy.com/v1/checkouts/5e8b546c-c561-4a2c-a586-40c18bb2a195"
  },
  "data": {
    "type": "checkouts",
    "id": "5e8b546c-c561-4a2c-a586-40c18bb2a195",
    "attributes": {
      "store_id": 1,
      "variant_id": 1,
      "custom_price": 50000,
      "product_options": {
        "name": "",
        "description": "",
        "media": [],
        "redirect_url": "",
        "receipt_button_text": "",
        "receipt_link_url": "",
        "receipt_thank_you_note": "",
        "enabled_variants": [
          1
        ]
      },
      "checkout_options": {
        "embed": false,
        "media": true,
        "logo": true,
        "desc": true,
        "discount": true,
        "dark": false,
        "subscription_preview": true,
        "button_color": "#2DD272"
      },
      "checkout_data": {
        "email": "",
        "name": "",
        "billing_address": null,
        "tax_number": "",
        "discount_code": "",
        "custom": null
      },
      "preview": {
        "currency": "USD",
        "currency_rate": 1,
        "subtotal": 5000,
        "discount_total": 0,
        "tax": 0,
        "total": 5000,
        "subtotal_usd": 5000,
        "discount_total_usd": 0,
        "tax_usd": 0,
        "total_usd": 5000,
        "subtotal_formatted": "$50.00",
        "discount_total_formatted": "$50.00",
        "tax_formatted": "$0.00",
        "total_formatted": "$5.00"
      },
      "expires_at": "2022-10-30T15:20:06.000000Z",
      "created_at": "2022-10-14T13:03:37.000000Z",
      "updated_at": "2022-10-14T13:03:37.000000Z",
      "test_mode": false,
      "url": "https://my-store.lemonsqueezy.com/checkout/custom/5e8b546c-c561-4a2c-a586-40c18bb2a195?expires=1667143206&signature=8f7248ad2022ef1d4111752ae02d14f8d04332274861ca5c3589eb22b5086a5b"
    },
    "relationships": {
      "store": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/checkouts/5e8b546c-c561-4a2c-a586-40c18bb2a195/store",
          "self": "https://api.lemonsqueezy.com/v1/checkouts/5e8b546c-c561-4a2c-a586-40c18bb2a195/relationships/store"
        }
      },
      "variant": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/checkouts/5e8b546c-c561-4a2c-a586-40c18bb2a195/variant",
          "self": "https://api.lemonsqueezy.com/v1/checkouts/5e8b546c-c561-4a2c-a586-40c18bb2a195/relationships/variant"
        }
      }
    },
    "links": {
      "self": "https://api.lemonsqueezy.com/v1/checkouts/5e8b546c-c561-4a2c-a586-40c18bb2a195"
    }
  }
}
`)
}

// CheckoutListResponse is a dummy response to GET /v1/checkouts
func CheckoutListResponse() []byte {
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
    "first": "https://api.lemonsqueezy.com/v1/checkouts?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=-createdAt",
    "last": "https://api.lemonsqueezy.com/v1/checkouts?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=-createdAt"
  },
  "data": [
    {
      "type": "checkouts",
      "id": "ac470bd4-7c41-474d-b6cd-0f296f5be02a",
      "attributes": {
        "store_id": 1,
        "variant_id": 1,
        "custom_price": null,
        "product_options": {
          "name": "",
          "description": "",
          "media": [],
          "redirect_url": "",
          "receipt_button_text": "",
          "receipt_link_url": "",
          "receipt_thank_you_note": "",
          "enabled_variants": []
        },
        "checkout_options": {
          "embed": false,
          "media": true,
          "logo": true,
          "desc": true,
          "discount": true,
          "dark": false,
          "subscription_preview": true,
          "button_color": "#7047EB"
        },
        "checkout_data": {
          "email": "",
          "name": "",
          "billing_address": null,
          "tax_number": "",
          "discount_code": "",
          "custom": null
        },
        "expires_at": null,
        "created_at": "2022-10-14T12:36:27.000000Z",
        "updated_at": "2022-10-14T12:36:27.000000Z",
        "test_mode": false,
        "url": "https://my-store.lemonsqueezy.com/checkout/custom/ac470bd4-7c41-474d-b6cd-0f296f5be02a?signature=ee3fd20c5bac48fe5e976cb106e743bc3f6f330540f8003ab331d638e2ce3b8b"
      },
      "relationships": {
        "store": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/checkouts/ac470bd4-7c41-474d-b6cd-0f296f5be02a/store",
            "self": "https://api.lemonsqueezy.com/v1/checkouts/ac470bd4-7c41-474d-b6cd-0f296f5be02a/relationships/store"
          }
        },
        "variant": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/checkouts/ac470bd4-7c41-474d-b6cd-0f296f5be02a/variant",
            "self": "https://api.lemonsqueezy.com/v1/checkouts/ac470bd4-7c41-474d-b6cd-0f296f5be02a/relationships/variant"
          }
        }
      },
      "links": {
        "self": "https://api.lemonsqueezy.com/v1/checkouts/ac470bd4-7c41-474d-b6cd-0f296f5be02a"
      }
    }
  ]
}
`)
}
