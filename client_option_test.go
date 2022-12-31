package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithHTTPClient(t *testing.T) {
	t.Run("httpClient is not set when the httpClient is nil", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()

		// Act
		WithHTTPClient(nil).apply(config)

		// Assert
		assert.NotNil(t, config.httpClient)
	})

	t.Run("httpClient is set when the httpClient is not nil", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()
		newClient := &http.Client{Timeout: 300}

		// Act
		WithHTTPClient(newClient).apply(config)

		// Assert
		assert.NotNil(t, config.httpClient)
		assert.Equal(t, newClient.Timeout, config.httpClient.Timeout)
	})
}

func TestWithBaseURL(t *testing.T) {
	t.Run("baseURL is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		baseURL := "https://example.com"
		config := defaultClientConfig()

		// Act
		WithBaseURL(baseURL).apply(config)

		// Assert
		assert.Equal(t, config.baseURL, config.baseURL)
	})

	t.Run("tailing / is trimmed from baseURL", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		baseURL := "https://example.com/"
		config := defaultClientConfig()

		// Act
		WithBaseURL(baseURL).apply(config)

		// Assert
		assert.Equal(t, "https://example.com", config.baseURL)
	})
}

func TestWithAPIKey(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	config := defaultClientConfig()
	apiKey := "api-key"

	// Act
	WithAPIKey(apiKey).apply(config)

	// Assert
	assert.Equal(t, apiKey, config.apiKey)
}

func TestWithSigningSecret(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	config := defaultClientConfig()
	signingSecret := "signing-secret"

	// Act
	WithSigningSecret(signingSecret).apply(config)

	// Assert
	assert.Equal(t, signingSecret, config.signingSecret)
}
