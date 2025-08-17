#!/bin/bash

# Glance Serverless Local Test Script
# This script helps test the serverless deployment locally

set -e

echo "🚀 Glance Serverless Local Test"
echo "==============================="

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.19+ to continue."
    exit 1
fi

echo "✅ Go version: $(go version)"

# Check if Vercel CLI is installed
if ! command -v vercel &> /dev/null; then
    echo "⚠️  Vercel CLI not found. Installing globally..."
    npm install -g vercel
fi

echo "✅ Vercel CLI: $(vercel --version)"

# Set up environment variables for testing
echo "📝 Setting up test environment variables..."

export GLANCE_PAGES='[{"name":"Test Dashboard","slug":"","columns":[{"size":"full","widgets":[{"type":"weather","data":{"location":"London"}}]}]}]'
export GLANCE_PROXIED="true"
export GLANCE_BASE_URL="http://localhost:3000"
export GLANCE_DISABLE_THEME_PICKER="false"

echo "✅ Environment variables configured"

# Test Go build
echo "🔨 Testing Go build..."
if go build ./api/...; then
    echo "✅ Go build successful"
else
    echo "❌ Go build failed"
    exit 1
fi

# Clean up binary
rm -f api/index

# Test local development
echo "🌐 Starting local development server..."
echo ""
echo "🎯 Test URLs:"
echo "   • Main Dashboard: http://localhost:3000"
echo "   • Health Check:   http://localhost:3000/api/healthz"
echo "   • Login Page:     http://localhost:3000/login"
echo ""
echo "💡 Press Ctrl+C to stop the server"
echo ""

# Start Vercel dev server
vercel dev