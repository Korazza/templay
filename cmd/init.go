package cmd

import (
	"io/ioutil"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [flags]",
	Short: "Initialize templay",
	Long:  `Initialize templay`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ioutil.WriteFile(".templays.yml", []byte("templays:"), 0644)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
