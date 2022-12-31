package client

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/NdoleStudio/go-http-client/internal/helpers"
	"github.com/stretchr/testify/assert"
)

func TestStatusService_Ok(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	client := New()

	// Act
	status, response, err := client.Status.Ok(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, &HTTPStatus{Code: 200, Description: "OK"}, status)
}

func TestBillsService_OkWithDelay(t *testing.T) {
	// Setup
	t.Parallel()
	start := time.Now()

	// Arrange
	client := New(WithDelay(500))

	// Act
	status, response, err := client.Status.Ok(context.Background())

	// Assert
	assert.Nil(t, err)
	assert.LessOrEqual(t, int64(100), time.Since(start).Milliseconds())
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, &HTTPStatus{Code: 200, Description: "OK"}, status)
}

func TestBillsService_OkWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, []byte("Internal Server Error"))
	client := New(WithBaseURL(server.URL))

	// Act
	status, response, err := client.Status.Ok(context.Background())

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, status)

	assert.Equal(t, "500: Internal Server Error, Body: Internal Server Error", err.Error())

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)
	assert.Equal(t, "Internal Server Error", string(*response.Body))

	// Teardown
	server.Close()
}
