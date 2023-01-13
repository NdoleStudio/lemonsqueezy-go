package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// UsersService is the API client for the `/v1/users` endpoint
type UsersService service

// Me retrieves the currently authenticated user.
// https://docs.lemonsqueezy.com/api/users#retrieve-the-authenticated-user
func (service *UsersService) Me(ctx context.Context) (*UserApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/users/me")
	if err != nil {
		return nil, response, err
	}

	user := new(UserApiResponse)
	if err = json.Unmarshal(*response.Body, user); err != nil {
		return nil, response, err
	}

	return user, response, nil
}
