package stubs

// LicenseActivateResponse is a dummy response to the POST /v1/licenses/activate endpoint
func LicenseActivateResponse() []byte {
	return []byte(`
	{
		"activated": true,
		"error": null,
		"license_key": {
			"id": 1,
			"status": "active",
			"key": "38b1460a-5104-4067-a91d-77b872934d51",
			"activation_limit": 1,
			"activation_usage": 5,
			"created_at": "2021-01-24T14:15:07.000000Z",
			"expires_at": null
		},
		"instance": {
			"id": "47596ad9-a811-4ebf-ac8a-03fc7b6d2a17",
			"name": "Test",
			"created_at": "2021-04-06T14:15:07.000000Z"
		},
		"meta": {
			"store_id": 1,
			"order_id": 2,
			"order_item_id": 3,
			"product_id": 4,
			"product_name": "Example Product",
			"variant_id": 5,
			"variant_name": "Default",
			"customer_id": 6,
			"customer_name": "Luke Skywalker",
			"customer_email": "luke@skywalker.com"
		}
	}
`)
}

// LicenseValidateResponse is a dummy response to the POST /v1/licenses/validate endpoint
func LicenseValidateResponse() []byte {
	return []byte(`
	{
		"valid": true,
		"error": null,
		"license_key": {
			"id": 1,
			"status": "active",
			"key": "38b1460a-5104-4067-a91d-77b872934d51",
			"activation_limit": 1,
			"activation_usage": 5,
			"created_at": "2021-01-24T14:15:07.000000Z",
			"expires_at": "2022-01-24T14:15:07.000000Z"
		},
		"instance": {
			"id": "f90ec370-fd83-46a5-8bbd-44a241e78665",
			"name": "Test",
			"created_at": "2021-02-24T14:15:07.000000Z"
		},
		"meta": {
			"store_id": 1,
			"order_id": 2,
			"order_item_id": 3,
			"product_id": 4,
			"product_name": "Example Product",
			"variant_id": 5,
			"variant_name": "Default",
			"customer_id": 6,
			"customer_name": "Luke Skywalker",
			"customer_email": "luke@skywalker.com"
		}
	}
`)
}

// LicenseDeactivateResponse is a dummy response to the POST /v1/licenses/deactivate endpoint
func LicenseDeactivateResponse() []byte {
	return []byte(`
	{
		"deactivated": true,
		"error": null,
		"license_key": {
			"id": 1,
			"status": "inactive",
			"key": "38b1460a-5104-4067-a91d-77b872934d51",
			"activation_limit": 5,
			"activation_usage": 0,
			"created_at": "2021-01-24T14:15:07.000000Z",
			"expires_at": null
		},
		"meta": {
			"store_id": 1,
			"order_id": 2,
			"order_item_id": 3,
			"product_id": 4,
			"product_name": "Example Product",
			"variant_id": 5,
			"variant_name": "Default",
			"customer_id": 6,
			"customer_name": "Luke Skywalker",
			"customer_email": "luke@skywalker.com"
		}
	}
`)
}
