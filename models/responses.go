package models

// StatusResponse for /countryinfo/v1/status/
type StatusResponse struct {
	RESTCountriesAPI int    `json:"restcountriesapi"`
	CurrenciesAPI    int    `json:"currenciesapi"`
	Version          string `json:"version"`
	Uptime           int    `json:"uptime"`
}

// CountryInfoResponse for /countryinfo/v1/info/{code}
type CountryInfoResponse struct {
	Name       string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Area       float64           `json:"area"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
}

// ExchangeResponse for /countryinfo/v1/exchange/{code}
type ExchangeResponse struct {
	Country       string               `json:"country"`
	BaseCurrency  string               `json:"base-currency"`
	ExchangeRates []map[string]float64 `json:"exchange-rates"`
}
