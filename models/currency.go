package models

// CurrencyResponse represents the response from Currency API
// Endpoint: /currency/{base}?symbols={target1},{target2}
type CurrencyResponse struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}
