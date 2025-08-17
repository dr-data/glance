# Vercel Deployment Best Practices & Troubleshooting

This guide covers best practices, common pitfalls, and troubleshooting tips for deploying Glance to Vercel, specifically designed for users with moderate technical experience.

## 🏆 Best Practices

### 🔧 Environment Variables

#### Naming Conventions
- ✅ **Use GLANCE_ prefix**: Keeps variables organized
- ✅ **ALL_CAPS**: Follow standard environment variable naming
- ✅ **Descriptive names**: `GLANCE_AUTH_SECRET` vs `SECRET`

#### Security
- ✅ **Never commit secrets**: Use Vercel environment variables only
- ✅ **Rotate regularly**: Update auth secrets periodically
- ✅ **Limit access**: Use minimal required permissions

#### Organization
```bash
# Group related variables
GLANCE_PAGES=...                    # Dashboard config
GLANCE_AUTH_SECRET=...              # Authentication
GLANCE_AUTH_USERS=...               # User accounts
GLANCE_BASE_URL=...                 # Site configuration
```

### 🚀 Deployment Strategy

#### Use Staging Environment
1. **Create staging project**: Separate Vercel project for testing
2. **Test changes first**: Deploy to staging before production
3. **Validate functionality**: Ensure all widgets work correctly

#### Version Control
- ✅ **Tag releases**: Use semantic versioning (v1.0.0)
- ✅ **Document changes**: Update README for major changes
- ✅ **Backup config**: Save environment variables externally

### 📊 Performance Optimization

#### Environment Variables
- ✅ **Use JSON minification**: Remove unnecessary whitespace
- ✅ **Validate JSON format**: Use online validators
- ✅ **Cache-friendly config**: Avoid frequent changes

#### Static Assets
- ✅ **Optimize images**: Compress icons and backgrounds
- ✅ **Use CDN**: Leverage Vercel's edge network
- ✅ **Set cache headers**: Let Vercel handle caching

### 🔍 Monitoring

#### Health Checks
```bash
# Regular monitoring endpoints
curl https://your-app.vercel.app/api/healthz
curl -I https://your-app.vercel.app
```

#### Vercel Analytics
1. **Enable analytics**: In Vercel dashboard
2. **Monitor performance**: Check page load times
3. **Track errors**: Review function logs regularly

## ⚠️ Common Pitfalls & Solutions

### 🚨 Configuration Errors

#### Invalid JSON in GLANCE_PAGES
**❌ Common Mistake:**
```json
GLANCE_PAGES=[{'name':'Dashboard','slug':'','columns':[]}]
```

**✅ Correct Format:**
```json
GLANCE_PAGES=[{"name":"Dashboard","slug":"","columns":[]}]
```

