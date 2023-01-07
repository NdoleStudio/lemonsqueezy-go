package helpers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MakeTestServer creates an api server for testing
func MakeTestServer(responseCode int, body []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(responseCode)
		_, err := res.Write(body)
		if err != nil {
			panic(err)
		}
	}))
}

// AssertObjectEqualsJSON checks if the JSON representation of an object matches the expected value
func AssertObjectEqualsJSON(t *testing.T, expectedJSON []byte, actual any) {
	actualJSON, err := json.Marshal(actual)
	assert.Nil(t, err)
	assert.JSONEq(t, string(expectedJSON), string(actualJSON))
}
