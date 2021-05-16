package countrycli

import (
	"encoding/json"
	"fmt"
)

// Country struct representation
type Country struct {
	Name        string     `json:"name"`
	Capital     string     `json:"capital"`
	Population  int        `json:"population"`
	Currencies  []Currency `json:"currencies"`
	Area        float64    `json:"area"`
	Demonym     string     `json:"demonym"`
	Languages   []Language `json:"languages"`
	NumericCode string     `json:"numericCode"`
	ISO2Code    string     `json:"alpha2Code"`
	ISO3Code    string     `json:"alpha3Code"`
	Region      Region     `json:"region"`
}

type Currency struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Symbol string `json:"symbol"`
}

type Language struct {
	Name       string `json:"name"`
	NativeName string `json:"nativeName"`
	IsoCode    string `json:"iso639_1"`
}

type Region int

const (
	Americas Region = iota
	Europe
	Asia
	Africa
	Oceania
)

var toID = map[string]Region{
	"Americas": Americas,
	"Europe":   Europe,
	"Asia":     Asia,
	"Africa":   Africa,
	"Oceania":  Oceania,
}

var toString = map[Region]string{
	Americas: "Americas",
	Europe:   "Europe",
	Asia:     "Asia",
	Africa:   "Africa",
	Oceania:  "Oceania",
}

func (r *Region) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*r = toID[j]
	return nil
}

type CountryRepo interface {
	GetCountries() ([]Country, error)
}

func (c Country) String() (s string) {

	s = fmt.Sprintf("Country: %s", c.Name)
	s += fmt.Sprintf("\nCapital: %s", c.Capital)
	s += fmt.Sprintf("\nPopulation: %d", c.Population)
	s += fmt.Sprintf("\nArea: %0.2f", c.Area)
	s += fmt.Sprintf("\nDemonym: %s", c.Demonym)
	s += fmt.Sprintf("\nRegion: %s", toString[c.Region])
	s += "\nCurrencies: "
	for _, currency := range c.Currencies {
		s += fmt.Sprintf("\n  %s (%s %s)", currency.Name, currency.Symbol, currency.Code)
	}
	s += "\nLanguages: "
	for _, language := range c.Languages {
		s += fmt.Sprintf("\n  %s: %s (%s)", language.IsoCode, language.Name, language.NativeName)
	}
	s += fmt.Sprintf("\nISO CODES: %s %s %s", c.NumericCode, c.ISO2Code, c.ISO3Code)
	return
}
