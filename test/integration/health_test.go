package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"consent-service-extensions/pkg/api"
)

func TestHealthEndpoint(t *testing.T) {
	router := api.NewRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	resp, err := http.Get(server.URL + "/health")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}

	var response map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if status, ok := response["status"]; !ok || status != "healthy" {
		t.Errorf("Expected status 'healthy', got %v", response)
	}
}
