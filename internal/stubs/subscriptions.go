package stubs

// SubscriptionGetResponse returns a dummy response to GET /v1/subscriptions/:id endpoint
func SubscriptionGetResponse() []byte {
	return []byte(`
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "self": "https://api.lemonsqueezy.com/v1/subscriptions/1"
  },
  "data": {
    "type": "subscriptions",
    "id": "1",
    "attributes": {
      "store_id": 1,
      "order_id": 1,
      "order_item_id": 1,
      "product_id": 1,
      "variant_id": 1,
      "product_name": "Example Product",
      "variant_name": "Example Variant",
      "user_name": "Darlene Daugherty",
      "user_email": "gernser@yahoo.com",
      "status": "active",
      "status_formatted": "Active",
      "pause": null,
      "cancelled": false,
      "trial_ends_at": null,
      "billing_anchor": 12,
      "urls": {
        "update_payment_method": "https://app.lemonsqueezy.com/my-orders/2ba92a4e-a00a-45d2-a128-16856ffa8cdf/subscription/8/update-payment-method?expires=1666869343&signature=9985e3bf9007840aeb3951412be475abc17439c449c1af3e56e08e45e1345413"
      },
      "renews_at": "2022-11-12T00:00:00.000000Z",
      "ends_at": null,
      "created_at": "2021-08-11T13:47:27.000000Z",
      "updated_at": "2021-08-11T13:54:19.000000Z",
      "test_mode": false
    },
    "relationships": {
      "store": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/store",
          "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/store"
        }
      },
      "order": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/order",
          "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/order"
        }
      },
      "order-item": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/order-item",
          "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/order-item"
        }
      },
      "product": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/product",
          "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/product"
        }
      },
      "variant": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/variant",
          "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/variant"
        }
      }
    },
    "links": {
      "self": "https://api.lemonsqueezy.com/v1/subscriptions/1"
    }
  }
}
`)
}

// SubscriptionUpdateResponse is a dummy response to the PATCH /v1/subscriptions/:id endpoint
func SubscriptionUpdateResponse() []byte {
	return []byte(`
{
  "jsonapi":{
    "version":"1.0"
  },
  "links":{
    "self":"https://api.lemonsqueezy.com/v1/subscriptions/1"
  },
  "data": {
    "type": "subscriptions",
    "id": "1",
    "attributes": {
      "store_id": 1,
      "order_id": 1,
      "order_item_id": 1,
      "product_id": 9,
      "variant_id": 11,
      "product_name": "New Plan Product",
      "variant_name": "New Plan Variant",
      "user_name": "Darlene Daugherty",
      "user_email": "gernser@yahoo.com",
      "status": "active",
      "status_formatted": "Active",
      "pause": null,
      "cancelled": false,
      "trial_ends_at": null,
      "billing_anchor": 29,
      "urls": {
        "update_payment_method": "https://app.lemonsqueezy.com/my-orders/2ba92a4e-a00a-45d2-a128-16856ffa8cdf/subscription/8/update-payment-method?expires=1666869343&signature=9985e3bf9007840aeb3951412be475abc17439c449c1af3e56e08e45e1345413"
      },
      "renews_at": "2022-11-12T00:00:00.000000Z",
      "ends_at": null,
      "created_at": "2021-08-11T13:47:27.000000Z",
      "updated_at": "2021-08-11T13:54:19.000000Z",
      "test_mode": false
    }
  }
}
`)
}

