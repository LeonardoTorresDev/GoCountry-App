package fetching

import (
	"github.com/LTSpark/Country-App/internal/domain"
)

//Write Service Definition
type writeService struct {
	writeCountryRepo domain.WriteCountryRepo
}

type WriteService interface {
	//Write csv countries
	WriteCountriesService(countries []domain.Country, csvName string) (err error)
	//Write new csv with all countries
	WriteAllCountriesService(countries []domain.Country, csvName string) (err error)
}

func NewWriteService(w domain.WriteCountryRepo) WriteService {
	return &writeService{w}
}

func (w *writeService) WriteCountriesService(countries []domain.Country, csvName string) (err error) {
	return w.writeCountryRepo.StoreCountryList(countries, csvName)
}

func (w *writeService) WriteAllCountriesService(countries []domain.Country, csvName string) (err error) {
	return w.writeCountryRepo.StoreAllCountriesList(countries, csvName)
}
