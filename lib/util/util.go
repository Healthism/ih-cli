package util

import (
	"bytes"
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

func ExecuteTemplate(templatePath string, input map[string]interface{}, outputPath string) error {
	log.Printf("[TEMPLATE] Creating Template: %s - %s", input["release_name"], input["command"])
	template, err := template.ParseFiles(templatePath)
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

func UpdateCLI() error {
	url := "https://github.com/rooneyl/temp/releases/download/1.0.0/main"

	//Todo github api to see if latest version is higher

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to rollback from bad update: %v", err)
		return err
	}

	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		if rerr := update.RollbackError(err); rerr != nil {
			log.Printf("Failed to rollback from bad update: %v", rerr)
		}
	}

	return err
}
