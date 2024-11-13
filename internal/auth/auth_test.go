package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Keep the previous test case...

	t.Run("should return api key when valid auth header", func(t *testing.T) {
		headers := make(http.Header)
		expectedKey := "test-api-key-123"
		headers.Set("Authorization", "ApiKey "+expectedKey)

		gotKey, err := GetAPIKey(headers)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if gotKey != expectedKey {
			t.Errorf("Expected key %q, got %q", expectedKey, gotKey)
		}
	})
	t.Run("should return error when malformed auth header", func(t *testing.T) {
		headers := make(http.Header)
		headers.Set("Authorization", "BadPrefix some-key-123")

		_, err := GetAPIKey(headers)

		if err == nil {
			t.Error("Expected error for malformed header, got nil")
		}
		if err.Error() != "malformed authorization header" {
			t.Errorf("Expected malformed header error, got %v", err)
		}
	})
}
