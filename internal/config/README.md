# Configuration Package

This package handles loading configuration from environment variables and `.env` files.

## Features

- ✅ Loads from `.env` file if present
- ✅ Falls back to environment variables
- ✅ Provides default values
- ✅ Environment variables take precedence over `.env` file

## Usage

```go
import "consent-service-extensions/internal/config"

func main() {
    cfg := config.Load()
    fmt.Printf("Server port: %s\n", cfg.Port)
    fmt.Printf("Log level: %s\n", cfg.LogLevel)
}
```

## Configuration Priority

1. **Environment variables** (highest priority)
2. **`.env` file** values
3. **Default values** (lowest priority)

## Available Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `3001` | Server port |
| `LOG_LEVEL` | `info` | Logging level (debug, info, warn, error) |

## Setup

1. **Copy the example file:**
   ```bash
   cp .env.example .env
   ```

2. **Edit `.env` with your values:**
   ```bash
   # .env
   PORT=3001
   LOG_LEVEL=debug
   ```

3. **Or use environment variables:**
   ```bash
   PORT=8080 go run cmd/server/main.go
   ```

## Development vs Production

### Development
Use `.env` file for local development:
```bash
# .env
PORT=3001
LOG_LEVEL=debug
```

### Production
Use environment variables (don't commit `.env` to git):
```bash
export PORT=8080
export LOG_LEVEL=info
./bin/server
```

## Adding New Configuration

1. Add to `Config` struct in `config.go`:
   ```go
   type Config struct {
       Port     string
       LogLevel string
       NewField string  // Add here
   }
   ```

2. Load in `Load()` function:
   ```go
   cfg := &Config{
       Port:     getEnv("PORT", "3001"),
       LogLevel: getEnv("LOG_LEVEL", "info"),
       NewField: getEnv("NEW_FIELD", "default"),  // Add here
   }
   ```

3. Update `.env.example`:
   ```bash
   # New Configuration
   NEW_FIELD=value
   ```

## Notes

- `.env` file is **gitignored** by default
- `.env.example` should be committed to git
- Always provide sensible defaults
- Environment variables override `.env` file values
