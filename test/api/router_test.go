package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"consent-service-extensions/pkg/api"
)

func TestHealthCheckEndpoint(t *testing.T) {
	router := api.NewRouter()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"status":"healthy"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestRouter_Routes(t *testing.T) {
	router := api.NewRouter()

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
		description    string
	}{
		{
			name:           "Health check endpoint",
			method:         http.MethodGet,
			path:           "/health",
			expectedStatus: http.StatusOK,
			description:    "Health check should return 200",
		},
		{
			name:           "Pre-process consent creation endpoint exists",
			method:         http.MethodPost,
			path:           "/api/services/pre-process-consent-creation",
			expectedStatus: http.StatusBadRequest,
			description:    "Endpoint exists but returns 400 for empty body",
		},
		{
			name:           "Non-existent endpoint",
			method:         http.MethodGet,
			path:           "/api/services/non-existent",
			expectedStatus: http.StatusNotFound,
			description:    "Should return 404 for non-existent route",
		},
		{
			name:           "Wrong HTTP method",
			method:         http.MethodGet,
			path:           "/api/services/pre-process-consent-creation",
			expectedStatus: http.StatusMethodNotAllowed,
			description:    "Should return 405 for wrong HTTP method",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("%s: handler returned wrong status code: got %v want %v - %s",
					tt.name, status, tt.expectedStatus, tt.description)
			}
		})
	}
}

func TestRouter_CORS(t *testing.T) {
	router := api.NewRouter()

	req := httptest.NewRequest(http.MethodOptions, "/health", nil)
	req.Header.Set("Origin", "http://example.com")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	// This test will pass as-is, but you can add CORS middleware later
	t.Log("CORS test placeholder - add CORS middleware if needed")
}

func TestRouter_ContentType(t *testing.T) {
	router := api.NewRouter()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	contentType := rr.Header().Get("Content-Type")
	expected := "application/json"
	if contentType != expected {
		t.Errorf("expected Content-Type %s, got %s", expected, contentType)
	}
}
