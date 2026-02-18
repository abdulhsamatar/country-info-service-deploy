package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"git.gvk.idi.ntnu.no/course/prog2005/prog2005-2026-workspace/abdulhas_1/assignment-1/models"
	"git.gvk.idi.ntnu.no/course/prog2005/prog2005-2026-workspace/abdulhas_1/assignment-1/services"
)

// StatusHandler handles GET /countryinfo/v1/status/
func StatusHandler(startTime time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Check both external APIs
		restCountriesStatus := services.CheckRESTCountriesStatus()
		currencyStatus := services.CheckCurrencyStatus()

		// Calculate uptime in seconds
		uptime := int(time.Since(startTime).Seconds())

		response := models.StatusResponse{
			RESTCountriesAPI: restCountriesStatus,
			CurrenciesAPI:    currencyStatus,
			Version:          "v1",
			Uptime:           uptime,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
