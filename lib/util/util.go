package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/inconshreveable/go-update"
	"ih/lib/log"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

func Exec(cmd string, options ...string) error {
	log.Printf("[EXEC] Running command: %s %s", cmd, strings.Join(options, " "))
	process := exec.Command(cmd, options...)

	process.Stdin = os.Stdin
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	err := process.Run()
	if err != nil {
		return log.Errorf("[EXEC] Failed to run command: %v", err)
	}

	log.Print("[EXEC] Exec Finished")
	return nil
}

func ExecuteTemplate(templateText string, input map[string]interface{}, outputPath string) error {
	log.Printf("[TEMPLATE] Creating Template: %s - %s", input["release_name"], input["command"])
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

	log.Print("[TEMPLATE] Template Created")
	return nil
}

func UpdateCLI(currVersion string) error {
	gitAPI := "https://api.github.com/repos/Healthism/ih-cli/releases/latest"
	resp, err := http.Get(gitAPI)
	if err != nil {
		return log.Errorf("[UPDATE] Failed to load latest version information: %v", err)
	}

	defer resp.Body.Close()
	jsonResp := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&jsonResp)
	latestVersion := jsonResp["name"]
	if latestVersion == currVersion {
		log.Print("[UPDATE] IH CLI is up to date! :)")
		return nil
	}

	log.Print("[UPDATE] IH CLI needs update! :(")
	binLink := fmt.Sprintf("https://github.com/Healthism/ih-cli/releases/download/%s/ih", latestVersion)
	log.Printf("[UPDATE] Fatching latest version: %s", binLink)
	bin, err := http.Get(binLink)
	if err != nil {
		return log.Errorf("[UPDATE] Failed to download latest update: %v", err)
	}

	defer bin.Body.Close()
	log.Print("[UPDATE] Applying update")
	err = update.Apply(bin.Body, update.Options{})
	if err != nil {
		if rerr := update.RollbackError(err); rerr != nil {
			return log.Errorf("[UPDATE] Failed to rollback from bad update: %v", rerr)
		}
		return log.Errorf("[UPDATE] Failed to rollback from bad update: %v", err)
	}

	return log.Error("[UPDATE] IH CLI has been updated succesfully :)")
}
