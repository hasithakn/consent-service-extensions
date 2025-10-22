# Consent Service Extensions

A Go web server implementing the Consent Management Service Extensions API based on the OpenAPI 3.0 specification.

> 📋 **See [STRUCTURE.md](STRUCTURE.md) for detailed project structure and organization**

## 📁 Project Structure

```
consent-service-extensions/
├── cmd/
│   └── server/              # Application entry points
│       └── main.go          # Main application
├── internal/                # Private application code
│   ├── handlers/            # HTTP request handlers
│   │   └── consent_handler.go
│   └── models/              # Data models
│       └── consent.go
├── pkg/                     # Public libraries
│   └── api/                 # API router configuration
│       └── router.go
├── test/                    # Test files (separate from source)
│   ├── handlers/            # Handler tests
│   │   └── consent_handler_test.go
│   └── api/                 # API/Router tests
│       └── router_test.go
├── docs/                    # Documentation
│   ├── swagger/             # OpenAPI/Swagger specifications
│   │   ├── consent-management-service-extensions.yaml
│   │   └── README.md
│   └── examples/            # API request/response examples
│       ├── basic-account-consent.json
│       ├── payment-consent.json
│       ├── recurring-consent.json
│       └── README.md
├── bin/                     # Compiled binaries (gitignored)
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums
├── .env.example            # Example environment variables
├── .gitignore              # Git ignore rules
├── Makefile                # Build automation
└── README.md               # This file
```

## 🚀 Getting Started

### Quick Start

1. **Clone and navigate to the project:**
   ```bash
   cd consent-service-extensions
   go mod download
   ```

2. **Start the server:**
   ```bash
   go run cmd/server/main.go
   ```
   Server starts on `http://localhost:8080`

3. **Test the API** (in another terminal):
   ```bash
   # Health check
   curl http://localhost:8080/health
   
   # Test consent creation
   curl -X POST http://localhost:8080/api/services/pre-process-consent-creation 
     -H "Content-Type: application/json" 
     -d @docs/examples/basic-account-consent.json
   ```

4. **Run integration tests** (server must be running):
   ```bash
   go test ./test/integration/... -v
   ```

### Prerequisites

- Go 1.21 or higher

### Running the Server

```bash
go run cmd/server/main.go
```

**With custom port:**
```bash
PORT=9090 go run cmd/server/main.go
```

The server will start on `http://localhost:8080` (or your specified port)

### Building the Binary

```bash
go build -o bin/server cmd/server/main.go
```

Then run:
```bash
./bin/server
```

## 🧪 Testing

### Run Integration Tests

**Important:** Start the server first in a separate terminal:
```bash
go run cmd/server/main.go
```

Then run the integration tests:
```bash
go test ./test/integration/... -v
```

### Run All Tests

```bash
go test ./... -v
```

### Run Tests with Coverage

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
open coverage.html
```

### Run Specific Test

```bash
go test ./test/integration/... -v -run TestHealthEndpoint
go test ./test/integration/... -v -run TestPreProcessConsentCreation_BasicAccountConsent
```

### Test with Custom Server URL

```bash
TEST_BASE_URL=http://localhost:9090 go test ./test/integration/... -v
```

## 📡 API Endpoints

### Health Check
**GET** `/health`

Returns the health status of the service.

```bash
curl http://localhost:8080/health
```

Response:
```json
{
  "status": "healthy"
}
```

### Pre-Process Consent Creation
**POST** `/api/services/pre-process-consent-creation`

Handle pre validations & obtain custom consent data to be stored.

**Request Example:**
```json
{
  "requestId": "Ec1wMjmiG8",
  "data": {
    "consentInitiationData": {
      "type": "accounts",
      "status": "AwaitingAuthorisation",
      "validityTime": 0,
      "recurringIndicator": false,
      "frequency": 0,
      "dataAccessValidityDuration": 86400,
      "requestPayload": {},
      "attributes": {},
      "authorizations": []
    },
    "requestHeaders": {}
  }
}
```

**Success Response (200):**
```json
{
  "responseId": "Ec1wMjmiG8",
  "status": "SUCCESS",
  "data": {
    "consentResource": {
      "type": "accounts",
      "status": "AwaitingAuthorisation",
      "validityTime": 0,
      "recurringIndicator": false,
      "frequency": 0,
      "dataAccessValidityDuration": 86400,
      "requestPayload": {},
      "attributes": {},
      "authorizations": []
    },
    "resolvedConsentPurposes": []
  }
}
```

**Error Response (400):**
```json
{
  "responseId": "Ec1wMjmiG8",
  "status": "ERROR",
  "errorMessage": "invalid_request",
  "errorDescription": "Invalid request body"
}
```

### Testing with cURL:

```bash
curl -X POST http://localhost:8080/api/services/pre-process-consent-creation \
  -H "Content-Type: application/json" \
  -d '{
    "requestId": "Ec1wMjmiG8",
    "data": {
      "consentInitiationData": {
        "type": "accounts",
        "status": "AwaitingAuthorisation",
        "validityTime": 0,
        "recurringIndicator": false,
        "frequency": 0,
        "requestPayload": {}
      },
      "requestHeaders": {}
    }
  }'
