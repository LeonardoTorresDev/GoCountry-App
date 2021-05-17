package cli

import (
	"github.com/LTSpark/Country-App/internal/storage/restcountries"
	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

var rootCmd = &cobra.Command{
	Use: "country-cli",
}

func Execute() {
	read := restcountries.NewCountriesRepository()
	rootCmd.AddCommand(InitCountriesCmd(read))
	rootCmd.Execute()
}
