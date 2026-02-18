package handlers

import (
	"encoding/json"
	"net/http"

	"git.gvk.idi.ntnu.no/course/prog2005/prog2005-2026-workspace/abdulhas_1/assignment-1/models"
)

// sendError sends a JSON error response
func sendError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResp := models.ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	}

	json.NewEncoder(w).Encode(errorResp)
}
