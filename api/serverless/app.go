package serverless

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// ServerlessApp wraps the glance application for serverless deployment
type ServerlessApp struct {
	config      *Config
	pages       []*Page  
	widgets     map[uint64]Widget
	requiresAuth bool
	secretKey   []byte
	users       map[string]*User
}

// Config represents the simplified configuration for serverless
type Config struct {
	Server ServerConfig           `json:"server"`
	Auth   AuthConfig            `json:"auth"`
	Theme  ThemeConfig           `json:"theme"`
	Pages  []PageConfig          `json:"pages"`
}

type ServerConfig struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Proxied bool   `json:"proxied"`
	BaseURL string `json:"base_url"`
}

type AuthConfig struct {
	SecretKey string              `json:"secret_key"`
	Users     map[string]*User    `json:"users"`
}

type ThemeConfig struct {
	DisablePicker bool `json:"disable_picker"`
}

type PageConfig struct {
	Name     string          `json:"name"`
	Slug     string          `json:"slug"`
	Columns  []ColumnConfig  `json:"columns"`
}

type ColumnConfig struct {
	Size    string          `json:"size"`
	Widgets []WidgetConfig  `json:"widgets"`
}

type WidgetConfig struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type User struct {
	PasswordHash string `json:"password_hash"`
}

type Page struct {
	Name string
	Slug string
}

type Widget interface {
	GetType() string
}

// NewServerlessApp creates a new serverless application from environment variables
func NewServerlessApp() (*ServerlessApp, error) {
	app := &ServerlessApp{
		widgets: make(map[uint64]Widget),
		users:   make(map[string]*User),
	}

	// Load configuration from environment
	if err := app.loadConfigFromEnv(); err != nil {
		return nil, fmt.Errorf("loading config from env: %w", err)
	}

	// Initialize pages and widgets
	if err := app.initializePagesAndWidgets(); err != nil {
		return nil, fmt.Errorf("initializing pages and widgets: %w", err)
	}

	return app, nil
}

func (app *ServerlessApp) loadConfigFromEnv() error {
	// Load basic server config
	config := &Config{
		Server: ServerConfig{
			Host:    getEnvWithDefault("GLANCE_HOST", ""),
			Port:    getEnvIntWithDefault("GLANCE_PORT", 8080),
			Proxied: getEnvBoolWithDefault("GLANCE_PROXIED", true),
			BaseURL: getEnvWithDefault("GLANCE_BASE_URL", ""),
		},
		Theme: ThemeConfig{
			DisablePicker: getEnvBoolWithDefault("GLANCE_DISABLE_THEME_PICKER", false),
		},
	}

	// Load auth config
	if secretKey := os.Getenv("GLANCE_AUTH_SECRET"); secretKey != "" {
		config.Auth.SecretKey = secretKey
		
		if usersJSON := os.Getenv("GLANCE_AUTH_USERS"); usersJSON != "" {
			var users map[string]*User
			if err := json.Unmarshal([]byte(usersJSON), &users); err != nil {
				return fmt.Errorf("parsing auth users: %w", err)
			}
			config.Auth.Users = users
			app.requiresAuth = true

			// Parse secret key
			secretBytes, err := base64.StdEncoding.DecodeString(secretKey)
			if err != nil {
				return fmt.Errorf("decoding secret key: %w", err)
			}
			app.secretKey = secretBytes
		}
	}

	// Load pages config
	if pagesJSON := os.Getenv("GLANCE_PAGES"); pagesJSON != "" {
		var pages []PageConfig
		if err := json.Unmarshal([]byte(pagesJSON), &pages); err != nil {
			return fmt.Errorf("parsing pages config: %w", err)
		}
		config.Pages = pages
	}

	app.config = config
	return nil
}

func (app *ServerlessApp) initializePagesAndWidgets() error {
	// Convert page configs to pages
	if len(app.config.Pages) == 0 {
		// Create a default dashboard page
		app.pages = []*Page{{
			Name: "Dashboard",
			Slug: "",
		}}
	} else {
		for _, pageConfig := range app.config.Pages {
			app.pages = append(app.pages, &Page{
				Name: pageConfig.Name,
				Slug: pageConfig.Slug,
			})
		}
	}
	
	return nil
}

// Handler is the main entry point for serverless requests
func (app *ServerlessApp) Handler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	
	switch {
	case path == "" || path == "/":
		app.handlePageRequest(w, r, "")
	case strings.HasPrefix(path, "api/"):
		app.handleAPIRequest(w, r)
	case strings.HasPrefix(path, "login"):
		app.handleLoginPageRequest(w, r)
	case strings.HasPrefix(path, "logout"):
		app.handleLogoutRequest(w, r)
	default:
		// Handle named page requests
		app.handlePageRequest(w, r, path)
	}
}

