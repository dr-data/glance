# Glance Serverless on Vercel 🚀

A serverless adaptation of the [Glance](https://github.com/glanceapp/glance) dashboard for deployment on Vercel's free tier.

## ✨ Features

- 🌐 **Serverless Functions**: Runs as stateless Vercel functions
- 📱 **Responsive Design**: Works on desktop and mobile
- ⚡ **Fast CDN**: Static assets served via Vercel Edge Network
- 🔧 **Environment Config**: Configuration via environment variables
- 🔐 **Authentication**: Simple user authentication system
- 🎨 **Customizable**: Themes and styling support

## 🚀 Quick Deploy

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2Fdr-data%2Fglance&env=GLANCE_PAGES&envDescription=Dashboard%20configuration%20in%20JSON%20format&envLink=https%3A%2F%2Fgithub.com%2Fdr-data%2Fglance%2Fblob%2Fmain%2FCONFIG_EXAMPLES.md)

## 📖 Documentation

- **[Deployment Guide](VERCEL_DEPLOYMENT.md)** - Complete setup instructions
- **[Configuration Examples](CONFIG_EXAMPLES.md)** - Sample configurations
- **[Original Glance Docs](https://github.com/glanceapp/glance/tree/main/docs)** - Widget and feature documentation

## 🎯 Quick Start

1. **Deploy to Vercel:**
   ```bash
   npx vercel --prod
   ```

2. **Set Environment Variables:**
   ```bash
   vercel env add GLANCE_PAGES production
   # Paste: [{"name":"Dashboard","slug":"","columns":[]}]
   ```

3. **Visit your dashboard:**
   ```
   https://your-app.vercel.app
   ```

## 🔧 Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `GLANCE_PAGES` | Dashboard configuration | `[{"name":"Home","slug":"","columns":[]}]` |
| `GLANCE_AUTH_SECRET` | Authentication secret | Generated via CLI |
| `GLANCE_AUTH_USERS` | User accounts | `{"admin":{"password_hash":"..."}}` |

## 🏗️ Architecture

```
glance/
├── api/
│   ├── index.go              # Main serverless handler
│   └── serverless/app.go     # Application logic
├── public/                   # Static assets
│   ├── css/, js/, fonts/    # Styling and scripts
│   └── favicon.svg          # Site icon
├── vercel.json              # Vercel configuration
└── docs/                    # Documentation
```

## 🔄 Migration from Docker

If you're migrating from the original Docker version:

1. Export your current `glance.yml` configuration
2. Convert YAML to JSON environment variables
3. Follow the [migration guide](VERCEL_DEPLOYMENT.md#migration-from-docker)

## ⚡ Local Development

```bash
# Install Vercel CLI
npm install -g vercel

# Clone repository
git clone https://github.com/dr-data/glance.git
cd glance

# Set environment variables
vercel env pull

# Run development server
vercel dev
```

## 🌐 API Endpoints

- `/` - Main dashboard
- `/api/healthz` - Health check
- `/api/pages/{page}/content` - Page content
- `/login` - User login
- `/logout` - User logout

## ⚠️ Limitations

Due to serverless constraints:

- ❌ No config file watching
- ❌ No in-memory caching between requests
- ❌ No background workers
- ⏱️ 10-second execution timeout per request
- 📦 50MB deployment size limit

## 🤝 Contributing

This is a serverless adaptation of the original Glance project. For general issues:

- **Original Project**: [glanceapp/glance](https://github.com/glanceapp/glance)
- **Serverless Issues**: Open issues in this repository

## 📄 License

This project maintains the same license as the original Glance project.

## 🙏 Acknowledgments

- [Glance](https://github.com/glanceapp/glance) - Original dashboard project
- [Vercel](https://vercel.com) - Serverless hosting platform