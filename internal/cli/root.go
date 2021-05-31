package cli

import (
	"github.com/LTSpark/Country-App/internal/fetching"
	"github.com/LTSpark/Country-App/internal/storage/csv"
	"github.com/LTSpark/Country-App/internal/storage/restcountries"
	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

var rootCmd = &cobra.Command{
	Use: "country-cli",
}

func Execute() {

	write := csv.NewWriteCountryRepository()
	read := restcountries.NewCountriesRepository()
	service := fetching.NewService(read, write)

	rootCmd.AddCommand(InitCountriesCmd(service))
	rootCmd.AddCommand(InitWriteCmd(service))

	rootCmd.Execute()
}
