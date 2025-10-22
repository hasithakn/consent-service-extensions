package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"consent-service-extensions/internal/handlers"
	"consent-service-extensions/internal/models"
)

func TestConsentHandler_PreProcessConsentCreation_Success(t *testing.T) {
	handler := handlers.NewConsentHandler()

	// Prepare test request
	reqBody := models.PreProcessConsentCreationRequest{
		RequestID: "test-123",
		Data: models.Request{
			ConsentInitiationData: models.DetailedConsentResourceData{
				Type:                       "accounts",
				Status:                     "AwaitingAuthorisation",
				ValidityTime:               0,
				RecurringIndicator:         false,
				Frequency:                  0,
				DataAccessValidityDuration: 86400,
				RequestPayload:             map[string]interface{}{},
				Attributes:                 map[string]interface{}{},
			},
			RequestHeaders: map[string]interface{}{},
		},
	}

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/api/services/pre-process-consent-creation", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	rr := httptest.NewRecorder()

	// Call handler
	handler.PreProcessConsentCreation(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response body
	var response models.SuccessResponsePreProcessConsentCreation
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	if response.Status != "SUCCESS" {
		t.Errorf("expected status SUCCESS, got %s", response.Status)
	}

	if response.ResponseID != "test-123" {
		t.Errorf("expected responseId test-123, got %s", response.ResponseID)
	}

	if response.Data.ConsentResource.Type != "accounts" {
		t.Errorf("expected type accounts, got %s", response.Data.ConsentResource.Type)
	}
}

func TestConsentHandler_PreProcessConsentCreation_BadRequest(t *testing.T) {
	handler := handlers.NewConsentHandler()

	// Send invalid JSON
	req := httptest.NewRequest(http.MethodPost, "/api/services/pre-process-consent-creation", bytes.NewReader([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.PreProcessConsentCreation(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	// Check response body
	var response models.ErrorResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("failed to decode error response: %v", err)
	}

	if response.Status != "ERROR" {
		t.Errorf("expected status ERROR, got %s", response.Status)
	}

	if response.ErrorMessage != "invalid_request" {
		t.Errorf("expected errorMessage invalid_request, got %s", response.ErrorMessage)
	}
}

func TestConsentHandler_PreProcessConsentCreation_WithAuthorizations(t *testing.T) {
	handler := handlers.NewConsentHandler()

	// Prepare test request with authorizations
	reqBody := models.PreProcessConsentCreationRequest{
		RequestID: "test-456",
		Data: models.Request{
			ConsentInitiationData: models.DetailedConsentResourceData{
				Type:                       "payments",
				Status:                     "AwaitingAuthorisation",
				ValidityTime:               3600,
				RecurringIndicator:         true,
				Frequency:                  5,
				DataAccessValidityDuration: 86400,
				RequestPayload: map[string]interface{}{
					"amount": 100.50,
				},
				Attributes: map[string]interface{}{
					"customField": "customValue",
				},
				Authorizations: []models.ConsentAuthorizationCreatePayload{
					{
						UserID: "user@example.com",
						Type:   "authorisation",
						Status: "created",
						Resource: map[string]interface{}{
							"key": "value",
						},
					},
				},
			},
			RequestHeaders: map[string]interface{}{
				"X-Request-ID": "req-123",
			},
		},
	}

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/api/services/pre-process-consent-creation", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.PreProcessConsentCreation(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response body
	var response models.SuccessResponsePreProcessConsentCreation
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	if response.Status != "SUCCESS" {
		t.Errorf("expected status SUCCESS, got %s", response.Status)
	}

	if len(response.Data.ConsentResource.Authorizations) != 1 {
		t.Errorf("expected 1 authorization, got %d", len(response.Data.ConsentResource.Authorizations))
	}

	if response.Data.ConsentResource.Type != "payments" {
		t.Errorf("expected type payments, got %s", response.Data.ConsentResource.Type)
	}
}

func TestConsentHandler_PreProcessConsentCreation_EmptyRequestID(t *testing.T) {
	handler := handlers.NewConsentHandler()

	// Prepare test request with empty request ID
	reqBody := models.PreProcessConsentCreationRequest{
		RequestID: "",
		Data: models.Request{
			ConsentInitiationData: models.DetailedConsentResourceData{
				Type:           "accounts",
				Status:         "AwaitingAuthorisation",
				RequestPayload: map[string]interface{}{},
			},
			RequestHeaders: map[string]interface{}{},
		},
	}

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/api/services/pre-process-consent-creation", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.PreProcessConsentCreation(rr, req)

	// Should still succeed as the service generates response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
