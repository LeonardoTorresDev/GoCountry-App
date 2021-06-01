package fetching

import (
	"fmt"

	"github.com/LTSpark/Country-App/internal/domain"
	"github.com/LTSpark/Country-App/internal/utils"
)

//Service provides country fetching operations
type CountryService interface {
	//Fetch countries by params
	FetchCountries(f domain.Flags) ([]domain.Country, error)
	//Fetch all countries
	FetchAllCountries() ([]domain.Country, error)
}

type countryService struct {
	countryRepo domain.CountryRepo
}

func NewCountryService(cr domain.CountryRepo) CountryService {
	return &countryService{cr}
}

func (s *countryService) FetchCountries(f domain.Flags) (countries []domain.Country, err error) {

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
	fmt.Printf("Total response: %d\n", len(countries))

	return

}

func (s *countryService) FetchAllCountries() ([]domain.Country, error) {
	return s.countryRepo.GetAllCountries()
}
