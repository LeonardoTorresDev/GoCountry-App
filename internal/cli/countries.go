package cli

import (
	"fmt"

	countrycli "github.com/LTSpark/Country-App/internal"
	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const nameFlag = "name"

func InitCountriesCmd(repository countrycli.CountryRepo) *cobra.Command {
	countryCmd := &cobra.Command{
		Use:   "country",
		Short: "Use:\n  Print data about a country by given its name",
		Run:   runCountriesCmd(repository),
	}

	countryCmd.Flags().StringP(nameFlag, "n", "", "name of the country")
	return countryCmd
}

func runCountriesCmd(repository countrycli.CountryRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {

		countries, _ := repository.GetCountries()
		name, _ := cmd.Flags().GetString(nameFlag)

		if name != "" {
			for id, country := range countries {
				if country.Name == name {
					fmt.Printf("%v \nIndex: %v", country, id)
					return
				}
			}
		} else {
			fmt.Println("Please enter a valid country name (in english) Example: Japan")
		}
	}
}
