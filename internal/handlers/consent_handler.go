package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"consent-service-extensions/internal/models"
)

// ConsentHandler handles consent-related operations
type ConsentHandler struct {
	// Add dependencies here (e.g., database, services)
}

// NewConsentHandler creates a new consent handler
func NewConsentHandler() *ConsentHandler {
	return &ConsentHandler{}
}

// PreProcessConsentCreation handles pre validations & obtains custom consent data to be stored
func (h *ConsentHandler) PreProcessConsentCreation(w http.ResponseWriter, r *http.Request) {
	var req models.PreProcessConsentCreationRequest

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request: %v", err)
		h.sendErrorResponse(w, http.StatusBadRequest, "invalid_request", "Invalid request body", req.RequestID)
		return
	}

	// Log the request
	log.Printf("Received pre-process-consent-creation request with ID: %s", req.RequestID)

	// TODO: Implement your business logic here
	// - Validate consent data
	// - Apply business rules
	// - Add custom attributes if needed

	// Extract resolved consent purposes from requestPayload.Data.Permissions
	resolvedPurposes := h.extractConsentPurposes(req.Data.ConsentInitiationData.RequestPayload)

	// For now, returning a success response with the received data
	response := models.SuccessResponsePreProcessConsentCreation{
		ResponseID: req.RequestID,
		Status:     "SUCCESS",
		Data: models.SuccessResponseWithDetailedConsentData{
			ConsentResource:         req.Data.ConsentInitiationData,
			ResolvedConsentPurposes: resolvedPurposes,
		},
	}

	// Send response
	h.sendJSONResponse(w, http.StatusOK, response)
}

// PreProcessConsentUpdate handles pre validations for consent updates & obtains custom consent data to be stored
func (h *ConsentHandler) PreProcessConsentUpdate(w http.ResponseWriter, r *http.Request) {
	var req models.PreProcessConsentUpdateRequest

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request: %v", err)
		h.sendErrorResponse(w, http.StatusBadRequest, "invalid_request", "Invalid request body", req.RequestID)
		return
	}

	// Log the request
	log.Printf("Received pre-process-consent-update request with ID: %s", req.RequestID)

	// TODO: Implement your business logic here
	// - Validate consent update data
	// - Apply business rules for updates
	// - Add custom attributes if needed

	// Extract resolved consent purposes from requestPayload.Data.Permissions
	resolvedPurposes := h.extractConsentPurposes(req.Data.ConsentInitiationData.RequestPayload)

	// Return a success response with the received data
	response := models.SuccessResponsePreProcessConsentCreation{
		ResponseID: req.RequestID,
		Status:     "SUCCESS",
		Data: models.SuccessResponseWithDetailedConsentData{
			ConsentResource:         req.Data.ConsentInitiationData,
			ResolvedConsentPurposes: resolvedPurposes,
		},
	}

	// Send response
	h.sendJSONResponse(w, http.StatusOK, response)
}

// extractConsentPurposes extracts the permissions from requestPayload.Data.Permissions
func (h *ConsentHandler) extractConsentPurposes(requestPayload map[string]interface{}) []string {
	var purposes []string

	// Check if requestPayload has a "Data" field
	if data, ok := requestPayload["Data"].(map[string]interface{}); ok {
		// Check if Data has a "Permissions" field
		if permissions, ok := data["Permissions"].([]interface{}); ok {
			// Convert each permission to string
			for _, perm := range permissions {
				if permStr, ok := perm.(string); ok {
					purposes = append(purposes, permStr)
				}
			}
		}
	}

	return purposes
}

// sendJSONResponse sends a JSON response
func (h *ConsentHandler) sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

// sendErrorResponse sends an error response
func (h *ConsentHandler) sendErrorResponse(w http.ResponseWriter, statusCode int, errorMessage, errorDescription, responseID string) {
	errorResp := models.ErrorResponse{
		ResponseID:       responseID,
		Status:           "ERROR",
		ErrorMessage:     errorMessage,
		ErrorDescription: errorDescription,
	}

	h.sendJSONResponse(w, statusCode, errorResp)
}
