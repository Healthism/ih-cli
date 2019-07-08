package cmd

import (
	"os"
	"os/signal"
	"syscall"

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
		RELEASE, _ = cmd.Flags().GetString("release")
		CLUSTER, _ = cmd.Flags().GetString("cluster")
		NAME_SPACE, _ = cmd.Flags().GetString("namespace")
		log.VERBOSE_OUTPUT, _ = cmd.Flags().GetBool("verbose")
		KUBE_ENV = []string{"--cluster", CLUSTER, "-n", NAME_SPACE}

		/** Intro Message **/
		cliInformation := "Staring IH CLI @ cluster: %s - release: %s"
		log.Printf(cliInformation, log.Yellow(CLUSTER), log.Yellow(RELEASE))

		/** UPDATE CLI **/
		if err := util.UpdateCLI(VERSION); err != nil {
			log.Positive("Please enter your command again")
			os.Exit(1)
			return
		}

		/** Loading Release Repository **/
		if err := git.Get(RELEASE); err != nil {
			log.Errorf("Invalid Release Target - [%s]", RELEASE)
			os.Exit(0)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		git.Reset()
		log.Print("Good Bye!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Errorf("%s", err)
		os.Exit(1)
	}
}

func init() {
	signal.Ignore(syscall.SIGINT)

	cobra.OnInitialize()

	rootCmd.PersistentFlags().Bool("verbose", false, "Verbose Output")
	rootCmd.PersistentFlags().Bool("update", false, "Update InputHealth CLI")
	rootCmd.PersistentFlags().StringP("release", "r", "", "Release Target (required)")
	rootCmd.PersistentFlags().StringP("namespace", "n", "chr-qa", "Release Name Space")
	rootCmd.PersistentFlags().StringP("cluster", "c", "gke_inputhealth-chr_northamerica-northeast1-a_staging", "Release Cluster")

	rootCmd.MarkPersistentFlagRequired("release")
}
