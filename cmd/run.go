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
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:   RUN,
	Short: RUN_DESCRIPTION_SHORT,
	Long:  RUN_DESCRIPTION_LONG,
	Run: func(cmd *cobra.Command, args []string) {
		command := "rails console"
		if len(args) > 0 {
			command = strings.Join(args," ");
		}

		viper.SetConfigName("values")
		viper.AddConfigPath(COFING_PATH)
		if err := viper.ReadInConfig(); err != nil {
			log.Errorf("[VIPER] Failed to read configuration: %v", err)
			return
		}

		imageURL, err := ioutil.ReadFile(IMAGE_PATH)
		if err != nil {
			log.Errorf("[IOUTIL] Failed to read app image url: %v", err)
			return
		}

		input := make(map[string]interface{})
		uuid := time.Now().Format("20060102150405")
		input["release_name"] = RELEASE
		input["command"] = command
		input["unique_id"] = uuid
		input["image_url"] = string(imageURL)

		
		if util.ExecuteTemplate(task.TASK, input, "/tmp/ih-launch.yaml") != nil {
			return
		}

		/** Kubectl: Apply Task File **/
		if util.Exec("kubectl", "-n", "chr-qa", "apply", "-f", "/tmp/ih-launch.yaml") != nil {
			return
		}

		/** Kubectl: Attach Work Unit **/
		var attachError error
		workId := fmt.Sprintf("job.batch/%s-console-%s", RELEASE, uuid)
		for retry := 3; retry > 0; retry-- {
			time.Sleep(time.Second)
			if attachError := util.Exec("kubectl", "-n", "chr-qa", "attach", "-it", workId); attachError == nil {
				break;
			}

			if retry != 1 {
				log.Error("Encountered error while connecting the pod, trying again...")
			}
		}
		
		if attachError != nil {
			log.Error("Unable to attach to the pod. Please submit a bug report.")
			log.Error("Deleting the unattachable pod...")
		}

		/** Close Kubectl After **/
		jobName := fmt.Sprintf("%s-console-%s", RELEASE, uuid)
		if util.Exec("kubectl", "-n", "chr-qa", "delete", "job", jobName) != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
