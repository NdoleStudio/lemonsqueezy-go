package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// statusService is the API client for the `/` endpoint
type statusService service

// Ok returns the 200 HTTP status Code.
//
// API Docs: https://httpstat.us
func (service *statusService) Ok(ctx context.Context) (*HTTPStatus, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodGet, "/200", nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	status := new(HTTPStatus)
	if err = json.Unmarshal(*response.Body, status); err != nil {
		return nil, response, err
	}

	return status, response, nil
}
