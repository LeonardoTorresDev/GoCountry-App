package csv

import (
	"encoding/csv"
	"fmt"
	"os"

	country "github.com/LTSpark/Country-App/internal/domain"
)

const (
	CsvExtension = ".csv"
)

type writeCountryRepo struct {
}

func NewWriteCountryRepository() country.WriteCountryRepo {
	return &writeCountryRepo{}
}

func (w *writeCountryRepo) StoreCountryList(c []country.Country, fileName string) (err error) {

	var file *os.File
	CsvFile := fileName + CsvExtension

	if _, err := os.Stat(CsvFile); err == nil {
		fmt.Printf("Appending data to existing file %s...\n", CsvFile)
		file, err = os.OpenFile(CsvFile, os.O_RDWR|os.O_APPEND, 0660)
		if nil != err {
			return err
		}
	} else if os.IsNotExist(err) {
		fmt.Printf("Creating new file %s...\n", CsvFile)
		file, err = os.Create(CsvFile)
		if nil != err {
			return err
		}
	}

	defer file.Close()
	writer := csv.NewWriter(file)

	for _, value := range c {
		err := writer.Write(value.ToArray())
		if nil != err {
			return err
		}
	}

	defer writer.Flush()
	return nil

}

func (w *writeCountryRepo) StoreAllCountriesList(c []country.Country, fileName string) (err error) {

	CsvFile := fileName + CsvExtension

	file, err := os.Create(CsvFile)
	if nil != err {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range c {
		err := writer.Write(value.ToArray())
		if nil != err {
			return err
		}
	}

	return nil

}
