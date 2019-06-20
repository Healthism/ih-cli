package cmd

import (
	"os"
	"fmt"
	"strings"

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
		log.VERBOSE_OUTPUT, _ = cmd.Flags().GetBool("verbose")
		clusterTarget := strings.Split(CLUSTER, " ")

		if err := util.UpdateCLI(VERSION); err != nil {
			log.Print("Please enter your command again")
			os.Exit(1)
			return
		}

		if len(clusterTarget) != 2 {
			log.Errorf("Invalid Cluster Target Length - [%s]", CLUSTER)
			os.Exit(0)
		}
		
		cliInformation := "Staring IH CLI @ cluster: %s - release: %s"
		log.Printf(cliInformation, log.Yellow(CLUSTER), log.Yellow(RELEASE))
		
		if err := git.Get(RELEASE); err != nil {
			log.Errorf("Invalid Release Target - [%s]", RELEASE)
			os.Exit(0)
		}

		cluster := fmt.Sprintf(CLUSTER_TEMPLATE,clusterTarget[0], clusterTarget[1])
		if err := util.Exec("kubectl","config", "use-context", cluster); err != nil {
			log.Errorf("Invalid Cluster Target - [%s]", CLUSTER)
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
	cobra.OnInitialize()

	rootCmd.PersistentFlags().Bool("verbose", false, "Verbose Output")
	rootCmd.PersistentFlags().Bool("update", false, "Update InputHealth CLI")
	rootCmd.PersistentFlags().StringP("release", "r", "", "Release Target (required)")
	rootCmd.PersistentFlags().StringP("cluster", "c", "inputhealth-chr staging", "Release Cluster")

	rootCmd.MarkPersistentFlagRequired("release")
}
