package cmd

import (
	"io/ioutil"
	"strconv"
	"time"

	"fmt"

	"ih/lib/log"
	"ih/lib/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:   RUN,
	Short: RUN_DESCRIPTION_SHORT,
	Long:  RUN_DESCRIPTION_LONG,
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("Initiating Run...")

		viper.SetConfigName("values")
		viper.AddConfigPath("./config")
		err = viper.ReadInConfig()
		if err != nil {
			log.Errorf("[VIPER] Failed to read configuration: %v", err)
			return
		}
		log.Print("[VIPER] Configuration loaded")

		imageURL, err := ioutil.ReadFile("config/substitutions/_APP_IMAGE_URL")
		if err != nil {
			log.Errorf("[IOUTIL] Failed to read app image url: %v", err)
			return
		}
		log.Print("[IOUTIL] App image url loaded")


		input := make(map[string]interface{})
		uuid := strconv.FormatInt(time.Now().Unix(), 36)
		input["release_name"] = RELEASE
		input["command"], _ = cmd.Flags().GetString("command")
		input["unique_id"] = uuid
		input["image_url"] = string(imageURL)

		err = util.ExecuteTemplate("tasks/task.yaml", input, "tasks/launch.yaml")
		if err != nil {
			return
		}

		/** Kubectl: Apply Task File **/
		err = util.Exec("kubectl", "-n", "chr-qa", "apply", "-f", "tasks/launch.yaml")
		if err != nil {
			return
		}

		/** Kubectl: Attach Work Unit **/
		workId := fmt.Sprintf("job.batch/%s-console-%s", RELEASE, uuid)
		err = util.Exec("kubectl", "-n", "chr-qa", "attach", "-it", workId)
		if err != nil {
			return
		}

		/** Close Kubectl After **/

		log.Print("Run Successful!")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("command", "c", "rails console", "Task Command")
}
