package cmd

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Healthism/ih-cli/config"
	"github.com/Healthism/ih-cli/util"
	"github.com/Healthism/ih-cli/util/console"
	"github.com/manifoldco/promptui"
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
	Run: func(cmd *cobra.Command, args []string) {
		console.AddLine()
		var cluster, nameSpace, release, action, command string
		cluster, nameSpace, release, action = getPodInformation()
		if action == "run" {
			defaultAction := "rails console"
			if strings.Contains(release, "up-patient") {
				defaultAction = "bash"
			}
			commandPrompt := promptui.Prompt{
				Label:   "Enter Your Command To Run",
				Default: defaultAction,
				Templates: &promptui.PromptTemplates{
					Prompt:  "{{ . }} ",
					Success: `  {{ "Command" | yellow }}    : `,
				},
			}
			var err error
			command, err = commandPrompt.Run()
			if err != nil {
				os.Exit(0)
			}
		}

		console.AddLine()
		console.Print("🏗  Running Command...")
		console.Print([]string{"ih", action, "--cluster", cluster, "--namespace", nameSpace, "--release", release, command})

		cmd.SetArgs([]string{action, "--cluster", cluster, "--namespace", nameSpace, "--release", release, command})
		cmd.Execute()
	},
}

func getPodInformation() (string, string, string, string) {
	type selectObject struct {
		Label string
		Value string
	}

	/** Deployment Target **/
	deployments := []selectObject{
		{"QA", "chr-qa"},
		{"Staging", "chr-staging"},
		{"Ontario Medical Spec", "chr-omd"},
	}
	deploymentPrompt := promptui.Select{
		Items: deployments,
		Templates: &promptui.SelectTemplates{
			Label:    "Select Deployment Environement ?",
			Active:   "✔ {{ .Label | cyan }}",
			Inactive: "  {{ .Label | cyan }}",
			Selected: `  {{ "Deployment" | yellow }} : {{ .Label }}`,
		},
	}
	deploymentIndex, _, err := deploymentPrompt.Run()
	if err != nil {
		os.Exit(0)
	}

	/** Release Target **/
	releases := []selectObject{
		{"Backend", "backend-1"},
		{"Desktop", "desktop"},
		{"Socket", "socket-1"},
		{"Patient App", "up-patient-backend"},
	}
	releasePrompt := promptui.Select{
		Items: releases,
		Templates: &promptui.SelectTemplates{
			Label:    "Select Release Target ?",
			Active:   "✔ {{ .Label | cyan }}",
			Inactive: "  {{ .Label | cyan }}",
			Selected: `  {{ "Release" | yellow }}    : {{ .Label }}`,
		},
	}
	releaseIndex, _, err := releasePrompt.Run()
	if err != nil {
		os.Exit(0)
	}

	/** Action **/
	actions := []selectObject{
		{"Execute Command on Server", "run"},
		{"Configure Environement Variables", "config"},
	}
	actionPrompt := promptui.Select{
		Items: actions,
		Templates: &promptui.SelectTemplates{
			Label:    "Select Action to Perform ?",
			Active:   "✔ {{ .Label | cyan }}",
			Inactive: "  {{ .Label | cyan }}",
			Selected: `  {{ "Action" | yellow }}     : {{ .Label }}`,
		},
	}
	actionIndex, _, err := actionPrompt.Run()
	if err != nil {
		os.Exit(0)
	}

	nameSpace := deployments[deploymentIndex].Value
	release := nameSpace + "-" + releases[releaseIndex].Value
	action := actions[actionIndex].Value

	return "gke_inputhealth-chr_northamerica-northeast1-a_staging", nameSpace, release, action
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
	rootCmd.PersistentFlags().Bool("verbose", false, "Verbose Output")
}
