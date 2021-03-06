package cmd

import (
	"os"

	"github.com/Korazza/templay/config"
	"github.com/spf13/cobra"
)

var Config config.Config

var rootCmd = &cobra.Command{
	Use:   "templay",
	Short: "A tool for generating templated directories",
	Long: `Templay
  A tool for generating templated directories`,
}

func Execute(c config.Config) {
	Config = c

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
