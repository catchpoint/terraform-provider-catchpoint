package client

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	rateLimit       = 7 // 7 requests per second
	bucketSize      = 7 // same as rate limit
	requestInterval = time.Second / rateLimit
)

// The generic Response function to use for calling web requests.
var doFunc = httpClient.Do

var (
	// A simple token bucket implementation to limit the rate of requests.
	// It allows up to 7 requests per second, with a maximum of 7 tokens in the bucket.
	// When a request is made, a token is consumed. If no tokens are available,
	// the request will be blocked until a token becomes available.
	// The bucket is refilled at a rate of 7 tokens per second.
	// This token bucket is used throughout the 'client' package.
	tokens = make(chan struct{}, bucketSize)

	// Re-use the HTTP client for all requests to avoid the overhead of creating a new client for each request.
	httpClient = &http.Client{}
)

// Init initializes the token bucket for all files in this package.
func init() {
	// Fill the bucket with initial tokens
	for range bucketSize {
		tokens <- struct{}{}
	}

	// Refill tokens at a rate of 7 per second
	go func() {
		ticker := time.NewTicker(requestInterval)
		defer ticker.Stop()
		for range ticker.C {
			select {
			case tokens <- struct{}{}:
			default:
				// The bucket is full; discard new tokens
			}
		}
	}()
}

// T is any type that can be unmarshaled into.
// If err, responseBody and responseStatus may be returned as empty strings.
func doAndUnmarshal[T any](req *http.Request, apiToken string) (unmarshaledResult T, responseBody, responseStatus string, err error) {
	resp, err := do(req, apiToken)
	if err != nil {
		// If there was an error making the request, attempt to return the body and status from the response anyway
		// as they have error context. They will be empty if the response is empty.
		responseBody, responseStatus = formatResponse(resp)
		return
	}
	defer resp.Body.Close()

	responseBody, responseStatus = formatResponse(resp)
	if err = json.Unmarshal([]byte(responseBody), &unmarshaledResult); err != nil {
		return
	}
	return
}

// Do sends an HTTP request with the provided API token and returns the response.
// It blocks until a token is available in the token bucket to enforce rate limiting.
// The request must have the "Authorization" header set with the API token.
// The "Content-Type" header is set to "application/json" and the "cp-integration" header is set to "1" for Catchpoint API requests.
func do(req *http.Request, apiToken string) (resp *http.Response, err error) {
	<-tokens

	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("cp-integration", "1")

	resp, err = doFunc(req)
	if resp != nil {
		log.Printf("Response status code: %s", strings.ToLower(resp.Status))
	} else {
		log.Printf("Response is nil")
	}
	return
}

// formatResponse reads the response body and returns it as a string along with the response status.
// It also closes the response body to prevent resource leaks.
func formatResponse(resp *http.Response) (responseBody, responseStatus string) {
	if resp == nil {
		return "", ""
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	responseBody = string(body)
	responseStatus = strings.ToLower(string(resp.Status))

	return
}
