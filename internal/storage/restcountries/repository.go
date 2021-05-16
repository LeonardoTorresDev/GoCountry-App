package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	countrycli "github.com/LTSpark/Country-App/internal"
)

const (
	restCountriesUrl = "https://restcountries.eu/rest/v2/all"
)

type countryRepo struct {
	url string
}

func NewCountriesRepository() countrycli.CountryRepo {
	return &countryRepo{url: restCountriesUrl}
}

func (c *countryRepo) GetCountries() (countries []countrycli.Country, err error) {

	response, err := http.Get(fmt.Sprintf("%v", restCountriesUrl))
	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &countries)

	if err != nil {
		return nil, err
	}

	return
}
