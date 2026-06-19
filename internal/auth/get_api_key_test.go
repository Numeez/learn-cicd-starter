package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		header      string
		expectedKey string
		expectErr   bool
	}{
		{
			name:      "missing authorization header",
			header:    "",
			expectErr: true,
		},
		{
			name:      "malformed header no key",
			header:    "ApiKey",
			expectErr: true,
		},
		{
			name:      "wrong authorization type",
			header:    "Bearer abc123",
			expectErr: true,
		},
		{
			name:        "valid api key",
			header:      "ApiKey abc123",
			expectedKey: "abc123",
			expectErr:   false,
		},
		{
			name:        "valid api key with extra values",
			header:      "ApiKey abc123 extra",
			expectedKey: "abc123",
			expectErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}

			if tt.header != "" {
				headers.Set("Authorization", tt.header)
			}

			key, err := GetAPIKey(headers)

			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if key != tt.expectedKey {
				t.Errorf("expected %q, got %q", tt.expectedKey, key)
			}
			t.Error("Oops")
		})
	}
}