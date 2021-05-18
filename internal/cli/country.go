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

func InitCountriesCmd(read country.CountryRepo, write country.WriteCountryRepo) *cobra.Command {

	countryCmd := &cobra.Command{
		Use:   "country",
		Short: "Print data about countries with similar name to given",
		Long: `This command prints a JSON which contains information about countries
which is received from an API called RestCountries, this info can be 
printed in console or written on a csv file`,
		Run: runCountriesCmd(read, write),
	}

	countryCmd.Flags().StringP(nameFlag, "n", "", "name of the country")
	countryCmd.Flags().StringP(regionFlag, "r", "", "region to search countries")
	countryCmd.Flags().IntP(skipFlag, "s", 0, "value to start pagination")
	countryCmd.Flags().IntP(limitFlag, "l", 5, "number of pages")

	return countryCmd

}

func runCountriesCmd(read country.CountryRepo, write country.WriteCountryRepo) CobraFn {

	return func(cmd *cobra.Command, args []string) {

		var countries []country.Country

		name, _ := cmd.Flags().GetString(nameFlag)
		region, _ := cmd.Flags().GetString(regionFlag)
		skip, _ := cmd.Flags().GetInt(skipFlag)
		limit, _ := cmd.Flags().GetInt(limitFlag)

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

		countries = printResponse(countries, skip, limit)
		write.StoreCountryList(countries)

	}

}

func printResponse(c []country.Country, skip, limit int) []country.Country {
	skip, limit = utils.ParseSkipLimit(len(c), skip, limit)
	if len(c) == 0 {
		fmt.Print("No countries founded")
		return nil
	}
	c = c[skip:limit]
	fmt.Println(c)
	fmt.Printf("Total results: %d", len(c))
	fmt.Printf("\nTotal response: %d", len(c[skip:limit]))
	return c
}
