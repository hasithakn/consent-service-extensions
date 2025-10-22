package models

// PreProcessConsentCreationRequest represents the request body for pre-process-consent-creation
type PreProcessConsentCreationRequest struct {
	RequestID string  `json:"requestId"`
	Data      Request `json:"data"`
}

// PreProcessConsentUpdateRequest represents the request body for pre-process-consent-update
type PreProcessConsentUpdateRequest struct {
	RequestID string        `json:"requestId"`
	Data      UpdateRequest `json:"data"`
}

// Request represents the data section of the request
type Request struct {
	ConsentInitiationData DetailedConsentResourceData `json:"consentInitiationData"`
	RequestHeaders        map[string]interface{}      `json:"requestHeaders"`
}

// UpdateRequest represents the data section of the update request
type UpdateRequest struct {
	ConsentInitiationData DetailedConsentResourceData `json:"consentInitiationData"`
	RequestHeaders        map[string]interface{}      `json:"requestHeaders"`
}

// DetailedConsentResourceData represents the consent resource data
type DetailedConsentResourceData struct {
	Type                       string                              `json:"type"`
	Status                     string                              `json:"status"`
	ValidityTime               int64                               `json:"validityTime"`
	RecurringIndicator         bool                                `json:"recurringIndicator"`
	Frequency                  int32                               `json:"frequency"`
	DataAccessValidityDuration int64                               `json:"dataAccessValidityDuration,omitempty"`
	RequestPayload             map[string]interface{}              `json:"requestPayload"`
	Attributes                 map[string]interface{}              `json:"attributes,omitempty"`
	Authorizations             []ConsentAuthorizationCreatePayload `json:"authorizations,omitempty"`
}

// ConsentAuthorizationCreatePayload represents an authorization object
type ConsentAuthorizationCreatePayload struct {
	UserID   string                 `json:"userId"`
	Type     string                 `json:"type"`
	Status   string                 `json:"status"`
	Resource map[string]interface{} `json:"resource,omitempty"`
}

// SuccessResponsePreProcessConsentCreation represents the success response
type SuccessResponsePreProcessConsentCreation struct {
	ResponseID string                                 `json:"responseId"`
	Status     string                                 `json:"status"`
	Data       SuccessResponseWithDetailedConsentData `json:"data"`
}

// SuccessResponseWithDetailedConsentData represents the data section of success response
type SuccessResponseWithDetailedConsentData struct {
	ConsentResource         DetailedConsentResourceData `json:"consentResource"`
	ResolvedConsentPurposes []string                    `json:"resolvedConsentPurposes"`
}

// FailedResponse represents a failed response
type FailedResponse struct {
	ResponseID string                 `json:"responseId"`
	Status     string                 `json:"status"`
	ErrorCode  int                    `json:"errorCode"`
	Data       map[string]interface{} `json:"data"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	ResponseID       string `json:"responseId,omitempty"`
	Status           string `json:"status"`
	ErrorMessage     string `json:"errorMessage"`
	ErrorDescription string `json:"errorDescription"`
}
