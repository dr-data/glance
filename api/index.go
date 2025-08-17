package handler

import (
	"net/http"

	"github.com/glanceapp/glance/api/serverless"
)

// Handler is the main entry point for Vercel serverless function
func Handler(w http.ResponseWriter, r *http.Request) {
	// Initialize serverless application
	app, err := serverless.NewServerlessApp()
	if err != nil {
		http.Error(w, "Failed to initialize application", http.StatusInternalServerError)
		return
	}

	// Delegate to the serverless app handler
	app.Handler(w, r)
}