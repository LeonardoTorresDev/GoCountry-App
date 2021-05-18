package domain

import (
	"encoding/json"
	"strconv"
)

// Country struct representation
type Country struct {
	Name           string          `json:"name"`
	AltNames       []string        `json:"altSpellings"`
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
	NameCountriesStrategy(name string) ([]Country, error)
	RegionCountriesStrategy(region string) ([]Country, error)
	AllCountriesStrategy() ([]Country, error)
}

type WriteCountryRepo interface {
	StoreCountryList(c []Country, fileName string) error
}

func (c Country) String() (s string) {
	out, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		panic(err)
	}
	s = string(out)
	return
}

func (c Country) ToArray() (arr []string) {

	arr = append(arr, c.Name)
	arr = append(arr, c.Capital)
	arr = append(arr, c.Region)
	arr = append(arr, c.Subregion)
	arr = append(arr, strconv.Itoa(c.Population))
	arr = append(arr, strconv.Itoa(int(c.Area)))
	arr = append(arr, c.Demonym)

	return

}
