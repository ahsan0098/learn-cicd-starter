package auth

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		header        http.Header
		expectedKey   string
		expectedError error
	}{
		{
			name:          "no header",
			header:        http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed header - no value",
			header: http.Header{
				"Authorization": []string{"Bearer12345"},
			},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name: "malformed header - missing key part",
			header: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name: "valid header",
			header: http.Header{
				"Authorization": []string{"ApiKey test12345"},
			},
			expectedKey:   "test12345",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.header)

			if key == tt.expectedKey {
				t.Errorf("expected key '%s', got '%s'", tt.expectedKey, key)
			}

			if (err != nil && tt.expectedError == nil) || (err == nil && tt.expectedError != nil) {
				t.Errorf("expected error '%v', got '%v'", tt.expectedError, err)
			} else if err != nil && tt.expectedError != nil && !strings.Contains(err.Error(), tt.expectedError.Error()) {
				t.Errorf("expected error '%v', got '%v'", tt.expectedError, err)
			}
		})
	}
}

func TestFailing(t *testing.T) {
	t.Fatal("go test")
}
