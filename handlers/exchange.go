package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"git.gvk.idi.ntnu.no/course/prog2005/prog2005-2026-workspace/abdulhas_1/assignment-1/models"
	"git.gvk.idi.ntnu.no/course/prog2005/prog2005-2026-workspace/abdulhas_1/assignment-1/services"
)

// ExchangeHandler handles GET /countryinfo/v1/exchange/{code}
func ExchangeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract country code from URL path
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

	// Fetch base country data
	country, err := services.GetCountryByCode(countryCode)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			sendError(w, "Country not found", http.StatusNotFound)
		} else {
			sendError(w, "Failed to fetch country data", http.StatusInternalServerError)
		}
		return
	}

	// Extract base currency (use first one if multiple)
	var baseCurrency string
	for code := range country.Currencies {
		baseCurrency = code
		break
	}

	if baseCurrency == "" {
		sendError(w, "Country has no currency information", http.StatusNotFound)
		return
	}

	// Handle countries with no borders
	if len(country.Borders) == 0 {
		response := models.ExchangeResponse{
			Country:       country.Name.Common,
			BaseCurrency:  baseCurrency,
			ExchangeRates: []map[string]float64{},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Collect currencies from bordering countries
	borderCurrencies := make(map[string]bool) // Use map to avoid duplicates
	for _, borderCode := range country.Borders {
		borderCountry, err := services.GetCountryByCCA3(borderCode)
		if err != nil {
			// Skip countries that can't be fetched
			continue
		}

		// Add all currencies from this border country
		for currencyCode := range borderCountry.Currencies {
			if currencyCode != baseCurrency {
				borderCurrencies[currencyCode] = true
			}
		}
	}

	// Convert map keys to slice
	var targetCurrencies []string
	for currency := range borderCurrencies {
		targetCurrencies = append(targetCurrencies, currency)
	}

	// Fetch exchange rates
	var exchangeRates []map[string]float64
	if len(targetCurrencies) > 0 {
		rates, err := services.GetExchangeRate(baseCurrency, targetCurrencies)
		if err != nil {
			sendError(w, "Failed to fetch exchange rates", http.StatusInternalServerError)
			return
		}

		// Convert to required format: array of single-key maps
		for currency, rate := range rates {
			exchangeRates = append(exchangeRates, map[string]float64{currency: rate})
		}
	}

	if exchangeRates == nil {
		exchangeRates = []map[string]float64{}
	}

	response := models.ExchangeResponse{
		Country:       country.Name.Common,
		BaseCurrency:  baseCurrency,
		ExchangeRates: exchangeRates,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
