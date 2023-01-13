package stubs

// UsersMeResponse is the response from the GET /v1/users/me endpoint
func UsersMeResponse() []byte {
	return []byte(`
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "self": "https://api.lemonsqueezy.com/v1/users/1"
  },
  "data": {
    "type": "users",
    "id": "1",
    "attributes": {
      "name": "Darlene Daugherty",
      "email": "gernser@yahoo.com",
      "color": "#898FA9",
      "avatar_url": "https://www.gravatar.com/avatar/1ace5b3965c59dbcd1db79d85da75048?d=blank",
      "has_custom_avatar": false,
      "createdAt": "2021-05-24T14:08:31.000000Z",
      "updatedAt": "2021-08-26T13:24:54.000000Z"
    },
    "links": {
      "self": "https://api.lemonsqueezy.com/v1/users/1"
    }
  }
}
`)
}