func (app *ServerlessApp) handlePageRequest(w http.ResponseWriter, r *http.Request, slug string) {
	// Find the page
	var page *Page
	if slug == "" {
		// Return first page as default
		if len(app.pages) > 0 {
			page = app.pages[0]
		}
	} else {
		for _, p := range app.pages {
			if p.Slug == slug {
				page = p
				break
			}
		}
	}

	if page == nil {
		http.NotFound(w, r)
		return
	}

	// For now, return a simple HTML response
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	
	html := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>%s - Glance</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="/css/main.css">
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
        <h1>%s</h1>
        <div class="status">
            ✅ Glance is successfully running on Vercel serverless!
        </div>
        <p>This is a serverless deployment of Glance dashboard. The application has been adapted to run as stateless serverless functions.</p>
        
        <h2>Available Features:</h2>
        <ul class="feature-list">
            <li>✅ Basic page routing</li>
            <li>✅ Static asset serving</li>
            <li>✅ Environment-based configuration</li>
            <li>✅ Health check endpoint (/api/healthz)</li>
            <li>🔄 Widget system (in progress)</li>
            <li>🔄 Authentication (in progress)</li>
            <li>🔄 Theme customization (in progress)</li>
        </ul>
        
        <h2>API Endpoints:</h2>
        <ul class="feature-list">
            <li><a href="/api/healthz">/api/healthz</a> - Health check</li>
            <li>/api/pages/{page}/content - Page content API</li>
            <li>/api/authenticate - Authentication endpoint</li>
        </ul>
        
        <p><strong>Note:</strong> This is a minimal serverless implementation. Some features from the original Docker version may not be available due to serverless constraints.</p>
    </div>
</body>
</html>`, page.Name, page.Name)
	
	w.Write([]byte(html))
}

func (app *ServerlessApp) handleAPIRequest(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/")
	
	switch {
	case path == "healthz":
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]interface{}{
			"status": "ok",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
			"version": "serverless",
			"deployment": "vercel",
		}
		json.NewEncoder(w).Encode(response)
	case strings.HasPrefix(path, "pages/"):
		app.handlePageContentRequest(w, r)
	case path == "authenticate":
		app.handleAuthenticationAttempt(w, r)
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "endpoint not found"})
	}
}

func (app *ServerlessApp) handlePageContentRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"content": "Page content API endpoint",
		"message": "This endpoint will provide page content for AJAX updates",
		"status": "placeholder",
	}
	json.NewEncoder(w).Encode(response)
}

func (app *ServerlessApp) handleLoginPageRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	html := `<!DOCTYPE html>
<html>
<head>
    <title>Login - Glance</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <style>
        body { font-family: Arial, sans-serif; margin: 0; padding: 20px; background: #f5f5f5; }
        .login-container { max-width: 400px; margin: 100px auto; background: white; padding: 30px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        h1 { text-align: center; color: #333; margin-bottom: 30px; }
        .form-group { margin-bottom: 15px; }
        label { display: block; margin-bottom: 5px; color: #555; }
        input { width: 100%; padding: 10px; border: 1px solid #ddd; border-radius: 4px; box-sizing: border-box; }
        button { width: 100%; padding: 10px; background: #007bff; color: white; border: none; border-radius: 4px; cursor: pointer; }
        button:hover { background: #0056b3; }
    </style>
</head>
<body>
    <div class="login-container">
        <h1>Login to Glance</h1>
        <form action="/api/authenticate" method="post">
            <div class="form-group">
                <label for="username">Username:</label>
                <input type="text" id="username" name="username" required>
            </div>
            <div class="form-group">
                <label for="password">Password:</label>
                <input type="password" id="password" name="password" required>
            </div>
            <button type="submit">Login</button>
        </form>
        <p style="text-align: center; margin-top: 20px; color: #666;">
            <small>Authentication system for serverless deployment</small>
        </p>
    </div>
</body>
</html>`
	w.Write([]byte(html))
}

func (app *ServerlessApp) handleLogoutRequest(w http.ResponseWriter, r *http.Request) {
	// Clear any auth cookies and redirect to home
	http.SetCookie(w, &http.Cookie{
		Name:   "auth",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})
	http.Redirect(w, r, "/", http.StatusFound)
}

func (app *ServerlessApp) handleAuthenticationAttempt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "method not allowed"})
		return
	}

	// For now, return a placeholder response
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"success": false,
		"message": "Authentication not fully implemented yet",
		"note": "This is a placeholder for serverless authentication",
	}
	json.NewEncoder(w).Encode(response)
}

// Utility functions
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

func getEnvBoolWithDefault(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}