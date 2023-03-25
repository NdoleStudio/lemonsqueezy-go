package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersService_Me(t *testing.T) {
	// Setup
	t.Parallel()

	user, response, err := client.Users.Me(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.NotNil(t, user)
}
