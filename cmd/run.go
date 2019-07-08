package cmd

import (
	"io/ioutil"
	"time"

	"fmt"
	"strings"

	"ih/lib/log"
	"ih/lib/util"
	"ih/task"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   RUN,
	Short: RUN_DESCRIPTION_SHORT,
	Long:  RUN_DESCRIPTION_LONG,
	Run: func(cmd *cobra.Command, args []string) {
		uuid := time.Now().Format("20060102150405")
		jobName := fmt.Sprintf("job.batch/%s-console-%s", RELEASE, uuid)
		podSelector := fmt.Sprintf("--selector=job-name=%s-console-%s", RELEASE, uuid)

		command := "rails console"
		if len(args) > 0 {
			command = strings.Join(args, " ")
		}

		imageURL, err := ioutil.ReadFile(IMAGE_PATH)
		if err != nil {
			log.Errorf("[IOUTIL] Failed to read app image url: %v", err)
			return
		}

		templateArgs := make(map[string]interface{})
		templateArgs["UUID"] = uuid
		templateArgs["COMMAND"] = command
		templateArgs["RELEASE_NAME"] = RELEASE
		templateArgs["IMAGE_URL"] = string(imageURL)

		if util.ExecuteTemplate(task.TASK, templateArgs, "/tmp/ih-launch.yaml") != nil {
			return
		}

		/** Kubectl: Apply Task File **/
		if util.Exec("kubectl", append(KUBE_ENV, "apply", "-f", "/tmp/ih-launch.yaml")...) != nil {
			return
		}

		/** Kubectl: Wait Container Ready **/
		podName, err := util.ExecWithStdBuffer("kubectl", append(KUBE_ENV, "get", "pods", podSelector, "-o=jsonpath='{.items[0].metadata.name}'")...)
		if err != nil {
			releaseKubeResources(err, jobName)
			return
		}
		if err := util.Exec("kubectl", append(KUBE_ENV, "wait", "--for=condition=ContainersReady", "--timeout=3m", "pod", podName)...); err != nil {
			releaseKubeResources(err, jobName)
			return
		}

		/** Kubectl: Attach Work Unit **/
		if err := util.Exec("kubectl", "-n", "chr-qa", "attach", "-it", jobName); err != nil {
			releaseKubeResources(err, jobName)
			return
		}

		releaseKubeResources(nil, jobName)
	},
}

func releaseKubeResources(err error, jobName string) {
	if err != nil {
		log.Error("Unable to attach to the pod. Please submit a bug report.")
		log.Error("Deleting the unattachable pod...")
	}

	if util.Exec("kubectl", append(KUBE_ENV, "delete", jobName)...) != nil {
		return
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
}
