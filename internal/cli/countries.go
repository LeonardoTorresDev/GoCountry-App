package cli

import (
	"fmt"

	country "github.com/LTSpark/Country-App/internal/domain"
	"github.com/LTSpark/Country-App/utils"
	"github.com/spf13/cobra"
)

const (
	nameFlag   = "name"
	regionFlag = "region"
	skipFlag   = "skip"
	limitFlag  = "limit"
)

func InitCountriesCmd(repository country.CountryRepo) *cobra.Command {

	countryCmd := &cobra.Command{
		Use:   "country",
		Short: "Print data about countries with similar name to given",
		Long: `This command prints a JSON which contains information about countries
which is received from an API called RestCountries, this info can be 
printed in console or written on a csv file`,
		Run: runCountriesCmd(repository),
	}

	countryCmd.Flags().StringP(nameFlag, "n", "", "name of the country")
	countryCmd.Flags().StringP(regionFlag, "r", "", "region to search countries")
	countryCmd.Flags().IntP(skipFlag, "s", 0, "value to start pagination")
	countryCmd.Flags().IntP(limitFlag, "l", 5, "number of pages")

	return countryCmd

}

func runCountriesCmd(repository country.CountryRepo) CobraFn {

	return func(cmd *cobra.Command, args []string) {

		var countries []country.Country

		name, _ := cmd.Flags().GetString(nameFlag)
		region, _ := cmd.Flags().GetString(regionFlag)
		skip, _ := cmd.Flags().GetInt(skipFlag)
		limit, _ := cmd.Flags().GetInt(limitFlag)

		if name != "" {
			if region != "" {
				countriesName, _ := repository.GetCountriesByName(name)
				countriesRegion, _ := repository.GetCountriesByRegion(region)
				countries = utils.IntersectCountrySlices(countriesName, countriesRegion)
			} else {
				countries, _ = repository.GetCountriesByName(name)
			}
		} else if region != "" {
			countries, _ = repository.GetCountriesByRegion(region)
		} else {
			countries, _ = repository.GetAllCountries()
		}

		printResponse(countries, skip, limit)

	}

}

func printResponse(c []country.Country, skip, limit int) {
	skip, limit = utils.ParseSkipLimit(len(c), skip, limit)
	fmt.Println(c[skip:limit])
	fmt.Printf("Total results: %d", len(c))
	fmt.Printf("\nTotal response: %d", len(c[skip:limit]))
}
