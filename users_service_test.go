package lemonsqueezy

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/NdoleStudio/lemonsqueezy-go/internal/helpers"
	"github.com/NdoleStudio/lemonsqueezy-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestUsersService_Me(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.UsersMeResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	user, response, err := client.Users.Me(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.UsersMeResponse(), *response.Body)
	assert.Equal(t, "gernser@yahoo.com", user.Data.Attributes.Email)

	// Teardown
	server.Close()
}

func TestUsersService_MeWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Users.Me(context.Background())

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestUsersService_MeCancelledContext(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// Act
	_, response, err := client.Users.Me(ctx)

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.True(t, errors.Is(err, context.Canceled))

	// Teardown
	server.Close()
}

func TestUsersService_MeWithInvalidJSONResponse(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, []byte("invalid JSON response"))
	client := New(WithBaseURL(server.URL))

	// Act
	_, _, err := client.Users.Me(context.Background())

	// Assert
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "invalid character"))

	// Teardown
	server.Close()
}
