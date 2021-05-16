package main

import (
	"github.com/LTSpark/Country-App/internal/cli"
	repository "github.com/LTSpark/Country-App/internal/storage/restcountries"
	"github.com/spf13/cobra"
)

func main() {

	repo := repository.NewCountriesRepository()
	rootCmd := &cobra.Command{Use: "country-cli"}
	rootCmd.AddCommand(cli.InitCountriesCmd(repo))
	rootCmd.Execute()

}
