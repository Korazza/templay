package cmd

import (
	"fmt"

	"github.com/Korazza/templay/flags"
	"github.com/Korazza/templay/logger"
	"github.com/Korazza/templay/templay"
	"github.com/spf13/cobra"
)

var templayVars flags.TemplayVars

var genCmd = &cobra.Command{
	Use:     "generate [-d destination] [-f file | -v key=value]... [flags] name",
	Aliases: []string{"gen"},
	Args:    cobra.ExactArgs(1),
	Short:   "Generate a templay",
	Long:    `Generate a templay`,
	RunE: func(cmd *cobra.Command, args []string) error {
		destination, err := cmd.Flags().GetString("dest")

		if err != nil {
			return err
		}

		varsFile, err := cmd.Flags().GetString("file")

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

		templayVars.Load(varsFile)

		if templayVars == nil {
			if err = templay.CopyDirectory(templayPath, destination); err != nil {
				return err
			}
		} else {
			if err = templay.ParseDirectory(templayPath, destination, templayVars); err != nil {
				return err
			}
		}

		logger.Response.Printf("Templay %s successfully generated in %s", templayName, destination)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().StringP("dest", "d", ".", "Destination of the templay")
	genCmd.Flags().StringP("file", "f", ".templayvars.yml", "Variable file")
	genCmd.Flags().VarP(&templayVars, "var", "v", "Variable to pass to templay")
}
