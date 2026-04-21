package api_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

type smokeCase struct {
	name           string
	method         string
	path           string
	body           string
	contentType    string
	expectedStatus []int
}

func TestAPISmoke(t *testing.T) {
	baseURL := strings.TrimSpace(os.Getenv("SMOKE_BASE_URL"))
	if baseURL == "" {
		t.Skip("SMOKE_BASE_URL is not set")
	}

	client := &http.Client{Timeout: 5 * time.Second}

	cases := []smokeCase{
		{
			name:           "health check",
			method:         http.MethodGet,
			path:           "/",
			expectedStatus: []int{http.StatusOK},
		},
		{
			name:           "categories list",
			method:         http.MethodGet,
			path:           "/api/categories",
			expectedStatus: []int{http.StatusOK},
		},
		{
			name:           "menus list",
			method:         http.MethodGet,
			path:           "/api/menus",
			expectedStatus: []int{http.StatusOK},
		},
		{
			name:           "tables list",
			method:         http.MethodGet,
			path:           "/api/tables",
			expectedStatus: []int{http.StatusOK},
		},
		{
			name:           "settings list",
			method:         http.MethodGet,
			path:           "/api/settings",
			expectedStatus: []int{http.StatusOK},
		},
		{
			name:           "analytics summary",
			method:         http.MethodGet,
			path:           "/api/analytics",
			expectedStatus: []int{http.StatusOK, http.StatusInternalServerError},
		},
		{
			name:           "orders requires table_session_id",
			method:         http.MethodGet,
			path:           "/api/orders",
			expectedStatus: []int{http.StatusBadRequest},
		},
		{
			name:           "orders reject invalid payload",
			method:         http.MethodPost,
			path:           "/api/orders",
			body:           `{}`,
			contentType:    "application/json",
			expectedStatus: []int{http.StatusBadRequest},
		},
		{
			name:           "auth login reject malformed json",
			method:         http.MethodPost,
			path:           "/api/auth/login",
			body:           `{`,
			contentType:    "application/json",
			expectedStatus: []int{http.StatusBadRequest},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
			defer cancel()

			var bodyReader io.Reader
			if tc.body != "" {
				bodyReader = bytes.NewBufferString(tc.body)
			}

			req, err := http.NewRequestWithContext(ctx, tc.method, baseURL+tc.path, bodyReader)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			if tc.contentType != "" {
				req.Header.Set("Content-Type", tc.contentType)
			}

			res, err := client.Do(req)
			if err != nil {
				t.Fatalf("request failed: %v", err)
			}
			defer func() {
				if closeErr := res.Body.Close(); closeErr != nil {
					t.Errorf("failed to close response body: %v", closeErr)
				}
			}()

			if !isExpectedStatus(res.StatusCode, tc.expectedStatus) {
				responseBody, _ := io.ReadAll(io.LimitReader(res.Body, 1024))
				t.Fatalf("unexpected status: got=%d expected=%v body=%s", res.StatusCode, tc.expectedStatus, string(responseBody))
			}
		})
	}
}

func isExpectedStatus(actual int, expected []int) bool {
	for _, status := range expected {
		if actual == status {
			return true
		}
	}
	return false
}
