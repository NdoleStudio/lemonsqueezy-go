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
      "subscription_id": 1,
      "customer_id": 1,
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
      "card_brand": "visa",
      "card_last_four": "42424",
      "pause": null,
      "cancelled": false,
      "trial_ends_at": null,
      "billing_anchor": 12,
      "first_subscription_item": {
        "id": 1,
        "subscription_id": 1,
        "price_id": 1,
        "quantity": 5,
        "created_at": "2021-08-11T13:47:28.000000Z",
        "updated_at": "2021-08-11T13:47:28.000000Z"
      },
      "urls": {
        "update_payment_method": "https://my-store.lemonsqueezy.com/subscription/1/payment-details?expires=1666869343&signature=9985e3bf9007840aeb3951412be475abc17439c449c1af3e56e08e45e1345413",
        "customer_portal": "https://my-store.lemonsqueezy.com/billing?expires=1666869343&signature=82ae290ceac8edd4190c82825dd73a8743346d894a8ddbc4898b97eb96d105a5"
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
      "customer": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/customer",
          "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/customer"
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
      },
      "subscription-items": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/subscription-items",
          "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/subscription-items"
        }
      },
      "subscription-invoices": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/subscription-invoices",
          "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/subscription-invoices"
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
      "customer_id": 1,
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
      "first_subscription_item": {
        "id": 1,
        "subscription_id": 1,
        "price_id": 1,
        "quantity": 5,
        "created_at": "2021-08-11T13:47:28.000000Z",
        "updated_at": "2021-08-11T13:47:28.000000Z"
      },
      "urls": {
        "update_payment_method": "https://my-store.lemonsqueezy.com/subscription/1/payment-details?expires=1666869343&signature=9985e3bf9007840aeb3951412be475abc17439c449c1af3e56e08e45e1345413",
        "customer_portal": "https://my-store.lemonsqueezy.com/billing?expires=1666869343&signature=82ae290ceac8edd4190c82825dd73a8743346d894a8ddbc4898b97eb96d105a5"
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
            "customer_id": 1,
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
            "first_subscription_item": {
              "id": 1,
              "subscription_id": 1,
              "price_id": 1,
              "quantity": 5,
              "created_at": "2021-08-11T13:47:28.000000Z",
              "updated_at": "2021-08-11T13:47:28.000000Z"
            },
            "urls":{
              "update_payment_method": "https://my-store.lemonsqueezy.com/subscription/1/payment-details?expires=1666869343&signature=9985e3bf9007840aeb3951412be475abc17439c449c1af3e56e08e45e1345413",
              "customer_portal": "https://my-store.lemonsqueezy.com/billing?expires=1666869343&signature=82ae290ceac8edd4190c82825dd73a8743346d894a8ddbc4898b97eb96d105a5"
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
            "customer": {
              "links": {
                "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/customer",
                "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/customer"
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
            },
            "subscription-items": {
              "links": {
                "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/subscription-items",
                "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/subscription-items"
              }
            },
            "subscription-invoices": {
              "links": {
                "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/subscription-invoices",
                "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/subscription-invoices"
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
            "customer_id": 2,
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
            "first_subscription_item": {
              "id": 2,
              "subscription_id": 2,
              "price_id": 2,
              "quantity": 5,
              "created_at": "2021-08-11T13:47:28.000000Z",
              "updated_at": "2021-08-11T13:47:28.000000Z"
            },
            "urls":{
              "update_payment_method": "https://my-store.lemonsqueezy.com/subscription/2/payment-details?expires=1666869343&signature=9985e3bf9007840aeb3951412be475abc17439c449c1af3e56e08e45e1345413",
              "customer_portal": "https://my-store.lemonsqueezy.com/billing?expires=1666869343&signature=82ae290ceac8edd4190c82825dd73a8743346d894a8ddbc4898b97eb96d105a5"
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
            "customer": {
              "links": {
                "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/customer",
                "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/customer"
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
            },
            "subscription-items": {
              "links": {
                "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/subscription-items",
                "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/subscription-items"
              }
            },
            "subscription-invoices": {
              "links": {
                "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/subscription-invoices",
                "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/subscription-invoices"
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
      "subscription_id": 1,
      "customer_id": 1,
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
      "first_subscription_item": {
        "id": 1,
        "subscription_id": 1,
        "price_id": 1,
        "quantity": 5,
        "created_at": "2021-08-11T13:47:28.000000Z",
        "updated_at": "2021-08-11T13:47:28.000000Z"
      },
      "urls": {
        "update_payment_method": "https://my-store.lemonsqueezy.com/subscription/1/payment-details?expires=1666869343&signature=9985e3bf9007840aeb3951412be475abc17439c449c1af3e56e08e45e1345413",
        "customer_portal": "https://my-store.lemonsqueezy.com/billing?expires=1666869343&signature=82ae290ceac8edd4190c82825dd73a8743346d894a8ddbc4898b97eb96d105a5"
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
      "customer": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/customer",
          "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/customer"
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
      },
      "subscription-items": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/subscription-items",
          "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/subscription-items"
        }
      },
      "subscription-invoices": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/subscriptions/1/subscription-invoices",
          "self": "https://api.lemonsqueezy.com/v1/subscriptions/1/relationships/subscription-invoices"
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
