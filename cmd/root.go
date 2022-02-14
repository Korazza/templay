package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "temply",
	Short: "A monorepo tool for generating templated folders",
	Long:  `A monorepo tool for generating templated folders`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