**📝 Key Points:**
- Use **double quotes** (not single quotes)
- Validate JSON with [jsonlint.com](https://jsonlint.com)
- Escape special characters properly

#### Malformed Widget Configuration
**❌ Common Mistake:**
```json
{"type":"weather","location":"London"}
```

**✅ Correct Format:**
```json
{"type":"weather","data":{"location":"London"}}
```

**📝 Key Points:**
- Widget data goes in `data` object
- Check widget documentation for required fields

### 🔧 Build & Deployment Issues

#### Framework Preset Selection
**❌ Wrong Choice:**
- Next.js, React, Vue, Angular

**✅ Correct Choice:**
- "Other" framework preset

**📝 Why:**
- Glance uses Go serverless functions
- JavaScript frameworks interfere with Go compilation

#### Root Directory Configuration
**❌ Common Mistake:**
- Setting root directory to `/api` or `/public`

**✅ Correct Setting:**
- Keep default `./` (project root)

**📝 Why:**
- Vercel needs access to all project files
- `vercel.json` must be in root directory

#### Build Command Confusion
**❌ Don't Add:**
```bash
# These commands are not needed
go build .
make build
npm run build
```

**✅ Leave Empty:**
- Build command should be empty
- Vercel handles Go compilation automatically

### 🔐 Authentication Problems

#### Secret Key Generation
**❌ Using Weak Secrets:**
```bash
# Don't use simple strings
GLANCE_AUTH_SECRET=mysecret123
```

**✅ Generate Properly:**
```bash
# Use the built-in generator
go run main.go secret
```

#### Password Hash Issues
**❌ Plain Text Passwords:**
```json
{"admin":{"password":"plaintext"}}
```

**✅ Hashed Passwords:**
```json
{"admin":{"password_hash":"$2a$10$..."}}
```

**📝 Generate Hash:**
```bash
go run main.go hash-password "your-secure-password"
```

### 🌐 Domain & URL Issues

#### Base URL Configuration
**❌ After Custom Domain Setup:**
```bash
# Don't keep the old URL
GLANCE_BASE_URL=https://project-abc123.vercel.app
```

**✅ Update to Custom Domain:**
```bash
GLANCE_BASE_URL=https://dashboard.yourdomain.com
```

#### SSL Certificate Problems
**📝 Wait Period:**
- SSL provisioning takes 5-15 minutes
- Don't panic if HTTPS doesn't work immediately
- Check Vercel domain settings for status

## 🔧 Troubleshooting Guide

### 🔍 Diagnosis Steps

#### Step 1: Check Deployment Status
1. **Vercel Dashboard**: Look for deployment errors
2. **Function Logs**: Check for runtime errors
3. **Build Logs**: Review compilation output

#### Step 2: Validate Configuration
```bash
# Test JSON validity
echo '$GLANCE_PAGES' | python -m json.tool

# Check environment variables
vercel env ls
```

#### Step 3: Test Endpoints
```bash
# Health check (should return 200)
curl -i https://your-app.vercel.app/api/healthz

# Main page (should return HTML)
curl -s https://your-app.vercel.app | grep "<title>"
```

### 🚨 Specific Error Solutions

#### Error: "Function exceeded timeout"
**Cause**: Function taking too long to execute
**Solutions:**
1. **Simplify config**: Reduce number of widgets
2. **Check external APIs**: Verify third-party services are responsive
3. **Optimize queries**: Review widget data sources

#### Error: "JSON parse error"
**Cause**: Invalid JSON in environment variables
**Solutions:**
1. **Validate JSON**: Use online JSON validator
2. **Check quotes**: Ensure double quotes throughout
3. **Escape characters**: Properly escape backslashes and quotes

#### Error: "404 Not Found"
**Cause**: Routing configuration issue
**Solutions:**
1. **Check vercel.json**: Ensure proper route configuration
2. **Verify file structure**: Confirm `api/index.go` exists
3. **Redeploy**: Sometimes fixes routing issues

#### Error: "500 Internal Server Error"
**Cause**: Runtime error in Go code
**Solutions:**
1. **Check logs**: Review Vercel function logs
2. **Test locally**: Run `vercel dev` for debugging
3. **Simplify config**: Start with minimal configuration

### 🧪 Local Testing

#### Setup Local Environment
```bash
# Install Vercel CLI
npm install -g vercel

# Clone and navigate
git clone https://github.com/your-username/glance.git
cd glance

# Set up environment
vercel env pull

# Start development server
vercel dev
```

#### Test Locally
```bash
# Test endpoints
curl http://localhost:3000/api/healthz
curl http://localhost:3000

# Check function logs
# Logs appear in terminal running `vercel dev`
```

### 📞 Getting Help

#### Self-Service Debug
1. **Enable debug logs**: Add `DEBUG=1` environment variable
2. **Check Vercel status**: Visit [vercel-status.com](https://vercel-status.com)
3. **Review documentation**: Check latest Vercel docs

#### Community Support
- **GitHub Issues**: Open issue in repository
- **Vercel Community**: [community.vercel.com](https://community.vercel.com)
- **Discord**: Join Vercel Discord server

#### Escalation Path
1. **Document the problem**: Error messages, configuration, steps taken
2. **Gather logs**: Both GitHub Actions and Vercel logs
3. **Create minimal reproduction**: Simplest config that shows the issue
4. **Open support ticket**: With Vercel or repository

## 📋 Pre-Deployment Checklist

### ✅ Before First Deployment
- [ ] Repository is public or accessible to Vercel
- [ ] `vercel.json` configuration is correct
- [ ] Required environment variables are defined
- [ ] JSON configuration is validated
- [ ] Framework preset is set to "Other"
- [ ] Root directory is set to `./`

### ✅ Before Production Deployment
- [ ] Tested in preview/staging environment
- [ ] All widgets load correctly
- [ ] Authentication works (if enabled)
- [ ] Custom domain is configured (if applicable)
- [ ] SSL certificate is active
- [ ] Health check endpoint responds

### ✅ After Deployment
- [ ] Main dashboard loads successfully
- [ ] All configured widgets display data
- [ ] Login/logout functions work
- [ ] Mobile responsiveness is acceptable
- [ ] Performance is satisfactory
- [ ] Analytics/monitoring is enabled

## 🔄 Maintenance Tasks

### Regular (Weekly)
- [ ] Check deployment health
- [ ] Review error logs
- [ ] Monitor performance metrics

### Periodic (Monthly)
- [ ] Update dependencies
- [ ] Rotate authentication secrets
- [ ] Review and optimize configuration
- [ ] Update documentation

### As Needed
- [ ] Scale resources if needed
- [ ] Update custom domain configuration
- [ ] Migrate to new Vercel features
- [ ] Backup configuration changes

---

*Following these best practices will help ensure a smooth, reliable deployment of your Glance dashboard on Vercel.*