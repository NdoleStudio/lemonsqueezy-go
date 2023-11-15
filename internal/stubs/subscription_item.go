package stubs

// SubscriptionItemGetResponse returns a dummy response to GET /v1/subscription-items/:id endpoint
func SubscriptionItemGetResponse() []byte {
	return []byte(`
{
	"jsonapi": {
		"version": "1.0"
	},
	"links": {
		"self": "https://api.lemonsqueezy.com/v1/subscription-item/1"
	},
	"data": {
		"type": "subscription-items",
		"id": "1",
		"attributes": {
			"subscription_id": 1,
			"price_id": 1,
			"quantity": 1,
			"is_usage_based": false,
			"created_at": "2023-07-18T12:16:24.000000Z",
			"updated_at": "2023-07-18T12:16:24.000000Z"
		},
		"relationships": {
			"subscription": {
				"links": {
					"related": "https://api.lemonsqueezy.com/v1/subscription-items/1/subscription",
					"self": "https://api.lemonsqueezy.com/v1/subscription-items/1/relationships/subscription"
				}
			},
			"price": {
				"links": {
					"related": "https://api.lemonsqueezy.com/v1/subscription-items/1/price",
					"self": "https://api.lemonsqueezy.com/v1/subscription-items/1/relationships/price"
				}
			},
			"usage-records": {
				"links": {
					"related": "https://api.lemonsqueezy.com/v1/subscription-items/1/usage-records",
					"self": "https://api.lemonsqueezy.com/v1/subscription-items/1/relationships/usage-records"
				}
			}
		},
		"links": {
			"self": "https://api.lemonsqueezy.com/v1/subscription-items/1"
		}
	}
}
`)
}

// SubscriptionItemUpdateResponse is a dummy response to the PATCH /v1/subscription-items/:id endpoint
func SubscriptionItemUpdateResponse() []byte {
	return []byte(`
{
	"jsonapi": {
		"version": "1.0"
	},
	"links": {
		"self": "https://api.lemonsqueezy.com/v1/subscription-item/1"
	},
	"data": {
		"type": "subscription-items",
		"id": "1",
		"attributes": {
		"subscription_id": 1,
		"price_id": 1,
		"quantity": 10,
		"is_usage_based": false,
		"created_at": "2023-07-18T12:16:24.000000Z",
		"updated_at": "2023-07-18T12:23:18.000000Z"
		},
		"relationships": {
			"subscription": {
				"links": {
					"related": "https://api.lemonsqueezy.com/v1/subscription-items/1/subscription",
					"self": "https://api.lemonsqueezy.com/v1/subscription-items/1/relationships/subscription"
				}
			},
			"price": {
				"links": {
					"related": "https://api.lemonsqueezy.com/v1/subscription-items/1/price",
					"self": "https://api.lemonsqueezy.com/v1/subscription-items/1/relationships/price"
				}
			},
			"usage-records": {
				"links": {
					"related": "https://api.lemonsqueezy.com/v1/subscription-items/1/usage-records",
					"self": "https://api.lemonsqueezy.com/v1/subscription-items/1/relationships/usage-records"
				}
			}
		},
		"links": {
			"self": "https://api.lemonsqueezy.com/v1/subscription-items/1"
		}
	}
}
`)
}

// SubscriptionItemsListResponse returns a dummy response to GET /v1/subscription-items endpoint
func SubscriptionItemsListResponse() []byte {
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
		"first": "https://api.lemonsqueezy.com/v1/subscription-items?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=-created_at",
		"last": "https://api.lemonsqueezy.com/v1/subscription-items?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=-created_at"
	},
	"data": [
		{
			"type": "subscription-items",
			"id": "1",
			"attributes": {
				"subscription_id": 1,
				"price_id": 1,
				"quantity": 1,
				"is_usage_based": false,
				"created_at": "2023-07-18T12:16:24.000000Z",
				"updated_at": "2023-07-18T12:16:24.000000Z"
			},
			"relationships": {
				"subscription": {
					"links": {
						"related": "https://api.lemonsqueezy.com/v1/subscription-items/1/subscription",
						"self": "https://api.lemonsqueezy.com/v1/subscription-items/1/relationships/subscription"
					}
				},
				"price": {
					"links": {
						"related": "https://api.lemonsqueezy.com/v1/subscription-items/1/price",
						"self": "https://api.lemonsqueezy.com/v1/subscription-items/1/relationships/price"
					}
				},
				"usage-records": {
					"links": {
						"related": "https://api.lemonsqueezy.com/v1/subscription-items/1/usage-records",
						"self": "https://api.lemonsqueezy.com/v1/subscription-items/1/relationships/usage-records"
					}
				}
			},
			"links": {
				"self": "https://api.lemonsqueezy.com/v1/subscription-items/1"
			}
		}
	]
}
`)
}

// SubscriptionItemCurrentUsageResponse returns a dummy response to GET /v1/subscription-items/:id/current-usage endpoint
func SubscriptionItemCurrentUsageResponse() []byte {
	return []byte(`
{
	"jsonapi": {
		"version": "1.0"
	},
	"meta": {
		"period_start": "2023-08-10T13:08:16.000000Z",
		"period_end": "2023-09-10T13:03:16.000000Z",
		"quantity": 5,
		"interval_unit": "month",
		"interval_quantity": 1
	}
}
`)
}
