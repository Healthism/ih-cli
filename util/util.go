package util

import (
	"os"
	"os/exec"
	"strings"

	"github.com/Healthism/ih-cli/util/console"
	"github.com/spf13/cobra"
)

func ParseFlags(cmd *cobra.Command) (string, string, string) {
	cluster, _ := cmd.Flags().GetString("cluster")
	nameSpace, _ := cmd.Flags().GetString("namespace")
	release, _ := cmd.Flags().GetString("release")

	if isQA, _ := cmd.Flags().GetString("qa"); isQA != "" {
		cluster = "gke_inputhealth-chr_northamerica-northeast1-a_staging"
		nameSpace = "chr-qa"
		release = nameSpace + "-" + isQA
	}

	if isStaging, _ := cmd.Flags().GetString("staging"); isStaging != "" {
		cluster = "gke_inputhealth-chr_northamerica-northeast1-a_staging"
		nameSpace = "chr-staging"
		release = nameSpace + "-" + isStaging
	}

	return cluster, nameSpace, release
}

func InteractiveExec(cmd string, options ...string) error {
	console.Verbosef("[EXEC] %s %s", cmd, strings.Join(options, " "))
	process := exec.Command(cmd, options...)

	process.Stdin = os.Stdin
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	return process.Run()
}

func BufferedExec(cmd string, options ...string) (string, error) {
	console.Verbosef("[EXEC] %s %s", cmd, strings.Join(options, " "))
	process := exec.Command(cmd, options...)

	stdoutStderr, err := process.CombinedOutput()
	std := string(stdoutStderr)
	if std != "" {
		console.Verbosef("%s", strings.Trim(std, "\n")+"\n")
	} else {
		console.Verbose("")
	}

	return strings.ReplaceAll(std, "'", ""), err
}
