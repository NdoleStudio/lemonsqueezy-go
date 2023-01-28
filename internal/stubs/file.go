package stubs

// FileGetResponse is a dummy response to the GET /v1/files/:id endpoint
func FileGetResponse() []byte {
	return []byte(`
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "self": "https://api.lemonsqueezy.com/v1/files/1"
  },
  "data": {
    "type": "files",
    "id": "1",
    "attributes": {
      "variant_id": 168,
      "identifier": "6dce5ba7-76f2-481f-ad1e-9c2bec6eb0e2",
      "name": "my_product.zip",
      "extension": "zip",
      "download_url": "https://app.lemonsqueezy.com/download/6dce5ba7-76f2-481f-ad1e-9c2bec6eb0e2?expires=1636384018&signature=f0a9bdec44ffabf143d4689594491f42a76d773d3cc88ec23ef84d6e903e8f11",
      "size": 874694,
      "size_formatted": "854 KB",
      "version": "1.0.0",
      "sort": 1,
      "status": "published",
      "createdAt": "2021-11-05T10:22:14.000000Z",
      "updatedAt": "2021-11-05T16:16:33.000000Z"
    },
    "relationships": {
      "variant": {
        "links": {
          "related": "https://api.lemonsqueezy.com/v1/files/1/variant",
          "self": "https://api.lemonsqueezy.com/v1/files/1/relationships/variant"
        }
      }
    },
    "links": {
      "self": "https://api.lemonsqueezy.com/v1/files/1"
    }
  }
}
`)
}

// FilesListResponse is a dummy response to GET /v1/files
func FilesListResponse() []byte {
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
    "first": "https://api.lemonsqueezy.com/v1/files?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=sort",
    "last": "https://api.lemonsqueezy.com/v1/files?page%5Bnumber%5D=1&page%5Bsize%5D=10&sort=sort"
  },
  "data": [
    {
      "type": "files",
      "id": "1",
      "attributes": {
        "variant_id": 168,
        "identifier": "6dce5ba7-76f2-481f-ad1e-9c2bec6eb0e2",
        "name": "my_product.zip",
        "extension": "zip",
        "download_url": "https://app.lemonsqueezy.com/download/6dce5ba7-76f2-481f-ad1e-9c2bec6eb0e2?expires=1636383388&signature=886a63faf7215c54011accfa08578b1b687def66f767092629f263061b3a253a",
        "size": 874694,
        "size_formatted": "854 KB",
        "version": "1.0.0",
        "sort": 1,
        "status": "published",
        "createdAt": "2021-11-05T10:22:14.000000Z",
        "updatedAt": "2021-11-05T16:16:33.000000Z"
      },
      "relationships": {
        "variant": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/files/1/variant",
            "self": "https://api.lemonsqueezy.com/v1/files/1/relationships/variant"
          }
        }
      },
      "links": {
        "self": "https://api.lemonsqueezy.com/v1/files/1"
      }
    },
	{
      "type": "files",
      "id": "2",
      "attributes": {
        "variant_id": 168,
        "identifier": "6dce5ba7-76f2-481f-ad1e-9c2bec6eb0e2",
        "name": "my_product.zip",
        "extension": "zip",
        "download_url": "https://app.lemonsqueezy.com/download/6dce5ba7-76f2-481f-ad1e-9c2bec6eb0e2?expires=1636383388&signature=886a63faf7215c54011accfa08578b1b687def66f767092629f263061b3a253a",
        "size": 874694,
        "size_formatted": "854 KB",
        "version": "1.0.0",
        "sort": 1,
        "status": "published",
        "createdAt": "2021-11-05T10:22:14.000000Z",
        "updatedAt": "2021-11-05T16:16:33.000000Z"
      },
      "relationships": {
        "variant": {
          "links": {
            "related": "https://api.lemonsqueezy.com/v1/files/1/variant",
            "self": "https://api.lemonsqueezy.com/v1/files/1/relationships/variant"
          }
        }
      },
      "links": {
        "self": "https://api.lemonsqueezy.com/v1/files/1"
      }
    }
  ]
}
`)
}
