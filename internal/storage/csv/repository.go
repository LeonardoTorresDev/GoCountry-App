package csv

import (
	"encoding/csv"
	"fmt"
	"os"

	country "github.com/LTSpark/Country-App/internal/domain"
)

const (
	CsvFile = "countries.csv"
)

type writeCountryRepo struct {
}

func NewWriteCountryRepository() country.WriteCountryRepo {
	return &writeCountryRepo{}
}

func (w *writeCountryRepo) StoreCountryList(c []country.Country) (err error) {

	var file *os.File

	if _, err := os.Stat(CsvFile); err == nil {
		file, err = os.OpenFile(CsvFile, os.O_APPEND|os.O_WRONLY, 0644)
		if nil != err {
			fmt.Println(err.Error())
			return err
		}
	} else if os.IsNotExist(err) {
		file, err = os.Create(CsvFile)
		if nil != err {
			fmt.Println(err.Error())
			return err
		}
	}

	writer := csv.NewWriter(file)

	defer func() {
		writer.Flush()
		e := file.Close()
		if nil != e {
			err = e
		}
	}()

	for _, value := range c {
		err := writer.Write(value.ToArray())
		if nil != err {
			return err
		}
	}

	return nil
}
