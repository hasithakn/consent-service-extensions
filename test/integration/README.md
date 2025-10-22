# Integration Tests

This directory contains HTTP integration tests for the Consent Service Extensions API.

## Overview

These tests make actual HTTP requests to a running server instance to verify the API behavior end-to-end.

## Running the Tests

### Prerequisites

1. **Start the server** in a separate terminal:
   ```bash
   go run cmd/server/main.go
   ```

2. **Wait for the server to be ready** (it should be listening on port 8080)

### Run All Integration Tests

```bash
go test ./test/integration/... -v
```

### Run Specific Test

```bash
go test ./test/integration/... -v -run TestHealthEndpoint
go test ./test/integration/... -v -run TestPreProcessConsentCreation_BasicAccountConsent
```

### Run with Custom Base URL

If your server is running on a different port or host:

```bash
TEST_BASE_URL=http://localhost:9090 go test ./test/integration/... -v
```

## Test Structure

### Files

- **`setup_test.go`** - Test setup, helper functions, and configuration
- **`health_test.go`** - Health check endpoint tests
- **`consent_creation_test.go`** - Pre-process consent creation endpoint tests

### Test Categories

1. **Health Tests**
   - Basic health check
   - Response time validation

2. **Consent Creation Tests**
   - Basic account consent
   - Payment consent with authorization
   - Recurring consent
   - Invalid JSON handling
   - Empty request ID
   - Wrong HTTP method
   - Missing content type

## Test Features

- ✅ Real HTTP requests to running server
- ✅ JSON request/response validation
- ✅ HTTP status code verification
- ✅ Content-Type header validation
- ✅ Error handling verification
- ✅ Server availability check with timeout
- ✅ Configurable base URL via environment variable

## Writing New Tests

1. Add new test files in `test/integration/`
2. Use the `makeRequest()` helper function for HTTP requests
3. Follow the naming convention: `Test<Feature>_<Scenario>`
4. Include both positive and negative test cases

Example:

```go
func TestNewEndpoint_SuccessCase(t *testing.T) {
    reqBody := map[string]interface{}{
        "key": "value",
    }
    
    resp, body := makeRequest(t, http.MethodPost, "/api/services/new-endpoint", reqBody)
    
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status 200, got %d. Body: %s", resp.StatusCode, string(body))
    }
    
    // Add more assertions...
}
```

## Continuous Integration

In CI/CD pipelines, you can:

1. Start the server in the background
2. Run integration tests
3. Stop the server

Example script:

```bash
#!/bin/bash

# Start server in background
go build -o bin/server cmd/server/main.go
./bin/server &
SERVER_PID=$!

# Wait for server to be ready
sleep 2

# Run tests
go test ./test/integration/... -v

# Capture test result
TEST_RESULT=$?

# Stop server
kill $SERVER_PID

# Exit with test result
exit $TEST_RESULT
```

## Troubleshooting

### Tests Fail with "connection refused"

Make sure the server is running:
```bash
go run cmd/server/main.go
```

### Tests Timeout

Check if the server is responding:
```bash
curl http://localhost:8080/health
```

### Wrong Base URL

Set the correct base URL:
```bash
export TEST_BASE_URL=http://localhost:8080
go test ./test/integration/... -v
```

## Test Coverage

Integration tests cover:
- ✅ HTTP endpoints
- ✅ Request/response formats
- ✅ Status codes
- ✅ Headers
- ✅ Error handling
- ✅ Edge cases

Unit tests (if needed) should be placed in the same package as the code being tested.
