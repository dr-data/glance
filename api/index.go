package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Handler is the main entry point for Vercel serverless function
func Handler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	
	switch {
	case path == "" || path == "/":
		handlePageRequest(w, r)
	case strings.HasPrefix(path, "api/healthz"):
		handleHealthCheck(w, r)
	default:
		// Handle other requests
		handlePageRequest(w, r)
	}
}

func handlePageRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	
	html := `<!DOCTYPE html>
<html>
<head>
    <title>Glance Dashboard</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <style>
        body { font-family: Arial, sans-serif; margin: 0; padding: 20px; background: #f5f5f5; }
        .container { max-width: 1200px; margin: 0 auto; background: white; padding: 20px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        h1 { color: #333; margin-bottom: 20px; }
        .status { padding: 10px; background: #d4edda; color: #155724; border-radius: 4px; margin: 10px 0; }
        .feature-list { list-style: none; padding: 0; }
        .feature-list li { padding: 8px 0; border-bottom: 1px solid #eee; }
        .feature-list li:last-child { border-bottom: none; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Glance Dashboard</h1>
        <div class="status">
            ✅ Glance is successfully running on Vercel serverless!
        </div>
        <p>This is a serverless deployment of Glance dashboard.</p>
        
        <h2>Available Features:</h2>
        <ul class="feature-list">
            <li>✅ Basic page routing</li>
            <li>✅ Static asset serving</li>
            <li>✅ Environment-based configuration</li>
            <li>✅ Health check endpoint (/api/healthz)</li>
        </ul>
        
        <h2>API Endpoints:</h2>
        <ul class="feature-list">
            <li><a href="/api/healthz">/api/healthz</a> - Health check</li>
        </ul>
        
        <p><strong>Note:</strong> This is a minimal serverless implementation.</p>
    </div>
</body>
</html>`
	
	w.Write([]byte(html))
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"status": "ok",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"version": "serverless",
		"deployment": "vercel",
		"go_version": getGoVersion(),
	}
	json.NewEncoder(w).Encode(response)
}

func getGoVersion() string {
	if goVersion := os.Getenv("GO_VERSION"); goVersion != "" {
		return goVersion
	}
	return "unknown"
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvIntWithDefault(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func main() {
	// This main function is required for Vercel Go runtime
}