# API Examples

This directory contains example requests and responses for testing the Consent Service Extensions API.

## Pre-Process Consent Creation

### Example 1: Basic Account Consent

**Request:**
```json
{
  "requestId": "req-001",
  "data": {
    "consentInitiationData": {
      "type": "accounts",
      "status": "AwaitingAuthorisation",
      "validityTime": 0,
      "recurringIndicator": false,
      "frequency": 0,
      "dataAccessValidityDuration": 86400,
      "requestPayload": {
        "Data": {
          "Permissions": [
            "ReadAccountsBasic",
            "ReadAccountsDetail",
            "ReadBalances"
          ]
        }
      },
      "attributes": {}
    },
    "requestHeaders": {
      "x-fapi-financial-id": "bank-123",
      "x-fapi-interaction-id": "interaction-456"
    }
  }
}
```

**Response:**
```json
{
  "responseId": "req-001",
  "status": "SUCCESS",
  "data": {
    "consentResource": {
      "type": "accounts",
      "status": "AwaitingAuthorisation",
      "validityTime": 0,
      "recurringIndicator": false,
      "frequency": 0,
      "dataAccessValidityDuration": 86400,
      "requestPayload": {
        "Data": {
          "Permissions": [
            "ReadAccountsBasic",
            "ReadAccountsDetail",
            "ReadBalances"
          ]
        }
      },
      "attributes": {}
    },
    "resolvedConsentPurposes": [
      "PURPOSE-ACCOUNT-ACCESS"
    ]
  }
}
```

### Example 2: Payment Consent with Authorization

**Request:**
```json
{
  "requestId": "req-002",
  "data": {
    "consentInitiationData": {
      "type": "payments",
      "status": "AwaitingAuthorisation",
      "validityTime": 3600,
      "recurringIndicator": false,
      "frequency": 0,
      "dataAccessValidityDuration": 3600,
      "requestPayload": {
        "Data": {
          "Initiation": {
            "InstructedAmount": {
              "Amount": "100.00",
              "Currency": "USD"
            },
            "CreditorAccount": {
              "SchemeName": "IBAN",
              "Identification": "GB76LOYD30949301273801"
            }
          }
        }
      },
      "attributes": {
        "paymentType": "single",
        "merchantId": "merchant-789"
      },
      "authorizations": [
        {
          "userId": "user@example.com",
          "type": "authorisation",
          "status": "created",
          "resource": {
            "authMethod": "SMS"
          }
        }
      ]
    },
    "requestHeaders": {
      "x-idempotency-key": "idempotency-123",
      "x-request-id": "request-456"
    }
  }
}
```

**Response:**
```json
{
  "responseId": "req-002",
  "status": "SUCCESS",
  "data": {
    "consentResource": {
      "type": "payments",
      "status": "AwaitingAuthorisation",
      "validityTime": 3600,
      "recurringIndicator": false,
      "frequency": 0,
      "dataAccessValidityDuration": 3600,
      "requestPayload": {
        "Data": {
          "Initiation": {
            "InstructedAmount": {
              "Amount": "100.00",
              "Currency": "USD"
            },
            "CreditorAccount": {
              "SchemeName": "IBAN",
              "Identification": "GB76LOYD30949301273801"
            }
          }
        }
      },
      "attributes": {
        "paymentType": "single",
        "merchantId": "merchant-789"
      },
      "authorizations": [
        {
          "userId": "user@example.com",
          "type": "authorisation",
          "status": "created",
          "resource": {
            "authMethod": "SMS"
          }
        }
      ]
    },
    "resolvedConsentPurposes": [
      "PURPOSE-PAYMENT-INITIATION"
    ]
  }
}
```

### Example 3: Error Response

**Request:**
```json
{
  "requestId": "req-003",
  "data": {
    "consentInitiationData": {
      "type": "invalid-type"
    }
  }
}
```

**Response:**
```json
{
  "responseId": "req-003",
  "status": "ERROR",
  "errorCode": 400,
  "data": {
    "errorMessage": "invalid_request",
    "errorDescription": "Invalid consent type provided"
  }
}
```

## Testing with cURL

### Basic Request
```bash
curl -X POST http://localhost:8080/api/services/pre-process-consent-creation \
  -H "Content-Type: application/json" \
  -H "Authorization: Basic $(echo -n 'user:pass' | base64)" \
  -d @examples/basic-account-consent.json
```

### With File
```bash
curl -X POST http://localhost:8080/api/services/pre-process-consent-creation \
  -H "Content-Type: application/json" \
  -d @docs/examples/basic-account-consent.json
```

## Testing with HTTPie

```bash
http POST http://localhost:8080/api/services/pre-process-consent-creation \
  Content-Type:application/json \
  < docs/examples/basic-account-consent.json
```

## Testing with Postman

1. Import the OpenAPI spec: `docs/swagger/consent-management-service-extensions.yaml`
2. Postman will automatically create a collection with all endpoints
3. Use the examples above as request bodies
