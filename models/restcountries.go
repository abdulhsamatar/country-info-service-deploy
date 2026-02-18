package models

// RestCountriesResponse represents the response from REST Countries API
// The API returns an array of countries
type RestCountriesResponse []RestCountry

type RestCountry struct {
	Name       Name                `json:"name"`
	Capital    []string            `json:"capital"`
	Population int                 `json:"population"`
	Area       float64             `json:"area"`
	Continents []string            `json:"continents"`
	Languages  map[string]string   `json:"languages"`
	Borders    []string            `json:"borders"`
	Flags      Flags               `json:"flags"`
	Currencies map[string]Currency `json:"currencies"`
	CCA2       string              `json:"cca2"` // 2-letter code
	CCA3       string              `json:"cca3"` // 3-letter code
}

type Name struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}

type Flags struct {
	PNG string `json:"png"`
	SVG string `json:"svg"`
	Alt string `json:"alt"`
}

type Currency struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
