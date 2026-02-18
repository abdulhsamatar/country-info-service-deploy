package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"git.gvk.idi.ntnu.no/course/prog2005/prog2005-2026-workspace/abdulhas_1/assignment-1/models"
	"git.gvk.idi.ntnu.no/course/prog2005/prog2005-2026-workspace/abdulhas_1/assignment-1/services"
)

// InfoHandler handles GET /countryinfo/v1/info/{code}
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract country code from URL path
	// Path format: /countryinfo/v1/info/{code}
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) != 4 {
		sendError(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	countryCode := pathParts[3]
	if len(countryCode) != 2 {
		sendError(w, "Country code must be 2 letters", http.StatusBadRequest)
		return
	}

	// Fetch country data from REST Countries API
	country, err := services.GetCountryByCode(countryCode)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			sendError(w, "Country not found", http.StatusNotFound)
		} else {
			sendError(w, "Failed to fetch country data", http.StatusInternalServerError)
		}
		return
	}

	// Extract capital - take first one if multiple exist
	capital := ""
	if len(country.Capital) > 0 {
		capital = country.Capital[0]
	}

	// Build response
	response := models.CountryInfoResponse{
		Name:       country.Name.Common,
		Continents: country.Continents,
		Population: country.Population,
		Area:       country.Area,
		Languages:  country.Languages,
		Borders:    country.Borders,
		Flag:       country.Flags.PNG,
		Capital:    capital,
	}

	// Handle nil slices
	if response.Continents == nil {
		response.Continents = []string{}
	}
	if response.Borders == nil {
		response.Borders = []string{}
	}
	if response.Languages == nil {
		response.Languages = make(map[string]string)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
