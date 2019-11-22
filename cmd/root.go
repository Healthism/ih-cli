package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Healthism/ih-cli/config"
	"github.com/Healthism/ih-cli/util"
	"github.com/Healthism/ih-cli/util/console"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     config.IH,
	Version: config.VERSION,
	Short:   config.ROOT_DESCRIPTION_SHORT,
	Long:    config.ROOT_DESCRIPTION_LONG,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		console.ENABLE_VERBOSE, _ = cmd.Flags().GetBool("verbose")

		updated, err := util.UpdateCLI(config.VERSION)
		if err != nil {
			os.Exit(1)
			return
		}

		if updated {
			os.Exit(0)
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		console.Errorf("%s", err)
		os.Exit(1)
	}
}

func init() {
	signal.Ignore(syscall.SIGINT)
	cobra.OnInitialize()

	rootCmd.PersistentFlags().Bool("qa", false, "QA Access")
	rootCmd.PersistentFlags().Bool("staging", false, "Staging Access")
	rootCmd.PersistentFlags().Bool("verbose", false, "Verbose Output")
	rootCmd.PersistentFlags().StringP("release", "r", "", "Release Target (required)")
	rootCmd.PersistentFlags().StringP("namespace", "n", "chr-qa", "Release Name Space")
	rootCmd.PersistentFlags().StringP("cluster", "c", "gke_inputhealth-chr_northamerica-northeast1-a_staging", "Release Cluster")
}
