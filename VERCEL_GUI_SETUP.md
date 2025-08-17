# Vercel GUI Setup Guide for Glance Dashboard

This guide provides step-by-step instructions for deploying Glance to Vercel using the web interface, specifically addressing the configuration options shown in the Vercel "New Project" screen.

## 🎯 Quick Deploy (One-Click)

The easiest way to deploy is using the one-click deploy button:

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2Fdr-data%2Fglance&env=GLANCE_PAGES&envDescription=Dashboard%20configuration%20in%20JSON%20format&envLink=https%3A%2F%2Fgithub.com%2Fdr-data%2Fglance%2Fblob%2Fmain%2FCONFIG_EXAMPLES.md)

This will automatically configure most settings for you. If you prefer manual setup or want to understand each step, continue with the manual setup below.

## 📋 Manual Setup Guide

### Step 1: Import Your Repository

1. **Navigate to Vercel**: Go to [vercel.com](https://vercel.com) and sign in
2. **Create New Project**: Click "New Project" or visit [vercel.com/new](https://vercel.com/new)
3. **Import from GitHub**: Select "Import Git Repository" and choose `dr-data/glance`

### Step 2: Configure Project Settings

Based on your screenshot, here's what to enter in each field:

#### **Project Name**
- **What to enter**: `glance` (or any name you prefer)
- **Purpose**: This will be part of your URL (e.g., `glance-abc123.vercel.app`)
- **Recommendation**: Use lowercase, no spaces, hyphens are allowed

#### **Vercel Team**
- **What to select**: Choose your personal account or team
- **Purpose**: Determines billing and access permissions
- **For beginners**: Use your personal account unless you're part of an organization

#### **Framework Preset**
- **What to select**: `Other` (Keep the default selection)
- **Why**: Glance is a Go application with custom serverless configuration
- **Important**: Do NOT select Next.js, React, or other JavaScript frameworks

#### **Root Directory**
- **What to enter**: `./` (Keep default - this means the project root)
- **Purpose**: Tells Vercel where your project files are located
- **Note**: Since all project files are in the repository root, leave this as default

### Step 3: Build and Output Settings

Click the "Build and Output Settings" dropdown and configure:

#### **Build Command**
- **Leave empty** or use: `echo "Using serverless functions"`
- **Why**: The Go serverless functions are built automatically by Vercel
- **Important**: Do NOT use `go build` here as it's not needed for serverless deployment

#### **Output Directory**
- **Leave empty**
- **Why**: Static assets are served from the `/public` directory automatically

#### **Install Command**
- **Leave empty**
- **Why**: Go dependencies are managed by `go.mod` and installed automatically

### Step 4: Environment Variables

Click the "Environment Variables" dropdown and add these **required** variables:

#### **GLANCE_PAGES** (Required)
- **Variable Name**: `GLANCE_PAGES`
- **Value**: `[{"name":"Dashboard","slug":"","columns":[{"size":"full","widgets":[{"type":"weather","data":{"location":"London"}}]}]}]`
- **Purpose**: Defines your dashboard layout and widgets
- **Note**: This is JSON format - be careful with quotes and brackets

#### **GLANCE_PROXIED** (Recommended)
- **Variable Name**: `GLANCE_PROXIED`
- **Value**: `true`
- **Purpose**: Enables proper handling of reverse proxy headers
- **Why needed**: Vercel acts as a reverse proxy

#### **GLANCE_BASE_URL** (Optional but recommended)
- **Variable Name**: `GLANCE_BASE_URL`
- **Value**: Leave empty initially (will be auto-set after deployment)
- **Purpose**: Used for generating absolute URLs
- **Note**: You can update this after deployment with your actual URL

### Step 5: Deploy

1. **Review Settings**: Double-check all your configurations
2. **Click Deploy**: Click the black "Deploy" button at the bottom
3. **Wait for Build**: The build process will take 1-3 minutes
4. **Get Your URL**: Once complete, you'll receive a URL like `https://glance-abc123.vercel.app`

## 🔧 Post-Deployment Configuration

### Update Base URL (Recommended)

After your first deployment:

1. Go to your Vercel dashboard
2. Click on your project
3. Go to "Settings" → "Environment Variables"
4. Add or update `GLANCE_BASE_URL` with your actual deployment URL
5. Redeploy the project

### Add Authentication (Optional)

To secure your dashboard:

1. **Generate a secret key locally**:
   ```bash
   # Run this on your local machine with Go installed
   go run main.go secret
   ```

2. **Create a password hash**:
   ```bash
   # Run this on your local machine
   go run main.go hash-password "your-secure-password"
   ```

3. **Add environment variables in Vercel**:
   - `GLANCE_AUTH_SECRET`: The secret from step 1
   - `GLANCE_AUTH_USERS`: `{"admin":{"password_hash":"$2a$...your-hash-from-step-2"}}`

## 🎨 Customization Options

### Custom Domain

1. In Vercel dashboard, go to "Settings" → "Domains"
2. Add your custom domain
3. Update DNS records as instructed by Vercel
4. Update `GLANCE_BASE_URL` environment variable to your custom domain

### Theme Customization

1. Fork the repository to your GitHub account
2. Modify files in `/public/css/` directory
3. Commit and push changes
4. Vercel will automatically redeploy

## 🚨 Common Mistakes to Avoid

### ❌ Wrong Framework Selection
- **Don't select**: Next.js, React, Vue, or other JavaScript frameworks
- **Select**: "Other" framework preset

### ❌ Incorrect Environment Variables
- **Don't forget**: JSON must be properly formatted for `GLANCE_PAGES`
- **Use double quotes**: Not single quotes in JSON
- **Validate JSON**: Use a JSON validator before pasting

### ❌ Build Command Errors
- **Don't add**: Custom build commands unless necessary
- **Let Vercel handle**: Go compilation automatically

### ❌ Root Directory Issues
- **Keep default**: `./` unless your project is in a subfolder
- **Don't change**: Unless you know what you're doing

## 🔧 Troubleshooting

### Build Fails with "Go not found"
- **Solution**: This shouldn't happen with proper serverless setup
- **Check**: Ensure `api/index.go` exists and `vercel.json` is configured correctly

### Environment Variables Not Working
- **Check**: JSON formatting in `GLANCE_PAGES`
- **Validate**: Use an online JSON validator
- **Redeploy**: After changing environment variables

### Dashboard Shows Error 500
- **Check**: Vercel logs (Functions tab in dashboard)
- **Common cause**: Malformed `GLANCE_PAGES` JSON
- **Solution**: Fix JSON formatting and redeploy

### Static Assets Not Loading
- **Check**: Files exist in `/public` directory
- **Verify**: `vercel.json` routing configuration
- **Solution**: Usually resolves automatically

## 📞 Getting Help

### Check Deployment Logs
1. Go to your Vercel dashboard
2. Click on your project
3. Go to "Functions" tab
4. Click on any failed function to see logs

### Test Locally
```bash
# Install Vercel CLI
npm install -g vercel

# Clone your repository
git clone https://github.com/your-username/glance.git
cd glance

# Test locally
vercel dev
```

### Common URLs to Test
After deployment, test these endpoints:
- `https://your-app.vercel.app` - Main dashboard
- `https://your-app.vercel.app/api/healthz` - Health check
- `https://your-app.vercel.app/login` - Login page (if auth enabled)

## 🎯 Next Steps

1. **Customize Your Dashboard**: Edit the `GLANCE_PAGES` environment variable
2. **Add Widgets**: See [CONFIG_EXAMPLES.md](CONFIG_EXAMPLES.md) for widget examples
3. **Set Up Authentication**: Follow the authentication setup above
4. **Configure Custom Domain**: Add your own domain for a professional look
5. **Monitor Usage**: Check Vercel analytics for performance metrics

## 📚 Additional Resources

- **[Complete Deployment Guide](VERCEL_DEPLOYMENT.md)** - Technical details and advanced configuration
- **[Configuration Examples](CONFIG_EXAMPLES.md)** - Sample widget configurations
- **[Original Glance Documentation](https://github.com/glanceapp/glance/tree/main/docs)** - Comprehensive widget documentation
- **[Vercel Documentation](https://vercel.com/docs)** - Platform-specific help

---

*This guide is designed for users with moderate technical experience. If you encounter issues, please refer to the troubleshooting section or seek help in the repository issues.*