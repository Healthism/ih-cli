package cmd

import (
	"fmt"
	"github.com/Healthism/ih-cli/config"
	"github.com/Healthism/ih-cli/util"
	"github.com/Healthism/ih-cli/util/console"
	"github.com/Healthism/ih-cli/util/git"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   config.CONFIG,
	Short: config.CONFIG_DESCRIPTION_SHORT,
	Long:  config.CONFIG_DESCRIPTION_LONG,
	Run: func(cmd *cobra.Command, args []string) {
		cluster, nameSpace, release := util.ParseFlags(cmd)

		console.AddTable([]string{
			fmt.Sprintf("%-20v: %s", console.SprintYellow("NameSpace"), nameSpace),
			fmt.Sprintf("%-20v: %s", console.SprintYellow("Cluster"), cluster),
			fmt.Sprintf("%-20v: %s", console.SprintYellow("Release"), release),
		})

		gitLoading := console.ShowLoading("Loading configuration resources", "[1/2]")
		err := git.Load(release)
		gitLoading.HideLoading(err)
		if err != nil {
			console.AddLine()
			console.Errorf("%s", err)
			console.Error("⚠️  Error while loading configuration resources ...")
			console.Error("This error propably was caused due to conflict within configuration repository")
			console.Error("Try resetting your resource repository by running commands:")
			console.Errorf("$ rm -rf %s", config.GIT_PATH)
			console.Errorf("$ gcloud source repos clone staging-deployment %s --project=inputhealth-chr", config.GIT_PATH)
			return
		}

		err = util.InteractiveExec(config.EDITOR, config.VALUE_PATH)
		if err != nil {
			console.AddLine()
			console.Errorf("%s", err)
			console.Error("⚠️  Error opening configuration resources ...")
			console.Error("This error propably was caused due to conflict within configuration repository")
			console.Error("Try resetting your resource repository by running commands:")
			console.Errorf("$ rm -rf %s", config.GIT_PATH)
			console.Errorf("$ gcloud source repos clone staging-deployment %s --project=inputhealth-chr", config.GIT_PATH)
		}

		gitLoading = console.ShowLoading("Updating configuration resources", "[2/2]")
		err = git.Save(release)
		gitLoading.HideLoading(err)
		if err != nil {
			console.AddLine()
			console.Errorf("%s", err)
			console.Error("⚠️  Error while Updating new configuration")
			console.Error("This error propably was caused by either")
			console.Error("1) Nothing has been changed")
			console.Error("2) Unable to update remote configuration repository\n")
			console.Error("The error(2) propably was caused due to conflict within configuration repository")
			console.Error("Try resetting your resource repository by running commands:")
			console.Errorf("$ rm -rf %s", config.GIT_PATH)
			console.Errorf("$ gcloud source repos clone staging-deployment %s --project=inputhealth-chr", config.GIT_PATH)
			return
		}

		console.AddLine()
		console.Print(console.SprintYellow("🚀 Configuration Succesfully Updated"))
		console.Print(console.SprintYellow("👋 Good Bye"))
		console.AddLine()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.PersistentFlags().String("qa", "", "Access QA Server")
	configCmd.PersistentFlags().String("staging", "", "Access Staging Server")

	configCmd.PersistentFlags().String("release", "r", "Release Target")
	configCmd.PersistentFlags().String("namespace", "n", "Release Name Space")
	configCmd.PersistentFlags().StringP("cluster", "c", "gke_inputhealth-chr_northamerica-northeast1-a_staging", "Release Cluster")
}
