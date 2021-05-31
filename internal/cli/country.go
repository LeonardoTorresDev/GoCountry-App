package cli

import (
	"fmt"
	"log"

	"github.com/LTSpark/Country-App/internal/domain"
	"github.com/LTSpark/Country-App/internal/errors"
	"github.com/LTSpark/Country-App/internal/fetching"
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

func InitCountriesCmd(service fetching.Service) *cobra.Command {

	countryCmd := &cobra.Command{
		Use:   "country",
		Short: "Print data about countries with similar name to given",
		Long: `This command prints a JSON which contains information about countries
which is received from an API called RestCountries, this info can be 
printed in console or written on a csv file`,
		Run: runCountriesCmd(service),
	}

	//Search flags
	countryCmd.Flags().StringP(nameFlag, "n", "NoNameGiven", "name of the country")
	countryCmd.Flags().StringP(regionFlag, "r", "NoRegionGiven", "region to search countries")

	//Pagination flags
	countryCmd.Flags().IntP(skipFlag, "s", 0, "value to start pagination")
	countryCmd.Flags().IntP(limitFlag, "l", 5, "number of pages")

	//CSV flags
	countryCmd.Flags().StringP(fileNameFlag, "f", "countries", "name of the csv file")

	//Console flags
	countryCmd.Flags().BoolP(consoleFlag, "c", false, "write countries info on console")

	return countryCmd

}

func InitWriteCmd(service fetching.Service) *cobra.Command {
	writeCmd := &cobra.Command{
		Use:   "write",
		Short: "Write data of countries around the world in a csv file",
		Long: `This command creates a csv file which contain information
about countries from all around the world, you can modify the name of file as well`,
		Run: runWriteCmd(service),
	}

	writeCmd.Flags().StringP(fileNameFlag, "f", "countries", "name of the csv file")
	return writeCmd
}

func runCountriesCmd(service fetching.Service) CobraFn {

	return func(cmd *cobra.Command, args []string) {

		name, _ := cmd.Flags().GetString(nameFlag)
		region, _ := cmd.Flags().GetString(regionFlag)
		skip, _ := cmd.Flags().GetInt(skipFlag)
		limit, _ := cmd.Flags().GetInt(limitFlag)

		csvName, _ := cmd.Flags().GetString(fileNameFlag)
		console, _ := cmd.Flags().GetBool(consoleFlag)

		flags := domain.Flags{
			Name:   name,
			Region: region,
			Skip:   skip,
			Limit:  limit,
		}

		countries, err := service.FetchCountries(flags)
		if errors.IsDataUnreacheable(err) {
			log.Fatal(err)
		}

		err = service.WriteCountries(countries, csvName)
		if errors.IsFileWritingFailed(err) {
			log.Fatal(err)
		}

		if console {
			fmt.Println(countries)
		}
	}
}

func runWriteCmd(service fetching.Service) CobraFn {
	return func(cmd *cobra.Command, args []string) {

		csvName, _ := cmd.Flags().GetString(fileNameFlag)
		countries, _ := service.FetchAllCountries()

		err := service.WriteAllCountries(countries, csvName)
		if errors.IsFileWritingFailed(err) {
			log.Fatal(err)
		}

	}
}
