package cmd

import (
	"fmt"

	"github.com/Korazza/templay/logger"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [flags]",
	Short: "List all templays",
	Long:  `List all templays`,
	Run: func(cmd *cobra.Command, args []string) {
		list := fmt.Sprintf("%-10s %s", "Name", "Path")
		for templay, path := range Config.Templays {
			list = fmt.Sprintf("%s\n%-10s %s", list, templay, path)
		}
		logger.Response.Print(list)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
