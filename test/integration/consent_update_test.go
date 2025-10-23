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

func TestPreProcessConsentUpdate_Success(t *testing.T) {
	router := api.NewRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	requestBody := models.PreProcessConsentUpdateRequest{
		RequestID: "UPD-123456",
		Data: models.UpdateRequest{
			ConsentInitiationData: models.DetailedConsentResourceData{
				Type:                       "accounts",
				Status:                     "AwaitingAuthorisation",
				ValidityTime:               86400,
				RecurringIndicator:         true,
				Frequency:                  5,
				DataAccessValidityDuration: 43200,
				RequestPayload: map[string]interface{}{
					"Data": map[string]interface{}{
						"Permissions": []interface{}{
							"accounts:read",
							"balance:read",
							"transactions:read",
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
	resp, err := http.Post(server.URL+"/api/services/pre-process-consent-update", "application/json", bytes.NewBuffer(body))
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

	// Verify resolved consent purposes
	expectedPurposes := []string{"accounts:read", "balance:read", "transactions:read"}
	if len(response.Data.ResolvedConsentPurposes) != len(expectedPurposes) {
		t.Errorf("Expected %d purposes, got %d", len(expectedPurposes), len(response.Data.ResolvedConsentPurposes))
	}
	for i, purpose := range expectedPurposes {
		if i >= len(response.Data.ResolvedConsentPurposes) || response.Data.ResolvedConsentPurposes[i] != purpose {
			t.Errorf("Expected purpose %s at index %d, got %s", purpose, i, response.Data.ResolvedConsentPurposes[i])
		}
	}
}

func TestPreProcessConsentUpdate_InvalidJSON(t *testing.T) {
	router := api.NewRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	resp, err := http.Post(server.URL+"/api/services/pre-process-consent-update", "application/json", bytes.NewBuffer([]byte("invalid json")))
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

func TestPreProcessConsentUpdate_EmptyPermissions(t *testing.T) {
	router := api.NewRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	requestBody := models.PreProcessConsentUpdateRequest{
		RequestID: "UPD-EMPTY",
		Data: models.UpdateRequest{
			ConsentInitiationData: models.DetailedConsentResourceData{
				Type:               "accounts",
				Status:             "AwaitingAuthorisation",
				ValidityTime:       86400,
				RecurringIndicator: false,
				Frequency:          0,
				RequestPayload: map[string]interface{}{
					"Data": map[string]interface{}{
						"Permissions": []interface{}{},
					},
				},
			},
		},
	}

	body, _ := json.Marshal(requestBody)
	resp, err := http.Post(server.URL+"/api/services/pre-process-consent-update", "application/json", bytes.NewBuffer(body))
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

	if len(response.Data.ResolvedConsentPurposes) != 0 {
		t.Errorf("Expected 0 purposes, got %d", len(response.Data.ResolvedConsentPurposes))
	}
}

func TestPreProcessConsentUpdate_NoPermissionsField(t *testing.T) {
	router := api.NewRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	requestBody := models.PreProcessConsentUpdateRequest{
		RequestID: "UPD-NO-PERMS",
		Data: models.UpdateRequest{
			ConsentInitiationData: models.DetailedConsentResourceData{
				Type:               "payments",
				Status:             "AwaitingAuthorisation",
				ValidityTime:       86400,
				RecurringIndicator: false,
				Frequency:          0,
				RequestPayload: map[string]interface{}{
					"Data": map[string]interface{}{},
				},
			},
		},
	}

	body, _ := json.Marshal(requestBody)
	resp, err := http.Post(server.URL+"/api/services/pre-process-consent-update", "application/json", bytes.NewBuffer(body))
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

	if len(response.Data.ResolvedConsentPurposes) != 0 {
		t.Errorf("Expected nil or empty purposes, got %v", response.Data.ResolvedConsentPurposes)
	}
}
