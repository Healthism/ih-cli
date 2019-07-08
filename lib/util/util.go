package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"errors"
	"github.com/inconshreveable/go-update"
	"ih/lib/log"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"text/template"
)

const (
	CLI_GIT_URL      = "https://api.github.com/repos/Healthism/ih-cli/releases/latest"
	CLI_DOWNLOAD_URL = "https://github.com/Healthism/ih-cli/releases/download/%s/ih"
)

func Exec(cmd string, options ...string) error {
	log.Printf("[EXEC] %s %s", cmd, strings.Join(options, " "))
	process := exec.Command(cmd, options...)

	process.Stdin = os.Stdin
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	err := process.Run()
	if err != nil {
		return log.Errorf("[EXEC] Failed to run command: %v", err)
	}

	return nil
}

func ExecWithStdBuffer(cmd string, options ...string) (string, error) {
	log.Printf("[EXEC] %s %s", cmd, strings.Join(options, " "))
	process := exec.Command(cmd, options...)

	var out bytes.Buffer
	process.Stdin = os.Stdin
	process.Stdout = &out
	process.Stderr = os.Stderr

	err := process.Run()
	if err != nil {
		return "", log.Errorf("[EXEC] Failed to run command: %v", err)
	}

	return strings.ReplaceAll(out.String(), "'", ""), nil
}

func ExecuteTemplate(templateText string, input map[string]interface{}, outputPath string) error {
	template, err := template.New("template").Parse(templateText)
	if err != nil {
		return log.Errorf("[TEMPLATE] Failed to parse template: %v", err)
	}

	var output = bytes.NewBuffer(nil)
	err = template.Execute(output, input)
	if err != nil {
		return log.Errorf("[TEMPLATE] Failed to execute template: %v", err)
	}

	err = ioutil.WriteFile(outputPath, output.Bytes(), 0644)
	if err != nil {
		return log.Errorf("[TEMPLATE] Failed to write template: %v", err)
	}

	return nil
}

func UpdateCLI(currVersion string) error {
	versionInfo, err := http.Get(CLI_GIT_URL)
	if err != nil {
		return log.Errorf("[UPDATE] Failed to load latest version information: %v", err)
	}

	defer versionInfo.Body.Close()
	version := make(map[string]interface{})
	json.NewDecoder(versionInfo.Body).Decode(&version)
	latestVersion := version["name"]
	if latestVersion == currVersion {
		return nil
	}

	log.Print("[UPDATE] IH CLI needs update! :(")
	whatsNew := version["body"]
	if whatsNew != "" {
		fmt.Println("*********************************************************************************************************************")
		fmt.Println(whatsNew)
		fmt.Println("*********************************************************************************************************************")
	}

	latestVersionUrl := fmt.Sprintf(CLI_DOWNLOAD_URL, latestVersion)
	if runtime.GOOS == "linux" {
		latestVersionUrl += "-linux"
	}

	log.Printf("[UPDATE] Fatching latest version: %s", latestVersionUrl)
	latestRelease, err := http.Get(latestVersionUrl)
	if err != nil {
		return log.Errorf("[UPDATE] Failed to download latest update: %v", err)
	}

	defer latestRelease.Body.Close()
	err = update.Apply(latestRelease.Body, update.Options{})
	if err != nil {
		if rerr := update.RollbackError(err); rerr != nil {
			return log.Errorf("[UPDATE] Failed to rollback from bad update: %v", rerr)
		}
		return log.Errorf("[UPDATE] Failed to rollback from bad update: %v", err)
	}

	log.Positive("IH CLI has been updated succesfully")
	return errors.New("")
}
