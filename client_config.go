package client

import "net/http"

type clientConfig struct {
	httpClient *http.Client
	delay      int
	baseURL    string
}

func defaultClientConfig() *clientConfig {
	return &clientConfig{
		httpClient: http.DefaultClient,
		delay:      0,
		baseURL:    "https://httpstat.us",
	}
}
