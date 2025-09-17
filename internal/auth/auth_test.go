package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("Testing malformed header", func(t *testing.T) {
		var header http.Header = make(http.Header)
		_, err := GetAPIKey(header)
		expectedError := ErrNoAuthHeaderIncluded
		if err.Error() != expectedError.Error() {
			t.Errorf("Expected error: %v Actual: %v", err, expectedError)
		}
	})
	t.Run("Testing malformed header", func(t *testing.T) {
		var header http.Header = make(http.Header)
		header.Add("Authorization", "malformedAuth")
		_, err := GetAPIKey(header)
		expectedError := errors.New("malformed authorization header")
		if err.Error() != expectedError.Error() {
			t.Errorf("Expected error: %v Actual: %v", err, expectedError)
		}
	})
}
