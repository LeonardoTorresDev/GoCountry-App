package restcountries

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/LTSpark/Country-App/internal/domain"
	"github.com/LTSpark/Country-App/internal/errors"
	"github.com/LTSpark/Country-App/internal/utils"
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

func NewCountriesRepository() domain.CountryRepo {
	return &countryRepo{url: RestCountriesUrl}
}

func (c *countryRepo) AllCountriesStrategy() (countries []domain.Country, err error) {

	url := fmt.Sprintf("%v%v", RestCountriesUrl, AllEndpoint)
	err = c.getJSONResponse(url, &countries)
	if err != nil {
		return nil, err
	}

	return

}

func (c *countryRepo) GetCountries(p domain.Params) (countries []domain.Country, err error) {

	var countriesByName []domain.Country
	var countriesByRegion []domain.Country

	urlName := fmt.Sprintf("%v%v%v", RestCountriesUrl, NameEndpoint, p.Name)
	urlRegion := fmt.Sprintf("%v%v%v", RestCountriesUrl, RegionEndpoint, p.Region)

	err = c.getJSONResponse(urlName, &countriesByName)
	if err != nil {
		return nil, err
	}

	err = c.getJSONResponse(urlRegion, &countriesByRegion)
	if err != nil {
		return nil, err
	}

	countries = utils.IntersectCountrySlices(countriesByName, countriesByRegion)
	return

}

func (c *countryRepo) getJSONResponse(url string, t *[]domain.Country) (err error) {

	response, err := http.Get(url)
	if err != nil {
		return errors.WrapDataUnreacheable(err, "Error getting response to %s", url)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.WrapDataUnreacheable(err, "Error reading the response to %s", url)
	}

	err = json.Unmarshal(contents, &t)
	if err != nil {
		return nil
	}

	return

}
