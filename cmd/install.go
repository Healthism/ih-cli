package cmd

import (
	"os/exec"

	"github.com/Healthism/ih-cli/config"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   config.INSTALL,
	Short: config.INSTALL_DESCRIPTION_SHORT,
	Long:  config.INSTALL_DESCRIPTION_LONG,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func init() {
	rootCmd.AddCommand(installCmd)
}
