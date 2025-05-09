package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := map[string]struct {
		input   http.Header
		want    string
		wantErr bool
	}{
		"valid header": {
			input: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "ApiKey abc123xyz")
				return h
			}(),
			want:    "abc123xyz",
			wantErr: false,
		},
		"missing header": {
			input:   http.Header{},
			want:    "",
			wantErr: true,
		},
		"malformed header": {
			input: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "Bearer abc123xyz")
				return h
			}(),
			want:    "",
			wantErr: true,
		},
		"empty ApiKey": {
			input: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "ApiKey")
				return h
			}(),
			want:    "",
			wantErr: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if (err != nil) != tc.wantErr {
				t.Fatalf("expected error: %v, got: %v", tc.wantErr, err)
			}
			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}
}
