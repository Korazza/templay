package cmd

import (
	"fmt"

	"github.com/Korazza/templay/utils"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen [-d destination] [flags] name",
	Args:  cobra.ExactArgs(1),
	Short: "Generate a templay",
	Long:  `Generate a templay`,
	RunE: func(cmd *cobra.Command, args []string) error {
		destination, err := cmd.Flags().GetString("dest")

		if err != nil {
			return err
		}

		templayName := args[0]

		var templayPath string

		for templay, path := range Config.Templays {
			if templayName == templay {
				templayPath = path
			}
		}

		if templayPath == "" {
			return fmt.Errorf("templay %s not found", templayName)
		}

		if err = utils.CopyDir(templayPath, destination); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().StringP("dest", "d", ".", "Destination of the templay")
}
