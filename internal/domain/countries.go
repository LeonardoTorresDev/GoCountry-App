package domain

import (
	"encoding/json"
)

// Country struct representation
type Country struct {
	Name           string          `json:"name"`
	Capital        string          `json:"capital"`
	Population     int             `json:"population"`
	Region         string          `json:"region"`
	Subregion      string          `json:"subregion"`
	RegionalBlocs  []RegionalBlocs `json:"regionalBlocs"`
	Currencies     []Currency      `json:"currencies"`
	Area           float64         `json:"area"`
	Demonym        string          `json:"demonym"`
	Languages      []Language      `json:"languages"`
	NumericCode    string          `json:"numericCode"`
	TopLevelDomain []string        `json:"topLevelDomain"`
	CallingCodes   []string        `json:"callingCodes"`
	ISO2Code       string          `json:"alpha2Code"`
	ISO3Code       string          `json:"alpha3Code"`
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

type RegionalBlocs struct {
	Acronym string `json:"acronym"`
	Name    string `json:"name"`
}

type CountryRepo interface {
	GetCountriesByName(name string) ([]Country, error)
	GetCountriesByRegion(region string) ([]Country, error)
	GetAllCountries() ([]Country, error)
}

func (c Country) String() (s string) {
	out, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		panic(err)
	}
	s = string(out)
	return
}
