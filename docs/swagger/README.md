# API Documentation

This directory contains the API documentation for the Consent Service Extensions.

## Files

- **`consent-management-service-extensions.yaml`** - Complete OpenAPI 3.0 specification
- **`README.md`** - This file

## Viewing the API Documentation

### Option 1: Swagger UI (Online)

Visit [Swagger Editor](https://editor.swagger.io/) and paste the contents of `consent-management-service-extensions.yaml`.

### Option 2: Swagger UI (Local)

You can run Swagger UI locally using Docker:

```bash
docker run -p 8081:8080 -e SWAGGER_JSON=/docs/consent-management-service-extensions.yaml \
  -v $(pwd)/docs/swagger:/docs swaggerapi/swagger-ui
```

Then open your browser to: http://localhost:8081

### Option 3: Redoc (Local)

Using Docker:

```bash
docker run -p 8081:80 -e SPEC_URL=docs/consent-management-service-extensions.yaml \
  -v $(pwd)/docs/swagger:/usr/share/nginx/html/docs redocly/redoc
```

Then open your browser to: http://localhost:8081

### Option 4: VS Code Extension

Install the "OpenAPI (Swagger) Editor" extension in VS Code:
1. Open VS Code
2. Go to Extensions (Cmd+Shift+X)
3. Search for "OpenAPI (Swagger) Editor"
4. Install and open the YAML file

## API Endpoints

### Currently Implemented

#### POST /api/services/pre-process-consent-creation
Handle pre validations & obtain custom consent data to be stored.

**Status**: ✅ Implemented

### Planned Endpoints

The following endpoints are defined in the OpenAPI spec and will be implemented:

- POST /api/services/enrich-consent-creation-response
- POST /api/services/pre-process-consent-retrieval
- POST /api/services/pre-process-consent-update
- POST /api/services/enrich-consent-update-response
- POST /api/services/pre-process-consent-revoke
- POST /api/services/pre-process-consent-file-upload
- POST /api/services/enrich-consent-file-response
- POST /api/services/validate-consent-file-retrieval
- POST /api/services/pre-process-consent-file-update
- POST /api/services/enrich-consent-file-update-response
- POST /api/services/map-accelerator-error-response

## Quick Reference

### Authentication

The API supports two authentication methods:
- Basic Authentication
- OAuth2

### Common Response Codes

- **200 OK** - Request succeeded
- **400 Bad Request** - Invalid request data
- **401 Unauthorized** - Authentication required
- **500 Internal Server Error** - Server error

### Request/Response Format

All requests and responses use JSON format with `Content-Type: application/json`.

## Examples

See the main [README.md](../../README.md) for detailed examples and usage instructions.

## OpenAPI Specification Version

- OpenAPI: 3.0.1
- API Version: v1.0.0

## Contact

For more information, visit: https://wso2.com/solutions/financial-services/open-banking/
