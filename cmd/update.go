package cmd

import (
	"ih/lib/git"
	"ih/lib/log"
	"ih/lib/util"
	"ih/lib/yaml"

	"github.com/spf13/cobra"
)

var err error
var updateCmd = &cobra.Command{
	Use:   UPDATE,
	Short: UPDATE_DESCRIPTION_SHORT,
	Long:  UPDATE_DESCRIPTION_LONG,
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("Initiating Update...")

		manualOption, _ := cmd.Flags().GetBool("manual")
		if manualOption {
			err = util.Exec(EDITOR, CONFIG_PATH)
		} else {
			err = yaml.UpdateValue(args)
		}

		if err != nil {
			return
		}

		err = git.Set(RELEASE)
		if err != nil {
			return
		}

		log.Print("Update Successful!")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().BoolP("manual", "m", false, "Manual Update Configuration File")
}
