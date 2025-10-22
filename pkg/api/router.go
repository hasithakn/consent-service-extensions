package api

import (
	"net/http"

	"consent-service-extensions/internal/handlers"

	"github.com/gorilla/mux"
)

// NewRouter creates and configures the main application router
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Create handlers
	consentHandler := handlers.NewConsentHandler()

	// Register routes
	api := router.PathPrefix("/api/services").Subrouter()

	// Consent endpoints
	api.HandleFunc("/pre-process-consent-creation", consentHandler.PreProcessConsentCreation).Methods(http.MethodPost)
	api.HandleFunc("/pre-process-consent-update", consentHandler.PreProcessConsentUpdate).Methods(http.MethodPost)

	// TODO: Add more endpoints as needed:
	// api.HandleFunc("/enrich-consent-creation-response", consentHandler.EnrichConsentCreationResponse).Methods(http.MethodPost)
	// api.HandleFunc("/pre-process-consent-retrieval", consentHandler.PreProcessConsentRetrieval).Methods(http.MethodPost)
	// api.HandleFunc("/enrich-consent-update-response", consentHandler.EnrichConsentUpdateResponse).Methods(http.MethodPost)
	// api.HandleFunc("/enrich-consent-update-response", consentHandler.EnrichConsentUpdateResponse).Methods(http.MethodPost)
	// api.HandleFunc("/pre-process-consent-revoke", consentHandler.PreProcessConsentRevoke).Methods(http.MethodPost)
	// api.HandleFunc("/pre-process-consent-file-upload", consentHandler.PreProcessConsentFileUpload).Methods(http.MethodPost)
	// api.HandleFunc("/enrich-consent-file-response", consentHandler.EnrichConsentFileResponse).Methods(http.MethodPost)
	// api.HandleFunc("/validate-consent-file-retrieval", consentHandler.ValidateConsentFileRetrieval).Methods(http.MethodPost)
	// api.HandleFunc("/pre-process-consent-file-update", consentHandler.PreProcessConsentFileUpdate).Methods(http.MethodPost)
	// api.HandleFunc("/enrich-consent-file-update-response", consentHandler.EnrichConsentFileUpdateResponse).Methods(http.MethodPost)
	// api.HandleFunc("/map-accelerator-error-response", errorHandler.MapAcceleratorErrorResponse).Methods(http.MethodPost)

	// Health check endpoint
	router.HandleFunc("/health", healthCheckHandler).Methods(http.MethodGet)

	return router
}

// healthCheckHandler returns service health status
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"healthy"}`))
}
