package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"git.gvk.idi.ntnu.no/course/prog2005/prog2005-2026-workspace/abdulhas_1/assignment-1/models"
)

const (
	RestCountriesBaseURL = "http://129.241.150.113:8080/v3.1"
	RequestTimeout       = 10 * time.Second
)

// GetCountryByCode fetches country data by 2-letter ISO code
func GetCountryByCode(code string) (*models.RestCountry, error) {
	// Normalize to lowercase
	code = strings.ToLower(code)

	// Build URL: /alpha/{code}
	url := fmt.Sprintf("%s/alpha/%s", RestCountriesBaseURL, code)

	client := &http.Client{Timeout: RequestTimeout}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch country data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("REST Countries API returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var countries models.RestCountriesResponse
	if err := json.Unmarshal(body, &countries); err != nil {
		return nil, fmt.Errorf("failed to parse country data: %w", err)
	}

	if len(countries) == 0 {
		return nil, fmt.Errorf("country not found")
	}

	return &countries[0], nil
}

// GetCountryByCCA3 fetches country data by 3-letter ISO code
func GetCountryByCCA3(code string) (*models.RestCountry, error) {
	code = strings.ToUpper(code)
	url := fmt.Sprintf("%s/alpha/%s", RestCountriesBaseURL, code)

	client := &http.Client{Timeout: RequestTimeout}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch country data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("REST Countries API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var countries models.RestCountriesResponse
	if err := json.Unmarshal(body, &countries); err != nil {
		return nil, fmt.Errorf("failed to parse country data: %w", err)
	}

	if len(countries) == 0 {
		return nil, fmt.Errorf("country not found")
	}

	return &countries[0], nil
}

// CheckRESTCountriesStatus checks if REST Countries API is reachable
func CheckRESTCountriesStatus() int {
	client := &http.Client{Timeout: RequestTimeout}
	resp, err := client.Get(RestCountriesBaseURL + "/all")
	if err != nil {
		return http.StatusServiceUnavailable
	}
	defer resp.Body.Close()

	return resp.StatusCode
}
