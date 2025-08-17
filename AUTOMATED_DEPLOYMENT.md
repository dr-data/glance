# Automated Vercel Deployment Setup

This guide explains how to set up automated deployment to Vercel using GitHub Actions. Once configured, your Glance dashboard will automatically deploy to Vercel whenever you push changes to your repository.

## 🎯 Benefits of Automated Deployment

- ✅ **Automatic deployments** on every push to main branch
- ✅ **Preview deployments** for pull requests
- ✅ **Automated testing** before deployment
- ✅ **Deployment status** comments on PRs and commits
- ✅ **Rollback capability** through Vercel dashboard
- ✅ **Zero downtime** deployments

## 📋 Prerequisites

1. **GitHub repository**: Your own fork/copy of the Glance repository
2. **Vercel account**: Free account at [vercel.com](https://vercel.com)
3. **Project deployed**: At least one manual deployment (see [VERCEL_GUI_SETUP.md](VERCEL_GUI_SETUP.md))

## 🔧 Step-by-Step Setup

### Step 1: Get Vercel Project Information

1. **Log in to Vercel dashboard**: Go to [vercel.com/dashboard](https://vercel.com/dashboard)
2. **Select your project**: Click on your Glance deployment
3. **Go to Settings**: Click "Settings" tab
4. **Find Project ID**: In "General" section, copy the "Project ID"
5. **Find Org ID**: 
   - Go to your account settings (click your avatar → Account Settings)
   - In "General" section, copy the "Team ID" (this is your Org ID)

### Step 2: Generate Vercel Token

1. **Go to Tokens page**: Visit [vercel.com/account/tokens](https://vercel.com/account/tokens)
2. **Create new token**: Click "Create Token"
3. **Configure token**:
   - **Token Name**: `GitHub Actions - Glance`
   - **Scope**: Full access (or limit to your specific team)
   - **Expiration**: 30 days or No expiration (your choice)
4. **Save the token**: Copy it immediately (you won't see it again)

### Step 3: Add GitHub Secrets

1. **Go to your GitHub repository**
2. **Navigate to Settings**: Click "Settings" tab (not the gear icon)
3. **Go to Secrets**: Click "Secrets and variables" → "Actions"
4. **Add Repository Secrets**: Click "New repository secret" for each of these:

#### Required Secrets:

| Secret Name | Value | Where to Find |
|-------------|-------|---------------|
| `VERCEL_TOKEN` | Your Vercel API token | From Step 2 above |
| `VERCEL_ORG_ID` | Your Vercel team/org ID | From Step 1 above |
| `VERCEL_PROJECT_ID` | Your Vercel project ID | From Step 1 above |

### Step 4: Create package.json (if needed)

The workflow expects a `package.json` file for Node.js caching. Create one in your repository root:

```json
{
  "name": "glance-serverless",
  "version": "1.0.0",
  "description": "Glance Dashboard - Serverless deployment",
  "private": true,
  "scripts": {
    "deploy": "vercel --prod",
    "dev": "vercel dev"
  },
  "devDependencies": {
    "vercel": "latest"
  }
}
```

### Step 5: Test the Workflow

1. **Make a small change**: Edit any file (like adding a comment to README.md)
2. **Commit and push**: 
   ```bash
   git add .
   git commit -m "Test automated deployment"
   git push origin main
   ```
3. **Check GitHub Actions**: Go to "Actions" tab in your repository
4. **Monitor deployment**: Watch the "Deploy to Vercel" workflow run
5. **Verify deployment**: Check the production URL in the workflow logs

## 🔄 How It Works

### Automatic Deployments

**On Push to Main Branch:**
- ✅ Runs tests (`go test ./...`)
- ✅ Builds application (`go build`)
- ✅ Deploys to Vercel production
- ✅ Posts deployment URL as comment on commit

**On Pull Request:**
- ✅ Runs tests
- ✅ Creates preview deployment
- ✅ Posts preview URL as comment on PR
- ✅ Updates comment when PR is updated

### Workflow Triggers

The automation runs on:
- **Push to main**: Production deployment
- **Pull requests**: Preview deployment
- **Manual trigger**: Via GitHub Actions UI
- **Workflow dispatch**: Programmatic trigger

## 🛠️ Customizing the Workflow

### Change Deployment Branch

To deploy from a different branch, edit `.github/workflows/vercel-deploy.yml`:

```yaml
on:
  push:
    branches: [ main, develop ]  # Add your branches here
```

### Add Environment-Specific Secrets

For different environments, add prefixed secrets:

```
VERCEL_TOKEN_STAGING
VERCEL_PROJECT_ID_STAGING
VERCEL_ORG_ID_STAGING
```

### Custom Build Steps

Add steps before deployment:

```yaml
- name: Run linting
  run: go vet ./...

- name: Run security checks
  run: gosec ./...
```

## 🚨 Troubleshooting

### Common Issues

#### ❌ "VERCEL_TOKEN not found"
- **Solution**: Ensure you've added all three required secrets
- **Check**: Secret names are exactly: `VERCEL_TOKEN`, `VERCEL_ORG_ID`, `VERCEL_PROJECT_ID`

#### ❌ "Project not found"
- **Solution**: Verify `VERCEL_PROJECT_ID` is correct
- **Check**: Project ID in Vercel dashboard settings

#### ❌ "Insufficient permissions"
- **Solution**: Regenerate Vercel token with full access
- **Check**: Token hasn't expired

#### ❌ Build fails on Node.js steps
- **Solution**: Add `package.json` file (see Step 4)
- **Alternative**: Remove Node.js cache steps if not needed

### Viewing Deployment Logs

1. **GitHub Actions logs**: Actions tab → Select workflow run
2. **Vercel deployment logs**: Vercel dashboard → Functions tab
3. **Runtime logs**: Vercel dashboard → View function logs

### Manual Deployment

If automation fails, you can always deploy manually:

```bash
# Install Vercel CLI
npm install -g vercel

# Deploy from command line
vercel --prod
```

## 🔒 Security Best Practices

### Token Security
- ✅ **Use repository secrets**: Never commit tokens to code
- ✅ **Limit token scope**: Only give necessary permissions
- ✅ **Rotate regularly**: Update tokens every 30-90 days
- ✅ **Monitor usage**: Check Vercel usage logs

### Branch Protection
- ✅ **Require PR reviews**: Protect main branch
- ✅ **Require status checks**: Ensure tests pass before merge
- ✅ **Restrict direct pushes**: Only allow through PRs

### Environment Variables
- ✅ **Use Vercel secrets**: For sensitive configuration
- ✅ **Separate environments**: Different configs for prod/staging
- ✅ **Validate inputs**: Check environment variables are set

## 📊 Monitoring Deployments

### Success Indicators
- ✅ **Green checkmark**: In GitHub Actions
- ✅ **Comment posted**: On commit/PR with URL
- ✅ **Site accessible**: Production URL responds
- ✅ **Health check passes**: `/api/healthz` returns 200

### Failure Recovery
1. **Check logs**: GitHub Actions and Vercel logs
2. **Fix issues**: Address any errors found
3. **Retry deployment**: Push a new commit or re-run workflow
4. **Rollback if needed**: Use Vercel dashboard to rollback

## 🎯 Next Steps

1. **Set up monitoring**: Configure Vercel analytics
2. **Add custom domain**: Set up your own domain
3. **Configure alerts**: Get notified of deployment failures
4. **Set up staging**: Create separate environment for testing
5. **Optimize performance**: Monitor and improve build times

## 📚 Additional Resources

- **[GitHub Actions Documentation](https://docs.github.com/en/actions)**
- **[Vercel CLI Documentation](https://vercel.com/docs/cli)**
- **[Vercel GitHub Integration](https://vercel.com/docs/concepts/git)**
- **[Managing Secrets in GitHub](https://docs.github.com/en/actions/security-guides/encrypted-secrets)**

---

*Once automated deployment is set up, you'll have a professional CI/CD pipeline that automatically tests and deploys your Glance dashboard whenever you make changes.*