package util

import (
	"os"
	"os/exec"
	"strings"

	"github.com/Healthism/ih-cli/util/console"
	"github.com/spf13/cobra"
)

func ParseFlags(cmd *cobra.Command) (string, string, string) {
	release, _ := cmd.Flags().GetString("release")
	cluster, _ := cmd.Flags().GetString("cluster")
	nameSpace, _ := cmd.Flags().GetString("namespace")

	if isQA, _ := cmd.Flags().GetBool("qa"); isQA {
		cluster = "gke_inputhealth-chr_northamerica-northeast1-a_staging"
		nameSpace = "chr-qa"
	}

	if isStaging, _ := cmd.Flags().GetBool("staging"); isStaging {
		cluster = "gke_inputhealth-chr_northamerica-northeast1-a_staging"
		nameSpace = "chr-staging"
	}

	return release, cluster, nameSpace
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
