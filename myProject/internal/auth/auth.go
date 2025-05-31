package auth

import (
	"fmt"
	"net/http"
	"strings"
)
// GetAPIKey extracts an API key from the request headers
// returns empty string if no API key is found
// returns error if the API key is not found or is malformed
// Example : Authorization: ApiKey {api_key}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", fmt.Errorf("no authorization header found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", fmt.Errorf("malformed authorization header")
	}
	if vals[0] != "ApiKey" {
		return "", fmt.Errorf("malformed authorization header")
	}
	return vals[1], nil
}
