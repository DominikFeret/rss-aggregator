package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIkey is a function that extracts an API key
// from the HTTP request headers
// Example:
// Authorization: ApiKey {key}
func getAPIkey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authorization header")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid Authorization header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("invalid first part of Authorization header")
	}
	return vals[1], nil
}
