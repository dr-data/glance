# Glance Vercel Serverless Deployment Guide

This guide provides complete instructions for deploying Glance on Vercel as serverless functions, replacing the original Docker-based deployment.

## 🏗️ Architecture Changes

### Original Docker Deployment
- Long-running HTTP server with embedded static assets
- File-based configuration with auto-reload
- In-memory state and caching
- Background workers and scheduled tasks

### New Serverless Deployment
- Stateless serverless functions in `/api/` directory
- Static assets served from `/public/` directory  
- Environment variable-based configuration
- Request-triggered processing only

## 📁 Project Structure

```
glance/
├── api/
│   ├── index.go              # Main serverless handler
│   └── serverless/
│       └── app.go           # Serverless application logic
├── public/                   # Static assets (CSS, JS, images)
│   ├── css/
│   ├── js/
│   ├── fonts/
│   ├── icons/
│   └── favicon.svg
├── vercel.json              # Vercel configuration
├── go.mod
└── go.sum
```

## 🚀 Quick Start Deployment

### 1. Clone and Prepare Repository

```bash
git clone https://github.com/dr-data/glance.git
cd glance
```

### 2. Install Vercel CLI

```bash
npm install -g vercel
```

### 3. Login to Vercel

```bash
vercel login
```

### 4. Deploy to Vercel

```bash
vercel --prod
```

### 5. Configure Environment Variables

In your Vercel dashboard or using the CLI:

```bash
# Basic configuration
vercel env add GLANCE_PAGES production
# Paste: [{"name":"Dashboard","slug":"","columns":[]}]

vercel env add GLANCE_PROXIED production  
# Value: true

vercel env add GLANCE_BASE_URL production
# Value: https://your-app.vercel.app (or your custom domain)
```

## ⚙️ Configuration Options

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `GLANCE_PAGES` | JSON array of page configurations | Default dashboard | No |
| `GLANCE_HOST` | Host binding (unused in serverless) | "" | No |
| `GLANCE_PORT` | Port (unused in serverless) | 8080 | No |
| `GLANCE_PROXIED` | Enable proxy headers | true | No |
| `GLANCE_BASE_URL` | Base URL for assets | "" | No |
| `GLANCE_DISABLE_THEME_PICKER` | Disable theme selection | false | No |
| `GLANCE_AUTH_SECRET` | Base64 encoded auth secret | "" | No |
| `GLANCE_AUTH_USERS` | JSON object of users | {} | No |

### Example Page Configuration

```json
[
  {
    "name": "Dashboard",
    "slug": "",
    "columns": [
      {
        "size": "full",
        "widgets": [
          {
            "type": "weather",
            "data": {
              "location": "London"
            }
          }
        ]
      }
    ]
  }
]
```

### Example Authentication Configuration

```bash
# Generate secret key (run locally)
go run main.go secret

# Create password hash (run locally)  
go run main.go hash-password "yourpassword"

# Set in Vercel
vercel env add GLANCE_AUTH_SECRET production
# Paste the generated secret

vercel env add GLANCE_AUTH_USERS production
# Paste: {"admin":{"password_hash":"$2a$...your_hash"}}
```

## 🛠️ Advanced Configuration

### Custom Domain

1. Add domain in Vercel dashboard
2. Update DNS records as instructed
3. Set `GLANCE_BASE_URL` to your domain

```bash
vercel env add GLANCE_BASE_URL production
# Value: https://your-domain.com
```

### Custom Styling

1. Modify files in `/public/css/`
2. Redeploy: `vercel --prod`

### Custom Favicon

Replace `/public/favicon.svg` with your icon and redeploy.

## 🔌 API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/` | GET | Main dashboard page |
| `/api/healthz` | GET | Health check |
| `/api/pages/{page}/content` | GET | Page content API |
| `/api/authenticate` | POST | User authentication |
| `/login` | GET | Login page |
| `/logout` | GET | Logout |

## 🚨 Limitations

### Removed Features
- ❌ Config file watching and auto-reload
- ❌ In-memory caching between requests
- ❌ Background workers and scheduled tasks  
- ❌ File-based storage
- ❌ Long-running processes

### Serverless Constraints
- ⏱️ 10-second execution timeout per request
- 💾 No persistent in-memory state
- 🔄 Cold starts may introduce latency
- 📦 50MB deployment size limit

### Available Features
- ✅ Basic dashboard rendering
- ✅ Static asset serving
- ✅ Environment-based configuration
- ✅ Authentication (simplified)
- ✅ API endpoints
- ✅ Responsive design

## 🐛 Troubleshooting

### Build Errors

```bash
# Check Go version compatibility
go version

# Verify dependencies
go mod tidy
go build ./api/...
```

### Deployment Issues

```bash
# Check Vercel logs
vercel logs

# Test locally
vercel dev
```

### Environment Variables

```bash
# List all environment variables
vercel env ls

# Pull environment to local
vercel env pull
```

## 📊 Performance Optimization

### Cold Start Reduction
- Keep functions small and focused
- Minimize import dependencies
- Use environment variables for configuration

### Static Asset Optimization
- Enable Vercel's automatic image optimization
- Use compressed assets
- Implement proper caching headers

## 🔄 Migration from Docker

### 1. Export Current Configuration

If migrating from Docker deployment:

```bash
# Export your current glance.yml
cat config/glance.yml > config-backup.yml
```

### 2. Convert to Environment Variables

Transform your YAML configuration to JSON environment variables:

```yaml
# Original glance.yml
pages:
  - name: "Home"
    slug: ""
    columns:
      - size: "full"
        widgets:
          - type: "weather"
            location: "London"
```

```bash
# Convert to env var
export GLANCE_PAGES='[{"name":"Home","slug":"","columns":[{"size":"full","widgets":[{"type":"weather","data":{"location":"London"}}]}]}]'
```

### 3. Test Before Switching

Deploy to a staging environment first:

```bash
vercel --target staging
```

## 📝 Code Changes Summary

### New Files
- `api/index.go` - Main serverless handler
- `api/serverless/app.go` - Serverless application logic
- `public/` - Static assets directory
- `vercel.json` - Vercel configuration

### Modified Files
- `.gitignore` - Added Vercel artifacts
- `go.mod` - Dependencies remain the same

### Removed Features
- Docker-specific code (file watching, long-running server)
- Embedded asset serving (moved to `/public/`)
- Config file parsing (replaced with env vars)

## 🎯 Next Steps

1. **Deploy**: Follow the quick start guide
2. **Configure**: Set up environment variables
3. **Test**: Verify all functionality works
4. **Monitor**: Check Vercel analytics and logs
5. **Optimize**: Tune performance based on usage

## 📞 Support

For issues specific to the serverless deployment:
1. Check Vercel logs: `vercel logs`
2. Review environment variables: `vercel env ls`
3. Test locally: `vercel dev`

For general Glance functionality:
- Original repository: https://github.com/glanceapp/glance
- Documentation: https://github.com/glanceapp/glance/tree/main/docs

---

**Note**: This serverless deployment is a simplified version optimized for Vercel. Some advanced features from the original Docker deployment may not be available due to serverless constraints.