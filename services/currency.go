package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"git.gvk.idi.ntnu.no/course/prog2005/prog2005-2026-workspace/abdulhas_1/assignment-1/models"
)

const CurrencyBaseURL = "http://129.241.150.113:9090/currency"

// GetExchangeRate fetches exchange rate from base currency to target currencies
// baseCurrency: e.g., "NOK"
// targetCurrencies: e.g., ["EUR", "SEK"]
func GetExchangeRate(baseCurrency string, targetCurrencies []string) (map[string]float64, error) {
	if len(targetCurrencies) == 0 {
		return make(map[string]float64), nil
	}

	// Build URL: /currency/{base}?symbols={target1},{target2}
	targets := strings.Join(targetCurrencies, ",")
	url := fmt.Sprintf("%s/%s?symbols=%s", CurrencyBaseURL, strings.ToUpper(baseCurrency), targets)

	client := &http.Client{Timeout: RequestTimeout}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch exchange rate: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Currency API returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var currencyResp models.CurrencyResponse
	if err := json.Unmarshal(body, &currencyResp); err != nil {
		return nil, fmt.Errorf("failed to parse currency data: %w", err)
	}

	return currencyResp.Rates, nil
}

// CheckCurrencyStatus checks if Currency API is reachable
func CheckCurrencyStatus() int {
	// Use a simple query that should always work
	client := &http.Client{Timeout: RequestTimeout}
	resp, err := client.Get(CurrencyBaseURL + "/EUR")
	if err != nil {
		return http.StatusServiceUnavailable
	}
	defer resp.Body.Close()

	return resp.StatusCode
}