```

## 🛠️ Development

### Adding New Endpoints

1. Define models in `internal/models/`
2. Create handler methods in `internal/handlers/`
3. Register routes in `pkg/api/router.go`
4. Add tests in corresponding `*_test.go` files

### Code Organization

- **`cmd/`**: Application entry points. Each subdirectory is a separate binary.
- **`internal/`**: Private application code that cannot be imported by other projects.
- **`pkg/`**: Public libraries that can be imported by other projects.
- **`internal/handlers/`**: HTTP request handlers with business logic.
- **`internal/models/`**: Data structures and models.
- **`pkg/api/`**: API routing and middleware configuration.

## 📝 Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8080` |

## 🔧 Development Commands

```bash
# Build the application
go build -o bin/server cmd/server/main.go

# Run the application
go run cmd/server/main.go

# Run tests (requires server to be running)
go test ./test/integration/... -v

# Run tests with coverage
go test -coverprofile=coverage.out ./...

# Format code
go fmt ./...

# Vet code
go vet ./...

# Tidy dependencies
go mod tidy
```

## 📖 API Documentation

Complete API documentation is available in the `docs/` directory:

- **[OpenAPI/Swagger Specification](docs/swagger/consent-management-service-extensions.yaml)** - Complete API contract
- **[Swagger Documentation Guide](docs/swagger/README.md)** - How to view and use the API docs
- **[API Examples](docs/examples/README.md)** - Request/response examples

### View API Documentation

You can view the API documentation using:

1. **Swagger Editor Online**: Visit [editor.swagger.io](https://editor.swagger.io/) and paste the OpenAPI spec
2. **Swagger UI (Docker)**:
   ```bash
   docker run -p 8081:8080 -e SWAGGER_JSON=/docs/consent-management-service-extensions.yaml \
     -v $(pwd)/docs/swagger:/docs swaggerapi/swagger-ui
   ```
   Then open: http://localhost:8081

3. **VS Code Extension**: Install "OpenAPI (Swagger) Editor" extension

## 📚 Future Endpoints (TODO)

Based on the OpenAPI spec, the following endpoints will be added:

- `/enrich-consent-creation-response`
- `/pre-process-consent-retrieval`
- `/pre-process-consent-update`
- `/enrich-consent-update-response`
- `/pre-process-consent-revoke`
- `/pre-process-consent-file-upload`
- `/enrich-consent-file-response`
- `/validate-consent-file-retrieval`
- `/pre-process-consent-file-update`
- `/enrich-consent-file-update-response`
- `/map-accelerator-error-response`

## 📄 License

This project follows the Apache 2.0 license as specified in the OpenAPI specification.

## 🤝 Contributing

1. Follow Go best practices and idioms
2. Write unit tests for new functionality
3. Update documentation as needed
4. Run `go fmt` before committing
5. Ensure all tests pass with `make test`
