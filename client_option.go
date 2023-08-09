package lemonsqueezy

import (
	"net/http"
	"strings"
)

// Option is options for constructing a client
type Option interface {
	apply(config *clientConfig)
}

type clientOptionFunc func(config *clientConfig)

func (fn clientOptionFunc) apply(config *clientConfig) {
	fn(config)
}

// WithHTTPClient sets the underlying HTTP client used for API requests.
// By default, http.DefaultClient is used.
func WithHTTPClient(httpClient *http.Client) Option {
	return clientOptionFunc(func(config *clientConfig) {
		if httpClient != nil {
			config.httpClient = httpClient
		}
	})
}

// WithBaseURL set's the base url for the flutterwave API
func WithBaseURL(baseURL string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		if baseURL != "" {
			config.baseURL = strings.TrimRight(baseURL, "/")
		}
	})
}

// WithAPIKey sets the lemonsqueezy API used to authenticate requests.
// https://docs.lemonsqueezy.com/api#authentication
func WithAPIKey(apiKey string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.apiKey = apiKey
	})
}

// WithSigningSecret sets the lemonsqueezy webhook signing secret used to authenticate webhook requests.
// https://docs.lemonsqueezy.com/api/webhooks#webhook-requests
func WithSigningSecret(signingSecret string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.signingSecret = signingSecret
	})
}
