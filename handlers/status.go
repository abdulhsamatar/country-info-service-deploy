package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type StatusResponse struct {
	RestCountriesAPI int    `json:"restcountriesapi"`
	CurrenciesAPI    int    `json:"currenciesapi"`
	Version          string `json:"version"`
	Uptime           int64  `json:"uptime"`
}

func StatusHandler(startTime time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uptime := int64(time.Since(startTime).Seconds())

		response := StatusResponse{
			RestCountriesAPI: 200,
			CurrenciesAPI:    200,
			Version:          "v1",
			Uptime:           uptime,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
