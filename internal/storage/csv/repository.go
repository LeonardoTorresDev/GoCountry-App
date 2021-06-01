package csv

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/LTSpark/Country-App/internal/domain"
	"github.com/LTSpark/Country-App/internal/errors"
)

const (
	CsvExtension = ".csv"
)

type writeCountryRepo struct {
}

func NewWriteCountryRepository() domain.WriteCountryRepo {
	return &writeCountryRepo{}
}

func (w *writeCountryRepo) StoreCountryList(c []domain.Country, fileName string) (err error) {

	var file *os.File
	csvFile := fileName + CsvExtension

	if _, err := os.Stat(csvFile); err == nil {
		fmt.Printf("Appending data to existing file %s...\n", csvFile)
		file, err = os.OpenFile(csvFile, os.O_RDWR|os.O_APPEND, 0660)
		if nil != err {
			return errors.WrapFileWritingFailed(err, "Error reading file %s", csvFile)
		}
	} else if os.IsNotExist(err) {
		fmt.Printf("Creating new file %s...\n", csvFile)
		file, err = os.Create(csvFile)
		if nil != err {
			return errors.WrapFileWritingFailed(err, "Error creating file %s", csvFile)
		}
	}

	defer file.Close()
	writer := csv.NewWriter(file)

	//Add overwriter verifier
	for _, value := range c {
		err := writer.Write(value.ToArray())
		if nil != err {
			return errors.WrapFileWritingFailed(err, "Error writing value %s on file %s", value, csvFile)
		}
	}

	defer writer.Flush()
	return nil

}

func (w *writeCountryRepo) StoreAllCountriesList(c []domain.Country, fileName string) (err error) {

	csvFile := fileName + CsvExtension

	file, err := os.Create(csvFile)
	if nil != err {
		return errors.WrapFileWritingFailed(err, "Error creating file %s", csvFile)
	}

	fmt.Printf("csv file '%s' created successfully", csvFile)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Name", "Capital", "Region", "Subregion", "Population", "Area", "Demonym"})
	for _, value := range c {
		err := writer.Write(value.ToArray())
		if nil != err {
			return errors.WrapFileWritingFailed(err, "Error writing value %s on file %s", value, csvFile)
		}
	}

	return nil

}
