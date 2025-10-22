# Consent Service Extensions

A Go web server implementing the Consent Management Service Extensions API based on the OpenAPI 3.0 specification.

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

### Prerequisites

- Go 1.21 or higher
- Make (optional, for using Makefile commands)

### Installation

1. Clone the repository and navigate to the project directory:
```bash
cd consent-service-extensions
```

2. Install dependencies:
```bash
go mod download
```

### Running the Server

**Using Go:**
```bash
go run cmd/server/main.go
```

**Using Make:**
```bash
make run
```

**With custom port:**
```bash
PORT=9090 go run cmd/server/main.go
```

The server will start on `http://localhost:8080` (or your specified port)

### Building the Binary

**Using Go:**
```bash
go build -o bin/server cmd/server/main.go
```

**Using Make:**
```bash
make build
```

Then run:
```bash
./bin/server
```

## 🧪 Testing

### Run all tests:
```bash
go test ./...
```

### Run tests with coverage:
```bash
go test -cover ./...
```

### Run tests with verbose output:
```bash
go test -v ./...
```

### Generate coverage report:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
open coverage.html
```

### Using Make:
```bash
make test
make test-coverage
```

### Test a specific package:
```bash
go test ./test/handlers/...
go test ./test/api/...
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

## 🔧 Available Make Commands

```bash
make build          # Build the application
make run            # Run the application
make test           # Run tests
make test-coverage  # Run tests with coverage report
make clean          # Clean build artifacts
make help           # Show available commands
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
