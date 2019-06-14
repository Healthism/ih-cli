package cmd

import (
	"os"

	"ih/lib/git"
	"ih/lib/log"
	"ih/lib/util"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     IH,
	Version: VERSION,
	Short:   ROOT_DESCRIPTION_SHORT,
	Long:    ROOT_DESCRIPTION_LONG,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := util.UpdateCLI(VERSION); err != nil {
			log.Print("Please enter your command again")
			os.Exit(1)
			return
		}
		log.Verbose, _ = cmd.Flags().GetBool("verbose")
		RELEASE, _ = cmd.Flags().GetString("release")
		log.Printf("Staring IH CLI, specified release - [%s]", RELEASE)
		if err := git.Get(RELEASE); err != nil {
			log.Errorf("Invalid Release Target - [%s]", RELEASE)
			os.Exit(0)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Errorf("%s", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().Bool("update", false, "Update InputHealth CLI")
	rootCmd.PersistentFlags().Bool("verbose", false, "Verbose Output")
	rootCmd.PersistentFlags().StringP("release", "r", "", "Release Target (required)")

	rootCmd.MarkPersistentFlagRequired("release")
}
