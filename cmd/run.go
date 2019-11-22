package cmd

import (
	"fmt"
	"strings"

	"github.com/Healthism/ih-cli/config"
	"github.com/Healthism/ih-cli/util"
	"github.com/Healthism/ih-cli/util/console"
	"github.com/Healthism/ih-cli/util/git"
	"github.com/Healthism/ih-cli/util/kubectl"
	"github.com/spf13/cobra"
)

var (
	job kubectl.Job
)

var runCmd = &cobra.Command{
	Use:   config.RUN,
	Short: config.RUN_DESCRIPTION_SHORT,
	Long:  config.RUN_DESCRIPTION_LONG,
	Run: func(cmd *cobra.Command, args []string) {
		release, cluster, nameSpace := util.ParseFlags(cmd)
		command := "rails console"
		if len(args) > 0 {
			command = strings.Join(args, " ")
		}

		console.AddTable([]string{
			fmt.Sprintf("NameSpace  : %s", console.SprintYellow(nameSpace)),
			fmt.Sprintf("Cluster    : %s", console.SprintYellow(cluster)),
			fmt.Sprintf("Release    : %s", console.SprintYellow(release)),
			fmt.Sprintf("Command    : %s", console.SprintYellow(command)),
		})

		gitLoading := console.ShowLoading("Loading configuration resources", "[1/3]")
		err := git.Load(release)
		gitLoading.HideLoading(err)
		if err != nil {
			showError(err)
			return
		}

		createLoading := console.ShowLoading("Request job creation", "[2/3]")
		job = kubectl.New(nameSpace, cluster, release, command)
		err = job.Create()
		createLoading.HideLoading(err)
		if err != nil {
			showError(err)
			return
		}

		waitLoading := console.ShowLoading("Waiting for the pod to be initialized", "[3/3]")
		err = job.Wait()
		waitLoading.HideLoading(err)
		if err != nil {
			showError(err)
			return
		}

		console.AddLine()
		console.Print(console.SprintYellow("🚀 Attaching to the pod ... \n"))
		err = job.Attach()
		if err != nil {
			showError(err)
		}

		console.AddLine()
		deleteLoading := console.ShowLoading("Deleting the pod ...", "")
		err = job.Delete()
		deleteLoading.HideLoading(err)
		if err != nil {
			showError(err)
			return
		}

		console.Print(console.SprintYellow("👋 Good Bye"))
		console.AddLine()
	},
}

func showError(err error) {
	console.AddLine()
	console.Errorf("%s", err)
	job.Delete()
	console.Error("⚠️  Error occured while running your command ...")
	console.Error("Please try to run with '--verbose' flag to identify the source of error.")
}

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.MarkPersistentFlagRequired("release")
}
