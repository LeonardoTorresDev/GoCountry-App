package fetching

import (
	"fmt"

	"github.com/LTSpark/Country-App/internal/domain"
	"github.com/LTSpark/Country-App/internal/utils"
)

//Service provides country fetching operations
type Service interface {
	//Fetch countries by params
	FetchCountries(f domain.Flags) ([]domain.Country, error)
	//Fetch all countries
	FetchAllCountries() ([]domain.Country, error)
	//Write csv countries
	WriteCountriesService(countries []domain.Country, csvName string) (err error)
	//Write new csv with all countries
	WriteAllCountriesService(countries []domain.Country, csvName string) (err error)
}

type service struct {
	countryRepo      domain.CountryRepo
	writeCountryRepo domain.WriteCountryRepo
}

func NewService(cr domain.CountryRepo, w domain.WriteCountryRepo) Service {
	return &service{cr, w}
}

func (s *service) FetchCountries(f domain.Flags) (countries []domain.Country, err error) {

	params := domain.Params{
		Name:   f.Name,
		Region: f.Region,
	}

	countries, err = s.countryRepo.GetCountries(params)
	if err != nil {
		return
	}

	countries, numberOfCountries := utils.ParseCountrySlice(countries, f.Skip, f.Limit)
	fmt.Printf("Founded %d countries...\n", numberOfCountries)
	fmt.Printf("\nTotal response: %d", len(countries))

	return

}

func (s *service) FetchAllCountries() ([]domain.Country, error) {
	return s.countryRepo.GetAllCountries()
}

func (s *service) WriteCountriesService(countries []domain.Country, csvName string) (err error) {
	return s.writeCountryRepo.StoreCountryList(countries, csvName)
}

func (s *service) WriteAllCountriesService(countries []domain.Country, csvName string) (err error) {
	return s.writeCountryRepo.StoreAllCountriesList(countries, csvName)
}
