package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type service struct {
	client *Client
}

// Client is the lemonsqueezy API client.
// Do not instantiate this client with Client{}. Use the New method instead.
type Client struct {
	httpClient    *http.Client
	common        service
	baseURL       string
	apiKey        string
	signingSecret string

	Webhooks      *webhooksService
	Subscriptions *subscriptionsService
}

// New creates and returns a new Client from a slice of Option.
func New(options ...Option) *Client {
	config := defaultClientConfig()

	for _, option := range options {
		option.apply(config)
	}

	client := &Client{
		httpClient:    config.httpClient,
		apiKey:        config.apiKey,
		signingSecret: config.signingSecret,
		baseURL:       config.baseURL,
	}

	client.common.client = client
	client.Subscriptions = (*subscriptionsService)(&client.common)
	client.Webhooks = (*webhooksService)(&client.common)

	return client
}

// newRequest creates an API request. A relative URL can be provided in uri,
// in which case it is resolved relative to the BaseURL of the Client.
// URI's should always be specified without a preceding slash.
func (client *Client) newRequest(ctx context.Context, method, uri string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, client.baseURL+uri, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+client.apiKey)
	req.Header.Set("Content-Type", "application/vnd.api+json")
	req.Header.Set("Accept", "application/vnd.api+json")

	return req, nil
}

// do carries out an HTTP request and returns a Response
func (client *Client) do(req *http.Request) (*Response, error) {
	if req == nil {
		return nil, fmt.Errorf("%T cannot be nil", req)
	}

	httpResponse, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = httpResponse.Body.Close() }()

	resp, err := client.newResponse(httpResponse)
	if err != nil {
		return resp, err
	}

	_, err = io.Copy(io.Discard, httpResponse.Body)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// newResponse converts an *http.Response to *Response
func (client *Client) newResponse(httpResponse *http.Response) (*Response, error) {
	if httpResponse == nil {
		return nil, fmt.Errorf("%T cannot be nil", httpResponse)
	}

	resp := new(Response)
	resp.HTTPResponse = httpResponse

	buf, err := io.ReadAll(resp.HTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	resp.Body = &buf

	return resp, resp.Error()
}
