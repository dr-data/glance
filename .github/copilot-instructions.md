# Glance Dashboard Application

Glance is a Go-based web application that provides a customizable dashboard with various widgets (RSS feeds, GitHub releases, weather, system stats, etc.). The application compiles to a single binary with embedded static assets and serves web dashboards via HTTP.

Always reference these instructions first and fallback to search or bash commands only when you encounter unexpected information that does not match the info here.

## Working Effectively

### Prerequisites and Setup
- Install Go >= 1.23 (we use Go 1.24+)
- Git for version control
- Optional: Docker for containerized builds (may fail in restricted environments due to proxy/TLS issues)

### Bootstrap, Build, and Test the Repository
1. Clone and navigate to repository:
   ```bash
   git clone <repository-url>
   cd glance
   ```

2. Build the application:
   ```bash
   go build -o build/glance .
   ```
   - Takes approximately 1-2 seconds on modern systems. NEVER CANCEL. Set timeout to 30+ seconds for safety.

3. Run tests:
   ```bash
   go test ./...
   ```
   - Takes approximately 1 second. NEVER CANCEL. Set timeout to 30+ seconds for safety.

4. Format and lint code:
   ```bash
   go fmt ./...
   go vet ./...
   ```
   - Takes less than 1 second each.

### Development Workflow
- For development/testing without creating binary:
  ```bash
  go run .
  ```
  - Starts immediately and serves on localhost:8080

- Validate configuration files:
  ```bash
  ./build/glance config:validate
  ./build/glance --config /path/to/config.yml config:validate
  ```

- Run diagnostics (useful for troubleshooting):
  ```bash
  ./build/glance diagnose
  ```

### Docker Build (Optional)
- Docker build may fail in environments with proxy/firewall restrictions:
  ```bash
  docker build -t glance:latest .
  ```
  - Takes 2-5 minutes when successful. NEVER CANCEL. Set timeout to 10+ minutes.
  - If build fails with TLS/certificate errors, document this limitation and use Go build instead.

## Validation

### ALWAYS run through complete end-to-end scenarios after making changes:

1. **Basic Functionality Test**:
   ```bash
   # Create minimal test config
   echo 'pages:
     - name: Test Page
       columns:
         - size: full
           widgets:
             - type: calendar' > test-config.yml
   
   # Validate config
   ./build/glance --config test-config.yml config:validate
   
   # Start server (background)
   ./build/glance --config test-config.yml &
   
   # Test HTTP response
   curl -I http://localhost:8080
   
   # Should return HTTP 200 with HTML content
   # Verify dashboard content
   curl -s http://localhost:8080 | grep -o '<title>.*</title>'
   
   # Stop server when done (if running in background)
   # Clean up: rm test-config.yml
   ```

2. **Configuration Validation Scenarios**:
   - Test config validation with the example config: `cp docs/glance.yml . && ./build/glance config:validate`
   - Test with invalid config to ensure proper error handling:
     ```bash
     echo 'invalid: yaml: content' > invalid-config.yml
     ./build/glance --config invalid-config.yml config:validate
     # Should exit with error code 1 and show YAML parsing error
     ```
   - Test environment variable substitution if making config changes

3. **Build Validation**:
   - Always run `go fmt ./...` and `go vet ./...` before committing
   - Ensure tests pass: `go test ./...`
   - Verify clean build: `rm -rf build/ && go build -o build/glance .`

4. **Manual Testing Requirements**:
   - Start the application and verify it serves content on localhost:8080
   - Verify HTTP 200 response: `curl -I http://localhost:8080`
   - Check dashboard title is present: `curl -s http://localhost:8080 | grep -o '<title>.*</title>'`
   - Verify widgets render properly in the HTML output
   - Test basic widget functionality like calendar display
   - Verify configuration reload if modifying config handling
   - Test error handling with invalid configuration files

## Common Tasks

### Repository Structure
```
/home/runner/work/glance/glance/
├── .github/workflows/release.yaml    # Release automation
├── Dockerfile                        # Container build
├── docs/                            # Documentation and example configs
│   ├── configuration.md            # Widget configuration reference
│   └── glance.yml                  # Example configuration
├── go.mod                          # Go module definition
├── main.go                         # Application entry point
├── internal/glance/                # Core application code
│   ├── main.go                     # Main application logic
│   ├── config.go                   # Configuration handling
│   ├── widget-*.go                 # Widget implementations
│   ├── auth.go                     # Authentication
│   └── auth_test.go               # Tests (currently minimal)
└── pkg/sysinfo/                    # System information utilities
```

### Key Files for Common Changes
- Widget implementations: `internal/glance/widget-*.go`
- Configuration schema: `internal/glance/config-fields.go` and `internal/glance/config.go`
- Main server logic: `internal/glance/main.go` and `internal/glance/glance.go`
- Templates and UI: `internal/glance/templates/` and `internal/glance/static/`
- Authentication: `internal/glance/auth.go`

### Configuration File Requirements
- Application requires a valid `glance.yml` configuration file to start
- Default config location: `./glance.yml`
- Override with: `--config /path/to/config.yml`
- Example config available at: `docs/glance.yml`
- Configuration supports environment variables: `${ENV_VAR}` syntax

### Build Artifacts and Dependencies
- Binary output: `build/glance` (single executable with embedded assets)
- Go modules: Dependencies defined in `go.mod`, downloaded automatically
- No external build tools required (no npm, Make, etc.)
- Static assets are embedded at build time

### Development Tips
- Application starts immediately (no lengthy startup)
- Changes require rebuild: `go build -o build/glance .`
- For rapid iteration use: `go run .`
- Default port: 8080 (configurable via config file)
- Application includes built-in help: `./build/glance --help`

### Troubleshooting Common Issues
1. **Build Errors**: Ensure Go >= 1.23 is installed and all dependencies can be downloaded
2. **Network Connectivity**: Use `./build/glance diagnose` to check external API connectivity
3. **Configuration Errors**: Use `./build/glance config:validate` to check YAML syntax and logic
4. **Docker Build Fails**: May fail in restricted environments; use native Go build instead
5. **Permission Errors**: Ensure binary has execute permissions: `chmod +x build/glance`

### Widget Development
- Each widget type has its own file: `internal/glance/widget-[name].go`
- Widget configuration structs defined in `internal/glance/config-fields.go`
- Widget templates in `internal/glance/templates/`
- Always test new widgets with minimal configuration first

## Critical Timing and Timeout Guidelines
- **Build Command**: Takes 1-2 seconds. Set timeout to 30+ seconds minimum.
- **Test Suite**: Takes ~1 second. Set timeout to 30+ seconds minimum.
- **Docker Build**: Takes 2-5 minutes when successful. Set timeout to 10+ minutes.
- **Application Startup**: Immediate. No extended startup time.
- **NEVER CANCEL** any build or test commands prematurely.

## CI/CD Integration
- GitHub Actions workflow: `.github/workflows/release.yaml`
- Automated releases use GoReleaser: `.goreleaser.yaml`
- No additional CI checks beyond Go build and test
- Always run `go fmt ./...` and `go vet ./...` before pushing changes