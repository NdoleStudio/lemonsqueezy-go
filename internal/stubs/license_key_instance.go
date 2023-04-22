package stubs

// LicenseKeyInstanceGetResponse is a dummy response to the GET /v1/license-key-instances/:id endpoint
func LicenseKeyInstanceGetResponse() []byte {
	return []byte(`
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "self": "https://api.lemonsqueezy.com/v1/license-key-instances/1"
  },
  "data": {
    "type": "license-key-instances",
    "id": "1",
    "attributes": {
      "license_key_id": 1,
      "identifier": "f70a79fa-6054-433e-9c1b-6075344292e4",
      "name": "example.com",
      "created_at": "2022-11-14T11:45:39.000000Z",
      "updated_at": "2022-11-14T11:45:39.000000Z"
    },
    "relationships": {
      "license-key": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/license-key-instances/1/license-key",
          "self": "https://api.lemonsqueezy.com/v1/license-key-instances/1/relationships/license-key"
        }
      }
    },
    "links": {
      "self": "https://api.lemonsqueezy.com/v1/license-key-instances/1"
    }
  }
}
`)
}

// LicenseKeyInstancesListResponse is a dummy response to GET /v1/license-key-instances
func LicenseKeyInstancesListResponse() []byte {
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
    "first": "https://api.lemonsqueezy.com/v1/license-key-instances?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=id",
    "last": "https://api.lemonsqueezy.com/v1/license-key-instances?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=id"
  },
  "data": [
    {
      "type": "license-key-instances",
      "id": "1",
      "attributes": {
        "license_key_id": 1,
        "identifier": "f70a79fa-6054-433e-9c1b-6075344292e4",
        "name": "example.com",
        "created_at": "2022-11-14T11:45:39.000000Z",
        "updated_at": "2022-11-14T11:45:39.000000Z"
      },
      "relationships": {
        "license-key": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/license-key-instances/1/license-key",
            "self": "https://api.lemonsqueezy.com/v1/license-key-instances/1/relationships/license-key"
          }
        }
      },
      "links": {
        "self": "https://api.lemonsqueezy.com/v1/license-key-instances/1"
      }
    }
  ]
}
`)
}
