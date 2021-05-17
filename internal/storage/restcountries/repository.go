package restcountries

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	country "github.com/LTSpark/Country-App/internal/domain"
)

const (
	RestCountriesUrl = "https://restcountries.eu/rest/v2"
	NameEndpoint     = "/name/"
	RegionEndpoint   = "/region/"
	AllEndpoint      = "/all/"
)

type countryRepo struct {
	url string
}

func NewCountriesRepository() country.CountryRepo {
	return &countryRepo{url: RestCountriesUrl}
}

func (c *countryRepo) GetAllCountries() (countries []country.Country, err error) {
	url := fmt.Sprintf("%v%v", RestCountriesUrl, AllEndpoint)
	err = c.getJSONResponse(url, &countries)
	return
}

func (c *countryRepo) GetCountriesByName(name string) (countries []country.Country, err error) {

	url := fmt.Sprintf("%v%v%v", RestCountriesUrl, NameEndpoint, name)
	err = c.getJSONResponse(url, &countries)
	return

}

func (c *countryRepo) GetCountriesByRegion(region string) (countries []country.Country, err error) {

	url := fmt.Sprintf("%v%v%v", RestCountriesUrl, RegionEndpoint, region)
	err = c.getJSONResponse(url, &countries)
	return

}

func (c *countryRepo) getJSONResponse(url string, t interface{}) (err error) {

	response, err := http.Get(url)
	if err != nil {
		return err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(contents, &t)
	if err != nil {
		return err
	}

	return
}
