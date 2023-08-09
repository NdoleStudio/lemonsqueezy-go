package lemonsqueezy

import "net/http"

type clientConfig struct {
	httpClient    *http.Client
	baseURL       string
	apiKey        string
	signingSecret string
}

func defaultClientConfig() *clientConfig {
	return &clientConfig{
		httpClient: http.DefaultClient,
		baseURL:    "https://api.lemonsqueezy.com",
	}
}
