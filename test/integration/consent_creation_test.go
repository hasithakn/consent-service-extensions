package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"consent-service-extensions/internal/models"
	"consent-service-extensions/pkg/api"
)

func TestPreProcessConsentCreation_Success(t *testing.T) {
	router := api.NewRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	requestBody := models.PreProcessConsentCreationRequest{
		RequestID: "REQ-123456",
		Data: models.Request{
			ConsentInitiationData: models.DetailedConsentResourceData{
				Type:                       "accounts",
				Status:                     "AwaitingAuthorisation",
				ValidityTime:               86400,
				RecurringIndicator:         false,
				Frequency:                  0,
				DataAccessValidityDuration: 43200,
				RequestPayload: map[string]interface{}{
					"Data": map[string]interface{}{
						"Permissions": []interface{}{
							"accounts:read",
							"balance:read",
						},
					},
				},
			},
			RequestHeaders: map[string]interface{}{
				"x-fapi-interaction-id": "93bac548-d2de-4546-b106-880a5018460d",
			},
		},
	}

	body, _ := json.Marshal(requestBody)
	resp, err := http.Post(server.URL+"/api/services/pre-process-consent-creation", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var response models.SuccessResponsePreProcessConsentCreation
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Status != "SUCCESS" {
		t.Errorf("Expected status SUCCESS, got %s", response.Status)
	}

	if response.ResponseID != requestBody.RequestID {
		t.Errorf("Expected responseId %s, got %s", requestBody.RequestID, response.ResponseID)
	}
}

func TestPreProcessConsentCreation_InvalidJSON(t *testing.T) {
	router := api.NewRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	resp, err := http.Post(server.URL+"/api/services/pre-process-consent-creation", "application/json", bytes.NewBuffer([]byte("invalid json")))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", resp.StatusCode)
	}

	var errorResponse models.ErrorResponse
	if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
		t.Fatalf("Failed to decode error response: %v", err)
	}

	if errorResponse.Status != "ERROR" {
		t.Errorf("Expected status ERROR, got %s", errorResponse.Status)
	}
}
