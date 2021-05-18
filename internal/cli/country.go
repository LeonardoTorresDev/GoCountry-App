package cli

import (
	"fmt"

	country "github.com/LTSpark/Country-App/internal/domain"
	"github.com/LTSpark/Country-App/utils"
	"github.com/spf13/cobra"
)

const (
	nameFlag     = "name"
	regionFlag   = "region"
	skipFlag     = "skip"
	limitFlag    = "limit"
	csvFlag      = "to-csv"
	fileNameFlag = "file-name"
	consoleFlag  = "console"
)

func InitCountriesCmd(read country.CountryRepo, write country.WriteCountryRepo) *cobra.Command {

	countryCmd := &cobra.Command{
		Use:   "country",
		Short: "Print data about countries with similar name to given",
		Long: `This command prints a JSON which contains information about countries
which is received from an API called RestCountries, this info can be 
printed in console or written on a csv file`,
		Run: runCountriesCmd(read, write),
	}

	//Search flags
	countryCmd.Flags().StringP(nameFlag, "n", "", "name of the country")
	countryCmd.Flags().StringP(regionFlag, "r", "", "region to search countries")

	//Pagination flags
	countryCmd.Flags().IntP(skipFlag, "s", 0, "value to start pagination")
	countryCmd.Flags().IntP(limitFlag, "l", 5, "number of pages")

	//CSV flags
	countryCmd.Flags().BoolP(csvFlag, "t", true, "permssion to write a csv file")
	countryCmd.Flags().StringP(fileNameFlag, "f", "countries", "name of the csv file")

	//Console flags
	countryCmd.Flags().BoolP(consoleFlag, "c", false, "write countries info on console")

	return countryCmd

}

func runCountriesCmd(read country.CountryRepo, write country.WriteCountryRepo) CobraFn {

	return func(cmd *cobra.Command, args []string) {

		var countries []country.Country

		name, _ := cmd.Flags().GetString(nameFlag)
		region, _ := cmd.Flags().GetString(regionFlag)
		skip, _ := cmd.Flags().GetInt(skipFlag)
		limit, _ := cmd.Flags().GetInt(limitFlag)
		csv, _ := cmd.Flags().GetBool(csvFlag)
		csvName, _ := cmd.Flags().GetString(fileNameFlag)
		console, _ := cmd.Flags().GetBool(consoleFlag)

		if name != "" {
			countries, _ = read.NameCountriesStrategy(name)
			if region != "" {
				countriesRegion, _ := read.RegionCountriesStrategy(region)
				countries = utils.IntersectCountrySlices(countries, countriesRegion)
			}
		} else if region != "" {
			countries, _ = read.RegionCountriesStrategy(region)
		} else {
			countries, _ = read.AllCountriesStrategy()
		}

		numberOfCountries := len(countries)
		if numberOfCountries == 0 {
			fmt.Print("No countries founded")
			return
		}

		skip, limit = utils.ParseSkipLimit(numberOfCountries, skip, limit)
		countries = countries[skip:limit]

		if csv {
			write.StoreCountryList(countries, csvName)
			fmt.Println("Data recovered on csv file correctly")
		}

		if console {
			fmt.Println(countries)
		}

		fmt.Printf("Total results: %d", numberOfCountries)
		fmt.Printf("\nTotal response: %d", len(countries))

	}
}