// SubscriptionsListResponse returns a list of subscription responses ordered by created at
func SubscriptionsListResponse() []byte {
	return []byte(`
{
   "meta":{
      "page":{
         "currentPage":1,
         "from":1,
         "lastPage":1,
         "perPage":10,
         "to":10,
         "total":10
      }
   },
   "jsonapi":{
      "version":"1.0"
   },
   "links":{
      "first":"https://api.lemonsqueezy.com/v1/subscriptions?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=-createdAt",
      "last":"https://api.lemonsqueezy.com/v1/subscriptions?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=-createdAt"
   },
   "data":[
      {
         "type":"subscriptions",
         "id":"1",
         "attributes":{
            "store_id":1,
            "order_id":1,
            "order_item_id":1,
            "product_id":1,
            "variant_id":1,
            "product_name":"Example Product",
            "variant_name":"Example Variant",
            "user_name":"Darlene Daugherty",
            "user_email":"gernser@yahoo.com",
            "status":"active",
            "status_formatted":"Active",
            "pause":null,
            "cancelled":false,
            "trial_ends_at":null,
            "billing_anchor":12,
            "urls":{
               "update_payment_method":"https://app.lemonsqueezy.com/my-orders/2ba92a4e-a00a-45d2-a128-16856ffa8cdf/subscription/8/update-payment-method?expires=1666869343&signature=9985e3bf9007840aeb3951412be475abc17439c449c1af3e56e08e45e1345413"
            },
            "renews_at":"2022-11-12T00:00:00.000000Z",
            "ends_at":null,
            "created_at":"2021-08-11T13:47:27.000000Z",
            "updated_at":"2021-08-11T13:54:19.000000Z",
            "test_mode":false
         },
         "relationships":{
            "store":{
               "links":{
                  "related":"https://api.lemonsqueezy.com/v1/subscriptions/1/store",
                  "self":"https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/store"
               }
            },
            "order":{
               "links":{
                  "related":"https://api.lemonsqueezy.com/v1/subscriptions/1/order",
                  "self":"https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/order"
               }
            },
            "order-item":{
               "links":{
                  "related":"https://api.lemonsqueezy.com/v1/subscriptions/1/order-item",
                  "self":"https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/order-item"
               }
            },
            "product":{
               "links":{
                  "related":"https://api.lemonsqueezy.com/v1/subscriptions/1/product",
                  "self":"https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/product"
               }
            },
            "variant":{
               "links":{
                  "related":"https://api.lemonsqueezy.com/v1/subscriptions/1/variant",
                  "self":"https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/variant"
               }
            }
         },
         "links":{
            "self":"https://api.lemonsqueezy.com/v1/subscriptions/1"
         }
      },
      {
         "type":"subscriptions",
         "id":"2",
         "attributes":{
            "store_id":2,
            "order_id":2,
            "order_item_id":2,
            "product_id":2,
            "variant_id":2,
            "product_name":"Example Product 2",
            "variant_name":"Example Variant 2",
            "user_name":"Darlene Daugherty 2",
            "user_email":"gernser2@yahoo.com",
            "status":"active",
            "status_formatted":"Active",
            "pause":null,
            "cancelled":false,
            "trial_ends_at":null,
            "billing_anchor":13,
            "urls":{
               "update_payment_method":"https://app.lemonsqueezy.com/my-orders/2ba92a4e-a00a-45d2-a128-16856ffa8cdf/subscription/8/update-payment-method?expires=1666869343&signature=9985e3bf9007840aeb3951412be475abc17439c449c1af3e56e08e45e1345413"
            },
            "renews_at":"2022-11-12T00:00:00.000000Z",
            "ends_at":null,
            "created_at":"2021-08-11T13:47:27.000000Z",
            "updated_at":"2021-08-11T13:54:19.000000Z",
            "test_mode":false
         },
         "relationships":{
            "store":{
               "links":{
                  "related":"https://api.lemonsqueezy.com/v1/subscriptions/1/store",
                  "self":"https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/store"
               }
            },
            "order":{
               "links":{
                  "related":"https://api.lemonsqueezy.com/v1/subscriptions/1/order",
                  "self":"https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/order"
               }
            },
            "order-item":{
               "links":{
                  "related":"https://api.lemonsqueezy.com/v1/subscriptions/1/order-item",
                  "self":"https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/order-item"
               }
            },
            "product":{
               "links":{
                  "related":"https://api.lemonsqueezy.com/v1/subscriptions/1/product",
                  "self":"https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/product"
               }
            },
            "variant":{
               "links":{
                  "related":"https://api.lemonsqueezy.com/v1/subscriptions/1/variant",
                  "self":"https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/variant"
               }
            }
         },
         "links":{
            "self":"https://api.lemonsqueezy.com/v1/subscriptions/1"
         }
      }
   ]
}
`)
}

// SubscriptionCancelResponse returns a dummy response to DELETE /v1/subscriptions/:id endpoint
func SubscriptionCancelResponse() []byte {
	return []byte(`
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "self": "https://api.lemonsqueezy.com/v1/subscriptions/1"
  },
  "data": {
    "type": "subscriptions",
    "id": "1",
    "attributes": {
      "store_id": 1,
      "order_id": 1,
      "order_item_id": 1,
      "product_id": 1,
      "variant_id": 1,
      "product_name": "Example Product",
      "variant_name": "Example Variant",
      "user_name": "Darlene Daugherty",
      "user_email": "gernser@yahoo.com",
      "status": "cancelled",
      "status_formatted": "Cancelled",
      "pause": null,
      "cancelled": true,
      "trial_ends_at": null,
      "billing_anchor": 12,
      "urls": {
        "update_payment_method": "https://app.lemonsqueezy.com/my-orders/2ba92a4e-a00a-45d2-a128-16856ffa8cdf/subscription/8/update-payment-method?expires=1666869343&signature=9985e3bf9007840aeb3951412be475abc17439c449c1af3e56e08e45e1345413"
      },
      "renews_at": "2022-11-12T00:00:00.000000Z",
      "ends_at": "2022-11-12T00:00:00.000000Z",
      "created_at": "2021-08-11T13:47:27.000000Z",
      "updated_at": "2021-08-11T13:54:19.000000Z",
      "test_mode": false
    },
    "relationships": {
      "store": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/store",
          "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/store"
        }
      },
      "order": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/order",
          "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/order"
        }
      },
      "order-item": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/order-item",
          "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/order-item"
        }
      },
      "product": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/product",
          "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/product"
        }
      },
      "variant": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/variant",
          "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/variant"
        }
      }
    },
    "links": {
      "self": "https://api.lemonsqueezy.com/v1/subscriptions/1"
    }
  }
}
`)
}
