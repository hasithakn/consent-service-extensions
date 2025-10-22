# API Endpoints

## Base URL
- Default: `http://localhost:3001`
- Configurable via `PORT` environment variable

## Endpoints

### 1. Health Check
**GET** `/health`

Returns the health status of the service.

**Response:**
```json
{
  "status": "healthy"
}
```

### 2. Pre-Process Consent Creation
**POST** `/api/services/pre-process-consent-creation`

Handles pre-validations and obtains custom consent data to be stored during consent creation.

**Request Body:**
```json
{
  "requestId": "REQ-123456",
  "data": {
    "consentInitiationData": {
      "type": "accounts",
      "status": "AwaitingAuthorisation",
      "validityTime": 0,
      "recurringIndicator": false,
      "frequency": 0,
      "requestPayload": {
        "Data": {
          "Permissions": ["accounts:read", "balance:read"],
          "ExpirationDateTime": "2025-12-01T12:03:49.936563+05:30"
        },
        "Risk": {}
      }
    },
    "requestHeaders": {
      "x-fapi-interaction-id": "93bac548-d2de-4546-b106-880a5018460d"
    }
  }
}
```

**Response:**
```json
{
  "responseId": "REQ-123456",
  "status": "SUCCESS",
  "data": {
    "consentResource": {
      "type": "accounts",
      "status": "AwaitingAuthorisation",
      "validityTime": 0,
      "recurringIndicator": false,
      "frequency": 0,
      "requestPayload": {
        "Data": {
          "Permissions": ["accounts:read", "balance:read"],
          "ExpirationDateTime": "2025-12-01T12:03:49.936563+05:30"
        },
        "Risk": {}
      }
    },
    "resolvedConsentPurposes": ["accounts:read", "balance:read"]
  }
}
```

**Features:**
- Extracts permissions from `requestPayload.Data.Permissions` array
- Returns them as `resolvedConsentPurposes` in the response

**Example:**
```bash
curl -X POST http://localhost:3001/api/services/pre-process-consent-creation \
  -H "Content-Type: application/json" \
  -d @docs/examples/test-purposes-extraction.json
```

### 3. Pre-Process Consent Update
**POST** `/api/services/pre-process-consent-update`

Handles pre-validations and obtains custom consent data to be stored during consent updates.

**Request Body:**
```json
{
  "requestId": "UPD-123456",
  "data": {
    "consentInitiationData": {
      "type": "accounts",
      "status": "AwaitingAuthorisation",
      "validityTime": 86400,
      "recurringIndicator": true,
      "frequency": 5,
      "dataAccessValidityDuration": 43200,
      "requestPayload": {
        "Data": {
          "Permissions": ["accounts:read", "balance:read", "transactions:read"],
          "ExpirationDateTime": "2024-12-31T23:59:59.000Z"
        }
      },
      "attributes": {
        "consentType": "account_update",
        "customerSegment": "retail"
      }
    },
    "requestHeaders": {
      "x-fapi-interaction-id": "93bac548-d2de-4546-b106-880a5018460d"
    }
  }
}
```

**Response:**
```json
{
  "responseId": "UPD-123456",
  "status": "SUCCESS",
  "data": {
    "consentResource": {
      "type": "accounts",
      "status": "AwaitingAuthorisation",
      "validityTime": 86400,
      "recurringIndicator": true,
      "frequency": 5,
      "dataAccessValidityDuration": 43200,
      "requestPayload": {
        "Data": {
          "Permissions": ["accounts:read", "balance:read", "transactions:read"],
          "ExpirationDateTime": "2024-12-31T23:59:59.000Z"
        }
      },
      "attributes": {
        "consentType": "account_update",
        "customerSegment": "retail"
      }
    },
    "resolvedConsentPurposes": ["accounts:read", "balance:read", "transactions:read"]
  }
}
```

**Features:**
- Similar to consent creation endpoint
- Extracts permissions from `requestPayload.Data.Permissions` array
- Returns them as `resolvedConsentPurposes` in the response

**Example:**
```bash
curl -X POST http://localhost:3001/api/services/pre-process-consent-update \
  -H "Content-Type: application/json" \
  -d @docs/examples/update-account-consent.json
```

## Error Responses

All endpoints return error responses in the following format:

```json
{
  "responseId": "REQ-123456",
  "status": "ERROR",
  "errorMessage": "invalid_request",
  "errorDescription": "Invalid request body"
}
```

**Common Error Codes:**
- `400 Bad Request` - Invalid JSON or malformed request
- `405 Method Not Allowed` - Wrong HTTP method used
- `500 Internal Server Error` - Server-side error

## Testing

### Run Tests
```bash
go test ./test/integration/... -v
```

### Start Server
```bash
# Default port (3001)
go run cmd/server/main.go

# Custom port
PORT=8080 go run cmd/server/main.go

# Or use the compiled binary
./bin/server
```

### Build
```bash
go build -o bin/server cmd/server/main.go
```

## Example Files

Test payload examples are available in `docs/examples/`:
- `basic-account-consent.json` - Basic account access consent
- `payment-consent.json` - Payment consent with authorization
- `recurring-consent.json` - Recurring consent
- `full-params-consent.json` - Consent with all parameters
- `minimal-consent.json` - Minimal consent payload
- `test-purposes-extraction.json` - Example for testing permission extraction
- `update-account-consent.json` - Example for consent update endpoint

## Configuration

Configuration is managed via environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `3001` | Server port |
| `LOG_LEVEL` | `info` | Logging level (debug, info, warn, error) |

You can also create a `.env` file in the project root:

```
PORT=3001
LOG_LEVEL=debug
```
