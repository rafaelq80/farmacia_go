package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	auth "github.com/rafaelq80/farmacia_go/security/service"
	"github.com/stretchr/testify/require"
)

func performRequestWithAuth(t *testing.T, app *fiber.App, method, path string, body ...interface{}) (*http.Response, map[string]interface{}) {
	t.Helper()

	// Prepare request
	req := prepareRequest(t, method, path, body...)

	// Add authorization
	token, err := auth.CreateToken("root@root.com.br")
	require.NoError(t, err, "Failed to create auth token")
	req.Header.Set("Authorization", "Bearer "+token)

	// Perform request
	resp, err := app.Test(req)
	require.NoError(t, err, "Failed to perform request")

	// Parse response
	responseBody, err := parseResponse(resp)
	require.NoError(t, err, "Failed to parse response")

	return resp, responseBody
}

func prepareRequest(t *testing.T, method, path string, body ...interface{}) *http.Request {
	t.Helper()

	var req *http.Request
	if len(body) > 0 {
		jsonBody, err := json.Marshal(body[0])
		require.NoError(t, err, "Failed to marshal request body")
		req = httptest.NewRequest(method, path, bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req = httptest.NewRequest(method, path, nil)
	}

	return req
}

func parseResponse(resp *http.Response) (map[string]interface{}, error) {
	if resp.ContentLength == 0 {
		return nil, nil
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Recreate the response body for future use
	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var responseBody map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &responseBody); err != nil {
		// If JSON parsing fails, store the raw body as a string
		responseBody = map[string]interface{}{"rawBody": string(bodyBytes)}
	}

	return responseBody, nil
}