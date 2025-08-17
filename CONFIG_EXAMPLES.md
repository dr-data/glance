# Glance Serverless Configuration Examples

This file contains example configurations for the Vercel serverless deployment of Glance.

## Basic Configuration

Set these environment variables in your Vercel dashboard:

### Minimal Setup
```bash
GLANCE_PAGES=[{"name":"Dashboard","slug":"","columns":[]}]
GLANCE_PROXIED=true
GLANCE_BASE_URL=https://your-app.vercel.app
```

### Complete Dashboard Example
```bash
GLANCE_PAGES=[
  {
    "name": "Home",
    "slug": "",
    "columns": [
      {
        "size": "full",
        "widgets": [
          {
            "type": "weather",
            "data": {
              "location": "London, UK",
              "units": "metric"
            }
          },
          {
            "type": "clock",
            "data": {
              "format": "12h",
              "timezone": "Europe/London"
            }
          }
        ]
      }
    ]
  },
  {
    "name": "Monitoring",
    "slug": "monitoring",
    "columns": [
      {
        "size": "half",
        "widgets": [
          {
            "type": "server-stats",
            "data": {
              "servers": [
                {
                  "name": "Production",
                  "url": "https://api.example.com",
                  "type": "remote"
                }
              ]
            }
          }
        ]
      }
    ]
  }
]
```

## Authentication Configuration

### 1. Generate Secret Key
Run locally to generate a secret:
```bash
go run main.go secret
```

### 2. Generate Password Hash
Run locally to hash a password:
```bash
go run main.go hash-password "your-secure-password"
```

### 3. Set Environment Variables
```bash
GLANCE_AUTH_SECRET=<your-generated-secret>
GLANCE_AUTH_USERS={"admin":{"password_hash":"<your-password-hash>"}}
```

### Multi-User Example
```bash
GLANCE_AUTH_USERS={
  "admin": {
    "password_hash": "$2a$10$..."
  },
  "viewer": {
    "password_hash": "$2a$10$..."
  }
}
```

## Widget Configuration Examples

### Weather Widget
```json
{
  "type": "weather",
  "data": {
    "location": "New York, NY",
    "units": "imperial",
    "api_key": "your-api-key"
  }
}
```

### RSS Feed Widget
```json
{
  "type": "rss",
  "data": {
    "url": "https://feeds.example.com/rss",
    "title": "Latest News",
    "limit": 10
  }
}
```

### Bookmarks Widget
```json
{
  "type": "bookmarks",
  "data": {
    "groups": [
      {
        "title": "Development",
        "bookmarks": [
          {
            "title": "GitHub",
            "url": "https://github.com"
          },
          {
            "title": "Stack Overflow",
            "url": "https://stackoverflow.com"
          }
        ]
      }
    ]
  }
}
```

### Server Stats Widget (Remote)
```json
{
  "type": "server-stats",
  "data": {
    "servers": [
      {
        "name": "Web Server",
        "url": "https://api.server.com",
        "type": "remote",
        "token": "bearer-token",
        "timeout": "5s"
      }
    ]
  }
}
```

## Theme Configuration

### Basic Theme Settings
```bash
GLANCE_DISABLE_THEME_PICKER=false
```

### Custom Styling
Modify `/public/css/main.css` and redeploy.

## Environment Variable Reference

### Server Configuration
- `GLANCE_HOST=""` - Host binding (unused in serverless)
- `GLANCE_PORT=8080` - Port (unused in serverless)  
- `GLANCE_PROXIED=true` - Enable proxy headers
- `GLANCE_BASE_URL=""` - Base URL for assets

### Authentication
- `GLANCE_AUTH_SECRET=""` - Base64 encoded secret key
- `GLANCE_AUTH_USERS={}` - JSON object of users

### Theme
- `GLANCE_DISABLE_THEME_PICKER=false` - Disable theme picker

### Pages
- `GLANCE_PAGES=[]` - JSON array of page configurations

## Local Development

### 1. Set Environment Variables
Create `.env` file:
```
GLANCE_PAGES=[{"name":"Dashboard","slug":"","columns":[]}]
GLANCE_PROXIED=true
GLANCE_BASE_URL=http://localhost:3000
```

### 2. Run Vercel Dev
```bash
vercel dev
```

### 3. Test Endpoints
- http://localhost:3000 - Main dashboard
- http://localhost:3000/api/healthz - Health check
- http://localhost:3000/login - Login page

## Production Deployment

### 1. Set Environment Variables
```bash
vercel env add GLANCE_PAGES production
vercel env add GLANCE_PROXIED production
vercel env add GLANCE_BASE_URL production
```

### 2. Deploy
```bash
vercel --prod
```

### 3. Verify
Check that all endpoints respond correctly:
- Main dashboard loads
- Static assets serve correctly
- API endpoints function
- Authentication works (if configured)

## Migration Notes

### From Docker
1. Export current `glance.yml` configuration
2. Convert YAML to JSON format for environment variables
3. Test in staging environment first
4. Update DNS/proxy configurations

### Limitations vs Docker Version
- No file watching/auto-reload
- No in-memory caching between requests
- No background workers
- Stateless operation only
- 10-second execution timeout per request

### Benefits of Serverless
- Zero-downtime deployments
- Automatic scaling
- No server management
- Built-in CDN for static assets
- Free tier available